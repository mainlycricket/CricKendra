import { IMatchInfo } from "./single-match.types";

export interface ISeriesHeader {
  series_id: number;
  series_name: string;
  season: string;
  tournament_id: number;
  tournament_name: string;

  top_batters?: ISeriesTopBatter[];
  top_bowlers?: ISeriesTopBowler[];
}

export interface ISeriesTopBatter {
  batter_id: number;
  batter_name: string;
  batter_image_url: string;

  innings_batted: number;
  runs_scored: string;
  average: number;
}

export interface ISeriesTopBowler {
  bowler_id: number;
  bowler_name: string;
  bowler_image_url: string;

  innings_bowled: number;
  wickets_taken: string;
  average: number;
}

export interface ISingleSeriesOverview {
  series_header: ISeriesHeader;

  winner_team_id: number;
  winner_team_name: string;
  final_status: string;

  fixture_matches?: IMatchInfo[];
  result_matches?: IMatchInfo[];
}

export interface ISingleSeriesMatches {
  series_header: ISeriesHeader;
  matches?: IMatchInfo[];
}

export interface ISingleSeriesTeams {
  series_header: ISeriesHeader;
  teams?: ISeriesTeam[];
}

export interface ISeriesTeam {
  team_id: number;
  team_name: string;
  team_image_url: string;
}

export interface ISingleSeriesSquadsList {
  series_header: ISeriesHeader;
  squad_list?: ISingleSeriesSquadEntry[];
}

export interface ISingleSeriesSquadEntry {
  squad_id: number;
  squad_label: string;
  team_image_url?: string;
}

export interface ISingleSeriesSingleSquad {
  series_header: ISeriesHeader;
  squad_list?: ISingleSeriesSquadEntry[];
  players?: ISingleSeriesSquadPlayer[];
}

export interface ISingleSeriesSquadPlayer {
  player_id: number;
  player_name: string;
  playing_role: string;
  date_of_birth: Date;
  is_rhb: boolean;
  primary_bowling_style: string;
  is_captain: boolean;
  is_vice_captain: boolean;
  is_wk: boolean;
}
