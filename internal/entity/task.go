package entity

import (
	"github.com/google/uuid"
)

type Task struct {
	TaskUUID    uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	UserUUID    uuid.UUID  `json:"user_id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}

type TaskStatus string

const (
	TaskStatusFinished TaskStatus = "finished"
	TaskStatusNew      TaskStatus = "new"
)
