package domain

import (
    "time"
)

type Subtask struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Completed bool `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}
