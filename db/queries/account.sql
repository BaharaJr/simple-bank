-- name: CreateAccount :one
INSERT INTO ACCOUNTS (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccountById :one
SELECT * FROM ACCOUNTS WHERE ID=$1 LIMIT 1;

-- name: GetAccountByOwner :one
SELECT * FROM ACCOUNTS WHERE OWNER=$1 LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM ACCOUNTS 
ORDER BY CREATED ASC
 LIMIT $1
 OFFSET $2;

 -- name: UpdateAccount :one
UPDATE ACCOUNTS 
SET BALANCE = $2
WHERE ID = $1
RETURNING *
;

-- name: GenericSearch :many
SELECT * FROM ACCOUNTS WHERE $1 = $2
ORDER BY CREATED ASC
 LIMIT $3
 OFFSET $4;

-- name: DeleteAccount :exec
DELETE FROM ACCOUNTS WHERE id = $1;