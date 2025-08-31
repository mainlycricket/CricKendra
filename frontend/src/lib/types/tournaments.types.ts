import { INextData } from "./shared.types";

export interface IAllTournaments {
  id: number;
  name: string;
  is_male: boolean;
  playing_level: string;
  playing_format: string;
}

export interface IAllTournamentsResponse extends INextData {
  tournaments: IAllTournaments[];
}
