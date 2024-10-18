package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task represents the task model
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

// Get all tasks
func GetTasks(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, title, description, due_date, status FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
}

// Create a new task
func CreateTask(c *gin.Context, db *sql.DB) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert task into the database
	stmt, err := db.Prepare("INSERT INTO tasks(title, description, due_date, status) VALUES(?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(newTask.Title, newTask.Description, newTask.DueDate, newTask.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	newTask.ID = int(id)
	c.JSON(http.StatusCreated, newTask)
}

// Update a task
func UpdateTask(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := db.Prepare("UPDATE tasks SET title=?, description=?, due_date=?, status=? WHERE id=?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedTask.Title, updatedTask.Description, updatedTask.DueDate, updatedTask.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedTask.ID, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, updatedTask)
}

// Delete a task
func DeleteTask(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	stmt, err := db.Prepare("DELETE FROM tasks WHERE id=?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// Register routes
func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/tasks", func(c *gin.Context) {
		GetTasks(c, db)
	})
	r.POST("/tasks", func(c *gin.Context) {
		CreateTask(c, db)
	})
	r.PUT("/tasks/:id", func(c *gin.Context) {
		UpdateTask(c, db)
	})
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		DeleteTask(c, db)
	})
}
