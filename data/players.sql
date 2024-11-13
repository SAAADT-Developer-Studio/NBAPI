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

-- name: GetAllStars :many
select * from all_stars where lower(playerFullName) like '%' || lower($1) || '%'  and season_year between $2 and $3;

-- name: GetPlayerAwards :many
SELECT * FROM player_awards where player_id = $1 and season_year BETWEEN $2 and $3;

-- name: GetPlayerAllTeams :many
SELECT * FROM all_teams JOIN player on player.id = all_teams.player_id where player_id = $1 and season_year BETWEEN $2 and $3;

-- name: GetAllTeams :many
SELECT * FROM all_teams JOIN player on player.id = all_teams.player_id and season_year BETWEEN $1 and $2;

-- name: GetAllTeamsType :many
SELECT * FROM all_teams JOIN player on player.id = all_teams.player_id where "type" = $1 and season_year BETWEEN $2 and $3;

-- name: GetPlayerAwardWinners :many
SELECT *
FROM player_awards as pa
WHERE (pa.season_year, pa.award, pa.share) IN (
    SELECT season_year, award, MAX(share) AS max_share
    FROM player_awards px
    GROUP BY px.season_year, px.award
    HAVING px.season_year BETWEEN $1 AND $2
);

