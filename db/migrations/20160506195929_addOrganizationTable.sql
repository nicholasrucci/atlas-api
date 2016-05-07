
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE organizations (
  id serial primary key,
  team_name varchar(255),
  contact_name varchar(255),
  contact_email varchar(100) unique,
  contact_phone varchar(255)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE organizations;
