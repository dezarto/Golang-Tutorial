package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type todo struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json: "completed"`
}

var todos = []todo{
	{ID: "1", Content: "Yatağını topla", Completed: false},
	{ID: "2", Content: "Sabah kahveni iç", Completed: false},
	{ID: "3", Content: "Toplantıyı kaçırma", Completed: false},
}

func addNewTodo(c *fiber.Ctx) error {
	var newTodo todo

	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	todos = append(todos, newTodo)
	return c.JSON(newTodo)
}

func add(c *fiber.Ctx) error {
	var newTodo todo

	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	todos = append(todos, newTodo)
	return c.JSON(todos)
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	for index, data := range todos {
		if data.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
		}
	}

	return c.JSON(todos)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateTodo todo

	if err := c.BodyParser(&updateTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	todo, err := getTodoById(id)

	if err != nil {
		return errors.New("id does not exist!")
	}

	todo.ID = updateTodo.ID
	todo.Content = updateTodo.Content
	todo.Completed = updateTodo.Completed

	return c.JSON(todos)
}

func getAllTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func getOneTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := getTodoById(id)

	if err != nil {
		return errors.New("something went wrong!")
	}
	return c.JSON(todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("something went wrong!")
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! This is main area :)")
	})

	app.Get("/api/getalltodos", getAllTodos)
	app.Get("/api/gettodo/:id", getOneTodo)
	app.Patch("/api/addTodo", addNewTodo)
	app.Patch("/api/test", add)
	app.Delete("/api/delete/:id", deleteTodo)
	app.Put("/api/update/:id", updateTodo)

	app.Listen(":3000")
}
