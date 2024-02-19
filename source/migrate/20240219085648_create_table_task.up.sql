CREATE TABLE IF NOT EXISTS tasks(
    id serial,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    user_id int NOT NULL,
    title varchar(50) NOT NULL,
    done_flag boolean NOT NULL,
    PRIMARY KEY (id)
)