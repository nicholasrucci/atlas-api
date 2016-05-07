
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE projects (
  id serial primary key,
  name varchar(255),
  client varchar(255),
  start_date varchar(100),
  organization_id integer references organizations(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE projects;
