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

insert into tbl_addresses (country, city, address, lat, lng) VALUES ('Russia', 'Moscow', 'Address Moscow', 77.777, 888.888);
insert into tbl_addresses (country, city, address, lat, lng) VALUES ('Monaco', 'Monaco', 'Address Monaco', 1.783763, 24.1430);
insert into tbl_addresses (country, city, address, lat, lng) VALUES ('Germany', 'Berlin', 'Address Berlin', 545.1234, 12345.12351);

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
    phone VARCHAR(255) DEFAULT '',
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(255) NOT NULL,
    address_id INT DEFAULT 1,
    removed INT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (role_id) REFERENCES tbl_role (id) ON DELETE SET DEFAULT,
    FOREIGN KEY (address_id) REFERENCES tbl_addresses (id) ON DELETE SET DEFAULT
);

-- admin admin
insert into tbl_users (username, email, role_id, password_hash, password_salt)
values ('admin', 'admin@mail.com', 1, 'JDJhJDA0JEVqeGpML3Z2T25QT2YvTXNUcEw2RC4vZ2hzZHRDcGdmdHVnZ0hkWHVvclRTanpwbDVFOUJL', 'lXpk1n+I9Xh0kRz8djIBCrOXYj8tvxZbn4J7xgyZ+mIWH51W69Sn/xoE3wlJCdAeuklIow==');

-- user user
insert into tbl_users (username, email, role_id, password_hash, password_salt)
values ('user', 'user@mail.com', 3, 'JDJhJDA0JFZmWFl3Y29PTWVjRmc0clpYZjhjRHV3SnhYYXpKbnVIV0tpV1liRnEzSnV2ZzZ2MDV6NEd5', 'kdFEO7zJJy94rKcAQhLAcOYxZ5lIb9FUXHJ2A2zEpDkGR+4hHrNOPgWHUvcn+SZUODeEhQ==');

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

-- insert into tbl_powerbanks (serial_number, capacity) VALUES ('1', 5400);
-- insert into tbl_powerbanks (serial_number, capacity) VALUES ('2', 5400);
-- insert into tbl_powerbanks (serial_number, capacity) VALUES ('3', 5400);
-- insert into tbl_powerbanks (serial_number, capacity) VALUES ('4', 5400);
-- insert into tbl_powerbanks (serial_number, capacity) VALUES ('5', 5400);
-- insert into tbl_powerbanks (serial_number, capacity) VALUES ('6', 5400);
-- insert into tbl_powerbanks (serial_number, capacity) VALUES ('7', 5400);
-- insert into tbl_powerbanks (serial_number, capacity) VALUES ('8', 5400);

CREATE TABLE tbl_stations (
    id serial PRIMARY KEY,
    serial_number VARCHAR(255) NOT NULL,
    address_id INT DEFAULT 1,
    capacity INT NULL NULL,
    free_capacity INT NULL NULL,
    removed INT DEFAULT 0,
    --price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ  DEFAULT NOW(),
    updated_at TIMESTAMPTZ  DEFAULT NOW(),
    deleted_at TIMESTAMPTZ  DEFAULT NULL,
    FOREIGN KEY (address_id) REFERENCES tbl_addresses (id) ON DELETE SET DEFAULT
);

-- insert into tbl_stations (serial_number, address_id, capacity, free_capacity) VALUES ('RL3H082111030142', 1, 8, 6);
-- insert into tbl_stations (serial_number, address_id, capacity, free_capacity)
--     VALUES ('st1', 2, 10, 8);
-- insert into tbl_stations (serial_number, address_id, capacity, free_capacity)
--     VALUES ('st1', 2, 10, 10);

CREATE TABLE tbl_station_powerbank (
    id serial PRIMARY KEY,
    station_id INT NOT NULL,
    powerbank_id INT NOT NULL,
    position INT,
    created_ad TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (station_id) REFERENCES tbl_stations (id) ON DELETE CASCADE,
    FOREIGN KEY (powerbank_id) REFERENCES tbl_powerbanks (id) ON DELETE CASCADE
);

-- insert into tbl_station_powerbank (station_id, powerbank_id) VALUES (1, 1);
-- insert into tbl_station_powerbank (station_id, powerbank_id, position) VALUES (1, 2, 2);
-- insert into tbl_station_powerbank (station_id, powerbank_id, position) VALUES (1, 3, 3);
-- insert into tbl_station_powerbank (station_id, powerbank_id, position) VALUES (1, 4, 4);
-- insert into tbl_station_powerbank (station_id, powerbank_id) VALUES (1, 5);
-- insert into tbl_station_powerbank (station_id, powerbank_id, position) VALUES (1, 6, 6);
-- insert into tbl_station_powerbank (station_id, powerbank_id, position) VALUES (1, 7, 7);
-- insert into tbl_station_powerbank (station_id, powerbank_id, position) VALUES (1, 8, 8);

CREATE TABLE tbl_user_powerbank (
    id serial PRIMARY KEY,
    user_id INT NOT NULL,
    powerbank_id INT NOT NULL,
    created_ad TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES tbl_users (id) ON DELETE CASCADE,
    FOREIGN KEY (powerbank_id) REFERENCES tbl_powerbanks (id) ON DELETE CASCADE
);

-- insert into tbl_user_powerbank (user_id, powerbank_id) VALUES (2, 5);

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

-- Триггеры на update - при установки поля remote
CREATE OR REPLACE FUNCTION trigger_set_timestamp_delete()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.removed = 1 AND OLD.removed = 0 THEN
        NEW.deleted_at = NOW();
    END IF;
    IF NEW.removed = 0 AND OLD.removed = 1 THEN
        NEW.deleted_at = NULL;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER set_timestamp_delete_powerbank
    BEFORE UPDATE OF removed ON tbl_powerbanks
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_delete();

CREATE TRIGGER set_timestamp_delete_station
    BEFORE UPDATE OF removed ON tbl_stations
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_delete();

CREATE TRIGGER set_timestamp_delete_user
    BEFORE UPDATE OF removed ON tbl_users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_delete();

