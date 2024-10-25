CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "description" text NOT NULL
);

CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "category_id" bigserial NOT NULL,
  "amount" real NOT NULL,
  "description" text NOT NULL,
  "date" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "categories" ("name");

CREATE INDEX ON "expenses" ("user_id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

INSERT INTO categories (name, description) VALUES
    ('Food', 'buying food, water'),
    ('Income', 'income from working'),
    ('Meds', 'medical drugs, pharmacy'),
    ('Hygiene', 'hygiene equipment, toothpaste, brush etc.');
