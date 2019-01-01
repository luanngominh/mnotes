-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- expires if user_id = Null, expires = 1 day

CREATE TABLE "public"."note"(
	"id" uuid NOT NULL,
	"user_id" uuid,
	"title" text,
	"body" text,
	"tags" text[],
	"expires" timestamptz,
	"created_at" timestamptz DEFAULT now(),
	"deleted_at" timestamptz,
) WITH (oids = false);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
