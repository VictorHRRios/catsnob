-- name: CreateReviewShort :one
insert into album_reviews (id, created_at, updated_at, user_id, album_id, title, review, score)
values (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2,
	NULL,
	NULL,
	$3
)
returning *;

-- name: CreateReviewLong :one
insert into album_reviews (id, created_at, updated_at, user_id, album_id, title, review, score)
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

-- name: GetReviewByUser :many
select album_reviews.*, albums.id, albums.name, albums.img_url
from album_reviews
join albums on albums.id = album_reviews.album_id
where user_id = $1;

-- name: GetReviewByAlbum :many
select * from album_reviews
where album_id = $1;

-- name: GetReview :one
select album_reviews.*, albums.id as album_id, albums.name as album_name, albums.img_url as album_img, users.name as username
from album_reviews
join albums on albums.id = album_reviews.album_id
join users on users.id = album_reviews.user_id
where album_reviews.id = $1;

-- name: DeleteReview :exec
delete from album_reviews
where album_reviews.id = $1;

-- name: UpdateReview :exec
update album_reviews
set 
title = $1,
review = $2,
score = $3
where 
id = $4;

