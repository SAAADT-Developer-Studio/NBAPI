-- name: GetPlayers :many
select * from player;

-- name: GetPlayer :one
select * from player where id = $1;