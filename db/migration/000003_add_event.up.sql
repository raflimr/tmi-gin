CREATE TABLE "event" (
  "id" SERIAL PRIMARY KEY,
  "url_foto" varchar,
  "judul_event" varchar,
  "deskripsi_event" varchar,
  "kriteria_event" varchar,
  "tanggal_event" varchar,
  "id_admin" integer NOT NULL,
  "created_at" timestamp,
  "finished_at" timestamp
);

CREATE TABLE "reaksi_event" (
  "id" SERIAL PRIMARY KEY,
  "reaksi_komentar" integer,
  "id_event" integer NOT NULL,
  "id_mahasiswa" integer NOT NULL
);

CREATE TABLE "review_event" (
  "id" SERIAL PRIMARY KEY,
  "komentar" varchar,
  "tanggal_komentar" varchar,
  "id_event" integer NOT NULL,
  "id_mahasiswa" integer NOT NULL
);

CREATE TABLE "status_event" (
  "id" SERIAL PRIMARY KEY,
  "skor_event" integer,
  "check_challenge" integer,
  "id_mahasiswa" integer NOT NULL,
  "id_event" integer NOT NULL,
  "finishedAt" timestamp
);

ALTER TABLE "event" ADD FOREIGN KEY ("id_admin") REFERENCES "admin" ("id");

ALTER TABLE "status_event" ADD FOREIGN KEY ("id_mahasiswa") REFERENCES "mahasiswa" ("id");

ALTER TABLE "status_event" ADD FOREIGN KEY ("id_event") REFERENCES "event" ("id");

ALTER TABLE "review_event" ADD FOREIGN KEY ("id_mahasiswa") REFERENCES "mahasiswa" ("id");

ALTER TABLE "review_event" ADD FOREIGN KEY ("id_event") REFERENCES "mahasiswa" ("id");

ALTER TABLE "reaksi_event" ADD FOREIGN KEY ("id_mahasiswa") REFERENCES "mahasiswa" ("id");

ALTER TABLE "reaksi_event" ADD FOREIGN KEY ("id_event") REFERENCES "mahasiswa" ("id");
