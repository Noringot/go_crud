package task

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Todos(c *fiber.Ctx) error {
	data := Task.All()
	return c.JSON(data)
}

func Todo(c *fiber.Ctx) error {
	data := Task.Where("id", c.Params("id"))
	return c.JSON(data)
}

func TodoUpdate(c *fiber.Ctx) error {
	f := new(Fillable)
	if err := c.BodyParser(f); err != nil {
		return err
	}
	mediate := map[string]string{
		"id":      c.Params("id"),
		"name":    f.Name,
		"text":    f.Text,
		"is_done": strconv.Itoa(f.IsDone),
	}

	err := Task.UpdateById(c.Params("id"), mediate)

	if err != nil {
		return err
	}
	return nil
}

func TodoStore(c *fiber.Ctx) error {
	f := new(Fillable)
	if err := c.BodyParser(f); err != nil {
		return err
	}

	mediate := map[string]string{
		"name":    f.Name,
		"text":    f.Text,
		"is_done": strconv.Itoa(f.IsDone),
	}

	err := Task.Store(mediate)

	if err != nil {
		return err
	}

	return nil
}
func TodoDelete(c *fiber.Ctx) error {
	err := Task.Delete("id", c.Params("id"))

	if err != nil {
		return err
	}

	return nil
}
