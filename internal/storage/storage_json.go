package storage

import (
	cli_tracker "cli-tracker"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	statusInProcess = "in progress"
)

type JsonStorage struct {
	fileName string
	file     *os.File
	tasks    []cli_tracker.Task
}

func NewJsonStorage(filename string) *JsonStorage {
	file := load(filename)
	return &JsonStorage{
		fileName: filename,
		file:     file,
		tasks:    loadTasks(filename),
	}
}
func (s *JsonStorage) Create(description string) (int, error) {

	task := cli_tracker.Task{
		ID:          len(s.tasks) + 1,
		Description: description,
		Status:      statusInProcess,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	s.tasks = append(s.tasks, task)

	err := s.save()

	return task.ID, err

}
func (s *JsonStorage) FindByStatus(status string) ([]cli_tracker.Task, error) {
	result := make([]cli_tracker.Task, 0)
	for i := 0; i < len(s.tasks); i++ {
		if s.tasks[i].Status == status {
			result = append(result, s.tasks[i])
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return result, nil
}

func (s *JsonStorage) Update(id int, status string) (cli_tracker.Task, error) {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Status = status
			s.tasks[i].UpdatedAt = time.Now()

			err := s.save()

			return s.tasks[i], err
		}
	}
	return cli_tracker.Task{}, errors.New("not found")
}

func loadTasks(filename string) []cli_tracker.Task {
	var tasks []cli_tracker.Task

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		return []cli_tracker.Task{}
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	return tasks
}

func (s *JsonStorage) save() error {
	data, err := json.MarshalIndent(s.tasks, "", " ")
	if err != nil {

		return err
	}
	return os.WriteFile(s.fileName, data, 0644)
}

func load(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	return file
}
