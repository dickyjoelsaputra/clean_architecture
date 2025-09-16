package handler

import (
	"clean_architecture/internal/dto"
	"clean_architecture/internal/dto/request"
	"clean_architecture/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.Services) *UserHandler {
	return &UserHandler{
		userService: service.User,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req request.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "Terjadi kesalahan", err, nil)
	}

	user, err := h.userService.CreateUser(c.Context(), &req)
	if err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "Gagal membuat user", err, nil)
	}

	return dto.JSONWithMessage(c, fiber.StatusCreated, "User created successfully!", user, nil)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "ID tidak valid", err, nil)
	}

	user, err := h.userService.GetUser(c.Context(), uint(id))
	if err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "User tidak ditemukan", err, nil)
	}

	return dto.JSONWithMessage(c, fiber.StatusOK, "OK", user, nil)
}

func (h *UserHandler) ListUser(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	usersResp, pagination, err := h.userService.ListUser(c.Context(), dto.PaginationRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "Gagal mengambil data user", err, nil)
	}

	// pagination := dto.CreatePaginationResponse(usersResp.Total, usersResp.Page, usersResp.Size)
	return dto.JSONWithMessage(c, fiber.StatusOK, "OK", usersResp, pagination)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "ID tidak valid", err, nil)
	}

	var req request.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "Terjadi kesalahan", err, nil)
	}

	user, err := h.userService.UpdateUser(c.Context(), uint(id), &req)
	if err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "Gagal update user", err, nil)
	}

	return dto.JSONWithMessage(c, fiber.StatusOK, "User updated successfully", user, nil)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "ID tidak valid", err, nil)
	}

	err = h.userService.DeleteUser(c.Context(), uint(id))
	if err != nil {
		return dto.ErrorWithMessage(c, fiber.StatusBadRequest, "Gagal hapus user", err, nil)
	}

	return dto.JSONWithMessage(c, fiber.StatusOK, "User deleted successfully", nil, nil)
}
