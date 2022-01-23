-- name: HomeDashboard :one
SELECT
	m.username, m.nama_lengkap, m.url_foto, m.university,
	(SELECT COUNT(sp.id) FROM "statusPractice" sp left join mahasiswa m on (m.id = sp.id_mahasiswa) where sp.id_mahasiswa = $1 GROUP BY m.nama_lengkap) as finished_practice,
	(SELECT COUNT(id) as jumlah_practice FROM practice),
	(SELECT COUNT(sc.id) FROM "statusChallenge" sc left join mahasiswa m on (m.id = sc.id_mahasiswa) where sc.id_mahasiswa = $1 GROUP BY m.nama_lengkap) as finished_challange,
	(SELECT COUNT(id) as jumlah_challange FROM challenge),
	SUM((SELECT SUM(skor_practice) FROM "statusPractice" WHERE id_mahasiswa = $1) + (SELECT SUM(sc.skor_challenge) FROM "statusChallenge" sc  WHERE id_mahasiswa = $1)) as total_skor
FROM
	mahasiswa m
WHERE
	id = $1
GROUP BY m.id;