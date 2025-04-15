-- name: CreateAlbum :one
insert into albums (id, created_at, updated_at, name, name_slug, genre, img_url, artist_id)
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

-- name: GetArtistAlbums :many
select albums.name, albums.name_slug, albums.genre, albums.img_url, artists.name as artist_name, artists.name_slug artist_name_slug
from albums
join artists
on albums.artist_id = artists.id
where artists.name_slug = $1;

