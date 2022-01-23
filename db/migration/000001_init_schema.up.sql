CREATE TABLE "mahasiswa" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "nomor_hp" varchar NOT NULL,
  "email" varchar NOT NULL,
  "url_foto" varchar NOT NULL,
  "nama_lengkap" varchar NOT NULL,
  "tanggal_lahir" varchar NOT NULL,
  "jenis_kelamin" varchar NOT NULL,
  "university" varchar NOT NULL,
  "nim" varchar NOT NULL,
  "jurusan" varchar NOT NULL,
  "tahun_masuk" varchar NOT NULL,
  "kota_kabupaten" varchar NOT NULL,
  "token" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW())
);