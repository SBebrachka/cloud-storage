CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE box_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    list_id int references box_lists(id) on delete cascade not null
);

CREATE TABLE box_items
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

CREATE TABLE lists_items
(
    id serial not null unique,
    item_id int references box_items(id) on delete cascade not null,
    list_id int references box_items(id) on delete cascade not null
);

CREATE TABLE files (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       filename VARCHAR(255) NOT NULL,
                       size BIGINT NOT NULL,
                       path VARCHAR(500) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_files_user_id ON files(user_id);