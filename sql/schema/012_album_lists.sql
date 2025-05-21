-- +goose Up
create table album_lists(
    id uuid primary key,
    user_lists_id uuid not null,
    album_id uuid not null,
    foreign key (album_id) references albums(id) on delete cascade,
    foreign key (user_lists_id) references user_lists(id_playlist_a) on delete cascade
);

-- +goose Down
drop table  album_lists CASCADE;

