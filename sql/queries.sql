-- name: CreateTodo :one
INSERT INTO todos (title, session, expires_at)
VALUES (?, ?, ?)
RETURNING *;

-- name: ListTodos :many
SELECT t.id, t.title FROM todos AS t
WHERE t.session = ?
ORDER BY t.created_at DESC;

-- name: ListTodoIDs :many
SELECT t.id FROM todos AS t
WHERE t.session = ?
ORDER BY t.created_at DESC;

-- name: ListCompleted :many
SELECT c.id, c.title FROM completed AS c
WHERE c.session = ?
ORDER BY c.created_at DESC;

-- name: ListCompletedIDs :many
SELECT c.id FROM completed AS c
WHERE c.session = ?
ORDER BY c.created_at DESC;

-- name: ListTrash :many
SELECT t.id, t.session FROM trash AS t
WHERE t.session = ?
ORDER BY t.created_at DESC;

-- name: ListTrashIDs :many
SELECT t.id FROM trash AS t
WHERE t.session = ?
ORDER BY t.created_at DESC;

-- name: CompleteTodoTransaction :exec
INSERT INTO completed (session, title)
SELECT t.session, t.title FROM todos AS t
WHERE t.id = ?;

-- name: TrashTodoTransaction :exec
INSERT INTO trash (session, title, created_at)
SELECT t.session, t.title, t.created_at FROM todos AS t
WHERE t.id = ?;

-- name: DeleteFromTodos :exec
DELETE FROM todos 
WHERE todos.id = ?;
