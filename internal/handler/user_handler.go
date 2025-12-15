package handler

import (
	"strconv"

	"go-projects/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
		Dob  string `json:"dob"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	user := h.userService.CreateUser(body.Name, body.Dob)
	return c.Status(fiber.StatusCreated).JSON(user)
}

// GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, found := h.userService.GetUserByID(id)
	if !found {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}
