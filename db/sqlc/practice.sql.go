// Code generated by sqlc. DO NOT EDIT.
// source: practice.sql

package db

import (
	"context"
	"database/sql"
)

const createPractice = `-- name: CreatePractice :one
INSERT INTO practice (
    thumbnail_practice, judul_practice, definisi, deskripsi, url_video_practice, task_practice, ujian_practice, id_admin, id_category
) VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, thumbnail_practice, judul_practice, definisi, deskripsi, url_video_practice, task_practice, ujian_practice, id_admin, id_category, created_at, finished_at
`

type CreatePracticeParams struct {
	ThumbnailPractice sql.NullString `json:"thumbnail_practice"`
	JudulPractice     sql.NullString `json:"judul_practice"`
	Definisi          sql.NullString `json:"definisi"`
	Deskripsi         sql.NullString `json:"deskripsi"`
	UrlVideoPractice  sql.NullString `json:"url_video_practice"`
	TaskPractice      sql.NullString `json:"task_practice"`
	UjianPractice     sql.NullString `json:"ujian_practice"`
	IDAdmin           NullInt32  `json:"id_admin"`
	IDCategory        int32          `json:"id_category"`
}

func (q *Queries) CreatePractice(ctx context.Context, arg CreatePracticeParams) (Practice, error) {
	row := q.db.QueryRowContext(ctx, createPractice,
		arg.ThumbnailPractice,
		arg.JudulPractice,
		arg.Definisi,
		arg.Deskripsi,
		arg.UrlVideoPractice,
		arg.TaskPractice,
		arg.UjianPractice,
		arg.IDAdmin,
		arg.IDCategory,
	)
	var i Practice
	err := row.Scan(
		&i.ID,
		&i.ThumbnailPractice,
		&i.JudulPractice,
		&i.Definisi,
		&i.Deskripsi,
		&i.UrlVideoPractice,
		&i.TaskPractice,
		&i.UjianPractice,
		&i.IDAdmin,
		&i.IDCategory,
		&i.CreatedAt,
		&i.FinishedAt,
	)
	return i, err
}

const listPractice = `-- name: ListPractice :many
SELECT id, thumbnail_practice, judul_practice, definisi, deskripsi, url_video_practice, task_practice, ujian_practice, id_admin, id_category, created_at, finished_at FROM practice
ORDER BY created_at
LIMIT $1
OFFSET $2
`

type ListPracticeParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPractice(ctx context.Context, arg ListPracticeParams) ([]Practice, error) {
	rows, err := q.db.QueryContext(ctx, listPractice, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Practice{}
	for rows.Next() {
		var i Practice
		if err := rows.Scan(
			&i.ID,
			&i.ThumbnailPractice,
			&i.JudulPractice,
			&i.Definisi,
			&i.Deskripsi,
			&i.UrlVideoPractice,
			&i.TaskPractice,
			&i.UjianPractice,
			&i.IDAdmin,
			&i.IDCategory,
			&i.CreatedAt,
			&i.FinishedAt,
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

const listPracticeByCategory = `-- name: ListPracticeByCategory :many
SELECT id, thumbnail_practice, judul_practice, definisi, deskripsi, url_video_practice, task_practice, ujian_practice, id_admin, id_category, created_at, finished_at FROM practice
WHERE id_category = $1
ORDER BY created_at
LIMIT $2
OFFSET $3
`

type ListPracticeByCategoryParams struct {
	IDCategory int32 `json:"id_category"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) ListPracticeByCategory(ctx context.Context, arg ListPracticeByCategoryParams) ([]Practice, error) {
	rows, err := q.db.QueryContext(ctx, listPracticeByCategory, arg.IDCategory, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Practice{}
	for rows.Next() {
		var i Practice
		if err := rows.Scan(
			&i.ID,
			&i.ThumbnailPractice,
			&i.JudulPractice,
			&i.Definisi,
			&i.Deskripsi,
			&i.UrlVideoPractice,
			&i.TaskPractice,
			&i.UjianPractice,
			&i.IDAdmin,
			&i.IDCategory,
			&i.CreatedAt,
			&i.FinishedAt,
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
