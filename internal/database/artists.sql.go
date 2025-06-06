// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: artists.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createArtist = `-- name: CreateArtist :one
insert into artists (id, created_at, updated_at, formed_at, name, biography, genre, img_url)
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
returning id, created_at, updated_at, formed_at, name, biography, genre, img_url
`

type CreateArtistParams struct {
	FormedAt  string
	Name      string
	Biography sql.NullString
	Genre     string
	ImgUrl    string
}

func (q *Queries) CreateArtist(ctx context.Context, arg CreateArtistParams) (Artist, error) {
	row := q.db.QueryRowContext(ctx, createArtist,
		arg.FormedAt,
		arg.Name,
		arg.Biography,
		arg.Genre,
		arg.ImgUrl,
	)
	var i Artist
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FormedAt,
		&i.Name,
		&i.Biography,
		&i.Genre,
		&i.ImgUrl,
	)
	return i, err
}

const deleteArtist = `-- name: DeleteArtist :exec
delete from artists
where artists.id = $1
`

func (q *Queries) DeleteArtist(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteArtist, id)
	return err
}

const getArtist = `-- name: GetArtist :one
select id, created_at, updated_at, formed_at, name, biography, genre, img_url from artists
where id = $1
`

func (q *Queries) GetArtist(ctx context.Context, id uuid.UUID) (Artist, error) {
	row := q.db.QueryRowContext(ctx, getArtist, id)
	var i Artist
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FormedAt,
		&i.Name,
		&i.Biography,
		&i.Genre,
		&i.ImgUrl,
	)
	return i, err
}

const getArtists = `-- name: GetArtists :many
select id, created_at, updated_at, formed_at, name, biography, genre, img_url from artists order by name
`

func (q *Queries) GetArtists(ctx context.Context) ([]Artist, error) {
	rows, err := q.db.QueryContext(ctx, getArtists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Artist
	for rows.Next() {
		var i Artist
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FormedAt,
			&i.Name,
			&i.Biography,
			&i.Genre,
			&i.ImgUrl,
		); err != nil {
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

const getTop12Artists = `-- name: GetTop12Artists :many
select id, created_at, updated_at, formed_at, name, biography, genre, img_url from artists limit 12
`

func (q *Queries) GetTop12Artists(ctx context.Context) ([]Artist, error) {
	rows, err := q.db.QueryContext(ctx, getTop12Artists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Artist
	for rows.Next() {
		var i Artist
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FormedAt,
			&i.Name,
			&i.Biography,
			&i.Genre,
			&i.ImgUrl,
		); err != nil {
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
