// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type NewTaskInput struct {
	ImplantUUID string   `json:"implantUuid"`
	Type        TaskType `json:"type"`
	Payload     string   `json:"payload"`
}

type TaskType string

const (
	TaskTypeNoop   TaskType = "NOOP"
	TaskTypeCmd    TaskType = "CMD"
	TaskTypeScript TaskType = "SCRIPT"
)

var AllTaskType = []TaskType{
	TaskTypeNoop,
	TaskTypeCmd,
	TaskTypeScript,
}

func (e TaskType) IsValid() bool {
	switch e {
	case TaskTypeNoop, TaskTypeCmd, TaskTypeScript:
		return true
	}
	return false
}

func (e TaskType) String() string {
	return string(e)
}

func (e *TaskType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskType", str)
	}
	return nil
}

func (e TaskType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}