package customers

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer Customer) (Customer, error)
	Update(ctx context.Context, customer Customer) (Customer, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindById(ctx context.Context, id uuid.UUID) (Customer, error)
	FindAll(ctx context.Context)([]Customer, error)	
	FindCustomerByUserID(ctx context.Context, userID uuid.UUID) (Customer, error)
}

type customerRepository struct{
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db:db,
	}
}

func(r *customerRepository) Create(ctx context.Context, customer Customer) (Customer, error){
	_, err := r.db.ExecContext(ctx, `
	INSERT INTO customers(id, name, surname, phone_number, email, user_id, is_active)
	VALUES($1, $2, $3, $4, $5, $6, $7)
	`,
	customer.ID, customer.Name, customer.Surname, customer.PhoneNumber, customer.Email, customer.User.ID, customer.IsActive,
	)
	if err != nil {
		return Customer{}, err
	}
	return customer, err
}

func(r *customerRepository) Update(ctx context.Context, customer Customer) (Customer, error){
	_, err := r.db.ExecContext(ctx, `
		UPDATE customers
		set name = $1,
		surname = $2,
		phone_number = $3,
		email = $4,
		is_active = $5
		WHERE id = $6`,
		customer.Name, customer.Surname, customer.PhoneNumber, customer.Email, customer.IsActive, customer.ID,
	)
	if err != nil {
		return Customer{}, err
	}
	return customer, err
}
func(r *customerRepository) Delete(ctx context.Context, id uuid.UUID) error{
	_, err := r.db.ExecContext(ctx, `DELETE FROM customers WHERE id = $1`, id)
	return err
}
func(r *customerRepository) FindById(ctx context.Context, id uuid.UUID) (Customer, error){
	var customer Customer
	err := r.db.QueryRowContext(ctx, `SELECT id, name, surname, phone_number, email, is_active, user_id FROM customers WHERE id = $1`, id,
).Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.PhoneNumber, &customer.Email, &customer.IsActive, &customer.User.ID)
if err != nil {
	return Customer{}, err
}
return customer, nil
}

func(r *customerRepository) FindAll(ctx context.Context)([]Customer, error){
	var customers []Customer
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, surname, phone_number, email, is_active, user_id FROM customers
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.PhoneNumber, &customer.Email, &customer.IsActive, &customer.User.ID); err != nil{
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func(r *customerRepository) FindCustomerByUserID(ctx context.Context, userID uuid.UUID) (Customer, error){
	var customer Customer
	err := r.db.QueryRowContext(ctx, `SELECT id, name, surname, phone_number, email, is_active, user_id FROM customers WHERE user_id = $1`, userID,
).Scan(&customer.ID, &customer.Name, &customer.Surname, &customer.PhoneNumber, &customer.Email, &customer.IsActive, &customer.User.ID)
if err != nil {
	return Customer{}, err
}
return customer, nil
}

