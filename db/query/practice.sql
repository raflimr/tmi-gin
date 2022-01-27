-- name: CreatePractice :one
INSERT INTO practice (
    thumbnail_practice, judul_practice, definisi, deskripsi, task_practice, ujian_practice, id_admin, id_category
) VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
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

-- name: GetPracticeById :one
SELECT * FROM practice WHERE id = $1;

-- name: InfoPracticeStatistik :one
SELECT COUNT(v.practice_id) as total_video, 
(SELECT sp.skor_practice FROM "statusPractice" as sp LEFT JOIN practice as p ON p.id=sp.id_practice WHERE p.id=$1) as point,
(SELECT COUNT(sp.id_mahasiswa) FROM "statusPractice" as sp WHERE sp.id_practice = $1)  as finish_student,
(SELECT c.name_category FROM category as c LEFT JOIN practice as p ON p.id_category=c.id WHERE p.id= $1) as category
FROM video as v
LEFT JOIN practice as p
ON v.practice_id=p.id
WHERE p.id=$1;

-- name: InfoPractice :one
SELECT * FROM practice WHERE id  = $1;

-- name: VideoInPractice :many
SELECT v.url_video, v.practice_id
FROM video as v
LEFT JOIN practice as p
ON p.id=v.practice_id
WHERE v.practice_id = $1;