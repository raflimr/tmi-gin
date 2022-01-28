-- name: ListEvent :many
SELECT
    e.url_foto,
    e.judul_event,
    e.deskripsi_event,
    e.tanggal_event,
    e.id,
    (
        SELECT
            COUNT(se.id_mahasiswa)
        FROM
            status_event as se
        where
            se.id_event = e.id
        GROUP BY
            id
    )
FROM
    event e
LIMIT
    $1 OFFSET $2;