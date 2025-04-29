-- name: CreateUser :one
insert into users (id, created_at, updated_at, name, img_url, hashed_password, is_admin)
values (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2,
	$3,
	false
)
returning *;

-- name: GetUser :one
select * from users
where name = $1;

-- name: GetUserFromID :one
select * from users
where id = $1;

-- name: GetUsers :many
select * from users;
