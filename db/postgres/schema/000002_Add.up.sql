-- +goose Up
ALTER TABLE "post"
ADD title text NOT NULL;
