package storage

import cli_tracker "cli-tracker"

type Storage interface {
	Create(description string) (int, error)
	FindByStatus(status string) ([]cli_tracker.Task, error)
	Delete(id int) error
	Update(id int, status string) (cli_tracker.Task, error)
}
