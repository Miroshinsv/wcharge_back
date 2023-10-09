-- Timezone
SET TIME ZONE 'Europe/Moscow';

-- Таблица адресов
CREATE TABLE tbl_addresses (
    id serial PRIMARY KEY,
    country VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    lat FLOAT NOT NULL,
    lng FLOAT NOT NULL
);

-- таблца ролей
CREATE TABLE tbl_role (
    id serial PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL UNIQUE,
    role_privileges INT NOT NULL
);

INSERT INTO tbl_role(role_name, role_privileges) VALUES ('admin', 0);
INSERT INTO tbl_role(role_name, role_privileges) VALUES ('employee', 1);
INSERT INTO tbl_role(role_name, role_privileges) VALUES ('user', 2);

-- Таблица пользователей
CREATE TABLE tbl_users (
    id serial PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    role_id INT DEFAULT 3,
    phone VARCHAR(255) DEFAULT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(255) NOT NULL,
    address_id INT DEFAULT NULL,
    removed INT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (role_id) REFERENCES tbl_role (id) ON DELETE SET DEFAULT,
    FOREIGN KEY (address_id) REFERENCES tbl_addresses (id) ON DELETE SET DEFAULT
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
    removed INT DEFAULT 0,
    -- price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE tbl_stations (
    id serial PRIMARY KEY,
    serial_number VARCHAR(255) NOT NULL,
    address_id INT DEFAULT NULL,
    capacity INT NULL NULL,
    free_capacity INT NULL NULL,
    removed INT DEFAULT 0,
    --price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ  DEFAULT NOW(),
    updated_at TIMESTAMPTZ  DEFAULT NOW(),
    deleted_at TIMESTAMPTZ  DEFAULT NULL,
    FOREIGN KEY (address_id) REFERENCES tbl_addresses (id) ON DELETE SET DEFAULT
);

CREATE TABLE tbl_station_powerbank (
    id serial PRIMARY KEY,
    station_id INT NOT NULL,
    powerbank_id INT NOT NULL,
    created_ad TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (station_id) REFERENCES tbl_stations (id) ON DELETE CASCADE,
    FOREIGN KEY (powerbank_id) REFERENCES tbl_powerbanks (id) ON DELETE CASCADE
);

CREATE TABLE tbl_user_powerbank (
    id serial PRIMARY KEY,
    user_id INT NOT NULL,
    powerbank_id INT NOT NULL,
    created_ad TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES tbl_users (id) ON DELETE CASCADE,
    FOREIGN KEY (powerbank_id) REFERENCES tbl_powerbanks (id) ON DELETE CASCADE
);

-- Триггеры на update
CREATE OR REPLACE FUNCTION trigger_set_timestamp_update()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp_update_users
    BEFORE UPDATE ON tbl_users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_update();

CREATE TRIGGER set_timestamp_update_powerbanks
    BEFORE UPDATE ON tbl_powerbanks
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_update();

CREATE TRIGGER set_timestamp_update_stations
    BEFORE UPDATE ON tbl_stations
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_update();

