-- +goose Up
create table user_lists(
    id_playlist_a uuid primary key,
    created_at timestamp not null,
	updated_at timestamp not null,
    name_ varchar(50),
    description_ varchar(100),
    type_ varchar(10),
    user_id uuid not null,
    foreign key (user_id) references users(id) on delete cascade
);

--
-- create table album_lists(
--     id uuid primary key,
--     created_at timestamp not null,
-- 	updated_at timestamp not null,
--     user_id uuid not null,
--     title varchar(50),
--     foreign key (user_id) references users(id) on delete cascade
-- );

-- +goose Down
drop table user_lists;