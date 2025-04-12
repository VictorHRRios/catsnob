-- +goose Up

create table artist_follows(
	id uuid,
	created_at timestamp not null,
	updated_at timestamp not null,
	user_id uuid not null,
	artist_id uuid not null,
	primary key (id),
	foreign key (user_id) references users(id) on delete cascade,
	foreign key (artist_id) references artists(id) on delete cascade,
	unique (user_id, artist_id)
);

-- +goose Down
drop table artist_follows;
