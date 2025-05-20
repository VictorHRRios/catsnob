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

-- name: CreateAlbumList_relKey :one
insert into album_lists_relKey (id, album_lists_id,album_id)
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

-- name: GetAlbumsFromList :many
SELECT albums.id, albums.name, albums.img_url
FROM album_lists
JOIN album_lists_relKey ON album_lists.id = album_lists_relKey.album_lists_id
JOIN albums ON albums.id = album_lists_relKey.album_id
WHERE album_lists.user_id = $1;
