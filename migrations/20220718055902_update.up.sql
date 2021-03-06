-- create "directories" table
CREATE TABLE "directories" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create "objects" table
CREATE TABLE "objects" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "url" character varying NOT NULL, "user_objects" bigint NULL, PRIMARY KEY ("id"));
-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT 'unknown', "password" character varying NOT NULL, "phone" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "user_phone_deleted_at" to table: "users"
CREATE UNIQUE INDEX "user_phone_deleted_at" ON "users" ("phone", "deleted_at");
