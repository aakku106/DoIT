-- name: CreateTodo :one
INSERT INTO todos (title, session, expires_at)
VALUES (?, ?, ?)
RETURNING *;

-- name: ListTodos :many
SELECT * FROM todos
WHERE session = ?
ORDER BY created_at DESC;

-- name: CompleteTodo :exec
UPDATE todos
SET completed = 1
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;
