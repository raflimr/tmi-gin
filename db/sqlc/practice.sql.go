// Code generated by sqlc. DO NOT EDIT.
// source: practice.sql

package db

import (
	"context"
	"database/sql"
)

const createPractice = `-- name: CreatePractice :one
INSERT INTO practice (
    thumbnail_practice, judul_practice, definisi, deskripsi, task_practice, ujian_practice, id_admin, id_category
) VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, thumbnail_practice, judul_practice, definisi, deskripsi, task_practice, ujian_practice, id_admin, id_category, created_at, finished_at
`

type CreatePracticeParams struct {
	ThumbnailPractice sql.NullString `json:"thumbnail_practice"`
	JudulPractice     sql.NullString `json:"judul_practice"`
	Definisi          sql.NullString `json:"definisi"`
	Deskripsi         sql.NullString `json:"deskripsi"`
	TaskPractice      sql.NullString `json:"task_practice"`
	UjianPractice     sql.NullString `json:"ujian_practice"`
	IDAdmin           sql.NullInt32  `json:"id_admin"`
	IDCategory        int32          `json:"id_category"`
}

func (q *Queries) CreatePractice(ctx context.Context, arg CreatePracticeParams) (Practice, error) {
	row := q.db.QueryRowContext(ctx, createPractice,
		arg.ThumbnailPractice,
		arg.JudulPractice,
		arg.Definisi,
		arg.Deskripsi,
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
		&i.TaskPractice,
		&i.UjianPractice,
		&i.IDAdmin,
		&i.IDCategory,
		&i.CreatedAt,
		&i.FinishedAt,
	)
	return i, err
}

const getPracticeById = `-- name: GetPracticeById :one
SELECT id, thumbnail_practice, judul_practice, definisi, deskripsi, task_practice, ujian_practice, id_admin, id_category, created_at, finished_at FROM practice WHERE id = $1
`

func (q *Queries) GetPracticeById(ctx context.Context, id int32) (Practice, error) {
	row := q.db.QueryRowContext(ctx, getPracticeById, id)
	var i Practice
	err := row.Scan(
		&i.ID,
		&i.ThumbnailPractice,
		&i.JudulPractice,
		&i.Definisi,
		&i.Deskripsi,
		&i.TaskPractice,
		&i.UjianPractice,
		&i.IDAdmin,
		&i.IDCategory,
		&i.CreatedAt,
		&i.FinishedAt,
	)
	return i, err
}

const infoPractice = `-- name: InfoPractice :one
SELECT id, thumbnail_practice, judul_practice, definisi, deskripsi, task_practice, ujian_practice, id_admin, id_category, created_at, finished_at FROM practice WHERE id  = $1
`

func (q *Queries) InfoPractice(ctx context.Context, id int32) (Practice, error) {
	row := q.db.QueryRowContext(ctx, infoPractice, id)
	var i Practice
	err := row.Scan(
		&i.ID,
		&i.ThumbnailPractice,
		&i.JudulPractice,
		&i.Definisi,
		&i.Deskripsi,
		&i.TaskPractice,
		&i.UjianPractice,
		&i.IDAdmin,
		&i.IDCategory,
		&i.CreatedAt,
		&i.FinishedAt,
	)
	return i, err
}

const infoPracticeStatistik = `-- name: InfoPracticeStatistik :one
SELECT COUNT(v.practice_id) as total_video, 
(SELECT sp.skor_practice FROM "statusPractice" as sp LEFT JOIN practice as p ON p.id=sp.id_practice WHERE p.id=$1) as point,
(SELECT COUNT(sp.id_mahasiswa) FROM "statusPractice" as sp WHERE sp.id_practice = $1)  as finish_student,
(SELECT c.name_category FROM category as c LEFT JOIN practice as p ON p.id_category=c.id WHERE p.id= $1) as category
FROM video as v
LEFT JOIN practice as p
ON v.practice_id=p.id
WHERE p.id=$1
`

type InfoPracticeStatistikRow struct {
	TotalVideo    int64          `json:"total_video"`
	Point         sql.NullInt32  `json:"point"`
	FinishStudent int64          `json:"finish_student"`
	Category      sql.NullString `json:"category"`
}

func (q *Queries) InfoPracticeStatistik(ctx context.Context, id int32) (InfoPracticeStatistikRow, error) {
	row := q.db.QueryRowContext(ctx, infoPracticeStatistik, id)
	var i InfoPracticeStatistikRow
	err := row.Scan(
		&i.TotalVideo,
		&i.Point,
		&i.FinishStudent,
		&i.Category,
	)
	return i, err
}

const listPractice = `-- name: ListPractice :many
SELECT id, thumbnail_practice, judul_practice, definisi, deskripsi, task_practice, ujian_practice, id_admin, id_category, created_at, finished_at FROM practice
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
SELECT id, thumbnail_practice, judul_practice, definisi, deskripsi, task_practice, ujian_practice, id_admin, id_category, created_at, finished_at FROM practice
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

const videoInPractice = `-- name: VideoInPractice :many
SELECT v.url_video, v.practice_id
FROM video as v
LEFT JOIN practice as p
ON p.id=v.practice_id
WHERE v.practice_id = $1
`

type VideoInPracticeRow struct {
	UrlVideo   sql.NullString `json:"url_video"`
	PracticeID int32          `json:"practice_id"`
}

func (q *Queries) VideoInPractice(ctx context.Context, practiceID int32) ([]VideoInPracticeRow, error) {
	rows, err := q.db.QueryContext(ctx, videoInPractice, practiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []VideoInPracticeRow{}
	for rows.Next() {
		var i VideoInPracticeRow
		if err := rows.Scan(&i.UrlVideo, &i.PracticeID); err != nil {
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
