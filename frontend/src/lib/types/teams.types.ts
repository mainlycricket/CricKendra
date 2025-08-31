import { INextData } from "./shared.types";

export interface IAllTeams {
  id: number;
  name: string;
  is_male: boolean;
  image_url?: string;
  playing_level: string;
  short_name: string;
}

export interface IAllTeamsResponse extends INextData {
  teams: IAllTeams[];
}
