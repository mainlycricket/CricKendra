import { StatsFiltersComponent } from "@/components/stat-filters/main-layout.component";
import { BattingStats } from "@/components/stats/batting-stats.component";
import { BowlingStats } from "@/components/stats/bowling-stats.component";
import { CommonStatsLayout } from "@/components/stats/common-layout.component";
import { doBackendRequest } from "@/lib/axiosFetch";
import { ICombinedBattingStatsType } from "@/lib/types/batting-stats.types";
import { ICombinedBowlingStatsType } from "@/lib/types/bowling-stats.types";
import { IStats } from "@/lib/types/common-stats.types";
import {
  EnumBatField,
  EnumHomeAway,
  EnumInningsNumber,
  EnumMatchResult,
  EnumPlayingFormat,
  EnumStatsType,
  EnumStatsView,
  EnumTossResult,
  isEnumDismissalType,
} from "@/lib/types/enums.types";
import { IStatsFilters, IStatsFiltersData } from "@/lib/types/filters-stats.types";
import { isValidIsoDate } from "@/lib/utils";

export default async function StatsFilters({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | string[] }>;
}) {
  try {
    const filters = await searchParams;
    const statsType = (filters["type"] || "batting") as EnumStatsType;
    const view = (filters["view"] || "overall") as EnumStatsView;
    const group = (filters["group"] || "batters") as string;

    const filterMap: IStatsFiltersData = {
      statsType,
      view,
      group,
    };

    for (const key in filters) {
      if (key === "type" || key === "view" || key === "group") continue;

      const values = filters[key];
      if (!values?.length) continue;

      if (key === "playing_format" && typeof values === "string") {
        filterMap[key] = values as EnumPlayingFormat;
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
        filterMap[key] = Array.isArray(values) ? (values as EnumHomeAway[]) : [values as EnumHomeAway];
      } else if (key === "match_result") {
        filterMap[key] = Array.isArray(values) ? (values as EnumMatchResult[]) : [values as EnumMatchResult];
      } else if (key === "toss_result" && (values === "won" || values === "lost")) {
        filterMap[key] = values;
      } else if (key === "bat_field_first" && (values === "bat" || values === "field")) {
        filterMap[key] = values;
      } else if (key === "innings_number") {
        filterMap[key] = Array.isArray(values)
          ? (values as EnumInningsNumber[])
          : [values as EnumInningsNumber];
      } else if (
        statsType === "batting" &&
        (key === "min__innings_runs_scored" ||
          key === "max__innings_runs_scored" ||
          key === "min__innings_batting_position" ||
          key === "max__innings_batting_position" ||
          key === "min__matches_played" ||
          key === "max__matches_played" ||
          key === "min__innings_batted" ||
          key === "max__innings_batted" ||
          key === "min__runs_scored" ||
          key === "max__runs_scored" ||
          key === "min__balls_faced" ||
          key === "max__balls_faced" ||
          key === "min__average" ||
          key === "max__average" ||
          key === "min__strike_rate" ||
          key === "max__strike_rate" ||
          key === "min__centuries" ||
          key === "max__centuries" ||
          key === "min__half_centuries" ||
          key === "max__half_centuries" ||
          key === "min__fifty_plus_scores" ||
          key === "max__fifty_plus_scores" ||
          key === "min__ducks" ||
          key === "max__ducks" ||
          key === "min__fours_scored" ||
          key === "max__fours_scored" ||
          key === "min__sixes_scored" ||
          key === "max__sixes_scored") &&
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
        (key === "min__innings_balls_bowled" ||
          key === "max__innings_balls_bowled" ||
          key === "min__innings_runs_conceded" ||
          key === "max__innings_runs_conceded" ||
          key === "min__innings_wickets_taken" ||
          key === "max__innings_wickets_taken" ||
          key === "min__innings_bowling_position" ||
          key === "max__innings_bowling_position" ||
          key === "min__matches_played" ||
          key === "max__matches_played" ||
          key === "min__innings_bowled" ||
          key === "max__innings_bowled" ||
          key === "min__overs_bowled" ||
          key === "max__overs_bowled" ||
          key === "min__maiden_overs" ||
          key === "max__maiden_overs" ||
          key === "min__runs_conceded" ||
          key === "max__runs_conceded" ||
          key === "min__wickets_taken" ||
          key === "max__wickets_taken" ||
          key === "min__average" ||
          key === "max__average" ||
          key === "min__strike_rate" ||
          key === "max__strike_rate" ||
          key === "min__economy" ||
          key === "max__economy" ||
          key === "min__fours_conceded" ||
          key === "max__fours_conceded" ||
          key === "min__sixes_conceded" ||
          key === "max__sixes_conceded" ||
          ((key === "min__four_wkt_hauls" ||
            key === "max__four_wkt_hauls" ||
            key === "min__five_wkt_hauls" ||
            key === "max__five_wkt_hauls" ||
            key === "min__ten_wkt_hauls" ||
            key === "max__ten_wkt_hauls") &&
            (view === "overall" ||
              [
                "series",
                "tournaments",
                "grounds",
                "host-nations",
                "oppositions",
                "years",
                "seasons",
              ].includes(group)))) &&
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
          (values === "innings_number" && (view === "individual" || group === "innings")))
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
              [
                "series",
                "tournaments",
                "grounds",
                "host-nations",
                "oppositions",
                "years",
                "seasons",
              ].includes(group))) ||
          (values === "player_name" && (view === "individual" || group === "bowlers")))
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

    const query: string[] = [];
    for (const key in filterMap) {
      if (key === "statsType" || key === "view" || key === "group") continue;

      const value = filterMap[key];
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

    let filtersUrl = `/stats/filter-options?`;
    if (filterMap["playing_format"]) filtersUrl += `playing_format=${filterMap["playing_format"]}&`;
    if (filterMap["is_male"]) filtersUrl += `is_male=${filterMap["is_male"]}`;
    const filtersResponse = await doBackendRequest<null, IStatsFilters>({ url: filtersUrl, method: "GET" });

    const statsUrl = `/stats/${statsType}/${view}/${group}?${query.join("&")}`;
    const statsResponse = await doBackendRequest<null, IStats>({
      url: statsUrl,
      method: "GET",
    });

    return (
      <CommonStatsLayout
        stats={statsResponse.data!}
        statsType={statsType}
        filtersMap={filterMap}
        filtersIdData={filtersResponse.data!}
      />
    );
  } catch (error) {
    console.error(error);
  }
}
