package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"go-projects/internal/service"
)

type UserHandler struct {
	userService *service.UserService
	validate    *validator.Validate
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
		validate:    validator.New(),
	}
}

// request body struct
type createUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	Dob  string `json:"dob"  validate:"required"`
}

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var body createUserRequest

	// parse JSON
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	// validate fields
	if err := h.validate.Struct(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate DOB format
	parsedDob, err := time.Parse("2006-01-02", body.Dob)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "dob must be YYYY-MM-DD")
	}

	if parsedDob.After(time.Now()) {
		return fiber.NewError(fiber.StatusBadRequest, "dob cannot be in the future")
	}

	user := h.userService.CreateUser(body.Name, body.Dob)
	return c.Status(fiber.StatusCreated).JSON(user)
}

// GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	user, found := h.userService.GetUserByID(id)
	if !found {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

// GET /users
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users := h.userService.GetAllUsers()
	return c.JSON(users)
}
