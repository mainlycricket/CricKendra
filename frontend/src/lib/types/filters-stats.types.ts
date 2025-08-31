import { ReadonlyURLSearchParams } from "next/navigation";
import { isValidIsoDate } from "../utils";
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
  isEnumBattingQualificationKey,
  isEnumBowlingQualificationKey,
  isEnumDismissalType,
  isEnumHomeAway,
  isEnumInningsNumber,
  isEnumMatchResult,
  isEnumPlayingFormat,
  isEnumStatsType,
  isEnumStatsView,
  isEnumTeamQualificationKey,
} from "./enums.types";

export interface IStatsFilters {
  teams?: ITeamAsForeignField[];
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

export interface IStatsFiltersMap {
  [key: string]: string | string[] | undefined;

  statsType: EnumStatsType;
  view: EnumStatsView;
  group: string;

  playing_format: EnumPlayingFormat;
  is_male: "true" | "false";
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

  /* bowling specific filters only */
  min__team_innings_runs?: string;
  max__team_innings_runs?: string;
  min__team_innings_wickets?: string;
  max__team_innings_wickets?: string;
  min__team_innings_balls?: string;
  max__team_innings_balls?: string;

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
  sort_order?: "default" | "reverse";
  __page?: string;
  __limit?: string;
}

export function prepareStatsFiltersMap(searchParams: ReadonlyURLSearchParams): IStatsFiltersMap {
  const statsTypeValue = searchParams.get("statsType"),
    viewValue = searchParams.get("view"),
    groupValue = searchParams.get("group");

  const statsType = statsTypeValue && isEnumStatsType(statsTypeValue) ? statsTypeValue : "batting";
  const view = viewValue && isEnumStatsView(viewValue) ? viewValue : "overall";

  const group = groupValue
    ? groupValue
    : view === "individual"
    ? "innings"
    : statsType === "batting"
    ? "batters"
    : "bowlers";

  const filterMap: IStatsFiltersMap = {
    statsType,
    view,
    group,
    playing_format: "ODI",
    is_male: "true",
  };

  for (const key of searchParams.keys()) {
    if (key === "statsType" || key === "view" || key === "group") continue;

    const values = searchParams.getAll(key).length > 1 ? searchParams.getAll(key) : searchParams.get(key);
    if (!values?.length) continue;

    if (key === "playing_format" && typeof values === "string") {
      filterMap[key] = isEnumPlayingFormat(values) ? values : "ODI";
    } else if (key === "is_male" && (values === "true" || values === "false")) {
      filterMap[key] = values;
    } else if (key === "min_start_date" && typeof values === "string" && isValidIsoDate(values)) {
      filterMap[key] = values;
    } else if (key === "max_start_date" && typeof values === "string" && isValidIsoDate(values)) {
      filterMap[key] = values;
    } else if (
      key === "season" ||
      key === "primary_team" ||
      key === "opposition_team" ||
      key === "continent" ||
      key === "host_nation" ||
      key === "ground" ||
      key === "series" ||
      key === "tournament"
    ) {
      filterMap[key] = Array.isArray(values) ? values : [values];
    } else if (key === "home_or_away") {
      filterMap[key] = Array.isArray(values)
        ? values.filter((value) => isEnumHomeAway(value))
        : isEnumHomeAway(values)
        ? [values]
        : undefined;
    } else if (key === "match_result") {
      filterMap[key] = Array.isArray(values)
        ? values.filter((value) => isEnumMatchResult(value))
        : isEnumMatchResult(values)
        ? [values]
        : undefined;
    } else if (key === "toss_result" && (values === "won" || values === "lost")) {
      filterMap[key] = values;
    } else if (key === "bat_field_first" && (values === "bat" || values === "field")) {
      filterMap[key] = values;
    } else if (key === "innings_number") {
      filterMap[key] = Array.isArray(values)
        ? values.filter((value) => isEnumInningsNumber(value))
        : isEnumInningsNumber(values)
        ? [values as EnumInningsNumber]
        : undefined;
    } else if (
      statsType === "batting" &&
      isEnumBattingQualificationKey(key) &&
      typeof values === "string" &&
      !isNaN(parseInt(values))
    ) {
      filterMap[key] = parseInt(values).toString();
    } else if (
      statsType === "batting" &&
      key === "innings_is_batter_dismissed" &&
      (values === "dismissed" || values === "not_out")
    ) {
      filterMap[key] = values;
    } else if (statsType === "batting" && key === "innings_batter_dismissal_type") {
      if (typeof values === "string" && isEnumDismissalType(values)) {
        filterMap[key] = [values];
      } else if (Array.isArray(values)) {
        filterMap[key] = values.filter((value) => isEnumDismissalType(value));
      }
    } else if (
      statsType === "bowling" &&
      (key === "min__four_wkt_hauls" ||
      key === "max__four_wkt_hauls" ||
      key === "min__five_wkt_hauls" ||
      key === "max__five_wkt_hauls" ||
      key === "min__ten_wkt_hauls" ||
      key === "max__ten_wkt_hauls"
        ? view === "overall" || (group !== "innings" && group !== "match-totals")
        : isEnumBowlingQualificationKey(key)) &&
      typeof values === "string" &&
      !isNaN(parseInt(values))
    ) {
      filterMap[key] = parseInt(values).toString();
    } else if (
      statsType === "team" &&
      isEnumTeamQualificationKey(key) &&
      typeof values === "string" &&
      !isNaN(parseInt(values))
    ) {
      filterMap[key] = parseInt(values).toString();
    } else if (
      key === "sort_by" &&
      statsType === "batting" &&
      (values === "runs_scored" ||
        values === "start_date" ||
        values === "matches_played" ||
        values === "innings_batted" ||
        values === "not_outs" ||
        values === "balls_faced" ||
        values === "average" ||
        values === "strike_rate" ||
        values === "centuries" ||
        values === "half_centuries" ||
        values === "fifty_plus_scores" ||
        values === "ducks" ||
        values === "fours_scored" ||
        values === "sixes_scored" ||
        (values === "player_name" && (view === "individual" || group === "batters")) ||
        (values === "innings_number" && view === "individual" && group === "innings"))
    ) {
      filterMap[key] = values;
    } else if (
      key === "sort_by" &&
      statsType === "bowling" &&
      (values === "wickets_taken" ||
        values === "matches_played" ||
        values === "innings_bowled" ||
        values === "overs_bowled" ||
        values === "maiden_overs" ||
        values === "runs_conceded" ||
        values === "average" ||
        values === "strike_rate" ||
        values === "economy" ||
        ((values === "four_wkt_hauls" ||
          values === "five_wkt_hauls" ||
          values === "ten_wkt_hauls" ||
          values === "best_bowling_match" ||
          values === "best bowling innings") &&
          (view === "overall" ||
            ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"].includes(
              group
            ))) ||
        (values === "player_name" && (view === "individual" || group === "bowlers")))
    ) {
      filterMap[key] = values;
    } else if (
      key === "sort_by" &&
      statsType === "team" &&
      (values === "matches_won" ||
        values === "matches_played" ||
        values === "matches_lost" ||
        values === "matches_tied" ||
        values === "matches_drawn" ||
        values === "matches_with_no_result" ||
        values === "win_loss_ratio" ||
        values === "innings_count" ||
        values === "total_runs" ||
        values === "total_balls" ||
        values === "total_wickets" ||
        values === "average" ||
        values === "scoring_rate" ||
        values === "highest_score" ||
        values === "lowest_score" ||
        values === "start_date" ||
        (values === "innings_number" && (view === "individual" || group === "innings")) ||
        (values === "team_name" && (view === "individual" || group === "teams")) ||
        (values === "player_name" && view === "overall" && group === "players"))
    ) {
      filterMap[key] = values;
    } else if (key === "sort_order" && (values === "default" || values === "reverse")) {
      filterMap[key] = values;
    } else if (
      (key === "__page" || key === "__limit") &&
      typeof values === "string" &&
      !isNaN(parseInt(values))
    ) {
      filterMap[key] = values;
    }
  }

  return filterMap;
}

export function stringifyFiltersMap(filtersMap: IStatsFiltersMap): string {
  const query: string[] = [];

  for (const key in filtersMap) {
    const value = filtersMap[key];
    if (value) {
      if (Array.isArray(value)) {
        for (const item of value) {
          if (item?.trim()?.length) query.push(`${key}=${item}`);
        }
      } else if (value?.trim()?.length) {
        query.push(`${key}=${value}`);
      }
    }
  }

  return query.join("&");
}
