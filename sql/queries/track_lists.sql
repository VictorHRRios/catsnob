-- name: AddTrackToList :one
INSERT INTO track_lists (id, added_at, user_lists_id, track_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetListByID :one
SELECT * FROM user_lists WHERE id_playlist_a = $1;

-- name: GetTrackListName :one
SELECT name_ FROM user_lists WHERE id_playlist_a = $1;

-- name: GetUserTrackLists :many
SELECT id_playlist_a, name_
FROM user_lists
WHERE user_id = $1 AND type_ = 'track';

-- name: GetTracksFromList :many
SELECT t.id, t.name, a.img_url
FROM tracks AS t
JOIN track_lists AS tl ON t.id = tl.track_id
JOIN albums AS a ON t.album_id = a.id
WHERE tl.user_lists_id = $1;

-- name: GetTracksNotInList :many
SELECT t.id, t.name, a.img_url
FROM tracks AS t
JOIN albums AS a ON t.album_id = a.id
WHERE t.id NOT IN (
    SELECT track_id
    FROM track_lists
    WHERE user_lists_id = $1
);

-- name: DeleteTrackFromList :exec
DELETE FROM track_lists
WHERE user_lists_id = $1 AND track_id = $2;
