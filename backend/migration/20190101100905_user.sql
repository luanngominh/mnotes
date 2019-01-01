-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE user (
	"id" uuid NOT NULL,
	
	
	"created_at" timestamptz DEFAULT now(),
	"deleted_at" timestamptz,
	CONSTRANINT "user_pkey" PRIMARY KEY ("id")
)

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE user
