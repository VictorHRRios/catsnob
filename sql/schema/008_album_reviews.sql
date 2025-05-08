-- +goose Up

create table album_reviews(
	id uuid,
	created_at timestamp not null,
	updated_at timestamp not null,
	user_id uuid not null,
	album_id uuid not null,
	title varchar(50),
	review text,
	score decimal(2,1) check (score in (0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5)) not null,
	primary key (id),
	foreign key (user_id) references users(id) on delete cascade,
	foreign key (album_id) references albums(id) on delete cascade,
	unique (user_id, album_id)
);

-- +goose Down
drop table album_reviews;
