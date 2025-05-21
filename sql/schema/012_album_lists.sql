-- +goose Up
create table album_lists(
    id uuid primary key,
    created_at timestamp not null,
	updated_at timestamp not null,
    user_id uuid not null,
    title varchar(50),
    foreign key (user_id) references users(id) on delete cascade
);

create table AlbumLists_Albums(
    id uuid primary key,
    album_lists_id uuid not null,
    album_id uuid not null,
    foreign key (album_id) references albums(id) on delete cascade,
    foreign key (album_lists_id) references album_lists(id) on delete cascade
);

-- +goose Down
drop table if exists album_lists_relKey;
drop table album_lists;

