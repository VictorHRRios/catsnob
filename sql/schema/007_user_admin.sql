-- +goose Up
alter table users
add column is_admin bool not null default false;

-- +goose Down
alter table users
drop column is_admin;
