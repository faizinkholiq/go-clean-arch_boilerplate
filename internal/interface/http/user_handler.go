package http

import (
	"strconv"

	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/domain"
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/interface/http/request"
	"github.com/faizinkholiq/go-clean-arch_boilerplate/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(usecase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: usecase}
}

func (h *UserHandler) GetUserList(c *fiber.Ctx) error {
	users, err := h.userUseCase.GetUserList()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to retrieve users")
	}
	return c.JSON(users)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user ID")
	}

	user, err := h.userUseCase.GetUserByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return c.JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userRequest request.UserRequest
	if err := c.BodyParser(&userRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
		// return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{Message: "Invalid request"})
	}

	user := domain.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	if err := h.userUseCase.CreateUser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to create user")
		// return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{Message: "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map(map[string]interface{}{"message": "success"}))
	// return c.Status(fiber.StatusCreated).JSON(response.UserResponse{User: &user})
}
