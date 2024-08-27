package task

import "test_case/internal/repo"

type Task struct {
	Id      int    `json:"id"`
	userID  int    `json:"userID"`
	Content string `json:"content"`
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Add(userid int, content string) (int, error) {
	return repo.Add(userid, content)
}

func (t *Task) Get(userid int) (map[int]string, error) {
	return repo.Get(userid)
}
