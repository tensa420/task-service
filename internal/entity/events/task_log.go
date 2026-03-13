package events

import (
	"time"

	"github.com/google/uuid"
)

type TaskLog struct {
	LogUUID    uuid.UUID
	Type       LogType
	UserUUID   uuid.UUID
	TaskUUID   uuid.UUID
	Created_at time.Time
}

type LogType string

var (
	LogTypeFinish LogType = "finish"
	LogTypeDelete LogType = "delete"
	LogTypeCreate LogType = "create"
)
