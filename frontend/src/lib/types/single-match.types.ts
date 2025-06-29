export interface IMatchInfo {
  match_id: number;
  playing_level: string;
  playing_format: string;
  match_type: string;
  event_match_number: number;
  match_state: "upcoming" | "live" | "break" | "completed";
  match_state_description: string;
  final_result?: "winner decided" | "tie" | "draw" | "no result" | "abandoned";

  match_winner_team_id?: number;
  match_loser_team_id?: number;
  is_won_by_innings?: boolean;
  is_won_by_runs?: boolean;
  win_margin?: number; // runs or wickets
  balls_margin?: number; // successful chases
  super_over_winner_id?: number;
  bowl_out_winner_id?: number;
  outcome_special_method?: "D/L" | "VJD" | "Awarded" | "1st innings score" | "Lost fewer wickets";
  toss_winner_team_id?: number;
  toss_loser_team_id?: number;
  is_toss_decision_bat?: boolean;

  season: string;
  start_date: Date;
  end_date: Date;
  start_date_time_utc: Date;
  is_day_night: boolean;
  ground_id: number;
  ground_name: string;
  main_series_id: number;
  main_series_name: string;

  team1_id: number;
  team1_name: string;
  team1_image_url: string;
  team2_id: number;
  team2_name: string;
  team2_image_url: string;

  innings_scores?: ITeamInningsShortInfo[];
}

export interface IMatchHeader extends IMatchInfo {
  player_awards?: IPlayerAwardInfo[];
}

export interface ITeamInningsShortInfo {
  innings_id: number;
  innings_number: number;
  batting_team_id: number;
  batting_team_name: string;

  total_runs: number;
  total_overs: number;
  total_wickets: number;
  innings_end: string;
  target_runs: number;
  max_overs: number;
}

export interface IPlayerAwardInfo {
  player_id: number;
  player_name: string;
  award_type: string;
}

/* Scorecard Summary */

export interface IMatchSummary {
  match_header: IMatchHeader;
  scorecard_summary?: IInningsScorecardSummary[];
  latest_commentary?: IBbbCommentary[];
}

export interface IInningsScorecardSummary {
  innings_id: number;
  innings_number: number;
  batting_team_id: number;
  batting_team_name: string;

  total_runs: number;
  total_wickets: number;
  total_overs: number;

  top_batters?: IScorecardSummaryBatter[];
  top_bowlers?: IScorecardSummaryBowler[];
}

export interface IScorecardSummaryBatter {
  batter_id: number;
  batter_name: string;
  runs_scored: number;
  balls_faced: number;
  fours_scored: number;
  sixes_scored: number;
}

export interface IScorecardSummaryBowler {
  bowler_id: number;
  bowler_name: string;
  overs_bowled: number;
  maiden_overs: number;
  wickets_taken: number;
  runs_conceded: number;
}

/* Scorecard */

export interface IMatchScorecard {
  match_header: IMatchHeader;
  innings_scorecards?: ITeamInningsScorecard[];
}

export interface IInningsExtrasData {
  byes: number;
  leg_byes: number;
  wides: number;
  noballs: number;
  penalty: number;
}

export interface IInningsTotalData {
  total_runs: number;
  total_overs: number;
  total_wickets: number;
  innings_end: string;
  target_runs: number;
  max_overs: number;
}

export interface ITeamInningsScorecard {
  innings_id: number;
  innings_number: number;
  batting_team_id: number;
  batting_team_name: string;

  total_runs: number;
  total_overs: number;
  total_wickets: number;
  byes: number;
  leg_byes: number;
  wides: number;
  noballs: number;
  penalty: number;

  innings_end: string;
  target_runs: number;
  max_overs: number;

  batter_scorecard_entries?: IBatterScorecardEntry[];
  bowler_scorecard_entries?: IBowlerScorecardEntry[];
  fall_of_wickets?: IFallOfWickets[];
}

export interface IBatterScorecardEntry {
  batter_id: number;
  batter_name: string;
  batting_position: number;
  has_batted: boolean;

  runs_scored: number;
  balls_faced: number;
  minutes_batted: number;
  fours_scored: number;
  sixes_scored: number;

  dismissal_type: string;
  dismissed_by_id: number;
  dismissed_by_name: string;
  fielder1_id: number;
  fielder1_name: string;
  fielder2_id: number;
  fielder2_name: string;
}

export interface IBowlerScorecardEntry {
  bowler_id: number;
  bowler_name: string;
  bowling_position: number;

  wickets_taken: number;
  runs_conceded: number;
  overs_bowled: number;
  maiden_overs: number;
  fours_conceded: number;
  sixes_conceded: number;
  wides_conceded: number;
  noballs_conceded: number;
}

export interface IFallOfWickets {
  batter_id: number;
  batter_name: string;
  ball_number: number;
  team_runs: number;
  wicket_number: number;
  dismissal_type: string;
}

/* Commentary */

export interface IMatchCommentary {
  match_header: IMatchHeader;
  commentary?: IBbbCommentary[];
}

export interface IBbbCommentary {
  innings_id: number;
  innings_delivery_number: number;
  ball_number: number;
  over_number: number;

  batter_id?: number;
  batter_name?: string;
  bowler_id?: number;
  bowler_name?: string;
  fielder1_id?: number;
  fielder1_name?: string;
  fielder2_id?: number;
  fielder2_name?: string;

  wides: number;
  noballs: number;
  legbyes: number;
  byes: number;
  total_runs: number;
  is_four: boolean;
  is_six: boolean;

  player1_dismissed_id?: number;
  player1_dismissed_name?: string;
  player1_dismissal_type?: string;
  player1_dismissed_runs?: number;
  player1_dismissed_balls?: number;
  player1_dismissed_fours?: number;
  player1_dismissed_sixes?: number;

  player2_dismissed_id?: number;
  player2_dismissed_name?: string;
  player2_dismissal_type?: string;
  player2_dismissed_runs?: number;
  player2_dismissed_balls?: number;
  player2_dismissed_fours?: number;
  player2_dismissed_sixes?: number;

  commentary: string;
}

export interface IOverSummary {
  battingTeamName: string;
  overNumber: number;
  overRuns: number;
  overWickets: number;
  totalRuns: number;
  totalWickets: number;

  batters: {
    id: number;
    name: string;
    runs: number;
    balls: number;
    fours: number;
    sixes: number;
    isActive: boolean;
  }[];

  bowlers: {
    id: number;
    name: string;
    runs: number;
    balls: number;
    maidens: number;
    wickets: number;
  }[];
}

/* Stats */

export interface IMatchStats {
  match_header: IMatchHeader;
  innings?: IInningsStats[];
}

export interface IInningsStats {
  innings_id: number;
  innings_number: number;
  batting_team_id: number;
  batting_team_name: string;
  partnerships?: IPartnershipStats[];
  overs?: IOverStats[];
}

export interface IPartnershipStats {
  for_wicket: number;
  is_unbeaten: boolean;
  start_ball_number: number;
  end_ball_number: number;
  partnership_runs: number;

  batter1_id: number;
  batter1_name: string;
  batter1_runs: number;
  batter1_balls: number;

  batter2_id: number;
  batter2_name: string;
  batter2_runs: number;
  batter2_balls: number;
}

export interface IOverStats {
  over_number: number;
  runs: number;
  balls: number;
  wickets: number;
}

/* Squads */

export interface IMatchSquad {
  match_header: IMatchHeader;
  team_squads?: ITeamSquadEntry[];
}

export interface ITeamSquadEntry {
  team_id: number;
  team_name: string;
  players?: IPlayerSquadEntry[];
}

export interface IPlayerSquadEntry {
  player_id: number;
  player_name: string;

  is_captain: boolean;
  is_wk: boolean;
  is_debut: boolean;
  is_vice_captain: boolean;
  playing_status: string;
}
