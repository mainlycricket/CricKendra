export interface ISinglePlayer {
  id: number;
  name: string;
  full_name: string;
  playing_role: string;
  nationality: string;
  is_male: boolean;
  date_of_birth: Date;
  image_url: string;
  biography: string;

  is_rhb: boolean;
  bowling_styles?: string[];
  primary_bowling_style: string;
  teams_represented?: { id: number; name: string }[];

  test_stats?: ICareerStats;
  odi_stats?: ICareerStats;
  t20i_stats?: ICareerStats;
  fc_stats?: ICareerStats;
  lista_stats?: ICareerStats;
  t20_stats?: ICareerStats;

  cricsheet_id: string;
  cricinfo_id: string;
  cricbuzz_id: string;
}

export interface ICareerStats {
  matches_played: number;

  innings_batted: number;
  runs_scored: number;
  not_outs: number;
  balls_faced: number;
  fours_scored: number;
  sixes_scored: number;
  centuries_scored: number;
  fifties_scored: number;
  highest_score: number;
  is_highest_not_out: boolean;

  innings_bowled: number;
  runs_conceded: number;
  wickets_taken: number;
  balls_bowled: number;
  fours_conceded: number;
  sixes_conceded: number;
  four_wkt_hauls: number;
  five_wkt_hauls: number;
  ten_wkt_hauls: number;
  best_inn_fig_runs: number;
  best_inn_fig_wkts: number;
  best_match_fig_runs: number;
  best_match_fig_wkts: number;
}
