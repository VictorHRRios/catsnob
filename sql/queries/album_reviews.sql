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
select * from album_reviews
where user_id = $1;

-- name: GetReviewByAlbum :many
select * from album_reviews
where album_id = $1;
