CREATE TABLE "practice" (
  "id" SERIAL PRIMARY KEY,
  "thumbnail_practice" varchar,
  "judul_practice" varchar,
  "definisi" varchar,
  "deskripsi" varchar,
  "task_practice" varchar,
  "ujian_practice" varchar,
  "id_admin" integer,
  "id_category" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "finished_at" timestamp
);

CREATE TABLE "challenge" (
  "id" SERIAL PRIMARY KEY,
  "tema" varchar,
  "judul_challenge" varchar,
  "latar" varchar,
  "url_video_challenge" varchar,
  "task_challenge" varchar,
  "ujian_challenge" varchar,
  "skor_challenge" varchar,
  "id_admin" integer,
  "id_category" integer,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "finishedAt" timestamp
);


CREATE TABLE "event" (
  "id" SERIAL PRIMARY KEY,
  "url_foto" varchar,
  "judul_event" varchar,
  "deskripsi_event" varchar,
  "kriteria_event" varchar,
  "tanggal_event" varchar,
  "id_mahasiswa" integer NOT NULL,
  "id_admin" integer,
  "created_at" timestamp,
  "finished_at" timestamp
);

CREATE TABLE "statusChallenge" (
  "id" SERIAL PRIMARY KEY,
  "skor_challenge" integer,
  "id_mahasiswa" integer NOT NULL,
  "id_challenge" integer,
  "check_challenge" integer,
  "finishedAt" timestamp
);

CREATE TABLE "statusPractice" (
  "id" SERIAL PRIMARY KEY,
  "skor_practice" integer,
  "id_mahasiswa" integer NOT NULL,
  "id_practice" integer,
  "check_practice" integer,
  "finishedAt" timestamp
);

CREATE TABLE "category" (
  "id" SERIAL PRIMARY KEY,
  "name_category" varchar,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "finished_at" timestamp
);

CREATE TABLE "admin" (
  "id" SERIAL PRIMARY KEY,
  "job" varchar,
  "photo" varchar,
  "bio" varchar,
  "status" varchar,
  "created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE "video" (
  "id" SERIAL PRIMARY KEY,
  "url_video" varchar,
  "practice_id" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW()),
  "update_at" timestamp
);

ALTER TABLE "practice" ADD FOREIGN KEY ("id_admin") REFERENCES "admin" ("id");

ALTER TABLE "statusPractice" ADD FOREIGN KEY ("id_mahasiswa") REFERENCES "mahasiswa" ("id");

ALTER TABLE "statusPractice" ADD FOREIGN KEY ("id_practice") REFERENCES "practice" ("id");

ALTER TABLE "statusChallenge" ADD FOREIGN KEY ("id_mahasiswa") REFERENCES "mahasiswa" ("id");

ALTER TABLE "statusChallenge" ADD FOREIGN KEY ("id_challenge") REFERENCES "challenge" ("id");

ALTER TABLE "practice" ADD FOREIGN KEY ("id_category") REFERENCES "category" ("id");

ALTER TABLE "video" ADD FOREIGN KEY ("practice_id") REFERENCES "practice" ("id");
