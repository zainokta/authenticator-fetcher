CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid NOT NULL,
    "name" character varying(255) NOT NULL,
    "phone" character varying(255) NOT NULL,
    "role" character varying(255) NOT NULL,
    "password" character varying(255) NOT NULL,
    "created_at" bigint NOT NULL,
    CONSTRAINT "users_id" PRIMARY KEY ("id"),
    CONSTRAINT "users_name" UNIQUE ("name"),
    CONSTRAINT "users_phone" UNIQUE ("phone")
) WITH (oids = false);