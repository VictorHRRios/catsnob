-- name: CreateAlbumTracks :one
insert into tracks (id, created_at, updated_at, name, duration, album_track_number, artist_id, album_id)
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

-- name: GetTrack :one
select tracks.*, albums.img_url
from tracks
join albums on albums.id = tracks.album_id
where tracks.id = $1;



-- name: GetTop12Tracks :many
select distinct on (albums.name)
	tracks.id, tracks.name, tracks.duration,
	albums.name as album_name, albums.img_url as img_url,
	artists.name as artist_name
from tracks
join albums on albums.id = tracks.album_id
join artists on artists.id = albums.artist_id
limit 12;
