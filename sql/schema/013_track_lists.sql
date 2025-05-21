-- +goose Up
create table track_lists(
    id uuid primary key,
    added_at timestamp,
    user_lists_id uuid not null,
    track_id uuid not null,
    foreign key (track_id) references tracks(id) on delete cascade,
    foreign key (user_lists_id) references user_lists(id_playlist_a) on delete cascade
);

-- +goose Down
drop table  track_lists CASCADE;

