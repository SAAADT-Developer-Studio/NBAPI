package main

import (
	"NBAPI/internal/database"
	"NBAPI/internal/sqlc"
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	database.Init()
	seedPlayers()
	seedTeams()
	seedTotals()
	seedTeamTotals()
	seedPlayerTotals()
	seedPer100Possesions()
	seedTeamPer100Possessions()
	seedPlayerPer100Possessions()
	seedPerGame()
	seedTeamPerGame()
	seedPlayerPerGame()
	seedOpponentsTotals()
	seedOpponentsPerGame()
	seedOpponentsPer100Possessions()
	seedAdvanced()
	seedPlayerAdvanced()
	seedAllTeamsVoting()
	seedAllTeams()
	seedPer36()
	seedPlayerShooting()
	seedAllStars()
	seedPlayerTeam()
	seedPlayerAwards()
	seedTeamSeason()
}

func getCSVRows(csvFile string) [][]string {
	logrus.Infof("Importing file %s", csvFile)
	file, err := os.Open("csv/" + csvFile)
	if err != nil {
		log.Fatalf("Failed to open the file: %s...", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %s", err)
	}
	return records
}

func parseInt(value string) int32 {
	num, _ := strconv.Atoi(value)
	return int32(num)
}

func parseOptInt(value string) pgtype.Int4 {
	num, err := strconv.Atoi(value)
	if err != nil {
		return pgtype.Int4{Int32: 0, Valid: false}
	}
	return pgtype.Int4{Int32: int32(num), Valid: true}
}

func parseFloat(value string) float32 {
	num, _ := strconv.ParseFloat(value, 32)
	return float32(num)
}

func parseOptFloat(value string) pgtype.Float4 {
	num, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return pgtype.Float4{Float32: 0, Valid: false}
	}
	return pgtype.Float4{Float32: float32(num), Valid: true}
}

func seedPlayers() {
	ctx := context.Background()
	for _, row := range getCSVRows("PSI.csv") {
		database.Queries.CreatePlayer(ctx, sqlc.CreatePlayerParams{ID: parseInt(row[0]), Fullname: row[1]})
	}
}

func seedTeams() {
	ctx := context.Background()
	for _, row := range getCSVRows("TA.csv") {
		database.Queries.CreateTeam(ctx, sqlc.CreateTeamParams{
			Abbr: row[4], Fullname: row[2],
		})
	}
}

func seedTotals() {
	ctx := context.Background()

	for _, row := range getCSVRows("PT.csv") {
		database.Queries.CreateTotals(ctx, sqlc.CreateTotalsParams{
			ID:  parseInt(row[0]),
			Gp:  parseInt(row[11]),
			Gs:  parseOptInt(row[12]),
			Mp:  parseInt(row[13]),
			Fg:  parseInt(row[14]),
			Fga: parseInt(row[15]),
			P3:  parseInt(row[16]),
			Pa3: parseInt(row[17]),
			P2:  parseInt(row[18]),
			Pa2: parseInt(row[20]),
			Ft:  parseInt(row[21]),
			Fta: parseInt(row[24]),
			Orb: parseInt(row[25]),
			Drb: parseInt(row[27]),
			Trb: parseInt(row[28]),
			Stl: parseInt(row[29]),
			Blk: parseInt(row[31]),
			Ast: parseInt(row[32]),
			Tov: parseInt(row[33]),
			Pf:  parseInt(row[34]),
			Pts: parseInt(row[35]),
		})
	}

	for _, row := range getCSVRows("TT.csv") {
		database.Queries.CreateTotals(ctx, sqlc.CreateTotalsParams{
			ID:  parseInt(row[0]),
			Gp:  parseInt(row[6]),
			Gs:  parseOptInt(row[6]),
			Mp:  parseInt(row[7]),
			Fg:  parseInt(row[8]),
			Fga: parseInt(row[9]),
			P3:  parseInt(row[11]),
			Pa3: parseInt(row[12]),
			P2:  parseInt(row[14]),
			Pa2: parseInt(row[15]),
			Ft:  parseInt(row[17]),
			Fta: parseInt(row[18]),
			Orb: parseInt(row[20]),
			Drb: parseInt(row[21]),
			Trb: parseInt(row[22]),
			Stl: parseInt(row[24]),
			Blk: parseInt(row[25]),
			Ast: parseInt(row[23]),
			Tov: parseInt(row[26]),
			Pf:  parseInt(row[27]),
			Pts: parseInt(row[28]),
		})
	}
	for _, row := range getCSVRows("OT.csv") {
		database.Queries.CreateTotals(ctx, sqlc.CreateTotalsParams{
			ID:  parseInt(row[0]),
			Gp:  parseInt(row[6]),
			Gs:  parseOptInt(row[6]),
			Mp:  parseInt(row[7]),
			Fg:  parseInt(row[8]),
			Fga: parseInt(row[9]),
			P3:  parseInt(row[11]),
			Pa3: parseInt(row[12]),
			P2:  parseInt(row[14]),
			Pa2: parseInt(row[15]),
			Ft:  parseInt(row[17]),
			Fta: parseInt(row[18]),
			Orb: parseInt(row[20]),
			Drb: parseInt(row[21]),
			Trb: parseInt(row[22]),
			Stl: parseInt(row[24]),
			Blk: parseInt(row[25]),
			Ast: parseInt(row[23]),
			Tov: parseInt(row[26]),
			Pf:  parseInt(row[27]),
			Pts: parseInt(row[28]),
		})
	}

}

func seedTeamTotals() {
	ctx := context.Background()
	for _, row := range getCSVRows("TT.csv") {
		database.Queries.CreateTeamTotals(ctx, sqlc.CreateTeamTotalsParams{
			TeamAbbr:   row[4],
			TotalID:    parseInt(row[0]),
			SeasonYear: parseInt(row[1]),
		})
	}
}

func seedPlayerTotals() {
	ctx := context.Background()
	for _, row := range getCSVRows("TT.csv") {
		database.Queries.CreatePlayerTotals(ctx, sqlc.CreatePlayerTotalsParams{
			PlayerID:   parseInt(row[4]),
			TotalID:    parseInt(row[0]),
			SeasonYear: parseInt(row[1]),
		})
	}
}

func seedPer100Possesions() {
	ctx := context.Background()
	for _, row := range getCSVRows("TSP100P.csv") {
		database.Queries.CreatePer100Possesions(ctx, sqlc.CreatePer100PossesionsParams{
			ID:   parseInt(row[0]),
			Fg:   parseFloat(row[8]),
			Fga:  parseFloat(row[9]),
			P3:   parseInt(row[11]),
			Pa3:  parseInt(row[12]),
			P2:   parseInt(row[14]),
			Pa2:  parseInt(row[15]),
			Ft:   parseFloat(row[17]),
			Fta:  parseFloat(row[18]),
			Orb:  parseFloat(row[20]),
			Drb:  parseFloat(row[21]),
			Trb:  parseFloat(row[22]),
			Stl:  parseFloat(row[24]),
			Blk:  parseFloat(row[25]),
			Ast:  parseFloat(row[23]),
			Tov:  parseFloat(row[26]),
			Pf:   parseFloat(row[27]),
			Pts:  parseFloat(row[28]),
			ORtg: parseOptFloat(row[19]), // Nullable
			DRtg: parseOptFloat(row[13]), // Nullable

		})
	}
	for _, row := range getCSVRows("PP100P.csv") {
		database.Queries.CreatePer100Possesions(ctx, sqlc.CreatePer100PossesionsParams{
			ID:   parseInt(row[0]),
			Fg:   parseFloat(row[14]),
			Fga:  parseFloat(row[15]),
			P3:   parseInt(row[17]),
			Pa3:  parseInt(row[18]),
			P2:   parseInt(row[20]),
			Pa2:  parseInt(row[21]),
			Ft:   parseFloat(row[23]),
			Fta:  parseFloat(row[24]),
			Orb:  parseFloat(row[26]),
			Drb:  parseFloat(row[27]),
			Trb:  parseFloat(row[28]),
			Stl:  parseFloat(row[30]),
			Blk:  parseFloat(row[31]),
			Ast:  parseFloat(row[29]),
			Tov:  parseFloat(row[32]),
			Pf:   parseFloat(row[33]),
			Pts:  parseFloat(row[34]),
			ORtg: parseOptFloat(row[35]), // Nullable
			DRtg: parseOptFloat(row[36]), // Nullable

		})
	}

	for _, row := range getCSVRows("OP100P.csv") {
		database.Queries.CreatePer100Possesions(ctx, sqlc.CreatePer100PossesionsParams{
			ID:   parseInt(row[0]),
			Fg:   parseFloat(row[8]),
			Fga:  parseFloat(row[9]),
			P3:   parseInt(row[11]),
			Pa3:  parseInt(row[12]),
			P2:   parseInt(row[14]),
			Pa2:  parseInt(row[15]),
			Ft:   parseFloat(row[17]),
			Fta:  parseFloat(row[18]),
			Orb:  parseFloat(row[20]),
			Drb:  parseFloat(row[21]),
			Trb:  parseFloat(row[22]),
			Stl:  parseFloat(row[24]),
			Blk:  parseFloat(row[25]),
			Ast:  parseFloat(row[23]),
			Tov:  parseFloat(row[26]),
			Pf:   parseFloat(row[27]),
			Pts:  parseFloat(row[28]),
			ORtg: parseOptFloat("NA"),
			DRtg: parseOptFloat("NA"),
		})
	}
}

func seedTeamPer100Possessions() {
	ctx := context.Background()
	for _, row := range getCSVRows("TSP100P.csv") {
		database.Queries.CreateTeamPer100Possesions(ctx, sqlc.CreateTeamPer100PossesionsParams{
			TeamAbbr:   (row[4]),
			Per100ID:   parseInt(row[0]),
			SeasonYear: parseInt(row[1]),
		})
	}
}

func seedPlayerPer100Possessions() {
	ctx := context.Background()
	for _, row := range getCSVRows("PP100P.csv") {
		database.Queries.CreatePlayerPer100Possesions(ctx, sqlc.CreatePlayerPer100PossesionsParams{
			PlayerID:   parseInt(row[3]),
			Per100ID:   parseInt(row[0]),
			SeasonYear: parseInt(row[2]),
		})
	}
}

func seedPerGame() {
	ctx := context.Background()
	for _, row := range getCSVRows("PPG.csv") {
		database.Queries.CreatePerGame(ctx, sqlc.CreatePerGameParams{
			ID:         parseInt(row[0]),
			Mp:         parseFloat(row[13]),
			Fg:         parseFloat(row[14]),
			Fga:        parseFloat(row[15]),
			FgPercent:  parseFloat(row[16]),
			P3:         parseInt(row[17]),
			Pa3:        parseInt(row[18]),
			PPercent3:  parseFloat(row[19]),
			P2:         parseInt(row[20]),
			Pa2:        parseInt(row[21]),
			PPercent2:  parseFloat(row[22]),
			EfgPercent: parseFloat(row[23]),
			Ft:         parseFloat(row[24]),
			Fta:        parseFloat(row[25]),
			FtPercent:  parseFloat(row[26]),
			Orb:        parseFloat(row[27]),
			Drb:        parseFloat(row[28]),
			Trb:        parseFloat(row[29]),
			Ast:        parseFloat(row[30]),
			Stl:        parseFloat(row[31]),
			Blk:        parseFloat(row[32]),
			Tov:        parseFloat(row[33]),
			Pf:         parseFloat(row[34]),
			Pts:        parseFloat(row[35]),
		})
	}
	for _, row := range getCSVRows("TPG.csv") {
		database.Queries.CreatePerGame(ctx, sqlc.CreatePerGameParams{
			ID:         parseInt(row[0]),
			Mp:         parseFloat(row[7]),
			Fg:         parseFloat(row[8]),
			Fga:        parseFloat(row[9]),
			FgPercent:  parseFloat(row[10]),
			P3:         parseInt(row[11]),
			Pa3:        parseInt(row[12]),
			PPercent3:  parseFloat(row[13]),
			P2:         parseInt(row[14]),
			Pa2:        parseInt(row[15]),
			PPercent2:  parseFloat(row[16]),
			EfgPercent: parseFloat(row[16]),
			Ft:         parseFloat(row[17]),
			Fta:        parseFloat(row[18]),
			FtPercent:  parseFloat(row[19]),
			Orb:        parseFloat(row[20]),
			Drb:        parseFloat(row[21]),
			Trb:        parseFloat(row[22]),
			Ast:        parseFloat(row[23]),
			Stl:        parseFloat(row[24]),
			Blk:        parseFloat(row[25]),
			Tov:        parseFloat(row[26]),
			Pf:         parseFloat(row[27]),
			Pts:        parseFloat(row[28]),
		})
	}

}

func seedTeamPerGame() {
	ctx := context.Background()
	for _, row := range getCSVRows("TPG.csv") {
		database.Queries.CreateTeamPerGame(ctx, sqlc.CreateTeamPerGameParams{
			TeamAbbr:   (row[4]),
			PerGameID:  parseInt(row[0]),
			SeasonYear: parseInt(row[1]),
		})
	}
}

func seedPlayerPerGame() {
	ctx := context.Background()
	for _, row := range getCSVRows("PPG.csv") {
		database.Queries.CreatePlayerPerGame(ctx, sqlc.CreatePlayerPerGameParams{
			PlayerID:   parseInt(row[3]),
			PerGameID:  parseInt(row[0]),
			SeasonYear: parseInt(row[2]),
		})
	}
}

func seedOpponentsTotals() {
	ctx := context.Background()
	for _, row := range getCSVRows("OT.csv") {
		database.Queries.CreateOpponentsTotals(ctx, sqlc.CreateOpponentsTotalsParams{
			TeamAbbr:   (row[4]),
			TotalID:    parseInt(row[0]),
			SeasonYear: parseInt(row[1]),
		})
	}
}

func seedOpponentsPerGame() {
	ctx := context.Background()
	for _, row := range getCSVRows("OPG.csv") {
		database.Queries.CreateOpponentsPerGame(ctx, sqlc.CreateOpponentsPerGameParams{
			TeamAbbr:   (row[4]),
			PerGameID:  parseInt(row[0]),
			SeasonYear: parseInt(row[1]),
		})
	}
}

func seedOpponentsPer100Possessions() {
	ctx := context.Background()
	for _, row := range getCSVRows("OPG.csv") {
		database.Queries.CreateOpponentsPer100Possessions(ctx, sqlc.CreateOpponentsPer100PossessionsParams{
			TeamAbbr:   (row[4]),
			Per100ID:   parseInt(row[0]),
			SeasonYear: parseInt(row[1]),
		})
	}
}

func seedAdvanced() {
	ctx := context.Background()
	for _, row := range getCSVRows("Advanced.csv") {
		database.Queries.CreateAdvanced(ctx, sqlc.CreateAdvancedParams{
			ID:         parseInt(row[0]),
			Per:        parseFloat(row[13]),
			TsPercent:  parseFloat(row[14]),
			PAr3:       parseFloat(row[15]),
			FTr:        parseFloat(row[16]),
			OrbPercent: parseFloat(row[17]),
			DrbPercent: parseFloat(row[18]),
			TrbPercent: parseFloat(row[19]),
			AstPercent: parseFloat(row[20]),
			StlPercent: parseFloat(row[21]),
			BlkPercent: parseFloat(row[22]),
			TovPercent: parseFloat(row[23]),
			UsgPercent: parseFloat(row[24]),
			Ows:        parseFloat(row[25]),
			Dws:        parseFloat(row[26]),
			Ws:         parseFloat(row[27]),
			Ws48:       parseFloat(row[28]),
			Obpm:       parseFloat(row[29]),
			Dbpm:       parseFloat(row[30]),
			Bpm:        parseFloat(row[31]),
			Vorp:       parseFloat(row[32]),
		})
	}
}

func seedPlayerAdvanced() {
	ctx := context.Background()
	for _, row := range getCSVRows("Advanced.csv") {
		database.Queries.CreatePlayerAdvanced(ctx, sqlc.CreatePlayerAdvancedParams{
			PlayerID:   parseInt(row[3]),
			AdvancedID: parseInt(row[0]),
			SeasonYear: parseInt(row[2]),
		})
	}
}

func seedAllTeamsVoting() {
	ctx := context.Background()
	for _, row := range getCSVRows("ATV.csv") {
		database.Queries.CreateAllTeamsVoting(ctx, sqlc.CreateAllTeamsVotingParams{
			PlayerID:   parseInt(row[15]),
			SeasonYear: parseInt(row[0]),
			Type:       (row[2]),
			PtsWon:     parseInt(row[8]),
			PtsMax:     parseInt(row[9]),
			Share:      parseFloat(row[10]),
			FirstTeam:  parseInt(row[11]),
			SecondTeam: parseInt(row[12]),
			ThirdTeam:  parseInt(row[13]),
		})
	}
}

func seedAllTeams() {
	ctx := context.Background()
	for _, row := range getCSVRows("AT.csv") {
		database.Queries.CreateAllTeams(ctx, sqlc.CreateAllTeamsParams{
			PlayerID:   parseInt(row[7]),
			SeasonYear: parseInt(row[0]),
			Type:       (row[2]),
			TeamNumber: (row[3]),
		})
	}
}

func seedPer36() {
	ctx := context.Background()
	for _, row := range getCSVRows("P36M.csv") {
		database.Queries.CreatePer36(ctx, sqlc.CreatePer36Params{
			PlayerID:   parseInt(row[2]),
			SeasonYear: parseInt(row[1]),
			Fg:         parseFloat(row[13]),
			Fga:        parseFloat(row[14]),
			P3:         parseInt(row[16]),
			Pa3:        parseInt(row[17]),
			P2:         parseInt(row[19]),
			Pa2:        parseInt(row[20]),
			Ft:         parseFloat(row[22]),
			Fta:        parseFloat(row[23]),
			Orb:        parseFloat(row[25]),
			Drb:        parseFloat(row[26]),
			Trb:        parseFloat(row[27]),
			Stl:        parseFloat(row[29]),
			Blk:        parseFloat(row[30]),
			Ast:        parseFloat(row[28]),
			Tov:        parseFloat(row[31]),
			Pf:         parseFloat(row[32]),
			Pts:        parseFloat(row[33]),
		})
	}
}

func seedPlayerShooting() {
	ctx := context.Background()
	for _, row := range getCSVRows("PS.csv") {
		database.Queries.CreatePlayerShooting(ctx, sqlc.CreatePlayerShootingParams{
			SeasonYear:              parseInt(row[1]),
			PlayerID:                parseInt(row[2]),
			AvgDistFga:              parseFloat(row[13]),
			PercentFgaFrom2pRange:   parseFloat(row[14]),
			PercentFgaFrom03Range:   parseFloat(row[15]),
			PercentFgaFrom310Range:  parseFloat(row[16]),
			PercentFgaFrom1016Range: parseFloat(row[17]),
			PercentFgaFrom163pRange: parseFloat(row[18]),
			PercentFgaFrom3pRange:   parseFloat(row[19]),
			FgPercentFrom2pRange:    parseFloat(row[20]),
			FgPercentFrom03Range:    parseFloat(row[21]),
			FgPercentFrom310Range:   parseFloat(row[22]),
			FgPercentFrom1016Range:  parseFloat(row[23]),
			FgPercentFrom163pRange:  parseFloat(row[24]),
			FgPercentFrom3pRange:    parseFloat(row[25]),
			PercentAssisted2pFg:     parseFloat(row[26]),
			PercentAssisted3pFg:     parseFloat(row[27]),
			PercentDunksOfFga:       parseFloat(row[28]),
			NumOfDunks:              parseFloat(row[29]),
			PercentCorner3sOf3pa:    parseFloat(row[30]),
			Corner3PointPercent:     parseFloat(row[31]),
			NumHeavesAttempted:      parseInt(row[32]),
			NumHeavesMade:           parseInt(row[33]),
		})
	}
}

func seedAllStars() {
	ctx := context.Background()
	for _, row := range getCSVRows("ASS.csv") {
		database.Queries.CreateAllStars(ctx, sqlc.CreateAllStarsParams{
			Playerfullname: (row[0]),
			SeasonYear:     parseInt(row[3]),
			Teamname:       pgtype.Text{String: row[1], Valid: true},
			Replaced:       pgtype.Bool{Bool: row[4] == "TRUE"},
		})
	}

}

func seedPlayerTeam() {
	ctx := context.Background()
	for _, row := range getCSVRows("PTI.csv") {
		database.Queries.CreatePlayerTeam(ctx, sqlc.CreatePlayerTeamParams{
			TeamAbbr:   (row[8]),
			PlayerID:   parseInt(row[2]),
			SeasonYear: parseInt(row[0]),
			Age:        parseInt(row[6]),
			Experience: parseInt(row[9]),
			Position:   (row[5]),
		})
	}

}

func seedPlayerAwards() {
	ctx := context.Background()
	for _, row := range getCSVRows("PAS.csv") {
		database.Queries.CreatePlayerAwards(ctx, sqlc.CreatePlayerAwardsParams{
			PlayerID:   parseInt(row[11]),
			SeasonYear: parseInt(row[0]),
			Award:      (row[1]),
			PtsWon:     parseInt(row[6]),
			PtsMax:     parseInt(row[7]),
			Share:      parseFloat(row[8]),
			Winner:     (row[9] == "TRUE"),
		})
	}

}

func seedTeamSeason() {
	ctx := context.Background()
	for _, row := range getCSVRows("TS.csv") {
		database.Queries.CreateTeamSeason(ctx, sqlc.CreateTeamSeasonParams{
			SeasonYear:    parseInt(row[0]),
			TeamAbbr:      row[3],
			Playoffs:      (row[4] == "TRUE"), // Note: BOOLEAN values may need special handling
			AvarageAge:    parseFloat(row[5]),
			W:             parseInt(row[6]),
			L:             parseInt(row[7]),
			Pw:            parseInt(row[8]),
			Pl:            parseInt(row[9]),
			Mov:           parseFloat(row[10]),
			Sos:           parseFloat(row[11]),
			Srs:           parseFloat(row[12]),
			ORtg:          parseFloat(row[13]),
			DRtg:          parseFloat(row[14]),
			NRtg:          parseFloat(row[15]),
			Pace:          parseFloat(row[16]),
			FTr:           parseFloat(row[17]),
			PAr3:          parseFloat(row[18]),
			TsPercent:     parseFloat(row[19]),
			EFgPercent:    parseFloat(row[20]),
			TovPercent:    parseFloat(row[21]),
			OrbPercent:    parseFloat(row[22]),
			FtFga:         parseFloat(row[23]),
			OppEFgPercent: parseFloat(row[24]),
			OppTovPercent: parseFloat(row[25]),
			OppDrbPercent: parseFloat(row[26]),
			OppFtFga:      parseFloat(row[27]),
			Arena:         (row[28]),
			Attend:        parseInt(row[29]),
			AttendG:       parseInt(row[30]),
		})
	}

}
