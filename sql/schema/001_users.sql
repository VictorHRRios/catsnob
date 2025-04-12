-- +goose Up

create table users(
	id uuid primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	name text unique not null,
	img_url text not null
);

-- +goose Down
drop table users;
