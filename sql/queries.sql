-- name: CreateTodo :exec
INSERT INTO todos (
    session,
    title,
    expires_at
) VALUES (?, ?, ?);

-- name: ListTodos :many
SELECT * FROM todos
WHERE session = ?
ORDER BY created_at;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;
