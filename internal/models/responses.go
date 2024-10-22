package models

import "github.com/jackc/pgx/v5/pgtype"

type TeamAsForeignField struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type MatchAsForeignField struct {
	Id    pgtype.Int8 `json:"id"`
	Team1 pgtype.Text `json:"team1"`
	Team2 pgtype.Text `json:"team2"`
	Date  pgtype.Date `json:"date"`
	City  pgtype.Text `json:"city"`
}

type AllPlayers struct {
	Id          pgtype.Int8 `json:"id"`
	Name        pgtype.Text `json:"name"`
	PlayingRole pgtype.Text `json:"playingRole"`
	Nationality pgtype.Text `json:"nationality"`
	IsMale      pgtype.Bool `json:"isMale"`
	DateOfBirth pgtype.Date `json:"dateOfBirth"`

	PrimaryBattingStyle pgtype.Text `json:"primaryBattingStyle"`
	PrimaryBowlingStyle pgtype.Text `json:"primaryBowlingStyle"`
}

type SinglePlayer struct {
	Id          pgtype.Int8 `json:"id"`
	Name        pgtype.Text `json:"name"`
	PlayingRole pgtype.Text `json:"playingRole"`
	Nationality pgtype.Text `json:"nationality"`
	IsMale      pgtype.Bool `json:"isMale"`
	DateOfBirth pgtype.Date `json:"dateOfBirth"`
	ImageURL    pgtype.Text `json:"imageURL"`
	Biography   pgtype.Text `json:"biography"`

	BattingStyles       []pgtype.Text        `json:"battingStyles"`
	PrimaryBattingStyle pgtype.Text          `json:"primaryBattingStyle"`
	BowlingStyles       []pgtype.Text        `json:"primaryBowlingStyles"`
	PrimaryBowlingStyle pgtype.Text          `json:"primaryBowlingStyle"`
	TeamsRepresented    []TeamAsForeignField `json:"teamsRepresented"`

	TestStats  CareerStats `json:"testStats"`
	OdiStats   CareerStats `json:"odiStats"`
	T20iStats  CareerStats `json:"t20iStats"`
	FcStats    CareerStats `json:"fcStats"`
	ListAStats CareerStats `json:"listAStats"`
	T20Stats   CareerStats `json:"t20Stats"`

	CricsheetId pgtype.Text `json:"cricsheetId"`
	CricinfoId  pgtype.Text `json:"cricinfoId"`
	CricbuzzId  pgtype.Text `json:"cricbuzzId"`
}

type CareerStatsResp struct {
	MatchesPlayed pgtype.Int8         `json:"matchesPlayed"`
	DebutMatchId  MatchAsForeignField `json:"debutMatch"`
	LastMatchId   MatchAsForeignField `json:"lastMatch"`

	InningsBatted     pgtype.Int8 `json:"inningsBatted"`
	RunsScored        pgtype.Int8 `json:"runsScored"`
	BattingDismissals pgtype.Int8 `json:"battingDismissals"`
	BallsFaced        pgtype.Int8 `json:"ballsFaced"`
	FoursScored       pgtype.Int8 `json:"foursScored"`
	SixesScored       pgtype.Int8 `json:"sixesScored"`
	CenturiesScored   pgtype.Int8 `json:"centuriesScored"`
	FiftiesScored     pgtype.Int8 `json:"fiftiesScored"`
	HighestScore      pgtype.Int8 `json:"highestScore"`
	IsHighestNotOut   pgtype.Bool `json:"isHighestScoreNotOut"`

	InningsBowled    pgtype.Int8 `json:"inningsBowled"`
	RunsConceded     pgtype.Int8 `json:"runsConceded"`
	WicketsTaken     pgtype.Int8 `json:"wicketsTaken"`
	BallsBowled      pgtype.Int8 `json:"ballsBowled"`
	FoursConceded    pgtype.Int8 `json:"foursConceded"`
	SixesConceded    pgtype.Int8 `json:"sixesConceded"`
	FourWktHauls     pgtype.Int8 `json:"fourWicketHauls"`
	FiveWktHauls     pgtype.Int8 `json:"fiveWicketHauls"`
	TenWktHauls      pgtype.Int8 `json:"tenWicketHauls"`
	BestInnFigRuns   pgtype.Int8 `json:"bestInningsFiguresRuns"`
	BestInnFigWkts   pgtype.Int8 `json:"bestInningsFiguresWickets"`
	BestMatchFigRuns pgtype.Int8 `json:"bestMatchFiguresRuns"`
	BestMatchFigWkts pgtype.Int8 `json:"bestMatchFiguresWickets"`
}
