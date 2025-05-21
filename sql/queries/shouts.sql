-- name: CreateShouts :one
insert into shouts (id, created_at, updated_at, user_id, review_id, title, shout_text)
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

-- name: GetShoutsByAlbum :many
SELECT * FROM shouts
WHERE review_id = $1
ORDER BY created_at DESC;

-- name: GetShoutByUserReview :one
SELECT * FROM shouts
WHERE review_id = $1 and user_id = $2;

-- name: GetShoutByReview :many
SELECT u.img_url, u.name, u.id AS user_id, s.title, s.shout_text, s.created_at, s.id
FROM shouts s
JOIN users u ON u.id = s.user_id
WHERE s.review_id = $1
ORDER BY s.created_at DESC;

-- name: UpdateShout :exec
UPDATE shouts
SET title = $1, 
shout_text = $2, 
updated_at = $3
WHERE id = $4;

-- name: DeleteShout :exec
DELETE FROM shouts 
WHERE shouts.id = $1;



