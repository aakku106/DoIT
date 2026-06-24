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
ORDER BY c.completed_at DESC;

-- name: ListCompletedIDs :many
SELECT c.id FROM completed AS c
WHERE c.session = ?
ORDER BY c.created_at DESC;

-- name: ListTrash :many
SELECT t.id, t.title FROM trash AS t
WHERE t.session = ?
ORDER BY t.created_at DESC;

-- name: ListTrashIDs :many
SELECT t.id FROM trash AS t
WHERE t.session = ?
ORDER BY t.created_at DESC;

-- name: ListIgnored :many
SELECT i.id, i.title FROM ignored AS i
WHERE i.session = ?
ORDER BY i.expired_at DESC;

-- name: ListIgnoredIDs :many
SELECT i.id FROM ignored AS i
WHERE i.session = ?
ORDER BY i.expired_at DESC;

-- name: CompleteTodoTransaction :exec
INSERT INTO completed (session, title)
SELECT t.session, t.title FROM todos AS t
WHERE t.id = ?;

-- name: TrashTodoTransaction :exec
INSERT INTO trash (session, title, created_at)
SELECT t.session, t.title, t.created_at FROM todos AS t
WHERE t.id = ?;

-- name: MoveCompletedTo :exec
INSERT INTO todos (session,title)
SELECT c.session, c.title FROM completed AS c
WHERE c.id = ?;

-- name: MoveCompletedToIgnored :exec
INSERT INTO ignored (session,title)
SELECT c.session, c.title FROM completed AS c
WHERE c.id = ?;

-- name: MoveCompletedToTrash :exec
INSERT INTO trash (session,title,created_at)
SELECT c.session, c.title, c.completed_at FROM completed AS c
WHERE c.id = ?;

-- name: DeleteFromTodos :exec
DELETE FROM todos 
WHERE todos.id = ?;

-- name: DeleteFromCompleted :exec
DELETE FROM completed
WHERE completed.id = ?;

-- name: DeleteFromIgnored :exec
DELETE FROM ignored
WHERE ignored.id = ?;
-- name: DeleteFromTrash :exec
DELETE FROM trash
WHERE trash.id = ?;

-- name: ClearTrash :exec
DELETE FROM trash;

-- name: ClearCompleted :exec
DELETE FROM completed;

-- name: ClearIgnored :exec
DELETE FROM ignored;
