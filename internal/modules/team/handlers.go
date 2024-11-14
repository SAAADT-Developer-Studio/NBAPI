package team

import (
	"NBAPI/internal/database"
	"NBAPI/internal/inputs"
	"NBAPI/internal/sqlc"
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
)

type TeamsResponseBody struct {
	Teams    []sqlc.Team `json:"teams"`
	NextPage *string     `json:"next_page" example:"NYK" doc:"Next page cursor"`
}

type TeamsResponse struct {
	Body TeamsResponseBody
}

type TeamsInput struct {
	inputs.PaginationParams
	Search string `query:"search" doc:"Filter results based on a search string."`
}

func TeamsHandler(ctx context.Context, input *TeamsInput) (*TeamsResponse, error) {
	search := input.Search
	pageCursor := input.Cursor
	pageSize := input.Limit
	teams, err := database.Queries.GetTeams(ctx, sqlc.GetTeamsParams{Search: search, PageSize: int32(pageSize), Cursor: pageCursor})
	fmt.Println("teams", teams)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if teams == nil {
		teams = []sqlc.Team{}
	}
	var nextPage *string
	if len(teams) > pageSize {
		nextPage = &teams[pageSize].Abbr
		teams = teams[:pageSize]
	}

	response := &TeamsResponse{
		Body: TeamsResponseBody{
			Teams:    teams,
			NextPage: nextPage,
		},
	}

	return response, nil
}

type TeamResponseBody struct {
	Team       sqlc.Team                `json:"team"`
	Totals     []sqlc.GetTeamTotalsRow  `json:"totals"`
	Per100Poss []sqlc.Per100Possesion   `json:"per_100_possesions"`
	PerGame    []sqlc.GetTeamPerGameRow `json:"per_game"`
}

type TeamResponse struct {
	Body TeamResponseBody
}

type TeamInput struct {
	Abbr string `path:"teamId" doc:"The team ID to fetch."`
	inputs.SeasonRangeParams
}

func TeamHandler(ctx context.Context, input *TeamInput) (*TeamResponse, error) {
	abbr := input.Abbr
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)
	team, teamErr := database.Queries.GetTeam(ctx, abbr)
	totalsRows, totalsErr := database.Queries.GetTeamTotals(ctx,
		sqlc.GetTeamTotalsParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)
	per100Rows, per100Err := database.Queries.GetTeamPer100Possesions(ctx,
		sqlc.GetTeamPer100PossesionsParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)
	perGameRows, perGameErr := database.Queries.GetTeamPerGame(ctx,
		sqlc.GetTeamPerGameParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if teamErr != nil {
		logrus.Error(teamErr)
		return nil, huma.Error500InternalServerError(fmt.Sprintf("Error fetching team with id %s", abbr))
	}

	if totalsErr != nil || per100Err != nil || perGameErr != nil {
		logrus.Error(totalsErr)
		return nil, huma.Error500InternalServerError("Error fetching subtables", teamErr, per100Err, perGameErr)
	}

	response := &TeamResponse{
		Body: TeamResponseBody{Team: team, Totals: totalsRows, Per100Poss: per100Rows, PerGame: perGameRows},
	}

	return response, nil
}

type TeamPerGameResponse struct {
	Body []sqlc.GetTeamPerGameRow
}

type TeamPerGameInput struct {
	TeamInput
}

func TeamPerGameStatsHandler(ctx context.Context, input *TeamPerGameInput) (*TeamPerGameResponse, error) {
	abbr := input.Abbr
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	perGame, err := database.Queries.GetTeamPerGame(ctx,
		sqlc.GetTeamPerGameParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	response := &TeamPerGameResponse{
		Body: perGame,
	}

	return response, nil
}

type TeamPer100PossResponse struct {
	Body []sqlc.Per100Possesion
}

type TeamPer100PossInput struct {
	TeamInput
}

func TeamPer100PossStatsHandler(ctx context.Context, input *TeamPer100PossInput) (*TeamPer100PossResponse, error) {
	abbr := input.Abbr
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	per100, err := database.Queries.GetTeamPer100Possesions(ctx,
		sqlc.GetTeamPer100PossesionsParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &TeamPer100PossResponse{Body: per100}, nil
}

type TeamTotalsResponse struct {
	Body []sqlc.GetTeamTotalsRow
}
type TeamTotalsInput struct {
	TeamInput
}

func TeamTotalsStatsHandler(ctx context.Context, input *TeamTotalsInput) (*TeamTotalsResponse, error) {
	abbr := input.Abbr
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	totals, err := database.Queries.GetTeamTotals(ctx,
		sqlc.GetTeamTotalsParams{Abbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &TeamTotalsResponse{Body: totals}, nil
}

type TeamOpponentsTotalsResponse struct {
	Body []sqlc.GetOpponentsTotalsRow
}

type TeamOpponentsTotalsInput struct {
	TeamInput
}

func TeamTotalsOpponentsHandler(ctx context.Context, input *TeamInput) (*TeamOpponentsTotalsResponse, error) {
	abbr := input.Abbr
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	totals, err := database.Queries.GetOpponentsTotals(ctx,
		sqlc.GetOpponentsTotalsParams{TeamAbbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &TeamOpponentsTotalsResponse{Body: totals}, nil
}

type TeamOpponentsPerGameResponse struct {
	Body []sqlc.GetOpponentsPerGameRow
}

type TeamOpponentsPerGameInput struct {
	TeamInput
}

func TeamPerGameOpponentsHandler(ctx context.Context, input *TeamOpponentsPerGameInput) (*TeamOpponentsPerGameResponse, error) {
	abbr := input.Abbr
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	perGame, err := database.Queries.GetOpponentsPerGame(ctx,
		sqlc.GetOpponentsPerGameParams{TeamAbbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &TeamOpponentsPerGameResponse{Body: perGame}, nil
}

type TeamOpponentsPer100PossResponse struct {
	Body []sqlc.GetOpponentsPer100PossesionsRow
}
type TeamOpponentsPer100PossInput struct {
	TeamInput
}

func TeamPer100PossOpponentsHandler(ctx context.Context, input *TeamOpponentsPer100PossInput) (*TeamOpponentsPer100PossResponse, error) {
	abbr := input.Abbr
	seasonFrom := int32(input.SeasonFrom)
	seasonTo := int32(input.SeasonTo)

	per100poss, err := database.Queries.GetOpponentsPer100Possesions(ctx,
		sqlc.GetOpponentsPer100PossesionsParams{TeamAbbr: abbr, SeasonYear: seasonFrom, SeasonYear_2: seasonTo},
	)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &TeamOpponentsPer100PossResponse{Body: per100poss}, nil
}
