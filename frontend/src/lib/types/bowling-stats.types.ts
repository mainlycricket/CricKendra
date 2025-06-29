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

export interface IOverall_Bowling_Year_Group extends IOverallBowlingStats {
  year: number;
  players_count: number;
}

export interface IOverall_Bowling_Season_Group extends IOverallBowlingStats {
  season: string;
  players_count: number;
}

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

export interface IOverall_Bowling_Tournament_Group extends IOverallBowlingStats {
  tournament_id: number;
  tournament_name: string;
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
