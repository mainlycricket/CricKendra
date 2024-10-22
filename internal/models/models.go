package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CareerStats struct {
	MatchesPlayed pgtype.Int8 `json:"matchesPlayed"`
	DebutMatchId  pgtype.Int8 `json:"debutMatchId"`
	LastMatchId   pgtype.Int8 `json:"lastMatchId"`

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

type Team struct {
	Id           pgtype.Int8 `json:"id"`
	Name         pgtype.Text `json:"name"`
	IsMale       pgtype.Bool `json:"isMale"`
	ImageURL     pgtype.Text `json:"imageURL"`
	PlayingLevel pgtype.Text `json:"playingLevel"`
}

type Player struct {
	Id          pgtype.Int8 `json:"id"`
	Name        pgtype.Text `json:"name"`
	PlayingRole pgtype.Text `json:"playingRole"`
	Nationality pgtype.Text `json:"nationality"`
	IsMale      pgtype.Bool `json:"isMale"`
	DateOfBirth pgtype.Date `json:"dateOfBirth"`
	ImageURL    pgtype.Text `json:"imageURL"`
	Biography   pgtype.Text `json:"biography"`

	BattingStyles       []pgtype.Text `json:"battingStyles"`
	PrimaryBattingStyle pgtype.Text   `json:"primaryBattingStyle"`
	BowlingStyles       []pgtype.Text `json:"primaryBowlingStyles"`
	PrimaryBowlingStyle pgtype.Text   `json:"primaryBowlingStyle"`
	TeamsRepresentedId  []pgtype.Int8 `json:"teamsRepresented"`

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

type Continents struct {
	Id   pgtype.Int8 `json:"id"`
	Name pgtype.Text `json:"name"`
}

type HostNation struct {
	Id          pgtype.Int8 `json:"id"`
	Name        pgtype.Text `json:"name"`
	ContinentId pgtype.Int8 `json:"continent"`
}

type Ground struct {
	Id           pgtype.Int8 `json:"id"`
	Name         pgtype.Text `json:"name"`
	HostNationId pgtype.Int8 `json:"hostNationId"`
	CityId       pgtype.Text `json:"cityId"`
}

type Tournament struct {
	Id            pgtype.Int8 `json:"id"`
	Name          pgtype.Text `json:"name"`
	IsMale        pgtype.Bool `json:"isMale"`
	PlayingLevel  pgtype.Text `json:"playingLevel"`
	PlayingFormat pgtype.Text `json:"playingFormat"`
}

type Series struct {
	Id             pgtype.Int8   `json:"id"`
	Name           pgtype.Text   `json:"name"`
	IsMale         pgtype.Bool   `json:"isMale"`
	PlayingLevel   pgtype.Text   `json:"playingLevel"`
	PlayingFormat  pgtype.Text   `json:"playingFormat"`
	SeasonId       pgtype.Text   `json:"season"`
	TeamsId        []pgtype.Int8 `json:"teams"`
	HostNationsId  []pgtype.Int8 `json:"hostNations"`
	TournamentId   pgtype.Int8   `json:"tournament"`
	ParentSeriesId pgtype.Int8   `json:"parentSeries"`
}

type Squad struct {
	PlayerId      pgtype.Int8 `json:"player"`
	SeriesId      pgtype.Int8 `json:"seriesId"`
	MatchId       pgtype.Int8 `json:"matchId"`
	IsCaptain     pgtype.Bool `json:"isCaptain"`
	IsWk          pgtype.Bool `json:"isWk"`
	IsDebut       pgtype.Bool `json:"isDebut"`
	PlayingStatus pgtype.Text `json:"playingStatus"` // withdrawn, playing XI, bench, substitute
}

type Match struct {
	Id            pgtype.Int8        `json:"id"`
	StartDateTime pgtype.Timestamptz `json:"startDateTime"`
	Team1Id       pgtype.Int8        `json:"team1"`
	Team2Id       pgtype.Int8        `json:"team2"`
	IsMale        pgtype.Bool        `json:"isMale"`
	TournamentId  pgtype.Int8        `json:"tournament"`
	SeriesId      pgtype.Int8        `json:"series"`
	HostNationId  pgtype.Int8        `json:"hostNation"`
	ContinentId   pgtype.Int8        `json:"continent"`
	GroundId      pgtype.Int8        `json:"groundId"`
	CurrentStatus pgtype.Text        `json:"currentStatus"`
	FinalResult   pgtype.Text        `json:"finalResult"` // completed, abandoned, no result
	HomeTeamId    pgtype.Int8        `json:"homeTeam"`
	AwayTeamId    pgtype.Int8        `json:"awayTeam"`
	MatchType     pgtype.Text        `json:"matchType"` // preliminary, semifinal, final
	PlayingLevel  pgtype.Text        `json:"playingLevel"`
	PlayingFormat pgtype.Text        `json:"playingFormat"`
	SeasonId      pgtype.Text        `json:"season"`
	IsDayNight    pgtype.Bool        `json:"isDayNight"`
	IsDLS         pgtype.Bool        `json:"isDLS"`
	TossWinnerId  pgtype.Int8        `json:"tossWinnerTeam"`
	TossLoserId   pgtype.Int8        `json:"tossLoserTeam"`
	TossDecision  pgtype.Text        `json:"tossDecision"`
	MatchWinnerId pgtype.Int8        `json:"matchWinnerTeam"`
	MatchLoserId  pgtype.Int8        `json:"matchLoserTeam"`
	IsWonByRuns   pgtype.Bool        `json:"isWonByRuns"`
	WinMargin     pgtype.Int8        `json:"winMargin"`             // runs or wickets
	BallsMargin   pgtype.Int8        `json:"ballsRemainingAferWin"` // successful chases
	PotmId        pgtype.Int8        `json:"playerOfTheMatch"`

	ScorersId      []pgtype.Int8 `json:"scorers"`
	CommentatorsId []pgtype.Int8 `json:"commentators"`
}

type Innings struct {
	Id            pgtype.Int8 `json:"id"`
	MatchId       pgtype.Int8 `json:"matchId"`
	IsMale        pgtype.Bool `json:"isMale"`
	InningsNumber pgtype.Int8 `json:"inningsNumber"`
	BattingTeamId pgtype.Int8 `json:"battingTeamId"`
	BowlingTeamId pgtype.Int8 `json:"bowlingTeamId"`
	TotalRuns     pgtype.Int8 `json:"totalRuns"`
	TotalWkts     pgtype.Int8 `json:"totalWickets"`
	Byes          pgtype.Int8 `json:"byes"`
	Legbyes       pgtype.Int8 `json:"legbyes"`
	Wides         pgtype.Int8 `json:"wides"`
	Noballs       pgtype.Int8 `json:"noballs"`
	Penalty       pgtype.Int8 `json:"penalty"`
	IsSuperOver   pgtype.Bool `json:"isSuperOver"`
	Status        pgtype.Text `json:"Status"` // live, completed, declared
}

type BattingScorecard struct {
	Id              pgtype.Int8 `json:"id"`
	InningsId       pgtype.Int8 `json:"inningsId"`
	BatterId        pgtype.Int8 `json:"batterId"`
	BattingPosition pgtype.Int8 `json:"battingPosition"`
	RunsScored      pgtype.Int8 `json:"runsScored"`
	BallsFaced      pgtype.Int8 `json:"ballsFaced"`
	MinutesBatted   pgtype.Int8 `json:"minutesBatted"`
	FoursScored     pgtype.Int8 `json:"foursScored"`
	SixesScored     pgtype.Int8 `json:"sixesScored"`
	DismissedById   pgtype.Int8 `json:"dismissedById"`
	WicketType      pgtype.Text `json:"wicketType"`
	DismissalBallId pgtype.Int8 `json:"dismissalBallId"`
	Fielder1Id      pgtype.Int8 `json:"fielder1Id"`
	Fielder2Id      pgtype.Int8 `json:"fielder2Id"`
}

type BowlingScorecard struct {
	Id              pgtype.Int8 `json:"id"`
	InningsId       pgtype.Int8 `json:"inningsId"`
	BowlerId        pgtype.Int8 `json:"bowlerId"`
	BowlingPosition pgtype.Int8 `json:"bowlingPosition"`
	WicketsTaken    pgtype.Int8 `json:"wicketsTaken"`
	RunsConceded    pgtype.Int8 `json:"runsConceded"`
	BallsBowled     pgtype.Int8 `json:"ballsBowled"`
	FoursConceded   pgtype.Int8 `json:"foursConceded"`
	SixesConceded   pgtype.Int8 `json:"sixesConceded"`
	WidesConceded   pgtype.Int8 `json:"widesConceded"`
	NoballsConceded pgtype.Int8 `json:"noballsConceded"`
}

type Deliveries struct {
	Id                   pgtype.Int8        `json:"id"`
	InningsId            pgtype.Int8        `json:"inningsId"`
	BallNumber           pgtype.Float8      `json:"ballNumber"`
	OverNumber           pgtype.Int8        `json:"overNumber"`
	BatterId             pgtype.Int8        `json:"batterId"`
	BowlerId             pgtype.Int8        `json:"bowlerId"`
	NonStrikerId         pgtype.Int8        `json:"nonStrikerId"`
	BattingTeamId        pgtype.Int8        `json:"battingTeamId"`
	BowlingTeamId        pgtype.Int8        `json:"bowlingTeamId"`
	BatterRuns           pgtype.Int8        `json:"batterRuns"`
	Wides                pgtype.Int8        `json:"wides"`
	Noballs              pgtype.Int8        `json:"noballs"`
	Legbyes              pgtype.Int8        `json:"legbyes"`
	Byes                 pgtype.Int8        `json:"byes"`
	Penalty              pgtype.Int8        `json:"penalty"`
	TotalExtras          pgtype.Int8        `json:"totalExtras"`
	TotalRuns            pgtype.Int8        `json:"totalRuns"`
	BowlerRuns           pgtype.Int8        `json:"bowlerRuns"`
	IsFour               pgtype.Bool        `json:"isFour"`
	IsSix                pgtype.Bool        `json:"isSix"`
	Player1DismissedId   pgtype.Int8        `json:"player1DismissedId"`
	Player1DismissalType pgtype.Text        `json:"player1DismissalType"`
	Player2DismissedId   pgtype.Int8        `json:"player2DismissedId"`
	Player2DismissalType pgtype.Text        `json:"player2DismissalType"`
	IsPace               pgtype.Bool        `json:"isPace"`          // true if pacer, false if spin
	BowlingStyle         pgtype.Text        `json:"bowlingStyle"`    // RAFM, LAFM, LAF etc
	IsBatterRHB          pgtype.Bool        `json:"isBatterRHB"`     // true if batter is RHB, false if LHB
	IsNonStrikerRHB      pgtype.Bool        `json:"isNonStrikerRHB"` // true if non-striker is RHB, false if LHB
	Line                 pgtype.Text        `json:"line"`
	Length               pgtype.Text        `json:"length"`
	BallType             pgtype.Text        `json:"ballType"`  // inswinger, googly
	BallSpeed            pgtype.Float8      `json:"ballSpeed"` // mph
	Misc                 pgtype.Text        `json:"misc"`      // edged, missed
	WwRegion             pgtype.Text        `json:"wwRegion"`  // cover, mid-wkt
	FootType             pgtype.Text        `json:"footType"`  // front foot, back foot, step out
	ShotType             pgtype.Text        `json:"shotType"`  // straight drive, pull shot
	Fielder1Id           pgtype.Int8        `json:"fielder1Id"`
	Fielder2Id           pgtype.Int8        `json:"fielder2Id"`
	Commentary           pgtype.Text        `json:"commentary"`
	CreatedAt            pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt            pgtype.Timestamptz `json:"updatedAt"`
}

type BlogArticles struct {
	Id             pgtype.Int8        `json:"id"`
	Title          pgtype.Text        `json:"title"`
	Content        pgtype.Text        `json:"content"`
	AuthorId       pgtype.Int8        `json:"authorId"`
	Category       pgtype.Int8        `json:"category"`
	Status         pgtype.Text        `json:"status"`
	PlayerTags     []pgtype.Int8      `json:"playerTags"`
	TeamTags       []pgtype.Int8      `json:"teamTags"`
	SeriesTags     []pgtype.Int8      `json:"seriesTags"`
	TournamentTags []pgtype.Int8      `json:"tournamentTags"`
	CreatedAt      pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt      pgtype.Timestamptz `json:"updatedAt"`
}

type User struct {
	Id       pgtype.Int8 `json:"id"`
	Name     pgtype.Text `json:"name"`
	Email    pgtype.Text `json:"email"`
	Password pgtype.Text `json:"password"`
	Role     pgtype.Text `json:"role"`
}
