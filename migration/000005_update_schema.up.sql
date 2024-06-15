-- Timezone
SET TIME ZONE 'Europe/Moscow';

-- Таблица адресов
CREATE TABLE addresses
(
    id      serial PRIMARY KEY,
    country text             NOT NULL,
    city    text             NOT NULL,
    address text             NOT NULL,
    lat     double precision NOT NULL,
    lng     double precision NOT NULL
);

-- insert into addresses (country, city, address, lat, lng) VALUES ('Russia', 'Moscow', 'Address Moscow', 77.777, 888.888);
-- insert into addresses (country, city, address, lat, lng) VALUES ('Monaco', 'Monaco', 'Address Monaco', 1.783763, 24.1430);
-- insert into addresses (country, city, address, lat, lng) VALUES ('Germany', 'Berlin', 'Address Berlin', 545.1234, 12345.12351);

-- таблца ролей
CREATE TABLE roles
(
    id         serial PRIMARY KEY,
    name       text    NOT NULL UNIQUE,
    privileges integer NOT NULL
);

INSERT into roles(name, privileges)
VALUES ('admin', 0);
INSERT into roles(name, privileges)
VALUES ('employee', 1);
INSERT into roles(name, privileges)
VALUES ('user', 2);

-- Таблица пользователей
CREATE TABLE users
(
    id            serial PRIMARY KEY,
    username      text NOT NULL UNIQUE,
    email         text NOT NULL UNIQUE,
    role          integer     DEFAULT 3,
    phone         text        DEFAULT '',
    password_hash text NOT NULL,
    password_salt text NOT NULL,
    address       integer,
    removed       boolean     DEFAULT false,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (role) REFERENCES roles (id) ON DELETE SET DEFAULT,
    FOREIGN KEY (address) REFERENCES addresses (id) ON DELETE SET DEFAULT
);

-- admin admin
insert into users (username, email, role, password_hash, password_salt)
values ('admin', 'admin@mail.com', 1,
        'JDJhJDA0JEVqeGpML3Z2T25QT2YvTXNUcEw2RC4vZ2hzZHRDcGdmdHVnZ0hkWHVvclRTanpwbDVFOUJL',
        'lXpk1n+I9Xh0kRz8djIBCrOXYj8tvxZbn4J7xgyZ+mIWH51W69Sn/xoE3wlJCdAeuklIow==');

-- user user
insert into users (username, email, role, password_hash, password_salt)
values ('user', 'user@mail.com', 3, 'JDJhJDA0JFZmWFl3Y29PTWVjRmc0clpYZjhjRHV3SnhYYXpKbnVIV0tpV1liRnEzSnV2ZzZ2MDV6NEd5',
        'kdFEO7zJJy94rKcAQhLAcOYxZ5lIb9FUXHJ2A2zEpDkGR+4hHrNOPgWHUvcn+SZUODeEhQ==');

-- Таблица заказов
--CREATE TABLE orders (
--    id serial PRIMARY KEY AUTO_INCREMENT,
--    user_id integer NOT NULL,
--    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--    status text NOT NULL,
--    FOREIGN KEY (user_id) REFERENCES users(id)
--);

-- Таблица powerbanks
CREATE TABLE powerbanks
(
    id            serial PRIMARY KEY,
    serial_number text             NOT NULL,
    capacity      double precision NULL NULL,
    used          boolean     DEFAULT false,
    removed       boolean     DEFAULT false,
    -- price DECIMAL(10, 2) NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ DEFAULT NULL
);

-- insert into powerbanks (serial_number, capacity) VALUES ('1', 5400);
-- insert into powerbanks (serial_number, capacity) VALUES ('2', 5400);
-- insert into powerbanks (serial_number, capacity) VALUES ('3', 5400);
-- insert into powerbanks (serial_number, capacity) VALUES ('4', 5400);
-- insert into powerbanks (serial_number, capacity) VALUES ('5', 5400);
-- insert into powerbanks (serial_number, capacity) VALUES ('6', 5400);
-- insert into powerbanks (serial_number, capacity) VALUES ('7', 5400);
-- insert into powerbanks (serial_number, capacity) VALUES ('8', 5400);

CREATE TABLE stations
(
    id            serial PRIMARY KEY,
    serial_number text             NOT NULL,
    address       integer     DEFAULT 1,
    capacity      double precision NULL NULL,
    free_capacity double precision NULL NULL,
    removed       boolean     DEFAULT false,
    --price DECIMAL(10, 2) NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (address) REFERENCES addresses (id) ON DELETE SET DEFAULT
);

-- insert into stations (serial_number, address_id, capacity, free_capacity) VALUES ('RL3H082111030142', 1, 8, 6);
-- insert into stations (serial_number, address_id, capacity, free_capacity)
--     VALUES ('st1', 2, 10, 8);
-- insert into stations (serial_number, address_id, capacity, free_capacity)
--     VALUES ('st1', 2, 10, 10);

CREATE TABLE rel__stations__powerbanks
(
--     id           serial PRIMARY KEY,
    station    integer NOT NULL,
    powerbank  integer NOT NULL,
    position   integer,
    created_ad TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (station) REFERENCES stations (id) ON DELETE CASCADE,
    FOREIGN KEY (powerbank) REFERENCES powerbanks (id) ON DELETE CASCADE
);

-- insert into station_powerbank (station_id, powerbank_id) VALUES (1, 1);
-- insert into station_powerbank (station_id, powerbank_id, position) VALUES (1, 2, 2);
-- insert into station_powerbank (station_id, powerbank_id, position) VALUES (1, 3, 3);
-- insert into station_powerbank (station_id, powerbank_id, position) VALUES (1, 4, 4);
-- insert into station_powerbank (station_id, powerbank_id) VALUES (1, 5);
-- insert into station_powerbank (station_id, powerbank_id, position) VALUES (1, 6, 6);
-- insert into station_powerbank (station_id, powerbank_id, position) VALUES (1, 7, 7);
-- insert into station_powerbank (station_id, powerbank_id, position) VALUES (1, 8, 8);

CREATE TABLE rel__users__powerbanks
(
--     id           serial PRIMARY KEY,
    "user"     integer NOT NULL,
    powerbank  integer NOT NULL,
    created_ad TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY ("user") REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (powerbank) REFERENCES powerbanks (id) ON DELETE CASCADE
);

-- insert into user_powerbank (user_id, powerbank_id) VALUES (2, 5);

-- Триггеры на update
CREATE OR REPLACE FUNCTION trigger_set_timestamp_update()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp_update_users
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_update();

CREATE TRIGGER set_timestamp_update_powerbanks
    BEFORE UPDATE
    ON powerbanks
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_update();

CREATE TRIGGER set_timestamp_update_stations
    BEFORE UPDATE
    ON stations
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_update();

-- Триггеры на update - при установки поля remote
CREATE OR REPLACE FUNCTION trigger_set_timestamp_delete()
    RETURNS TRIGGER AS
$$
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
    BEFORE UPDATE OF removed
    ON powerbanks
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_delete();

CREATE TRIGGER set_timestamp_delete_station
    BEFORE UPDATE OF removed
    ON stations
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_delete();

CREATE TRIGGER set_timestamp_delete_user
    BEFORE UPDATE OF removed
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp_delete();

