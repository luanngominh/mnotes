-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."users" (
	"id" uuid NOT NULL,
	"name" text NOT NULL,
	"email" text NOT NULL,
	"hashed_pass" text NOT NULL,
	"verify" text,
	"status" smallint DEFAULT 0,
	"access_code" text,
	"created_at" timestamptz DEFAULT now(),
	"deleted_at" timestamptz,
	CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;