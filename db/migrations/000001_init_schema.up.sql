CREATE TABLE "users" (
  "id"          bigserial  PRIMARY KEY,
  "username"    varchar    NOT NULL UNIQUE,
  "password"    varchar    NOT NULL,
  "role"        varchar    NOT NULL,
  "created_at"  timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "listings" (
  "id"         bigserial  PRIMARY KEY,
  "title"      varchar    NOT NULL,
  "body"       text       NOT NULL,
  "user_id"    bigint     NOT NULL,
  "is_settled" boolean    NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT "fk_listings_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET DEFAULT
);

CREATE TABLE "posts" (
  "id"         bigserial  PRIMARY KEY,
  "title"      varchar    NOT NULL,
  "body"       text       NOT NULL,
  "user_id"    bigint     NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT "fk_post_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET DEFAULT
);

CREATE TABLE "comments" (
  "id"             bigserial  PRIMARY KEY,
  "body"           text       NOT NULL,
  "user_id"        bigint     NOT NULL,
  "post_id"        bigint     NOT NULL,
  "parent_comment" bigint,
  "created_at"     timestamptz NOT NULL DEFAULT (now()),
  CONSTRAINT "fk_comment_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET DEFAULT,
  CONSTRAINT "fk_comment_post" FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE SET DEFAULT,
  CONSTRAINT "fk_comment_parent" FOREIGN KEY ("parent_comment") REFERENCES "comments" ("id") ON DELETE SET DEFAULT

);

COMMENT ON COLUMN "users"."password" IS 'hashed password';

COMMENT ON COLUMN "listings"."body" IS 'Content of the listing';

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

