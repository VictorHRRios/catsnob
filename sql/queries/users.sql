-- name: CreateUser :one
insert into users (id, created_at, updated_at, name, img_url)
values (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2
)
returning *;

-- name: GetUser :one
select * from users
where name = $1;

-- name: GetUsers :many
select * from users;
