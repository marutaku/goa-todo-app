-- Create "auth_records" table
CREATE TABLE "auth_records" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "password" text NOT NULL,
  "created_at" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "auth_records_name_key" to table: "auth_records"
CREATE UNIQUE INDEX "auth_records_name_key" ON "auth_records" ("name");
-- Create "task_records" table
CREATE TABLE "task_records" (
  "id" bigserial NOT NULL,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "done" boolean NOT NULL,
  "done_at" text NULL,
  "done_by" bigint NULL,
  "created_at" text NOT NULL,
  "created_by" bigint NOT NULL,
  PRIMARY KEY ("id")
);
