DROP SCHEMA IF EXISTS postgres CASCADE;

CREATE SCHEMA postgres;

CREATE TABLE player (
    id INT PRIMARY KEY,
    fullName VARCHAR(100) NOT NULL
);

CREATE TABLE team (
    fullName VARCHAR(100) NOT NULL,
    abbr VARCHAR(3) PRIMARY KEY
);

CREATE TABLE totals (
    id INT PRIMARY KEY,
    gp INT NOT NULL,
    gs INT,
    mp INT NOT NULL,
    fg INT NOT NULL,
    fga INT NOT NULL,
    "3p" INT NOT NULL,
    "3pa" INT NOT NULL,
    "2p" INT NOT NULL,
    "2pa" INT NOT NULL,
    ft INT NOT NULL,
    fta INT NOT NULL,
    orb INT NOT NULL,
    drb INT NOT NULL,
    trb INT NOT NULL,
    stl INT NOT NULL,
    blk INT NOT NULL,
    ast INT NOT NULL,
    tov INT NOT NULL,
    pf INT NOT NULL,
    pts INT NOT NULL
);

CREATE Table per_100_possesions (
    id INT PRIMARY KEY,
    fg REAL NOT NULL,
    fga REAL NOT NULL,
    "3p" INT NOT NULL,
    "3pa" INT NOT NULL,
    "2p" INT NOT NULL,
    "2pa" INT NOT NULL,
    ft REAL NOT NULL,
    fta REAL NOT NULL,
    orb REAL NOT NULL,
    drb REAL NOT NULL,
    trb REAL NOT NULL,
    stl REAL NOT NULL,
    blk REAL NOT NULL,
    ast REAL NOT NULL,
    tov REAL NOT NULL,
    pf REAL NOT NULL,
    pts REAL NOT NULL,
    o_rtg REAL,
    d_rtg REAL
);

CREATE Table per_game (
    id INT PRIMARY KEY,
    mp REAL NOT NULL,
    fg REAL NOT NULL,
    fga REAL NOT NULL,
    fg_percent REAL NOT NULL,
    "3p" INT NOT NULL,
    "3pa" INT NOT NULL,
    "3p_percent" REAL NOT NULL,
    "2p" INT NOT NULL,
    "2pa" INT NOT NULL,
    "2p_percent" REAL NOT NULL,
    efg_percent REAL NOT NULL
    ft REAL NOT NULL,
    fta REAL NOT NULL,
    ft_percent REAL NOT NULL,
    orb REAL NOT NULL,
    drb REAL NOT NULL,
    trb REAL NOT NULL,
    ast REAL NOT NULL,
    stl REAL NOT NULL,
    blk REAL NOT NULL,
    tov REAL NOT NULL,
    pf REAL NOT NULL,
    pts REAL NOT NULL,
);

CREATE Table advanced (
    id INT PRIMARY KEY,
    "per" REAL NOT NULL,
    ts_percent REAL NOT NULL,
    "3p_ar" REAL NOT NULL,
    f_tr REAL NOT NULL,
    orb_percent REAL NOT NULL,
    drb_percent REAL NOT NULL,
    trb_percent REAL NOT NULL,
    ast_percent REAL NOT NULL,
    stl_percent REAL NOT NULL,
    blk_percent REAL NOT NULL,
    tov_percent REAL NOT NULL,
    usg_percent REAL NOT NULL,
    ows REAL NOT NULL,
    dws REAL NOT NULL,
    ws REAL NOT NULL,
    ws48 REAL NOT NULL,
    obpm REAL NOT NULL,
    dbpm REAL NOT NULL,
    bpm REAL NOT NULL,
    vorp REAL NOT NULL
);
-- analyst voting for top 15 players in all categories ()
CREATE TABLE all_teams_voting (
    player_id INT NOT NULL REFERENCES player (id),
    season_year INT NOT NULL,
    "type" VARCHAR(50) NOT NULL,
    pts_won INT NOT NULL,
    pts_max INT NOT NULL,
    "share" REAL NOT NULL,
    "1st_team" INT NOT NULL,
    "2nd_team" INT NOT NULL,
    "3rd_team" INT NOT NULL,
    PRIMARY KEY (player_id, season_year),
);

CREATE TABLE per_36 (
    player_id INT NOT NULL REFERENCES player (id),
    season_year INT NOT NULL,
    fg REAL NOT NULL,
    fga REAL NOT NULL,
    "3p" REAL NOT NULL,
    "3pa" REAL NOT NULL,
    "2p" REAL NOT NULL,
    "2pa" REAL NOT NULL,
    ft REAL NOT NULL,
    fta REAL NOT NULL,
    orb REAL NOT NULL,
    drb REAL NOT NULL,
    trb REAL NOT NULL,
    stl REAL NOT NULL,
    blk REAL NOT NULL,
    ast REAL NOT NULL,
    tov REAL NOT NULL,
    pf REAL NOT NULL,
    pts REAL NOT NULL,
    PRIMARY KEY (player_id, season_year)
);

CREATE TABLE all_teams (
    player_id INT NOT NULL REFERENCES player (id),
    season_year INT NOT NULL,
    "type" VARCHAR(50) NOT NULL,
    team_number VARCHAR(3) NOT NULL,
    PRIMARY KEY (player_id, season_year)
);

CREATE TABLE player_team (
    team_abbr VARCHAR(3) NOT NULL REFERENCES team (abbr),
    player_id INT NOT NULL REFERENCES player (id),
    season_year INT NOT NULL,
    age INT NOT NULL,
    experience INT NOT NULL,
    position VARCHAR(5) NOT NULL,
    PRIMARY KEY (
        player_id,
        team_abbr,
        season_year
    )
);

CREATE TABLE team_totals (
    team_abbr VARCHAR(3) NOT NULL REFERENCES team (abbr),
    total_id INT NOT NULL REFERENCES totals (id),
    season_year INT NOT NULL,
    PRIMARY KEY (team_abbr, season_year)
);

CREATE TABLE team_per_100_possesions (
    team_abbr VARCHAR(3) NOT NULL REFERENCES team (abbr),
    per_100_id INT NOT NULL REFERENCES per_100_possesions (id),
    season_year INT NOT NULL,
    PRIMARY KEY (team_abbr, season_year)
);

CREATE TABLE team_per_game (
    team_abbr VARCHAR(3) NOT NULL REFERENCES team (abbr),
    per_game_id INT NOT NULL REFERENCES per_game (id),
    season_year INT NOT NULL,
    PRIMARY KEY (team_abbr, season_year)
);

CREATE TABLE player_totals (
    player_id INT NOT NULL REFERENCES player (id),
    total_id INT NOT NULL REFERENCES totals (id),
    season_year INT NOT NULL,
    PRIMARY KEY (player_id, season_year)
);

CREATE TABLE player_per_100_possesions (
    player_id INT NOT NULL REFERENCES player (id),
    per_100_id INT NOT NULL REFERENCES per_100_possesions (id),
    season_year INT NOT NULL,
    PRIMARY KEY (player_id, season_year)
);

CREATE TABLE player_per_game (
    player_id INT NOT NULL REFERENCES player (id),
    per_game_id INT NOT NULL REFERENCES per_game (id),
    season_year INT NOT NULL,
    PRIMARY KEY (player_id, season_year)
);

CREATE TABLE opponents_totals (
    team_abbr VARCHAR(3) NOT NULL REFERENCES team (abbr),
    total_id INT NOT NULL REFERENCES totals (id),
    season_year INT NOT NULL,
    PRIMARY KEY (team_abbr, season_year)
);

CREATE TABLE opponents_per_100_possesions (
    team_abbr VARCHAR(3) NOT NULL REFERENCES team (abbr),
    per_100_id INT NOT NULL REFERENCES per_100_possesions (id),
    season_year INT NOT NULL,
    PRIMARY KEY (team_abbr, season_year)
);

CREATE TABLE opponents_per_game (
    team_abbr VARCHAR(3) NOT NULL REFERENCES team (abbr),
    per_game_id INT NOT NULL REFERENCES per_game (id),
    season_year INT NOT NULL,
    PRIMARY KEY (team_abbr, season_year)
);

CREATE TABLE all_stars (
    playerFullName VARCHAR(100) NOT NULL,
    season_year INT NOT NULL,
    teamName VARCHAR(255),
    replaced BOOLEAN
);

CREATE TABLE player_advanced (
    player_id INT NOT NULL REFERENCES player (id),
    advanced_id INT NOT NULL REFERENCES advanced (id),
    season_year INT NOT NULL,
    PRIMARY KEY (player_id, season_year)
);

CREATE TABLE player_shooting (
    season_year INT NOT NULL,
    player_id INT NOT NULL REFERENCES player (id),
    avg_dist_fga REAL NOT NULL REAL NOT NULL,
    percent_fga_from_2p_range REAL NOT NULL,
    percent_fga_from_0_3_range REAL NOT NULL,
    percent_fga_from_3_10_range REAL NOT NULL,
    percent_fga_from_10_16_range REAL NOT NULL,
    percent_fga_from_16_3p_range REAL NOT NULL,
    percent_fga_from_3p_range REAL NOT NULL,
    fg_percent_from_2p_range REAL NOT NULL,
    fg_percent_from_0_3_range REAL NOT NULL,
    fg_percent_from_3_10_range REAL NOT NULL,
    fg_percent_from_10_16_range REAL NOT NULL,
    fg_percent_from_16_3p_range REAL NOT NULL,
    fg_percent_from_3p_range REAL NOT NULL,
    percent_assisted_2p_fg REAL NOT NULL,
    percent_assisted_3p_fg REAL NOT NULL,
    percent_dunks_of_fga REAL NOT NULL,
    num_of_dunks REAL NOT NULL,
    percent_corner_3s_of_3pa REAL NOT NULL,
    corner_3_point_percent REAL NOT NULL,
    num_heaves_attempted INT NOT NULL,
    num_heaves_made INT NOT NULL,
    PRIMARY KEY(season_year, player_id)
);

CREATE TABLE player_awards (
    player_id INT NOT NULL REFERENCES player (id),
    season_year INT NOT NULL,
    award VARCHAR(50) NOT NULL,
    pts_won INT NOT NULL,
    pts_max INT NOT NULL,
    share REAL NOT NULL,
    winner BOOLEAN NOT NULL,
    PRIMARY KEY(season_year, player_id)
);

CREATE TABLE team_season (
    season_year INT NOT NULL,
    team_abbr VARCHAR(3) NOT NULL REFERENCES team (abbr),
    playoffs BOOLEAN NOT NULL,
    avarage_age REAL NOT NULL,
    w INT NOT NULL,
    l INT NOT NULL,
    pw INT NOT NULL,
    pl INT NOT NULL,
    mov REAL NOT NULL,
    sos REAL NOT NULL,
    srs REAL NOT NULL,
    o_rtg REAL NOT NULL,
    d_rtg REAL NOT NULL,
    n_rtg REAL NOT NULL,
    pace REAL NOT NULL,
    f_tr REAL NOT NULL,
    3p_ar REAL NOT NULL,     
    ts_percent REAL NOT NULL,
    e_fg_percent REAL NOT NULL,
    tov_percent REAL NOT NULL,
    orb_percent REAL NOT NULL,
    ft_fga REAL NOT NULL,
    opp_e_fg_percent REAL NOT NULL,
    opp_tov_percent REAL NOT NULL,
    opp_drb_percent REAL NOT NULL,
    opp_ft_fga REAL NOT NULL,
    arena VARCHAR(100) NOT NULL,
    attend INT NOT NULL,
    attend_g INT NOT NULL,
);