-- name: CreateArtist :one
insert into artists (id, created_at, updated_at, formed_at, name, biography, genre, img_url)
values (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2,
	$3,
	$4,
	$5
)
returning *;

-- name: GetArtist :one
select * from artists
where id = $1;

-- name: GetTop12Artists :many
select * from artists limit 12;
