-- name: CreateAuthor :one
INSERT INTO authors (
    owner,
    balance,
    currency
) VALUES (
    $1, $2, $3
)