package model

import "time"

type Task struct {
    ID int
    Title string
    Description string
    AssignedUser int
    Status string
    DueDate time.Time
}