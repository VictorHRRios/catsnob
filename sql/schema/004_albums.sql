-- +goose Up

create table albums(
	id uuid primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	name text not null,
	genre text not null,
	img_url text not null,
	artist_id uuid not null,
	foreign key (artist_id) references artists(id) on delete cascade
);

-- +goose Down
drop table albums;
