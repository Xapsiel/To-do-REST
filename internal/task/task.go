package task

import (
	"test_case/internal/repo"
)

type Task struct {
	Id      int    `json:"id"`
	UserID  int    `json:"userID"`
	Content string `json:"content"`
	repo    *repo.Repo
}

func NewTask(userid int, content string) *Task {
	return &Task{UserID: userid, Content: content, repo: repo.New()}
}

func (t *Task) Add() (int, error) {
	return t.repo.Add(t.UserID, t.Content)
}

func (t *Task) Get(userid int) (map[int]string, error) {

	return t.repo.Get(userid)
}
