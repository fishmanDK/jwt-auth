CREATE TABLE IF NOT EXISTS services
(
    id serial PRIMARY KEY,
    name varchar(255) not null,

    CONSTRAINT unique_name UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    id_service int not null,
    user_name varchar(63) not null,
    first_name varchar(63),
    last_name varchar(63),
    email varchar(255) not null,
    password_hash varchar(255) not null,
    role varchar(63) not null,

    CONSTRAINT unique_idService_email UNIQUE (id_service, email),
    CONSTRAINT unique_idService_userName UNIQUE (id_service, user_name),
    CONSTRAINT fk_sirvice_id FOREIGN KEY (id_service) REFERENCES services(id)
);

CREATE TABLE IF NOT EXISTS jwt
(
    -- id serial,
    user_id int unique,
    refresh_token varchar(255),
    expiresAt timestamp with time zone,

    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);