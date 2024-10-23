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


-- name: CreatePlayer :exec
insert into player (id, fullName) values ($1, $2);

-- name: CreateTeam :exec
insert into team (abbr, fullName) values ($1, $2);

-- name: CreateTotals :exec
insert into
    totals (
        id,
        gp,
        gs,
        mp,
        fg,
        fga,
        p3,
        pa3,
        p2,
        pa2,
        ft,
        fta,
        orb,
        drb,
        trb,
        stl,
        blk,
        ast,
        tov,
        pf,
        pts
    )
values (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21
    );

-- name: CreateTeamTotals :exec
insert into
    team_totals (
        team_abbr,
        total_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreatePlayerTotals :exec
insert into
    player_totals (
        player_id,
        total_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreatePer100Possesions :exec
INSERT INTO
    per_100_possesions (
        id,
        fg,
        fga,
        p3,
        pa3,
        p2,
        pa2,
        ft,
        fta,
        orb,
        drb,
        trb,
        stl,
        blk,
        ast,
        tov,
        pf,
        pts,
        o_rtg,
        d_rtg
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20
    );

-- name: CreateTeamPer100Possesions :exec
insert into
    team_per_100_possesions (
        team_abbr,
        per_100_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreatePlayerPer100Possesions :exec
insert into
    player_per_100_possesions (
        player_id,
        per_100_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreatePerGame :exec
insert into
    per_game (
        id,
        mp,
        fg,
        fga,
        fg_percent,
        p3,
        pa3,
        p_percent3,
        p2,
        pa2,
        p_percent2,
        efg_percent,
        ft,
        fta,
        ft_percent,
        orb,
        drb,
        trb,
        ast,
        stl,
        blk,
        tov,
        pf,
        pts
    )
values (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21,
        $22,
        $23,
        $24
    );

-- name: CreateTeamPerGame :exec
insert into
    team_per_game (
        team_abbr,
        per_game_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreatePlayerPerGame :exec
insert into
    player_per_game (
        player_id,
        per_game_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreateOpponentsTotals :exec
insert into
    opponents_totals (
        team_abbr,
        total_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreateOpponentsPerGame :exec
insert into
    opponents_per_game (
        team_abbr,
        per_game_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreateOpponentsPer100Possessions :exec
insert into
    opponents_per_100_possesions (
        team_abbr,
        per_100_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreateAdvanced :exec
insert into
    advanced (
        id,
        "per",
        ts_percent,
        p_ar3,
        f_tr,
        orb_percent,
        drb_percent,
        trb_percent,
        ast_percent,
        stl_percent,
        blk_percent,
        tov_percent,
        usg_percent,
        ows,
        dws,
        ws,
        ws48,
        obpm,
        dbpm,
        bpm,
        vorp
    )
values (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21
    );

-- name: CreatePlayerAdvanced :exec
insert into
    player_advanced (
        player_id,
        advanced_id,
        season_year
    )
values ($1, $2, $3);

-- name: CreateAllTeamsVoting :exec
insert into
    all_teams_voting (
        player_id,
        season_year,
        "type",
        pts_won,
        pts_max,
        "share",
        first_team,
        second_team,
        third_team
    )
values (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9
    );

-- name: CreateAllTeams :exec
insert into
    all_teams (
        player_id,
        season_year,
        "type",
        team_number
    )
values ($1, $2, $3, $4);


-- name: GetPlayerAwards :many
SELECT * FROM player_awards where player_id = $1 and season_year BETWEEN $2 and $3;

-- name: GetPlayerAllTeams :many
SELECT * FROM all_teams JOIN player on player.id = all_teams.player_id where player_id = $1 and season_year BETWEEN $2 and $3;

-- name: GetAllTeams :many
SELECT * FROM all_teams JOIN player on player.id = all_teams.player_id and season_year BETWEEN $1 and $2;

-- name: GetAllTeamsType :many
SELECT * FROM all_teams JOIN player on player.id = all_teams.player_id where "type" = $1 and season_year BETWEEN $2 and $3;

-- name: GetAwardWinners :many
SELECT * FROM player_awards where winner = true AND season_year BETWEEN $1 and $2 ORDER BY season_year DESC;

-- name: GetSpecificAwardWinners :many
SELECT * FROM player_awards where winner = true AND award = $3 AND season_year BETWEEN $1 and $2;

-- name: CreatePer36 :exec
INSERT INTO
    player_per_36 (
        player_id,
        season_year,
        fg,
        fga,
        p3,
        pa3,
        p2,
        pa2,
        ft,
        fta,
        orb,
        drb,
        trb,
        stl,
        blk,
        ast,
        tov,
        pf,
        pts
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19
    );

-- name: CreatePlayerShooting :exec
insert into
    player_shooting (
        season_year,
        player_id,
        avg_dist_fga,
        percent_fga_from_2p_range,
        percent_fga_from_0_3_range,
        percent_fga_from_3_10_range,
        percent_fga_from_10_16_range,
        percent_fga_from_16_3p_range,
        percent_fga_from_3p_range,
        fg_percent_from_2p_range,
        fg_percent_from_0_3_range,
        fg_percent_from_3_10_range,
        fg_percent_from_10_16_range,
        fg_percent_from_16_3p_range,
        fg_percent_from_3p_range,
        percent_assisted_2p_fg,
        percent_assisted_3p_fg,
        percent_dunks_of_fga,
        num_of_dunks,
        percent_corner_3s_of_3pa,
        corner_3_point_percent,
        num_heaves_attempted,
        num_heaves_made
    )
values (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21,
        $22,
        $23
    );

-- name: CreateAllStars :exec
insert into
    all_stars (
        playerFullName,
        season_year,
        teamName,
        replaced
    )
values ($1, $2, $3, $4);

-- name: CreatePlayerTeam :exec
insert into player_team (
  team_abbr,
  player_id,
  season_year,
  age,
  experience,
  position


) values ( $1, $2, $3, $4, $5, $6);

-- name: GetAllStars :many
select * from all_stars where lower(playerFullName) like '%' || lower($1) || '%'  and season_year between $2 and $3;

-- name: CreatePlayerAwards :exec
insert into
    player_awards (
        player_id,
        season_year,
        award,
        pts_won,
        pts_max,
        "share",
        winner
    )
values ($1, $2, $3, $4, $5, $6, $7);

-- name: CreateTeamSeason :exec
insert into
    team_season (
        season_year,
        team_abbr,
        playoffs,
        avarage_age,
        w,
        l,
        pw,
        pl,
        mov,
        sos,
        srs,
        o_rtg,
        d_rtg,
        n_rtg,
        pace,
        f_tr,
        p_ar3,
        ts_percent,
        e_fg_percent,
        tov_percent,
        orb_percent,
        ft_fga,
        opp_e_fg_percent,
        opp_tov_percent,
        opp_drb_percent,
        opp_ft_fga,
        arena,
        attend,
        attend_g
    )
values (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21,
        $22,
        $23,
        $24,
        $25,
        $26,
        $27,
        $28,
        $29
    );