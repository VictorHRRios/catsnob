-- +goose Up

create table artists(
	id uuid primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	born_at timestamp not null,
	name text unique not null,
	biography text,
	type text not null,
	genre text not null,
	img_url text not null
);

-- +goose Down
drop table artists;
