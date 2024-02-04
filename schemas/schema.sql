CREATE TABLE "list" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp
);

CREATE TABLE "task" (
  "id" SERIAL PRIMARY KEY,
  "listid" integer,
  "name" varchar,
  "description" text,
  "completed" bool,
  "created_at" timestamp
);

CREATE TABLE "subtask" (
  "id" SERIAL,
  "name" varchar,
  "created_at" timestamp,
  "completed" bool
);

ALTER TABLE "task" ADD FOREIGN KEY ("id") REFERENCES "list" ("id");

ALTER TABLE "subtask" ADD FOREIGN KEY ("id") REFERENCES "task" ("id");

