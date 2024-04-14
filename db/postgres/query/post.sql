-- name: ListPosts :many
SELECT * FROM "post"
LIMIT $1
OFFSET $2;
