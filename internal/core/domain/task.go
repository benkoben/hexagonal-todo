package domain


import (
    "time"
)

type Task struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Completed bool `json:"completed"`
    Due time.Time `json:"due"`
    CreatedAt time.Time `json:"created_at"`
}
