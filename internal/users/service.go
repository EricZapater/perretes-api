package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, request UserRequest) (User, error)
	Update(ctx context.Context, id string, request UserRequest)(User, error)
	Delete(ctx context.Context, id string) (error)
	ChangePassword(ctx context.Context, request ChangePasswordRequest) (User, error)	
	FindByUsername(ctx context.Context, username string) (User, error)
	FindByID(ctx context.Context, id string) (User, error)	
	FindAll(ctx context.Context) ([]User, error)	
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func(s *userService) Create(ctx context.Context, request UserRequest) (User, error) {
	// Validate the request
	if request.Username == "" || request.Password == ""  {
		return User{} , ErrInvalidRequest
	}

	// Check if the username is already taken
	_, err := s.repo.FindByUsername(ctx, request.Username)	
	if err == nil {		
		return User{}, ErrUsernameTaken
	}

	if err != ErrUserNotFound {		
		return User{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
    if err != nil {
        return User{}, err
    }
	// Create a new User instance
	now := time.Now()
	user := User{
		ID:       uuid.New(),
		Username: request.Username,
		Password: string(hashedPassword),		
		IsActive: true,		
		IsCustomer: *request.IsCustomer,		
		PasswordChangedAt: &now,		
	}

	// Insert the user into the database
	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return User{}, err
	}

	//Create a validation string store in redis and send it to the user via email or sms

	return createdUser, nil
}

func(s *userService) Update(ctx context.Context,id string,  request UserRequest)(User, error){
	if id == "" || request.Username == ""  {
		return User{} , ErrInvalidRequest
	}

	existingUser, err := s.repo.FindByID(ctx, uuid.MustParse(id))
	if err != nil && !errors.Is(err, ErrUserNotFound){
		return User{}, fmt.Errorf("something went wrong getting the user")
	}
	if errors.Is(err, ErrUserNotFound) {
		return User{}, err
	}
	if !existingUser.IsActive {
		return User{}, ErrInactiveUser
	}
	now := time.Now()
	user := User{
		ID:       uuid.MustParse(id),		
		Username: request.Username,	
		IsCustomer: *request.IsCustomer,			
		PasswordChangedAt: &now,		
	}

	response, err := s.repo.Update(ctx, user)
	if err != nil {
		return User{}, err
	}
	response.IsCustomer = existingUser.IsCustomer	
	return response, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return ErrInvalidID
	}
	existingUser, err := s.repo.FindByID(ctx, uuid.MustParse(id))
	if err != nil && !errors.Is(err, ErrUserNotFound){
		return  fmt.Errorf("something went wrong getting the user")
	}
	if errors.Is(err, ErrUserNotFound) {
		return err
	}
	if !existingUser.IsActive {
		return ErrInactiveUser
	}

	err = s.repo.Delete(ctx, parsedID)
	if err != nil {		
		return err
	}

	return nil
}

func (s *userService) ChangePassword(ctx context.Context, request ChangePasswordRequest) (User, error) {
	if request.ID == "" || request.Password == "" {
		return User{}, ErrInvalidRequest
	}

	existingUser, err := s.repo.FindByID(ctx, uuid.MustParse(request.ID))
	if err != nil && !errors.Is(err, ErrUserNotFound){
		return User{}, fmt.Errorf("something went wrong getting the user")
	}
	if errors.Is(err, ErrUserNotFound) {
		return User{}, err
	}
	if !existingUser.IsActive {
		return User{}, ErrInactiveUser
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	request.Password = string(hashedPassword)

	response, err := s.repo.ChangePassword(ctx, request)
	if err != nil {
		return User{}, err
	}
	return response, nil
}

func (s *userService) FindByUsername(ctx context.Context, username string) (User, error) {
	if username == "" {
		return User{}, ErrInvalidRequest
	}

	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, err
	}

	return user, nil
}

func (s *userService) FindByID(ctx context.Context, id string) (User, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return User{}, ErrInvalidID
	}

	user, err := s.repo.FindByID(ctx, parsedID) // Usa l'ID validat
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, err
	}

	return user, nil
}

func (s *userService) FindAll(ctx context.Context) ([]User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var Users []User
	for _, user := range users {
		Users = append(Users, user)
	}

	return Users, nil
}
