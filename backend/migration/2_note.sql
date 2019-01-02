-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."notes"(
	"id" uuid NOT NULL,
	"user_id" uuid,
	"title" text,
	"body" text,
	"tags" text[],
	"expire" timestamptz,
	"created_at" timestamptz DEFAULT now(),
	"deleted_at" timestamptz,
	CONSTRAINT "notes_pkey" PRIMARY KEY("id"),
	CONSTRAINT "notes_users" FOREIGN kEY ("id") REFERENCES users("id")
) WITH (oids = false);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
