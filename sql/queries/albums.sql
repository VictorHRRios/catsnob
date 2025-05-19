-- name: CreateAlbum :one
insert into albums (id, created_at, updated_at, name, genre, img_url, artist_id)
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

-- name: GetAlbum :one
select * from albums where id = $1;

-- name: GetAlbums :many
select * from albums;

-- name: GetArtistAlbums :many
select albums.id, albums.name, albums.genre, albums.img_url, artists.name as artist_name
from albums
join artists
on albums.artist_id = artists.id
where artists.id = $1;

-- name: GetAlbumTracks :many
select tracks.id, tracks.name, tracks.duration, tracks.album_track_number, albums.name as album_name, albums.img_url as img_url
from tracks
join albums
on tracks.album_id = albums.id
where albums.id = $1;

-- name: GetTop12Albums :many
select 
	albums.id, albums.name, albums.genre, albums.img_url, artists.name as artist_name
from albums
join artists
on albums.artist_id = artists.id
limit 12;

