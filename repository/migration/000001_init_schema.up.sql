CREATE TABLE "accounts" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "moneyGoal" bigint,
  "startSum" bigint
);

CREATE TABLE "entries" (
  "id" SERIAL PRIMARY KEY,
  "account_id" int,
  "salaryPerMonth" bigint,
  "outcomePerMonth" bigint,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transactions" (
  "id" SERIAL PRIMARY KEY,
  "account_id" int,
  "value" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transactions" ("account_id");
