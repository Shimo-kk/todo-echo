CREATE TABLE IF NOT EXISTS users(
    id serial,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name varchar(50) NOT NULL,
    email varchar(256) NOT NULL,
    password varchar(128) NOT NULL,
    PRIMARY KEY (id)
)