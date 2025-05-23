// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: track_lists.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addTrackToList = `-- name: AddTrackToList :one
INSERT INTO track_lists (id, added_at, user_lists_id, track_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    $1,
    $2
)
RETURNING id, added_at, user_lists_id, track_id
`

type AddTrackToListParams struct {
	UserListsID uuid.UUID
	TrackID     uuid.UUID
}

func (q *Queries) AddTrackToList(ctx context.Context, arg AddTrackToListParams) (TrackList, error) {
	row := q.db.QueryRowContext(ctx, addTrackToList, arg.UserListsID, arg.TrackID)
	var i TrackList
	err := row.Scan(
		&i.ID,
		&i.AddedAt,
		&i.UserListsID,
		&i.TrackID,
	)
	return i, err
}

const deleteTrackFromList = `-- name: DeleteTrackFromList :exec
DELETE FROM track_lists
WHERE user_lists_id = $1 AND track_id = $2
`

type DeleteTrackFromListParams struct {
	UserListsID uuid.UUID
	TrackID     uuid.UUID
}

func (q *Queries) DeleteTrackFromList(ctx context.Context, arg DeleteTrackFromListParams) error {
	_, err := q.db.ExecContext(ctx, deleteTrackFromList, arg.UserListsID, arg.TrackID)
	return err
}

const getListByID = `-- name: GetListByID :one
SELECT id_playlist_a, created_at, updated_at, name_, description_, type_, user_id FROM user_lists WHERE id_playlist_a = $1
`

func (q *Queries) GetListByID(ctx context.Context, idPlaylistA uuid.UUID) (UserList, error) {
	row := q.db.QueryRowContext(ctx, getListByID, idPlaylistA)
	var i UserList
	err := row.Scan(
		&i.IDPlaylistA,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.Type,
		&i.UserID,
	)
	return i, err
}

const getTrackListName = `-- name: GetTrackListName :one
SELECT name_ FROM user_lists WHERE id_playlist_a = $1
`

func (q *Queries) GetTrackListName(ctx context.Context, idPlaylistA uuid.UUID) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getTrackListName, idPlaylistA)
	var name_ sql.NullString
	err := row.Scan(&name_)
	return name_, err
}

const getTracksFromList = `-- name: GetTracksFromList :many
SELECT t.id, t.name, a.img_url
FROM tracks AS t
JOIN track_lists AS tl ON t.id = tl.track_id
JOIN albums AS a ON t.album_id = a.id
WHERE tl.user_lists_id = $1
`

type GetTracksFromListRow struct {
	ID     uuid.UUID
	Name   string
	ImgUrl string
}

func (q *Queries) GetTracksFromList(ctx context.Context, userListsID uuid.UUID) ([]GetTracksFromListRow, error) {
	rows, err := q.db.QueryContext(ctx, getTracksFromList, userListsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTracksFromListRow
	for rows.Next() {
		var i GetTracksFromListRow
		if err := rows.Scan(&i.ID, &i.Name, &i.ImgUrl); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTracksNotInList = `-- name: GetTracksNotInList :many
SELECT t.id, t.name, a.img_url
FROM tracks AS t
JOIN albums AS a ON t.album_id = a.id
WHERE t.id NOT IN (
    SELECT track_id
    FROM track_lists
    WHERE user_lists_id = $1
)
`

type GetTracksNotInListRow struct {
	ID     uuid.UUID
	Name   string
	ImgUrl string
}

func (q *Queries) GetTracksNotInList(ctx context.Context, userListsID uuid.UUID) ([]GetTracksNotInListRow, error) {
	rows, err := q.db.QueryContext(ctx, getTracksNotInList, userListsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTracksNotInListRow
	for rows.Next() {
		var i GetTracksNotInListRow
		if err := rows.Scan(&i.ID, &i.Name, &i.ImgUrl); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTrackLists = `-- name: GetUserTrackLists :many
SELECT id_playlist_a, name_
FROM user_lists
WHERE user_id = $1 AND type_ = 'track'
`

type GetUserTrackListsRow struct {
	IDPlaylistA uuid.UUID
	Name        sql.NullString
}

func (q *Queries) GetUserTrackLists(ctx context.Context, userID uuid.UUID) ([]GetUserTrackListsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserTrackLists, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserTrackListsRow
	for rows.Next() {
		var i GetUserTrackListsRow
		if err := rows.Scan(&i.IDPlaylistA, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
