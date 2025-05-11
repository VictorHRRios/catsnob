-- +goose Up

create table artists(
	id uuid primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	formed_at text not null,
	name text not null,
	biography text,
	genre text not null,
	img_url text not null
);

-- +goose Down
drop table artists;
