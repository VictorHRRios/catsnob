-- name: CreateAlbumList :one
insert into album_lists (id, created_at, updated_at, user_id, title)
values (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2
)
returning *;

-- name: AddAlbumToList :one
insert into AlbumLists_Albums (id, album_lists_id,album_id)
values (
	gen_random_uuid(),
	$1,
	$2
)
returning *;

-- name: GetUserLists :many
SELECT list.id, list.title
FROM album_lists as list
WHERE list.user_id = $1;

-- name: GetListName :many
SELECT list.title
FROM album_lists as list
WHERE list.id = $1;

-- name: GetAlbumsFromList :many
SELECT a.id, a.name, a.img_url
FROM albums as a
JOIN AlbumLists_Albums as ala ON a.id = ala.album_id
WHERE ala.album_lists_id = $1;

-- name: GetAlbumsNotInList :many
SELECT a.id, a.name, a.img_url
FROM albums as a
WHERE id NOT IN (
	SELECT album_id 
	FROM AlbumLists_Albums
	WHERE album_lists_id = $1
);

-- name: DeleteAlbumFromList :exec
DELETE FROM AlbumLists_Albums
WHERE album_lists_id = $1 AND album_id = $2;
