-- name: CreateTodo :one
INSERT INTO todos (title, session, expires_at)
VALUES (?, ?, ?)
RETURNING *;

-- name: ListTodos :many
SELECT id, title FROM todos
WHERE session = ? AND completed = 0
ORDER BY created_at DESC;

-- name: ListTodoIDs :many
SELECT id FROM todos
WHERE session = ?
ORDER BY created_at DESC;

-- name: CompleteTodoTransaction :exec
INSERT INTO completed (session, title)
SELECT session, title FROM todos WHERE id = ?;

-- name: TrashTodoTransaction :exec
INSERT INTO trash (session, title, created_at)
SELECT session, title, created_at FROM todos WHERE id = ?;

-- name: DeleteFromTodos :exec
DELETE FROM todos WHERE id = ?;
