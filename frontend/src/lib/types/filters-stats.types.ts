import {
  EnumBatField,
  EnumDismissalType,
  EnumHomeAway,
  EnumInningsNumber,
  EnumMatchResult,
  EnumPlayingFormat,
  EnumStatsType,
  EnumStatsView,
  EnumTossResult,
} from "./enums.types";

export interface IStatsFilters {
  primary_teams?: ITeamAsForeignField[];
  opposition_teams?: ITeamAsForeignField[];
  host_nations?: IHostNationAsForeignField[];
  continents?: IContinentAsForeignField[];
  grounds?: IGroundAsForeignField[];
  min_date?: string;
  max_date?: string;
  seasons?: string[];
  series?: ISeriesAsForeignField[];
  tournaments?: ITournamentAsForeignField[];
}

export interface ITeamAsForeignField {
  id: number;
  name: string;
}

export interface IHostNationAsForeignField {
  id: number;
  name: string;
}

export interface IContinentAsForeignField {
  id: number;
  name: string;
}

export interface IGroundAsForeignField {
  id: number;
  name: string;
  city_name: string;
  host_nation_name: string;
}

export interface ISeriesAsForeignField {
  id: number;
  name: string;
  season: string;
}

export interface ITournamentAsForeignField {
  id: number;
  name: string;
}

export interface IStatsFiltersData {
  [key: string]: string | string[] | undefined;

  statsType: EnumStatsType;
  view: EnumStatsView;
  group: string;

  playing_format?: EnumPlayingFormat;
  is_male?: "true" | "false";
  min_start_date?: string; // YYYY-MM-DD
  max_start_date?: string; // YYYY-MM-DD

  season?: string[];
  primary_team?: string[];
  opposition_team?: string[];
  continent?: string[];
  host_nation?: string[];
  ground?: string[];
  series?: string[];
  tournament?: string[];

  home_or_away?: EnumHomeAway[];
  match_result?: EnumMatchResult[];
  toss_result?: EnumTossResult;
  bat_field_first?: EnumBatField;
  innings_number?: EnumInningsNumber[];

  /* batting specific fitlers only */
  min__innings_runs_scored?: string;
  max__innings_runs_scored?: string;
  min__innings_batting_position?: string;
  max__innings_batting_position?: string;
  innings_is_batter_dismissed?: "dismissed" | "not_out";
  innings_batter_dismissal_type?: EnumDismissalType[];

  /* bowling specific fitlers only */
  min__innings_balls_bowled?: string;
  max__innings_balls_bowled?: string;
  min__innings_runs_conceded?: string;
  max__innings_runs_conceded?: string;
  min__innings_wickets_taken?: string;
  max__innings_wickets_taken?: string;
  min__innings_bowling_position?: string;
  max__innings_bowling_position?: string;

  /* common qualification filters */
  min__matches_played?: string;
  max__matches_played?: string;
  min__average?: string;
  max__average?: string;
  min__strike_rate?: string;
  max__strike_rate?: string;

  /* batting specific qualification fitlers only */
  min__innings_batted?: string;
  max__innings_batted?: string;
  min__not_outs?: string;
  max__not_outs?: string;
  min__runs_scored?: string;
  max__runs_scored?: string;
  min__balls_faced?: string;
  max__balls_faced?: string;
  // average & strike rate covered in common
  min__centuries?: string;
  max__centuries?: string;
  min__half_centuries?: string;
  max__half_centuries?: string;
  min__fifty_plus_scores?: string;
  max__fifty_plus_scores?: string;
  min__ducks?: string;
  max__ducks?: string;
  min__fours_scored?: string;
  max__fours_scored?: string;
  min__sixes_scored?: string;
  max__sixes_scored?: string;

  /* bowling specific qualification fitlers only */
  min__innings_bowled?: string;
  max__innings_bowled?: string;
  min__overs_bowled?: string;
  max__overs_bowled?: string;
  min__maiden_overs?: string;
  max__maiden_overs?: string;
  min__runs_conceded?: string;
  max__runs_conceded?: string;
  min__wickets_taken?: string;
  max__wickets_taken?: string;
  // average & strike rate covered in common
  min__economy?: string;
  max__economy?: string;
  min__fours_conceded?: string;
  max__fours_conceded?: string;
  min__sixes_conceded?: string;
  max__sixes_conceded?: string;
  min__four_wkt_hauls?: string;
  max__four_wkt_hauls?: string;
  min__five_wkt_hauls?: string;
  max__five_wkt_hauls?: string;
  min__ten_wkt_hauls?: string;
  max__ten_wkt_hauls?: string;

  sort_by?: string;
  sort_order?: string;
  __page?: string;
  __limit?: string;
}
