-- name: CreateMahasiswa :one
INSERT INTO mahasiswa (
  username, password, email, nomor_hp, url_foto, nama_lengkap, tanggal_lahir, jenis_kelamin, university, nim, jurusan, tahun_masuk,kota_kabupaten, token
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
)
RETURNING *;

-- name: GetMahasiswa :one
SELECT * FROM mahasiswa
WHERE username = $1 OR email =$2 LIMIT 1;

-- name: UpdateMahasiswa :exec
UPDATE
    mahasiswa
SET
    username = $2,
    password = $3,
    nomor_hp = $4,
    email = $5,
    url_foto = $6,
    nama_lengkap = $7,
    tanggal_lahir = $8,
    jenis_kelamin = $9,
    university = $10,
    nim = $11,
    jurusan = $12,
    tahun_masuk = $13,
    kota_kabupaten = $14
WHERE
    id = $1;

-- name: ChangePassword :exec
UPDATE
    mahasiswa
SET
    password = $2
WHERE
    email = $1;

-- name: CheckUsername :one
SELECT * FROM mahasiswa WHERE username = $1;

-- name: CheckEmail :one
SELECT * FROM mahasiswa WHERE email = $1;

-- name: UpdateOTPInDB :exec
UPDATE
    mahasiswa
SET
    token = $2
WHERE
    email = $1;
