-- name: CreateEntry :one
INSERT INTO ENTRIES (
  amount,
  account_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntryId :one
SELECT * FROM ENTRIES WHERE ID=$1 LIMIT 1;

-- name: GetEntries :many
SELECT * FROM ENTRIES 
ORDER BY CREATED ASC
 LIMIT $1
 OFFSET $2;

 -- name: UpdateEntry :one
UPDATE ENTRIES 
SET AMOUNT = $2
WHERE ID = $1
RETURNING *
;

-- name: DeleteEntry :exec
DELETE FROM ENTRIES WHERE id = $1;