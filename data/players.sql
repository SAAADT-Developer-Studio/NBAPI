-- name: GetPlayers :many
select
    *
from
    player
where
    lower(fullName) like '%' || lower(sqlc.arg(search)) || '%' and
    id >= sqlc.arg(cursor)
    order by id
    limit sqlc.arg(page_size) + 1;

-- name: GetPlayerById :one
select * from player where id = $1;

-- name: GetPlayerTotals :many
select totals.* from player
inner join player_totals on player.id = player_totals.player_id
inner join totals on totals.id = player_totals.total_id
where player.id = $1
and player_totals.season_year between $2 and $3;

-- name: GetPlayerPer100 :many
select per_100_possesions.* from player
  inner join player_per_100_possesions on player.id = player_per_100_possesions.player_id
  inner join per_100_possesions on player_per_100_possesions.per_100_id = per_100_possesions.id
  where player.id = $1
  and player_per_100_possesions.season_year between $2 and $3;

-- name: GetPlayerPerGame :many
select per_game.* from player
  inner join player_per_game on player.id = player_per_game.player_id
  inner join per_game on player_per_game.per_game_id = per_game.id
  where player.id = $1
  and player_per_game.season_year between $2 and $3;

-- name: GetPlayerAdvanced :many
select advanced.* from player
  inner join player_advanced on player.id = player_advanced.player_id
  inner join advanced on player_advanced.advanced_id = advanced.id
  where player.id = $1
  and player_advanced.season_year between $2 and $3;

-- name: GetPlayerPer36 :many
select player_per_36.* from player
  inner join player_per_36 on player_per_36.player_id = player.id
  where player.id = $1
  and player_per_36.season_year between $2 and $3;

-- name: GetPlayerShooting :many
select player_shooting.* from player
  inner join player_shooting on player_shooting.player_id = player.id
  where player.id = $1
  and player_shooting.season_year between $2 and $3;
