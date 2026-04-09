package handlers

import (
	"strconv"
	"strings"

	"stability-test-task-api/models"
	"stability-test-task-api/store"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	tasks := store.GetAllTasks()
	return c.Status(fiber.StatusOK).JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Validasi ID tidak boleh kosong
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task id is required",
		})
	}

	// Validasi ID harus angka
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id, must be a number",
		})
	}

	// Validasi ID tidak boleh negatif
	if id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task id cannot be negative",
		})
	}

	task := store.GetTaskByID(id)

	if task == nil {
		// return c.Status(200).JSON(fiber.Map{ // Harusnya status 404 Not Found atau kalau di fiber.StatusNotfound karena data tidak ada
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	//Parse JSON body
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Validasi Title tidak boleh kosong
	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "title is required",
		})
	}

	// Validasi Title tidak boleh hanya spasi
	trimmedTitle := strings.TrimSpace(task.Title)
	if trimmedTitle == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "title cannot be empty or just spaces",
		})
	}

	// Validasi panjang Title 
	if len(trimmedTitle) > 100 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "title too long (max 100 characters)",
		})
	}

	// Set title yang sudah di-trim
	task.Title = trimmedTitle

	// Set default value untuk Done jika tidak disediakan
	// (Done sudah default false dari Go)

	// Generate ID otomatis
	task.ID = store.GetNextID()

	store.AddTask(task)

	return c.Status(fiber.StatusCreated).JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Validasi ID tidak boleh kosong
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task id is required",
		})
	}

	// Validasi ID harus angka
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id, must be a number",
		})
	}

	// Validasi ID tidak boleh negatif
	if id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task id cannot be negative",
		})
	}

	// Cek apakah task exist sebelum delete
	task := store.GetTaskByID(id)
	if task == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	store.DeleteTask(id)

	// Return 200 OK dengan pesan sukses
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task deleted successfully",
		"id":      id,
	})
}

// FUNGSI BARU: Update task (PUT)
func UpdateTask(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Validasi ID
	if idParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task id is required",
		})
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id, must be a number",
		})
	}

	if id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task id cannot be negative",
		})
	}

	// Cek apakah task exists
	existingTask := store.GetTaskByID(id)
	if existingTask == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	// Parse request body
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Update title jika ada
	if title, ok := updates["title"]; ok {
		if titleStr, ok := title.(string); ok {
			trimmedTitle := strings.TrimSpace(titleStr)
			if trimmedTitle == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "title cannot be empty",
				})
			}
			if len(trimmedTitle) > 100 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "title too long (max 100 characters)",
				})
			}
			existingTask.Title = trimmedTitle
		}
	}

	// Update done jika ada
	if done, ok := updates["done"]; ok {
		if doneBool, ok := done.(bool); ok {
			existingTask.Done = doneBool
		}
	}

	// Simpan perubahan
	store.UpdateTask(id, *existingTask)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "task updated successfully",
		"task":    existingTask,
	})
}
