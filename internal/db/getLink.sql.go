// Code generated by sqlc. DO NOT EDIT.
// source: getLink.sql

package db

import (
	"context"
)

const getLink = `-- name: GetLink :one
SELECT short_link, original_link
FROM links
WHERE short_link = $1
LIMIT 1
`

func (q *Queries) GetLink(ctx context.Context, shortLink string) (Link, error) {
	row := q.db.QueryRowContext(ctx, getLink, shortLink)
	var i Link
	err := row.Scan(&i.ShortLink, &i.OriginalLink)
	return i, err
}