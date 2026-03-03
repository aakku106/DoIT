-- name: CreateCat :one
insert into cat(
name,desc
) values (
  ?,?
  )
RETURNING *;

-- GetCat :many
select * from cat
order by id;

-- DeleteCat :exec
delete from cat where id=?
