-- name: CreatePractice :one
INSERT INTO practice (
    thumbnail_practice, judul_practice, definisi, deskripsi, url_video_practice, task_practice, ujian_practice, id_admin, id_category
) VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: ListPracticeByCategory :many
SELECT * FROM practice
WHERE id_category = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: ListPractice :many
SELECT * FROM practice
ORDER BY created_at
LIMIT $1
OFFSET $2;
