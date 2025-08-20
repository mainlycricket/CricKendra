export type ICombinedTeamStatsType =
  | IOverall_Team_Players_Group
  | IOverall_Team_Matches_Group
  | IOverall_Team_Teams_Group
  | IOverall_Team_Grounds_Group
  | IOverall_Team_HostNations_Group
  | IOverall_Team_Continents_Group
  | IOverall_Team_Years_Group
  | IOverall_Team_Seasons_Group
  | IOverall_Team_Series_Group
  | IOverall_Team_Tournament_Group
  | IOverall_Team_Decades_Group
  | IOverall_Team_Aggregate_Group
  | IIndividual_Team_Series_Group
  | IIndividual_Team_Tournaments_Group
  | IIndividual_Team_Grounds_Group
  | IIndividual_Team_HostNations_Group
  | IIndividual_Team_Years_Group
  | IIndividual_Team_Seasons_Group;

export function isT<T extends IOverallTeamStats>(obj: IOverallTeamStats, fields: string[]): obj is T {
  return fields.every((field) => field in obj);
}

export type ICombinedTeamStatsType2 =
  | IIndividual_Team_Innings_Group
  | IIndividual_Team_MatchTotals_Group
  | IIndividual_Team_MatchResults_Group;

export function isT2<T extends IIndividualMatchInfo>(obj: IIndividualMatchInfo, fields: string[]): obj is T {
  return fields.every((field) => field in obj);
}

/* Overall Stats */

export interface IOverall_Team_Teams_Group extends IOverallTeamStats {
  team_id: number;
  team_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
}

export interface IOverall_Team_Players_Group extends IOverallTeamStats {
  player_id: number;
  player_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
  teams_count: number;
}

export interface IOverall_Team_Matches_Group extends IOverallTeamStats {
  match_id: number;
  team1_id: number;
  team1_name: string;
  team2_id: number;
  team2_name: string;
  ground_id: number;
  city_name: string;
  season: string;
  start_date: string; // YYYY-MM-DD
  teams_count: number;
}

export interface IOverall_Team_Series_Group extends IOverallTeamStats {
  series_id: number;
  series_name: string;
  series_season: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
  teams_count: number;
}

export interface IOverall_Team_Tournament_Group extends IOverallTeamStats {
  tournament_id: number;
  tournament_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
  teams_count: number;
}

export interface IOverall_Team_Grounds_Group extends IOverallTeamStats {
  ground_id: number;
  ground_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
  teams_count: number;
}

export interface IOverall_Team_HostNations_Group extends IOverallTeamStats {
  host_nation_id: number;
  host_nation_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
  teams_count: number;
}

export interface IOverall_Team_Continents_Group extends IOverallTeamStats {
  continent_id: number;
  continent_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
  teams_count: number;
}

export interface IOverall_Team_Years_Group extends IOverallTeamStats {
  year: number;
  teams_count: number;
}

export interface IOverall_Team_Seasons_Group extends IOverallTeamStats {
  season: string;
  teams_count: number;
}

export interface IOverall_Team_Decades_Group extends IOverallTeamStats {
  decade: number;
  teams_count: number;
}

export interface IOverall_Team_Aggregate_Group extends IOverallTeamStats {
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
  teams_count: number;
}

/* Individual Stats */

export interface IIndividual_Team_Innings_Group extends IIndividualMatchInfo {
  innings_id: number;
  innings_number: number;
  innings_end: string;
  total_runs: number;
  total_wickets: number;
  total_overs: number;
  scoring_rate: number;
}

export interface IIndividual_Team_MatchTotals_Group extends IIndividualMatchInfo {
  total_runs: number;
  total_balls: number;
  total_wickets: number;
  average: number;
  scoring_rate: number;
}

export interface IIndividual_Team_MatchResults_Group extends IIndividualMatchInfo {
  toss_winner_id: number;
  innings_number: number;
  win_margin: number;
  balls_remaining_after_win: number;
  is_won_by_runs: boolean;
  is_won_by_innings: boolean;
}

export interface IIndividual_Team_Series_Group extends IOverallTeamStats {
  team_id: number;
  team_name: string;
  series_id: number;
  series_name: string;
  series_season: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
}

export interface IIndividual_Team_Tournaments_Group extends IOverallTeamStats {
  team_id: number;
  team_name: string;
  tournament_id: number;
  tournament_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
}

export interface IIndividual_Team_Grounds_Group extends IOverallTeamStats {
  team_id: number;
  team_name: string;
  ground_id: number;
  ground_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
}

export interface IIndividual_Team_HostNations_Group extends IOverallTeamStats {
  team_id: number;
  team_name: string;
  host_nation_id: number;
  host_nation_name: string;
  min_date: string; // YYYY-MM-DD
  max_date: string; // YYYY-MM-DD
}

export interface IIndividual_Team_Years_Group extends IOverallTeamStats {
  team_id: number;
  team_name: string;
  year: number;
}

export interface IIndividual_Team_Seasons_Group extends IOverallTeamStats {
  team_id: number;
  team_name: string;
  season: string;
}

// Embedded in other structs
export interface IOverallTeamStats {
  matches_played: number;
  matches_won: number;
  matches_lost: number;
  win_loss_ratio: number;
  matches_drawn: number;
  matches_tied: number;
  matches_no_result: number;

  innings_count: number;
  total_runs: number;
  total_balls: number;
  total_wickets: number;
  average: number;
  scoring_rate: number;
  highest_score: number;
  lowest_score: number;
}

export interface IIndividualMatchInfo {
  match_id: number;
  team_id: number;
  team_name: string;
  opposition_id: number;
  opposition_name: string;
  ground_id: number;
  city_name: string;
  start_date: string; // YYYY-MM-DD
  final_result: string;
  match_winner_id: number;
}
