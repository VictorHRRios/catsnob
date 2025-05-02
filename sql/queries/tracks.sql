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

-- name: GetTrack :one
select
	tracks.*,
	albums.name_slug as album_name,
	artists.name_slug as artist_name
from tracks
join albums on albums.id = tracks.album_id
join artists on artists.id = albums.artist_id
where tracks.name_slug = $1;

-- name: GetAlbumTracks :many
select tracks.name, tracks.name_slug, tracks.duration, tracks.album_track_number, albums.name as album_name, albums.name_slug as album_name_slug, albums.img_url as img_url
from tracks
join albums
on tracks.album_id = albums.id
where albums.name_slug = $1;


-- name: GetTop12Tracks :many
select distinct on (albums.name)
	tracks.name, tracks.name_slug, tracks.duration,
	albums.name as album_name, albums.name_slug as album_name_slug, albums.img_url as img_url,
	artists.name as artist_name, artists.name_slug as artist_name_slug
from tracks
join albums on albums.id = tracks.album_id
join artists on artists.id = albums.artist_id
limit 12;
