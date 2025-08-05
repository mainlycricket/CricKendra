export type ICombinedBowlingStatsType =
  | IOverall_Bowling_Bowler_Group
  | IOverall_Bowling_TeamInnings_Group
  | IOverall_Bowling_Match_Group
  | IOverall_Bowling_Team_Group
  | IOverall_Bowling_Opposition_Group
  | IOverall_Bowling_Ground_Group
  | IOverall_Bowling_HostNation_Group
  | IOverall_Bowling_Continent_Group
  | IOverall_Bowling_Series_Group
  | IOverall_Bowling_Tournament_Group
  | IOverall_Bowling_Year_Group
  | IOverall_Bowling_Season_Group
  | IOverall_Bowling_Decade_Group
  | IOverall_Bowling_Aggregate_Group
  | IOverall_Bowling_Summary_HomeAway_Group
  | IOverall_Bowling_Summary_TossResult_Group
  | IOverall_Bowling_Summary_TossDecision_Group
  | IOverall_Bowling_Summary_BatBowlFirst_Group
  | IOverall_Bowling_Summary_InningsNumber_Group
  | IOverall_Bowling_Summary_MatchResult_Group
  | IOverall_Bowling_Summary_MatchResultBatBowlFirst_Group
  | IOverall_Bowling_Summary_SeriesTeamsCount_Group
  | IOverall_Bowling_Summary_SeriesMatchNumber_Group
  | IOverall_Bowling_Summary_BowlingPosition_Group
  | IIndividual_Bowling_Series_Group
  | IIndividual_Bowling_Tournament_Group
  | IIndividual_Bowling_Ground_Group
  | IIndividual_Bowling_HostNation_Group
  | IIndividual_Bowling_Opposition_Group
  | IIndividual_Bowling_Year_Group
  | IIndividual_Bowling_Season_Group;

export type ICombinedBowlingStatsType2 =
  | IIndividual_Bowling_Innings_Group
  | IIndividual_Bowling_MatchTotals_Group;

export interface IOverall_Bowling_Summary_Group {
  teams?: IOverall_Bowling_Team_Group[];
  oppositions?: IOverall_Bowling_Opposition_Group[];
  host_nations?: IOverall_Bowling_HostNation_Group[];
  continents?: IOverall_Bowling_Continent_Group[];
  years?: IOverall_Bowling_Year_Group[];
  seasons?: IOverall_Bowling_Season_Group[];
  home_away?: IOverall_Bowling_Summary_HomeAway_Group[];
  toss_won_lost?: IOverall_Bowling_Summary_TossResult_Group[];
  toss_decision?: IOverall_Bowling_Summary_TossDecision_Group[];
  bat_bowl_first?: IOverall_Bowling_Summary_BatBowlFirst_Group[];
  innings_number?: IOverall_Bowling_Summary_InningsNumber_Group[];
  match_result?: IOverall_Bowling_Summary_MatchResult_Group[];
  match_result_bat_bowl_first?: IOverall_Bowling_Summary_MatchResultBatBowlFirst_Group[];
  series_teams_count?: IOverall_Bowling_Summary_SeriesTeamsCount_Group[];
  series_event_match_number?: IOverall_Bowling_Summary_SeriesMatchNumber_Group[];
  tournaments?: IOverall_Bowling_Tournament_Group[];
  bowling_positions?: IOverall_Bowling_Summary_BowlingPosition_Group[];
}

export function isT<T extends IOverallBowlingStats>(obj: IOverallBowlingStats, fields: string[]): obj is T {
  return fields.every((field) => field in obj);
}

export function isT2<T extends ICombinedBowlingStatsType2>(obj: ICombinedBowlingStatsType2, fields: string[]): obj is T {
  return fields.every((field) => field in obj);
}

/* Overall Stats */

export interface IOverall_Bowling_Bowler_Group extends IOverallBowlingStats {
  bowler_id: number;
  bowler_name: string;
  teams_represented: string[];
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_TeamInnings_Group extends IOverallBowlingStats {
  match_id: number;
  innings_number: number;
  bowling_team_id: number;
  bowling_team_name: string;
  batting_team_id: number;
  batting_team_name: string;
  season: string;
  city_name: string;
  start_date: string;
  players_count: number;
}

export interface IOverall_Bowling_Match_Group extends IOverallBowlingStats {
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

export interface IOverall_Bowling_Team_Group extends IOverallBowlingStats {
  team_id: number;
  team_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Opposition_Group extends IOverallBowlingStats {
  opposition_team_id: number;
  opposition_team_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Ground_Group extends IOverallBowlingStats {
  ground_id: number;
  ground_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_HostNation_Group extends IOverallBowlingStats {
  host_nation_id: number;
  host_nation_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Continent_Group extends IOverallBowlingStats {
  continent_id: number;
  continent_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Series_Group extends IOverallBowlingStats {
  series_id: number;
  series_name: string;
  series_season: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Tournament_Group extends IOverallBowlingStats {
  tournament_id: number;
  tournament_name: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Year_Group extends IOverallBowlingStats {
  year: number;
  players_count: number;
}

export interface IOverall_Bowling_Season_Group extends IOverallBowlingStats {
  season: string;
  players_count: number;
}

export interface IOverall_Bowling_Decade_Group extends IOverallBowlingStats {
  decade: number;
  players_count: number;
}

export interface IOverall_Bowling_Aggregate_Group extends IOverallBowlingStats {
  players_count: number;
  min_date: string;
  max_date: string;
}

/* Overall Summary Stats */

export interface IOverall_Bowling_Summary_HomeAway_Group extends IOverallBowlingStats {
  home_away_label: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_TossResult_Group extends IOverallBowlingStats {
  toss_result: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_TossDecision_Group extends IOverallBowlingStats {
  toss_result: string;
  is_toss_decision_bat: boolean;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_BatBowlFirst_Group extends IOverallBowlingStats {
  bat_bowl_first: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_InningsNumber_Group extends IOverallBowlingStats {
  innings_number: number;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_MatchResult_Group extends IOverallBowlingStats {
  match_result: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_MatchResultBatBowlFirst_Group extends IOverallBowlingStats {
  match_result: string;
  bat_bowl_first: string;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_SeriesTeamsCount_Group extends IOverallBowlingStats {
  teams_count: number;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_SeriesMatchNumber_Group extends IOverallBowlingStats {
  event_match_number: number;
  players_count: number;
  min_date: string;
  max_date: string;
}

export interface IOverall_Bowling_Summary_BowlingPosition_Group extends IOverallBowlingStats {
  bowling_position: number;
  players_count: number;
  min_date: string;
  max_date: string;
}

/* Individual Stats */

export interface IIndividual_Bowling_Innings_Group {
  match_id: number;
  start_date: string;
  ground_id: number;
  city_name: string;

  innings_number: number;
  bowler_id: number;
  bowler_name: string;
  batting_team_id: number;
  batting_team_name: string;
  bowling_team_id: number;
  bowling_team_name: string;

  overs_bowled: number;
  maiden_overs: number;
  runs_conceded: number;
  wickets_taken: number;
  economy: number;
  fours_conceded: number;
  sixes_conceded: number;
}

export interface IIndividual_Bowling_MatchTotals_Group {
  match_id: number;
  start_date: string;
  ground_id: number;
  city_name: string;

  bowler_id: number;
  bowler_name: string;
  batting_team_id: number;
  batting_team_name: string;
  bowling_team_id: number;
  bowling_team_name: string;

  overs_bowled: number;
  maiden_overs: number;
  runs_conceded: number;
  wickets_taken: number;
  average: number;
  economy: number;
  strike_rate: number;
  fours_conceded: number;
  sixes_conceded: number;
}

export interface IIndividual_Bowling_Ground_Group extends IOverall_Bowling_Bowler_Group {
  ground_id: number;
  ground_name: string;
}

export interface IIndividual_Bowling_Series_Group extends IOverall_Bowling_Bowler_Group {
  series_id: number;
  series_name: string;
  series_season: string;
}

export interface IIndividual_Bowling_Tournament_Group extends IOverall_Bowling_Bowler_Group {
  tournament_id: number;
  tournament_name: string;
}

export interface IIndividual_Bowling_HostNation_Group extends IOverall_Bowling_Bowler_Group {
  host_nation_id: number;
  host_nation_name: string;
}

export interface IIndividual_Bowling_Opposition_Group extends IOverall_Bowling_Bowler_Group {
  opposition_team_id: number;
  opposition_team_name: string;
}

export interface IIndividual_Bowling_Year_Group extends IOverall_Bowling_Bowler_Group {
  year: number;
}

export interface IIndividual_Bowling_Season_Group extends IOverall_Bowling_Bowler_Group {
  season: string;
}

/* Extended By Others */

export interface IOverallBowlingStats {
  matches_played: number;
  innings_bowled: number;
  overs_bowled: number;
  maiden_overs: number;
  runs_conceded: number;
  wickets_taken: number;
  average: number;
  strike_rate: number;
  economy: number;
  four_wicket_hauls: number;
  five_wicket_hauls: number;
  ten_wicket_hauls: number;
  best_match_wickets: number;
  best_match_runs: number;
  best_innings_wickets: number;
  best_innings_runs: number;
  fours_conceded: number;
  sixes_conceded: number;
}
