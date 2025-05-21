-- name: CreateUserList :one
insert into user_lists (id_playlist_a, created_at, updated_at,name_,description_,type_, user_id)
values (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2,
	$3,
	$4
)
returning *;

-- name: AddAlbumToList :one
insert into album_lists (id, user_lists_id,album_id)
values (
	gen_random_uuid(),
	$1,
	$2
)
returning *;

-- name: GetUserLists :many
SELECT list.id_playlist_a, list.name_
FROM user_lists as list
WHERE list.user_id = $1;

-- name: GetListName :many
SELECT list.name_
FROM user_lists as list
WHERE list.id_playlist_a = $1;

-- name: GetAlbumsFromList :many
SELECT a.id, a.name, a.img_url
FROM albums as a
JOIN album_lists as al ON a.id = al.album_id
WHERE al.user_lists_id = $1;

-- name: GetAlbumsNotInList :many
SELECT a.id, a.name, a.img_url
FROM albums as a
WHERE id NOT IN (
	SELECT album_id 
	FROM album_lists
	WHERE user_lists_id = $1
);

-- name: DeleteAlbumFromList :exec
DELETE FROM album_lists
WHERE user_lists_id = $1 AND album_id = $2;
