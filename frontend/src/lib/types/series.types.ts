import { EnumPlayingFormat, EnumPlayingLevel } from "./enums.types";
import { ITeamAsForeignField } from "./filters-stats.types";
import { INextData } from "./shared.types";

export interface ISeries {
  id: number;
  name: string;
  is_male: boolean;
  playing_level: EnumPlayingLevel;
  playing_format: EnumPlayingFormat;
  season: string;
  teams_id: number[];
  tournament_id: number;
  tour_flag: string;
  start_date: string; // YYYY-MM-DDD
  end_date: string; // YYYY-MM-DDD
  winner_team_id: number;
  final_status: string;
}

export interface IAllSeries {
  id: number;
  name: string;
  is_male: boolean;
  playing_level: string;
  playing_format: string;
  season: string;
  teams: ITeamAsForeignField[];
  start_date: string;
  end_date: string;
  winner_team_id: number;
  final_status: string;
  tour_flag: string;
}

export interface IAllSeriesResponse extends INextData {
  series: IAllSeries[];
}
