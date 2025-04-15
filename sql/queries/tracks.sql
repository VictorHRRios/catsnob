-- name: CreateAlbumTracks :one
insert into tracks (id, created_at, updated_at, name, name_slug, duration, album_track_number, artist_id, album_id)
values (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
returning *;

