package service

import (
	"clean_architecture/internal/dto"
	"clean_architecture/internal/dto/request"
	"clean_architecture/internal/dto/response"
	"clean_architecture/internal/entity"
	"clean_architecture/internal/repository"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo    repository.UserRepository
	productRepo repository.ProductRepository
}

func NewUserService(repos *repository.Repositories) *UserService {
	return &UserService{
		userRepo:    repos.User,
		productRepo: repos.Product,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *request.CreateUserRequest) (interface{}, error) {
	// Check if email already exists
	if _, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Manual mapping dari DTO ke Entity
	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Manual mapping dari Entity ke Response DTO

	return "User created!", nil
}

func (s *UserService) GetUser(ctx context.Context, id uint) (interface{}, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return user, nil
}

func (s *UserService) ListUser(ctx context.Context, paginate dto.PaginationRequest) (interface{}, *dto.MetaPagination, error) {

	users, count, err := s.userRepo.GetAll(ctx, paginate.GetLimit(), paginate.GetOffset())
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, fmt.Errorf("failed to get users: %w", err)
	}

	var userResponses = make([]response.UserResponse, 0)
	for _, user := range users {
		userResponses = append(userResponses, response.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			// Empty address since entity doesn't have it
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	// panic(true)

	// return nil, nil, errors.New("failed to list user")

	return userResponses, dto.CreatePaginationResponse(count, paginate.GetPage(), paginate.GetLimit()), nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, req *request.UpdateUserRequest) (interface{}, error) {
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Update fields from DTO
	if req.Username != "" {
		existingUser.Username = req.Username
	}
	if req.Email != "" {
		existingUser.Email = req.Email
	}

	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return "User updated!", nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	if _, err := s.userRepo.GetByID(ctx, id); err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	return s.userRepo.Delete(ctx, id)
}
