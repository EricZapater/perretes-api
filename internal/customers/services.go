package customers

import (
	"context"
	"database/sql"
	"perretes-api/internal/users"

	"github.com/google/uuid"
)

type CustomerService interface {
	Create(ctx context.Context, request CustomerRequest)(Customer, error)
	Update(ctx context.Context,id string, request CustomerRequest)(Customer, error)
	Delete(ctx context.Context, id string)(error)
	FindByID(ctx context.Context, id string)(Customer, error)
	FindAll(ctx context.Context)([]Customer, error)
	FindCustomerByUserID(ctx context.Context, userID string) (Customer, error)	
}

type customerService struct {
	repo CustomerRepository
	usersService users.UserService
}

func NewCustomerService(repo CustomerRepository, usersService users.UserService) CustomerService{
	return &customerService{repo, usersService}
}

func(s *customerService) Create(ctx context.Context, request CustomerRequest)(Customer, error){
	if request.Name == "" || request.PhoneNumber == "" || request.Surname == "" || request.Email == "" || 
	request.Username == "" || request.Password == "" {
		return Customer{}, ErrInvalidRequest
	}
	iscustomer := true
	user := users.UserRequest {		
		Username: request.Username,
		Password: request.Password,		
		IsCustomer: &iscustomer,	
	}	
	createdUser, err := s.usersService.Create(ctx, user)
	if err != nil {
		return Customer{}, err
	}

	customer := Customer{
		ID: uuid.New(),
		Name: request.Name,
		Surname: request.Surname,
		PhoneNumber: request.PhoneNumber,
		Email: request.Email,
		User: createdUser,
		IsActive: true,
	}

	return s.repo.Create(ctx, customer)
}

func(s *customerService) Update(ctx context.Context, id string, request CustomerRequest)(Customer, error){
	customerID, err := uuid.Parse(id)
	if err != nil {
		return Customer{}, ErrInvalidID
	}
	if request.Name == "" || request.PhoneNumber == "" || request.Surname == "" || request.Email == "" || 
	request.Username == "" || request.Password == "" {
		return Customer{}, ErrInvalidRequest
	}
	customer := Customer{
		ID: customerID,
		Name: request.Name,
		Surname: request.Surname,
		PhoneNumber: request.PhoneNumber,
		Email: request.Email,
		IsActive: true,
	}

	return s.repo.Update(ctx, customer)	
}

func(s *customerService) Delete(ctx context.Context, id string)(error){
	customerID, err := uuid.Parse(id)
	if err != nil {
		return ErrInvalidID
	}
	return s.repo.Delete(ctx, customerID)
}
func(s *customerService) FindByID(ctx context.Context, id string)(Customer, error){
	customerID, err := uuid.Parse(id)
	if err != nil {
		return Customer{}, ErrInvalidID
	}
	return s.repo.FindById(ctx, customerID)
}
func(s *customerService) FindAll(ctx context.Context)([]Customer, error){
	return s.repo.FindAll(ctx)
}

func(s *customerService) FindCustomerByUserID(ctx context.Context, userID string) (Customer, error){
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return Customer{}, ErrInvalidID
	}

	customer, err := s.repo.FindCustomerByUserID(ctx, userUUID)
	if err != nil && err != sql.ErrNoRows {
		return Customer{}, err
	}
	if err == sql.ErrNoRows {
		return Customer{}, nil
	}
	return customer, nil
}