// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AccessToken struct {
	Token     string           `json:"token"`
	ValidTill pgtype.Timestamp `json:"valid_till"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Breeder struct {
	ID   int64       `json:"id"`
	Name pgtype.Text `json:"name"`
}

type Grow struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Germination pgtype.Timestamp `json:"germination"`
	Seedling    pgtype.Timestamp `json:"seedling"`
	Vegetative  pgtype.Timestamp `json:"vegetative"`
	Flowering   pgtype.Timestamp `json:"flowering"`
	Harvesting  pgtype.Timestamp `json:"harvesting"`
	WeightWet   pgtype.Int4      `json:"weight_wet"`
	WeightDry   pgtype.Int4      `json:"weight_dry"`
	Rating      pgtype.Int4      `json:"rating"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
}

type GrowStrain struct {
	Grow   pgtype.Int8 `json:"grow"`
	Strain pgtype.Int8 `json:"strain"`
	Count  int32       `json:"count"`
}

type Image struct {
	ID        int64            `json:"id"`
	File      []byte           `json:"file"`
	Grow      pgtype.Int8      `json:"grow"`
	Strain    pgtype.Int8      `json:"strain"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Post struct {
	ID        int64            `json:"id"`
	Content   string           `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	Title     string           `json:"title"`
}

type Strain struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	Type      pgtype.Text      `json:"type"`
	Effects   pgtype.Text      `json:"effects"`
	Price     pgtype.Float8    `json:"price"`
	Count     pgtype.Int4      `json:"count"`
	Rating    pgtype.Int4      `json:"rating"`
	Breeder   pgtype.Int8      `json:"breeder"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}