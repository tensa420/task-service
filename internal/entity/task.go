package entity

import (
	"github.com/google/uuid"
)

type Task struct {
	TaskUUID    uuid.UUID  `json:"task_uuid"`
	Title       string     `json:"title"`
	UserUUID    uuid.UUID  `json:"user_uuid"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}

type TaskStatus string

const (
	TaskStatusFinished TaskStatus = "finished"
	TaskStatusNew      TaskStatus = "new"
)
