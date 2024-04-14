-- name: CreateAccessToken :exec
INSERT INTO "accessToken" (
  token, valid_till
) VALUES (
  $1, $2
);

-- name: GetToken :one
SELECT token, valid_till FROM "accessToken"
WHERE token = $1
LIMIT 1;
