package controllers

import (
	"net/http"

	"github.com/HardwareAndro/go-kanban-api/app/api/services"
	models "github.com/HardwareAndro/go-kanban-api/app/models"
	constants "github.com/HardwareAndro/go-kanban-api/app/shared/constants"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

// AddTask creates a new task
func (tc *TaskController) AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ERR_INVALID_INPUT})
		return
	}
	createdTask, err := tc.taskService.AddTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_ADD_TASK})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

// GetTasks retrieves all tasks
func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_RETRIEVE_TASKS})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskById retrieves a task by ID
func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskService.GetTaskById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": constants.ERR_TASK_NOT_FOUND})
		return
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTaskById updates a task by ID
func (tc *TaskController) UpdateTaskById(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ERR_INVALID_INPUT})
		return
	}
	updatedTask, err := tc.taskService.UpdateTaskById(&task, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_UPDATE_TASK})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTaskById deletes a task by ID
func (tc *TaskController) DeleteTaskById(c *gin.Context) {
	id := c.Param("id")
	if _, err := tc.taskService.DeleteTaskById(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_DELETE_TASK})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
