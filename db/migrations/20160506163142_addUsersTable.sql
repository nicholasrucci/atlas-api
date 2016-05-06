
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id serial primary key,
  first_name varchar(255),
  last_name varchar(255),
  email varchar(100) unique,
  password_hash varchar(255),
  password_salt varchar(255),
  diabled boolean
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;

