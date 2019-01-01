-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- In status field, 0 is default. Because 0 is inactive, 1 is active and 2 is deactive.

CREATE TABLE "public"."user" (
	"id" uuid NOT NULL,
	"name" text NOT NULL,
	"email" text NOT NULL,
	"hashed_pass" text NOT NULL,
	"status" smallint DEFAULT 0,
	"access_code" text,
	"created_at" timestamptz DEFAULT now(),
	"deleted_at" timestamptz,
	CONSTRANINT "user_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
