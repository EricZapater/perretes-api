CREATE TABLE users(

	ID uuid primary key not null,
	username varchar(250) not null,	
	password varchar(250) not null,
	is_active bool default true,
	is_customer bool default false,
	password_changed_at timestamptz default now()
);
CREATE INDEX idx_user_username ON users(username);
CREATE INDEX idx_user_active_username ON users(is_active, username);

CREATE TABLE customers(
	ID uuid primary key not null,
	name varchar(250) not null,
	surname varchar(250) not null,
	phone varchar(20),
	email varchar(250),
	user_id uuid references users(id) on delete cascade,
	is_active bool default true
);

CREATE INDEX idx_customers_user_id ON customers(user_id);
CREATE INDEX idx_customers_email ON customers(email);
CREATE INDEX idx_customers_phone ON customers(phone);
CREATE INDEX idx_customers_is_active ON customers(is_active);

