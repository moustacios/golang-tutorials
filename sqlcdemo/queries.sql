-- ~/Workspace/sqlcdemo/queries.sql

-- name: GetComment :one
SELECT * FROM comments 
WHERE id = ? LIMIT 1;

-- name: ListComments :many
SELECT * FROM comments 
ORDER BY bot_probability DESC;

-- name: SaveComment :execresult
INSERT INTO comments (
  email, comment_text
) VALUES (
  ?, ?
);

-- name: DeleteComment :execresult
DELETE FROM comments 
WHERE id = ?;

-- name: FlagBotComment :execresult
UPDATE comments
SET bot_probability = ?
WHERE id = ?;

-- name: PurgeBotComments :execresult 
DELETE FROM comments 
WHERE bot_probability > ?;
