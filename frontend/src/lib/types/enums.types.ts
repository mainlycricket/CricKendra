export type EnumPlayingFormat = "Test" | "ODI" | "T20I" | "first_class" | "list_a" | "T20";
export function isEnumPlayingFormat(value: string): value is EnumPlayingFormat {
  return ["Test", "ODI", "T20I", "first_class", "list_a", "T20"].includes(value);
}

export type EnumStatsType = "batting" | "bowling" | "team";
export function isEnumStatsType(value: string): value is EnumStatsType {
  return ["batting", "bowling", "team"].includes(value);
}

export type EnumStatsView = "overall" | "individual";
export function isEnumStatsView(value: string): value is EnumStatsView {
  return value === "overall" || value === "individual";
}

export type EnumHomeAway = "home" | "away" | "neutral";
export function isEnumHomeAway(value: string): value is EnumHomeAway {
  return ["home", "away", "neutral"].includes(value);
}

export type EnumMatchResult = "won" | "lost" | "drawn" | "tied" | "no result";
export function isEnumMatchResult(value: string): value is EnumMatchResult {
  return ["won", "lost", "drawn", "tied", "no result"].includes(value);
}

export type EnumTossResult = "won" | "lost";
export function isEnumTossResult(value: string): value is EnumStatsView {
  return value === "won" || value === "lost";
}

export type EnumBatField = "bat" | "field";
export function isEnumBatField(value: string): value is EnumStatsView {
  return value === "bat" || value === "field";
}

export type EnumInningsNumber = "1" | "2" | "3" | "4";
export function isEnumInningsNumber(value: string): value is EnumInningsNumber {
  return ["1", "2", "3", "4"].includes(value);
}

export type EnumDismissalType =
  | "caught"
  | "bowled"
  | "lbw"
  | "run out"
  | "stumped"
  | "hit wicket"
  | "handled the ball"
  | "obstructing the field"
  | "timed out"
  | "retired hurt"
  | "hit the ball twice"
  | "caught and bowled"
  | "retired out"
  | "retired not out";

export function isEnumDismissalType(value: string): value is EnumDismissalType {
  const validValues = [
    "caught",
    "bowled",
    "lbw",
    "run out",
    "stumped",
    "hit wicket",
    "handled the ball",
    "obstructing the field",
    "timed out",
    "retired hurt",
    "hit the ball twice",
    "caught and bowled",
    "retired out",
    "retired not out",
  ];

  return validValues.includes(value);
}

export type EnumBattingQualificationKey =
  | "innings_runs_scored"
  | "innings_batting_position"
  | "matches_played"
  | "innings_batted"
  | "not_outs"
  | "runs_scored"
  | "balls_faced"
  | "average"
  | "strike_rate"
  | "centuries"
  | "half_centuries"
  | "fifty_plus_scores"
  | "ducks"
  | "fours_scored"
  | "sixes_scored";

export function isEnumBattingQualificationKey(value: string): value is EnumBattingQualificationKey {
  if (!value.startsWith("min__") && !value.startsWith("max__")) return false;

  const key = value.slice(5);
  return [
    "innings_runs_scored",
    "innings_batting_position",
    "matches_played",
    "innings_batted",
    "not_outs",
    "runs_scored",
    "balls_faced",
    "average",
    "strike_rate",
    "centuries",
    "half_centuries",
    "fifty_plus_scores",
    "ducks",
    "fours_scored",
    "sixes_scored",
  ].includes(key);
}

export type EnumBowlingQualificationKey =
  | "innings_balls_bowled"
  | "innings_runs_conceded"
  | "innings_wickets_taken"
  | "innings_bowling_position"
  | "matches_played"
  | "innings_bowled"
  | "overs_bowled"
  | "maiden_overs"
  | "runs_conceded"
  | "wickets_taken"
  | "average"
  | "strike_rate"
  | "economy"
  | "fours_conceded"
  | "sixes_conceded"
  | "four_wkt_hauls"
  | "five_wkt_hauls"
  | "ten_wkt_hauls";

export function isEnumBowlingQualificationKey(value: string): value is EnumBowlingQualificationKey {
  if (!value.startsWith("min__") && !value.startsWith("max__")) return false;

  const key = value.slice(5);
  return [
    "innings_balls_bowled",
    "innings_runs_conceded",
    "innings_wickets_taken",
    "innings_bowling_position",
    "matches_played",
    "innings_bowled",
    "overs_bowled",
    "maiden_overs",
    "runs_conceded",
    "wickets_taken",
    "average",
    "strike_rate",
    "fours_conceded",
    "sixes_conceded",
    "four_wkt_hauls",
    "five_wkt_hauls",
    "ten_wkt_hauls",
  ].includes(key);
}

export type EnumTeamQualificationKey =
  | "team_innings_runs"
  | "team_innings_wickets"
  | "team_innings_balls"
  | "matches_played"
  | "matches_won"
  | "matches_lost"
  | "matches_tied"
  | "matches_drawn"
  | "matches_with_no_result"
  | "average"
  | "scoring_rate"
  | "win_loss_ratio"
  | "innings_count"
  | "total_runs"
  | "total_balls"
  | "total_wickets";

export function isEnumTeamQualificationKey(value: string): value is EnumBowlingQualificationKey {
  if (!value.startsWith("min__") && !value.startsWith("max__")) return false;

  const key = value.slice(5);
  return [
    "team_innings_runs",
    "team_innings_wickets",
    "team_innings_balls",
    "matches_played",
    "matches_won",
    "matches_lost",
    "matches_tied",
    "matches_drawn",
    "matches_with_no_result",
    "average",
    "scoring_rate",
    "win_loss_ratio",
    "innings_count",
    "total_runs",
    "total_balls",
    "total_wickets",
  ].includes(key);
}
