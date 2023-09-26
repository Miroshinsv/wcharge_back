-- Таблица адресов
CREATE TABLE tbl_addresses (
                               id serial PRIMARY KEY,
                               contry VARCHAR(255) NOT NULL,
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
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           deleted_at TIMESTAMP DEFAULT NULL
    --password VARCHAR(255) NOT NULL
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
                              price DECIMAL(10, 2) NOT NULL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE tbl_role (
                          id serial PRIMARY KEY,
                          role_name VARCHAR(255) NOT NULL,
                          priv VARCHAR(255) NOT NULL
);

-- добавим юзеров
INSERT INTO postgres.public.tbl_users (username, email, role_id) VALUES ('user1', 'user1@mail,ru', 1);
INSERT INTO postgres.public.tbl_users (username, email, role_id) VALUES ('user2', 'user2@mail,ru', 2);
INSERT INTO postgres.public.tbl_users (username, email, role_id) VALUES ('user3', 'user3@mail,ru', 3);

-- добавим роли
INSERT INTO postgres.public.tbl_role (role_name, priv) VALUES ('root', 'rwx');
INSERT INTO postgres.public.tbl_role (role_name, priv) VALUES ('administrator', 'rw');
INSERT INTO postgres.public.tbl_role (role_name, priv) VALUES ('user', 'r');