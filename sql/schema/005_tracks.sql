-- +goose Up

create table tracks(
	id uuid primary key,
	created_at timestamp not null,
	updated_at timestamp not null,
	name text unique not null,
	name_slug text unique not null,
	duration int not null,
	album_track_number int not null,
	artist_id uuid not null,
	album_id uuid not null,
	foreign key (artist_id) references artists(id) on delete cascade,
	foreign key (album_id) references albums(id) on delete cascade
);

-- +goose Down
drop table tracks;
