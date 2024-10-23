-- name: GetTeams :many
select
  *
from
  team
where
  lower(fullname) like '%' || lower($1) || '%';

-- name: GetTeam :one
select
  *
from
  team
where
  abbr = $1;

-- name: GetTeamTotals :many
select totals.*, team_totals.season_year from team
  inner join team_totals on team.abbr = team_totals.team_abbr
  inner join totals on team_totals.total_id = totals.id
  where team.abbr = $1 and
    team_totals.season_year between $2 and $3
  order by team_totals.season_year desc;


-- name: GetTeamPer100Possesions :many
select per_100_possesions.* from team
  inner join team_per_100_possesions on team.abbr = team_per_100_possesions.team_abbr
  inner join per_100_possesions on team_per_100_possesions.per_100_id = per_100_possesions.id
  where team.abbr = $1 and
    team_per_100_possesions.season_year between $2 and $3
  order by team_per_100_possesions.season_year desc;

-- name: GetTeamPerGame :many
select per_game.*, team_per_game.season_year from team
  inner join team_per_game on team.abbr = team_per_game.team_abbr
  inner join per_game on team_per_game.per_game_id = per_game.id
  where team.abbr = $1 and
    team_per_game.season_year between $2 and $3
  order by team_per_game.season_year desc;