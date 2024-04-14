CREATE TABLE "accessToken" (
  "token" text PRIMARY KEY,
  "valid_till" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "post" (
  "id" BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "content" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "image" (
  "id" BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "file" bytea NOT NULL,
  "grow" bigint,
  "strain" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "grow" (
  "id" BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" text NOT NULL,
  "germination" timestamp,
  "seedling" timestamp,
  "vegetative" timestamp,
  "flowering" timestamp,
  "harvesting" timestamp,
  "weight_wet" int,
  "weight_dry" int,
  "rating" int,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "strain" (
  "id" BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" text NOT NULL,
  "type" text,
  "effects" text,
  "price" float,
  "count" int,
  "rating" int,
  "breeder" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "grow_strain" (
  "grow" bigint,
  "strain" bigint,
  "count" int NOT NULL
);

CREATE TABLE "breeder" (
  "id" BIGINT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" text
);

ALTER TABLE "image" ADD FOREIGN KEY ("grow") REFERENCES "grow" ("id");

ALTER TABLE "image" ADD FOREIGN KEY ("strain") REFERENCES "strain" ("id");

ALTER TABLE "strain" ADD FOREIGN KEY ("breeder") REFERENCES "breeder" ("id");

ALTER TABLE "grow_strain" ADD FOREIGN KEY ("grow") REFERENCES "grow" ("id");

ALTER TABLE "grow_strain" ADD FOREIGN KEY ("strain") REFERENCES "strain" ("id");
