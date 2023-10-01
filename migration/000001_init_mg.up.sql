-- Таблица адресов
CREATE TABLE tbl_addresses (
    id serial PRIMARY KEY,
    country VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    lat FLOAT NOT null,
    lng FLOAT NOT null
);


-- Таблица пользователей
CREATE TABLE tbl_users (
    id serial PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(255) NOT NULL,
    address_id INT NOT NULL,
    suspended_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Таблица заказов
--CREATE TABLE orders (
--    id INT PRIMARY KEY AUTO_INCREMENT,
--    user_id INT NOT NULL,
--    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--    status VARCHAR(255) NOT NULL,
--    FOREIGN KEY (user_id) REFERENCES users(id)
--);

-- Таблица powerbanks
CREATE TABLE tbl_powerbanks (
    id serial PRIMARY KEY,
    serial_number VARCHAR(255) NOT NULL,
    capacity INT NULL NULL,
    used INT DEFAULT 0,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE tbl_stations (
    id serial PRIMARY KEY,
    serial_number VARCHAR(255) NOT NULL,
    address_id INT not null,
    capacity INT NULL NULL,
    free_capacity INT NULL NULL,
    --price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE tbl_role (
    id serial PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL,
    priv VARCHAR(255) NOT NULL
);
