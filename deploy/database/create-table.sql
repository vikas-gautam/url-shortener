CREATE TABLE "users" (
  "id" serial UNIQUE PRIMARY KEY,
  "first_name" varchar(255),
  "last_name" varchar(255),
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "is_active" varchar(10) DEFAULT 'true',
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "urls" (
  "id" serial UNIQUE PRIMARY KEY,
  "user_id" integer,
  "actual_url" varchar(255) NOT NULL,
  "short_url" varchar(255) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "counters" (
  "id" serial UNIQUE PRIMARY KEY,
  "user_id" integer,
  "short_url" varchar(255) NOT NULL,
  "hit_counter" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "counters" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "urls" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");