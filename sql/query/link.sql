-- name: GetLink :one
SELECT *
FROM links
WHERE id = $1
LIMIT 1;

-- name: CreateLink :one
INSERT INTO links(id,url)
VALUES($1, $2)
RETURNING *;