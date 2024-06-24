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

insert into addresses (country, city, address, lat, lng)
VALUES ('Georgia', 'Kutaisi', '37 Tbilisi St', 42.2702501, 42.7098063);
insert into addresses (country, city, address, lat, lng)
VALUES ('Georgia', 'Tbilisi', '44 Davit Aghmashenebeli Ave', 41.7057606, 44.7837527);
insert into addresses (country, city, address, lat, lng)
VALUES ('Georgia', 'Tbilisi', '99 Akaki Beliashvili St', 41.775668, 44.7678351);

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
    phone         text,
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

-- Таблица powerbanks
CREATE TABLE powerbanks
(
    id            serial PRIMARY KEY,
    serial_number text             NOT NULL,
    capacity      double precision NULL NULL,
    used          boolean     DEFAULT false,
    removed       boolean     DEFAULT false,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE stations
(
    id            serial PRIMARY KEY,
    serial_number text             NOT NULL,
    address       integer,
    capacity      double precision NULL NULL,
    free_capacity double precision NULL NULL,
    removed       boolean     DEFAULT false,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (address) REFERENCES addresses (id) ON DELETE SET DEFAULT
);

CREATE TABLE rel__stations__powerbanks
(
    station    integer NOT NULL,
    powerbank  integer NOT NULL,
    position   integer,
    created_ad TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (station) REFERENCES stations (id) ON DELETE CASCADE,
    FOREIGN KEY (powerbank) REFERENCES powerbanks (id) ON DELETE CASCADE
);

CREATE TABLE rel__users__powerbanks
(
    "user"     integer NOT NULL,
    powerbank  integer NOT NULL,
    created_ad TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY ("user") REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (powerbank) REFERENCES powerbanks (id) ON DELETE CASCADE
);

create table history_powerbank
(
    powerbank      integer not null,
    take_station   integer not null,
    return_station integer not null,
    take           TIMESTAMPTZ DEFAULT NOW(),
    return         TIMESTAMPTZ,
    FOREIGN KEY (powerbank) REFERENCES powerbanks (id) ON DELETE CASCADE,
    FOREIGN KEY (take_station) REFERENCES stations (id) ON DELETE CASCADE,
    FOREIGN KEY (return_station) REFERENCES stations (id) ON DELETE CASCADE
);

-- TODO
-- create table history_stations ();

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
    IF NEW.removed = true AND OLD.removed = false THEN
        NEW.deleted_at = NOW();
    END IF;
    IF NEW.removed = false AND OLD.removed = true THEN
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

