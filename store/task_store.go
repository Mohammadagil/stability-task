package store

import "stability-test-task-api/models"

var Tasks = []models.Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: false},
}

func GetAllTasks() []models.Task {
	return Tasks
}

// PERBAIKAN: Fix pointer bug
func GetTaskByID(id int) *models.Task {
	for i := range Tasks {
		if Tasks[i].ID == id {
			return &Tasks[i]
		}
	}
	return nil
}


func AddTask(task models.Task) {
	Tasks = append(Tasks, task)
}

func DeleteTask(id int) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return // Stop setelah delete
		}
	}
}

// FUNGSI BARU: Generate ID otomatis
func GetNextID() int {
	maxID := 0
	for _, t := range Tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

// FUNGSI BARU: Update task
func UpdateTask(id int, updatedTask models.Task) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks[i] = updatedTask
			return
		}
	}
}