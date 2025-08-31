import {
  EnumBattingQualificationKey,
  EnumBowlingQualificationKey,
  EnumTeamQualificationKey,
} from "@/lib/types/enums.types";
import { IStatsFilters, IStatsFiltersMap } from "@/lib/types/filters-stats.types";
import { capitalizeFirstLetter, getDismissalTypeDisplayText, getDisplayDate } from "@/lib/utils";

export function FiltersList({
  filtersMap,
  filtersIdData,
}: {
  filtersMap: IStatsFiltersMap;
  filtersIdData: IStatsFilters;
}) {
  type qualKeyEnum = EnumBattingQualificationKey | EnumBowlingQualificationKey | EnumTeamQualificationKey;

  const qualFilters: Record<qualKeyEnum, string> = {
    /* Batting */
    innings_runs_scored: "Runs scored in an inns",
    innings_batting_position: "Batting Position",

    /* Bowling */
    innings_balls_bowled: "Balls Bowled in an inns",
    innings_runs_conceded: "Runs conceded in an inns",
    innings_wickets_taken: "Wkts Taken in an inns",
    innings_bowling_position: "Bowling Position",

    /* Team */
    team_innings_runs: "Innings runs",
    team_innings_wickets: "Innings wickets",
    team_innings_balls: "Innings balls",

    // Common
    matches_played: "Matches Played",

    // Batting
    innings_batted: "Innings Batted",
    runs_scored: "Runs Scored",
    not_outs: "Not Outs",
    balls_faced: "Balls Faced",

    // Bowling
    innings_bowled: "Innings Bowled",
    overs_bowled: "Overs Bowled",
    maiden_overs: "Maiden Overs",
    runs_conceded: "Runs Conceded",
    wickets_taken: "Wickets Taken",

    // Team
    matches_won: "Matches Won",
    matches_lost: "Matches Lost",
    matches_tied: "Matches Tied",
    matches_drawn: "Matches Drawn",
    matches_with_no_result: "N/R Matches",

    // Batting & Bowling
    average: "Average",
    strike_rate: "Strike Rate",
    scoring_rate: "Scoring Rate",

    // Batting
    centuries: "Centuries",
    half_centuries: "Half Centuries",
    fifty_plus_scores: "50+ scores",
    ducks: "Ducks",
    fours_scored: "4s scored",
    sixes_scored: "6s scored",

    // Bowling
    economy: "Economy",
    fours_conceded: "4s conceded",
    sixes_conceded: "6s conceded",
    four_wkt_hauls: "4-wkt hauls",
    five_wkt_hauls: "5-wkt hauls",
    ten_wkt_hauls: "10-wkt hauls",

    // Team
    win_loss_ratio: "W/L Ratio",
    innings_count: "Innings Count",
    total_runs: "Total Runs",
    total_balls: "Total Balls",
    total_wickets: "Total Wickets",
  };

  return (
    <div className="flex flex-col gap-2">
      <FilterItem label="Stats Type" value1={capitalizeFirstLetter(filtersMap["statsType"])} />
      <FilterItem label="Playing Format" value1={filtersMap["playing_format"]} />
      <FilterItem label="Gender" value1={filtersMap["is_male"] === "true" ? "Male" : "Female"} />

      <FilterItem
        label="Primary Team"
        value1={filtersMap?.primary_team
          ?.map((id) => {
            const name = filtersIdData?.teams?.find((item) => item.id.toString() === id)?.name || "";
            return name;
          })
          ?.join(" or ")}
      />

      <FilterItem
        label="Opposition Team"
        value1={filtersMap?.opposition_team
          ?.map((id) => {
            const name = filtersIdData?.teams?.find((item) => item.id.toString() === id)?.name || "";
            return name;
          })
          ?.join(" or ")}
      />

      <FilterItem
        label="Start Date"
        value1={filtersMap?.min_start_date ? getDisplayDate(new Date(filtersMap?.min_start_date)) : ""}
        value2={filtersMap?.max_start_date ? getDisplayDate(new Date(filtersMap?.max_start_date)) : ""}
        displayOn="laptop"
      />

      <FilterItem
        label="Min. Start Date"
        value1={filtersMap?.min_start_date ? getDisplayDate(new Date(filtersMap?.min_start_date)) : ""}
        displayOn="mobile"
      />

      <FilterItem
        label="Max. Start Date"
        value1={filtersMap?.max_start_date ? getDisplayDate(new Date(filtersMap?.max_start_date)) : ""}
        displayOn="mobile"
      />

      <FilterItem label="Season" value1={filtersMap?.season?.join(" or ")} />

      <FilterItem
        label="Home or Away"
        value1={filtersMap?.home_or_away?.map((value) => capitalizeFirstLetter(value))?.join(" or ")}
      />

      <FilterItem
        label="Continents"
        value1={filtersMap?.continent
          ?.map((id) => {
            const name = filtersIdData?.continents?.find((item) => item.id.toString() === id)?.name;
            if (name) return name;
          })
          ?.join(" or ")}
      />

      <FilterItem
        label="Host Nations"
        value1={filtersMap?.host_nation
          ?.map((id) => {
            const name = filtersIdData?.host_nations?.find((item) => item.id.toString() === id)?.name;
            if (name) return name;
          })
          ?.join(" or ")}
      />

      <FilterItem
        label="Grounds"
        value1={filtersMap?.ground
          ?.map((id) => {
            const name = filtersIdData?.grounds?.find((item) => item.id.toString() === id)?.name;
            if (name) return name;
          })
          ?.join(" or ")}
      />

      <FilterItem
        label="Tournaments"
        value1={filtersMap?.tournament
          ?.map((id) => {
            const name = filtersIdData?.tournaments?.find((item) => item.id.toString() === id)?.name;
            if (name) return name;
          })
          ?.join(" or ")}
      />

      <FilterItem
        label="Series"
        value1={filtersMap?.series
          ?.map((id) => {
            const item = filtersIdData?.series?.find((item) => item.id.toString() === id);
            if (item) {
              return `${item.name}, ${item.season}`;
            }
          })
          ?.join(" or ")}
      />

      <FilterItem
        label="Match Result"
        value1={filtersMap?.match_result?.map((value) => capitalizeFirstLetter(value))?.join(" or ")}
      />

      <FilterItem label="Toss Result" value1={filtersMap?.toss_result} />
      <FilterItem label="Bat or Field First" value1={filtersMap?.bat_field_first} />
      <FilterItem label="Innings Number" value1={filtersMap?.innings_number?.sort()?.join(" or ")} />

      {/* Batting Specific filters */}
      <FilterItem
        label="Dismissed"
        value1={
          filtersMap?.innings_is_batter_dismissed
            ? capitalizeFirstLetter(filtersMap?.innings_is_batter_dismissed?.replaceAll("_", " "))
            : ""
        }
      />

      <FilterItem
        label="Dismissal Type"
        value1={filtersMap?.innings_batter_dismissal_type
          ?.map((value) => getDismissalTypeDisplayText(value))
          ?.join(" or ")}
      />

      {(Object.keys(qualFilters) as qualKeyEnum[])?.map((key) => {
        return (
          <FilterItem
            key={key}
            label={qualFilters[key]}
            value1={filtersMap?.[`min__${key}`] ? `Min. ${filtersMap?.[`min__${key}`]}` : ""}
            value2={filtersMap?.[`max__${key}`] ? `Max. ${filtersMap?.[`max__${key}`]}` : ""}
          />
        );
      })}

      {/* View & Sorting */}

      <FilterItem label="View" value1={capitalizeFirstLetter(filtersMap.view)} />
      <FilterItem label="Group By" value1={capitalizeFirstLetter(filtersMap.group)} />

      <FilterItem
        label="Sort By"
        value1={filtersMap?.sort_by ? capitalizeFirstLetter(filtersMap?.sort_by?.replaceAll("_", " ")) : ""}
      />

      <FilterItem
        label="Sort Order"
        value1={filtersMap?.sort_order ? capitalizeFirstLetter(filtersMap?.sort_order) : ""}
      />
    </div>
  );
}

function FilterItem({
  label,
  displayOn = "both",
  value1,
  value2,
}: {
  label: string;
  displayOn?: "mobile" | "laptop" | "both";
  value1?: string;
  value2?: string;
}) {
  if (!value1 && !value2) return <></>;

  const displayClass =
    displayOn === "mobile"
      ? "flex md:hidden gap-2"
      : displayOn === "laptop"
      ? "hidden md:flex gap-2"
      : "flex gap-2";

  return (
    <div className={displayClass}>
      <p className="w-[180px] font-medium">{label}</p>
      <p>
        {value1 && <span>{value1}</span>}
        {value1 && value2 && <span> and </span>}
        {value2 && <span>{value2}</span>}
      </p>
    </div>
  );
}
