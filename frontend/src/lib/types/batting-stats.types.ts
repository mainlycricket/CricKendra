export type ICombinedBattingStatsType =
  | IOverall_Batting_Batter_Group
  | IOverall_Batting_TeamInnings_Group
  | IOverall_Batting_Match_Group
  | IOverall_Batting_Team_Group
  | IOverall_Batting_Opposition_Group
  | IOverall_Batting_Ground_Group
  | IOverall_Batting_HostNation_Group
  | IOverall_Batting_Continent_Group
  | IOverall_Batting_Series_Group
  | IOverall_Batting_Tournament_Group
  | IOverall_Batting_Year_Group
  | IOverall_Batting_Season_Group
  | IOverall_Batting_Decade_Group
  | IOverall_Batting_Aggregate_Group
  | IOverall_Batting_Summary_HomeAway_Group
  | IOverall_Batting_Summary_TossResult_Group
  | IOverall_Batting_Summary_TossDecision_Group
  | IOverall_Batting_Summary_BatBowlFirst_Group
  | IOverall_Batting_Summary_InningsNumber_Group
  | IOverall_Batting_Summary_MatchResult_Group
  | IOverall_Batting_Summary_MatchResultBatBowlFirst_Group
  | IOverall_Batting_Summary_SeriesTeamsCount_Group
  | IOverall_Batting_Summary_SeriesMatchNumber_Group
  | IOverall_Batting_Summary_BattingPosition_Group
  | IIndividual_Batting_Series_Group
  | IIndividual_Batting_Tournaments_Group
  | IIndividual_Batting_Ground_Group
  | IIndividual_Batting_HostNation_Group
  | IIndividual_Batting_Opposition_Group
  | IIndividual_Batting_Year_Group
  | IIndividual_Batting_Season_Group;

export type ICombinedBattingStatsType2 =
  | IIndividual_Batting_Innings_Group
  | IIndividual_Batting_MatchTotals_Group;

export interface IOverall_Batting_Summary_Group {
  teams?: IOverall_Batting_Team_Group[];
  oppositions?: IOverall_Batting_Opposition_Group[];
  host_nations?: IOverall_Batting_HostNation_Group[];
  continents?: IOverall_Batting_Continent_Group[];
  years?: IOverall_Batting_Year_Group[];
  seasons?: IOverall_Batting_Season_Group[];
  home_away?: IOverall_Batting_Summary_HomeAway_Group[];
  toss_won_lost?: IOverall_Batting_Summary_TossResult_Group[];
  toss_decision?: IOverall_Batting_Summary_TossDecision_Group[];
  bat_bowl_first?: IOverall_Batting_Summary_BatBowlFirst_Group[];
  innings_number?: IOverall_Batting_Summary_InningsNumber_Group[];
  match_result?: IOverall_Batting_Summary_MatchResult_Group[];
  match_result_bat_bowl_first?: IOverall_Batting_Summary_MatchResultBatBowlFirst_Group[];
  series_teams_count?: IOverall_Batting_Summary_SeriesTeamsCount_Group[];
  series_event_match_number?: IOverall_Batting_Summary_SeriesMatchNumber_Group[];
  tournaments?: IOverall_Batting_Tournament_Group[];
  batting_positions?: IOverall_Batting_Summary_BattingPosition_Group[];
}

export function isT<T extends IOverallBattingStats>(obj: IOverallBattingStats, fields: string[]): obj is T {
  return fields.every((field) => field in obj);
}

export function isT2<T extends ICombinedBattingStatsType2>(
  obj: ICombinedBattingStatsType2,
  fields: string[]
): obj is T {
  return fields.every((field) => field in obj);
}

/* Overall Stats */

export interface IOverall_Batting_Batter_Group extends IOverallBattingStats {
  batter_id: number;
  batter_name: string;
  teams_represented: string[];
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_TeamInnings_Group extends IOverallBattingStats {
  match_id: number;
  innings_number: number;
  batting_team_id: number;
  batting_team_name: string;
  bowling_team_id: number;
  bowling_team_name: string;
  season: string;
  city_name: string;
  start_date: string;
  players_count: number;
}

export interface IOverall_Batting_Match_Group extends IOverallBattingStats {
  match_id: number;
  team1_id: number;
  team1_name: string;
  team2_id: number;
  team2_name: string;
  season: string;
  city_name: string;
  start_date: string;
  players_count: number;
}

export interface IOverall_Batting_Team_Group extends IOverallBattingStats {
  team_id: number;
  team_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Opposition_Group extends IOverallBattingStats {
  opposition_team_id: number;
  opposition_team_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Ground_Group extends IOverallBattingStats {
  ground_id: number;
  ground_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_HostNation_Group extends IOverallBattingStats {
  host_nation_id: number;
  host_nation_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Continent_Group extends IOverallBattingStats {
  continent_id: number;
  continent_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Series_Group extends IOverallBattingStats {
  series_id: number;
  series_name: string;
  series_season: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Tournament_Group extends IOverallBattingStats {
  tournament_id: number;
  tournament_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Year_Group extends IOverallBattingStats {
  year: number;
  players_count: number;
}

export interface IOverall_Batting_Season_Group extends IOverallBattingStats {
  season: string;
  players_count: number;
}

export interface IOverall_Batting_Decade_Group extends IOverallBattingStats {
  decade: number;
  players_count: number;
}

export interface IOverall_Batting_Aggregate_Group extends IOverallBattingStats {
  players_count: number;
  min_date: string;
  max_date: string;
}

/* Overall Summary Stats */

export interface IOverall_Batting_Summary_HomeAway_Group extends IOverallBattingStats {
  home_away_label: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_TossResult_Group extends IOverallBattingStats {
  toss_result: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_TossDecision_Group extends IOverallBattingStats {
  toss_result: string;
  is_toss_decision_bat: boolean;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_BatBowlFirst_Group extends IOverallBattingStats {
  bat_bowl_first: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_InningsNumber_Group extends IOverallBattingStats {
  innings_number: number;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_MatchResult_Group extends IOverallBattingStats {
  match_result: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_MatchResultBatBowlFirst_Group extends IOverallBattingStats {
  match_result: string;
  bat_bowl_first: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_SeriesTeamsCount_Group extends IOverallBattingStats {
  teams_count: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_SeriesMatchNumber_Group extends IOverallBattingStats {
  event_match_number: number;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Batting_Summary_BattingPosition_Group extends IOverallBattingStats {
  batting_position: number;
  players_count: number;
  min_date: string;
  max_date: string;
}

/* Individual Stats */

export interface IIndividual_Batting_Innings_Group {
  match_id: number;
  start_date: string;
  ground_id: number;
  city_name: string;

  innings_number: number;
  batter_id: number;
  batter_name: string;
  batting_team_id: number;
  batting_team_name: string;
  bowling_team_id: number;
  bowling_team_name: string;

  runs_scored: number;
  balls_faced: number;
  is_not_out: boolean;
  strike_rate: number;
  fours_scored: number;
  sixes_scored: number;
}

export interface IIndividual_Batting_MatchTotals_Group {
  match_id: number;
  start_date: string;
  ground_id: number;
  city_name: string;

  batter_id: number;
  batter_name: string;
  batting_team_id: number;
  batting_team_name: string;
  bowling_team_id: number;
  bowling_team_name: string;

  innings: {
    runs_scored: number;
    is_not_out: boolean;
  }[];

  runs_scored: number;
  balls_faced: number;
  strike_rate: number;
  fours_scored: number;
  sixes_scored: number;
}

export interface IIndividual_Batting_Series_Group extends IOverall_Batting_Batter_Group {
  series_id: number;
  series_name: string;
  series_season: string;
}

export interface IIndividual_Batting_Tournaments_Group extends IOverall_Batting_Batter_Group {
  tournament_id: number;
  tournament_name: string;
}

export interface IIndividual_Batting_Ground_Group extends IOverall_Batting_Batter_Group {
  ground_id: number;
  ground_name: string;
}

export interface IIndividual_Batting_HostNation_Group extends IOverall_Batting_Batter_Group {
  host_nation_id: number;
  host_nation_name: string;
}

export interface IIndividual_Batting_Opposition_Group extends IOverall_Batting_Batter_Group {
  opposition_team_id: number;
  opposition_team_name: string;
}

export interface IIndividual_Batting_Year_Group extends IOverallBattingStats {
  year: number;
  batter_id: number;
  batter_name: string;
  teams_represented: string[];
}

export interface IIndividual_Batting_Season_Group extends IOverallBattingStats {
  season: string;
  batter_id: number;
  batter_name: string;
  teams_represented: string[];
}

/* Extended by Others */

export interface IOverallBattingStats {
  matches_played: number;
  innings_batted: number;
  runs_scored: number;
  balls_faced: number;
  not_outs: number;
  average: number;
  strike_rate: number;
  highest_score: number;
  highest_not_out_score: number;
  centuries: number;
  half_centuries: number;
  fifty_plus_scores: number;
  ducks: number;
  fours_scored: number;
  sixes_scored: number;
}
