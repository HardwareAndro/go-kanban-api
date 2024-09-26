package services

import (
	"github.com/HardwareAndro/go-kanban-api/internal/config"
	"github.com/HardwareAndro/go-kanban-api/internal/models"
	repository "github.com/HardwareAndro/go-kanban-api/internal/repositories"
	"github.com/HardwareAndro/go-kanban-api/internal/shared/constants"
	"log"
	"os"
)

type TaskService struct {
	taskRepository *repository.GenericRepository[model.Task]
	App            config.GoAppTools
}

func NewTaskService(taskRepository *repository.GenericRepository[model.Task]) *TaskService {
	InfoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
	var app config.GoAppTools
	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger
	return &TaskService{
		taskRepository: taskRepository,
		App:            app,
	}
}

func (ts *TaskService) AddTask(task *model.Task) (*model.Task, error) {
	_, err := ts.taskRepository.Create(task)
	if err != nil {
		ts.App.ErrorLogger.Fatalln(constants.ERR_ADD_TASK, err)
		return nil, err
	}
	ts.App.InfoLogger.Println(constants.SUCCESS_ADD_TASK, task.ID)
	return task, nil
}

func (ts *TaskService) GetTasks() ([]model.Task, error) {
	tasks, err := ts.taskRepository.FindAll()
	if err != nil {
		ts.App.ErrorLogger.Println(constants.ERR_TASK_NOT_FOUND, err)
		return []model.Task{}, nil // Return an empty slice instead of nil
	}
	return tasks, nil
}

func (ts *TaskService) GetTaskById(id string) (*model.Task, error) {
	task, err := ts.taskRepository.FindById(id)
	if err != nil {
		ts.App.ErrorLogger.Println(constants.ERR_TASK_NOT_FOUND, err)
		return nil, err
	}
	return task, nil
}

func (ts *TaskService) UpdateTaskById(task *model.Task, id string) (*model.Task, error) {
	err := ts.taskRepository.Update(id, task)
	if err != nil {
		ts.App.ErrorLogger.Println(constants.ERR_UPDATE_TASK, err)
		return nil, err
	}
	ts.App.InfoLogger.Println(constants.SUCCESS_UPDATE_TASK, id)
	return task, nil
}

func (ts *TaskService) DeleteTaskById(id string) (int64, error) {
	deleteResult, err := ts.taskRepository.Delete(id)
	if err != nil {
		ts.App.ErrorLogger.Println(constants.ERR_DELETE_TASK, err)
		return 0, err
	}
	ts.App.InfoLogger.Println(constants.SUCCESS_DELETE_TASK, id)
	return deleteResult, nil
}
