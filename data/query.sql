-- name: GetPlayers :many
select
    *
from
    player;

-- name: GetPlayer :one
select
    player.id,
    player.fullName,
    player_totals.season_year,
    totals.gp,
    totals.gs,
    totals.mp as total_mp,
    totals.fg as totals_fg,
    totals.fga as totals_fga,
    totals.p3 as totals_p3,
    totals.pa3 as totals_pa3,
    totals.p2 as totals_p2,
    totals.pa2 as totals_pa2,
    totals.ft as totals_ft,
    totals.fta as totals_fta,
    totals.orb as totals_orb,
    totals.drb as totals_drb,
    totals.trb as totals_trb,
    totals.stl as totals_stl,
    totals.blk as totals_blk,
    totals.ast as totals_ast,
    totals.tov as totals_tov,
    totals.pf as totals_pf,
    totals.pts as totals_pts,
    per_game.mp as per_game_mp,
    per_game.fg as per_game_fg,
    per_game.fga as per_game_fga,
    per_game.fg_percent as per_game_fg_percent,
    per_game.p3 as per_game_p3,
    per_game.pa3 as per_game_pa3,
    per_game.p_percent3 as per_game_p_percent3,
    per_game.p2 as per_game_p2,
    per_game.pa2 as per_game_pa2,
    per_game.p_percent2 as per_game_p_percent2,
    per_game.efg_percent as per_game_efg_percent,
    per_game.ft as per_game_ft,
    per_game.fta as per_game_fta,
    per_game.ft_percent as per_game_ft_percent,
    per_game.orb as per_game_orb,
    per_game.drb as per_game_drb,
    per_game.trb as per_game_trb,
    per_game.ast as per_game_ast,
    per_game.stl as per_game_stl,
    per_game.blk as per_game_blk,
    per_game.tov as per_game_tov,
    per_game.pf as per_game_pf,
    per_game.pts as per_game_pts,
    per_100_possesions.fg as per_100_possesions_fg,
    per_100_possesions.fga as per_100_possesions_fga,
    per_100_possesions.p3 as per_100_possesions_p3,
    per_100_possesions.pa3 as per_100_possesions_pa3,
    per_100_possesions.p2 as per_100_possesions_p2,
    per_100_possesions.pa2 as per_100_possesions_pa2,
    per_100_possesions.ft as per_100_possesions_ft,
    per_100_possesions.fta as per_100_possesions_fta,
    per_100_possesions.orb as per_100_possesions_orb,
    per_100_possesions.drb as per_100_possesions_drb,
    per_100_possesions.trb as per_100_possesions_trb,
    per_100_possesions.stl as per_100_possesions_stl,
    per_100_possesions.blk as per_100_possesions_blk,
    per_100_possesions.ast as per_100_possesions_ast,
    per_100_possesions.tov as per_100_possesions_tov,
    per_100_possesions.pf as per_100_possesions_pf,
    per_100_possesions.pts as per_100_possesions_pts,
    per_100_possesions.o_rtg as per_100_possesions_o_rtg,
    per_100_possesions.d_rtg as per_100_possesions_d_rtg,
    per_36.fg as per_36_fg,
    per_36.fga as per_36_fga,
    per_36.p3 as per_36_p3,
    per_36.pa3 as per_36_pa3,
    per_36.p2 as per_36_p2,
    per_36.pa2 as per_36_pa2,
    per_36.ft as per_36_ft,
    per_36.fta as per_36_fta,
    per_36.orb as per_36_orb,
    per_36.drb as per_36_drb,
    per_36.trb as per_36_trb,
    per_36.stl as per_36_stl,
    per_36.blk as per_36_blk,
    per_36.ast as per_36_ast,
    per_36.tov as per_36_tov,
    per_36.pf as per_36_pf,
    per_36.pts as per_36_pts,
    player_shooting.avg_dist_fga,
    player_shooting.percent_fga_from_2p_range,
    player_shooting.percent_fga_from_0_3_range,
    player_shooting.percent_fga_from_3_10_range,
    player_shooting.percent_fga_from_10_16_range,
    player_shooting.percent_fga_from_16_3p_range,
    player_shooting.percent_fga_from_3p_range,
    player_shooting.fg_percent_from_2p_range,
    player_shooting.fg_percent_from_0_3_range,
    player_shooting.fg_percent_from_3_10_range,
    player_shooting.fg_percent_from_10_16_range,
    player_shooting.fg_percent_from_16_3p_range,
    player_shooting.fg_percent_from_3p_range,
    player_shooting.percent_assisted_2p_fg,
    player_shooting.percent_assisted_3p_fg,
    player_shooting.percent_dunks_of_fga,
    player_shooting.num_of_dunks,
    player_shooting.percent_corner_3s_of_3pa,
    player_shooting.corner_3_point_percent,
    player_shooting.num_heaves_attempted,
    player_shooting.num_heaves_made,
    advanced."per",
    advanced.ts_percent,
    advanced.p_ar3,
    advanced.f_tr,
    advanced.orb_percent,
    advanced.drb_percent,
    advanced.trb_percent,
    advanced.ast_percent,
    advanced.stl_percent,
    advanced.blk_percent,
    advanced.tov_percent,
    advanced.usg_percent,
    advanced.ows,
    advanced.dws,
    advanced.ws,
    advanced.ws48,
    advanced.obpm,
    advanced.dbpm,
    advanced.bpm,
    advanced.vorp
from
    player
    left join player_totals on player.id = totals.player_id
    left join totals on totals.id = player_totals.total_id
    left join player_per_game on player.id = player_per_game.player_id
    left join per_game on per_game.id = player_per_game.per_game_id
    left join player_per_100_possesions on player.id = player_per_100_possesions.player_id
    left join per_100_possesions on per_100_possesions.id = player_per_100_possesions.per_100_id
    left join per_36 on per_36.player_id = player.id
    left join player_shooting on player_shooting.player_id = player.id
    left join player_advanced on player_advanced.player_id = player_id
    left join advanced on advanced.id = player_advanced.advanced_id
where
    player.id = $1
    and per_36.season_year between $2
    and $3
    and player_totals.season_year between $2
    and $3
    and player_per_game.season_year between $2
    and $3
    and player_per_100_possesions.season_year between $2
    and $3
    and player_shooting.season_year between $2
    and $3
    and player_advanced.season_year between $2
    and $3;

-- name: GetPlayerBySearch :many
select
    *
from
    player
where
    name like '%' || $1 || '%'; 

-- name: CreatePlayer :exec
insert into
    player (id, fullName)
values
    ($1, $2);

-- name: CreateTeam :exec
insert into
    team (abbr, fullName)
values
    ($1, $2);

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
values
    (
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
values
    ($1,$2,$3);

-- name: CreatePlayerTotals :exec
insert into
    player_totals (
        player_id,
        total_id,
        season_year
    )
values
    ($1,$2,$3);

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
VALUES
    (
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
values
    ($1,$2,$3);

-- name: CreatePlayerPer100Possesions :exec
insert into
    player_per_100_possesions (
        player_id,
        per_100_id,
        season_year
    )
values
    ($1,$2,$3);

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
values
    (
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
values
    ($1,$2,$3);

-- name: CreatePlayerPerGame :exec
insert into
    player_per_game (
        player_id,
        per_game_id,
        season_year
    )
values
    ($1,$2,$3);

-- name: CreateOpponentsTotals :exec
insert into
    opponents_totals (
        team_abbr,
        total_id,
        season_year
    )
values
    ($1,$2,$3);

-- name: CreateOpponentsPerGame :exec
insert into
    opponents_per_game (
        team_abbr,
        per_game_id,
        season_year
    )
values
    ($1,$2,$3);

-- name: CreateOpponentsPer100Possessions :exec
insert into
    opponents_per_100_possesions (
        team_abbr,
        per_100_id,
        season_year
    )
values
    ($1,$2,$3);

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
values
    (
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
values
    ($1,$2,$3);

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
values
    (
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
values
    ($1,$2,$3,$4);

-- name: CreatePer36 :exec
INSERT INTO
    per_36 (
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
VALUES
    (
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
values
    (
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
values
    ($1,$2,$3,$4);

-- name: CreatePlayerTeam :exec
insert into
    player_team (
        team_abbr,
        player_id,
        season_year,
        age,
        experience,
        position
    )
values
    ($1,$2,$3,$4,$5,$6);

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
values
    ($1,$2,$3,$4,$5,$6,$7);

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
values
    (
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