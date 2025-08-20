import { EnumStatsType, EnumStatsView } from "@/lib/types/enums.types";
import { RadioInput } from "./radio-input.component";
import { SelectInput } from "./select-input.component";

export function SortFilters({
  sortOptions,
  defaultSortKey,
  defaultSortOrder,
}: {
  sortOptions: { label: string; value: string }[];
  defaultSortKey?: string;
  defaultSortOrder: "default" | "reverse";
}) {
  return (
    <div className="flex flex-col gap-4">
      {/* Sort By */}
      <SelectInput name="sort_by" label="sort by" options={sortOptions} defaultValue={defaultSortKey} />

      {/* Sort Order */}
      <RadioInput
        name="sort_order"
        label="sort order"
        options={[
          { label: "default", value: "default" },
          { label: "reverse", value: "reverse" },
        ]}
        defaultValue={defaultSortOrder}
      />
    </div>
  );
}

export function getSortOptions({
  statsType,
  viewValue,
  groupValue,
}: {
  statsType: EnumStatsType;
  viewValue: EnumStatsView;
  groupValue: string;
}) {
  const sortOptionsMap: Record<
    EnumStatsType,
    Record<
      string,
      {
        label: string;
        commonAll: boolean;
        commonOverall: boolean;
        commonIndividual: boolean;
        overallGroups: string[];
        individualGroups: string[];
      }
    >
  > = {
    batting: {
      runs_scored: { label: "Runs Scored", ...getQualificationFilterObject(true) },
      start_date: { label: "Start Date", ...getQualificationFilterObject(true) },
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      innings_batted: { label: "Innings Batted", ...getQualificationFilterObject(true) },
      not_outs: { label: "Not Outs", ...getQualificationFilterObject(true) },
      balls_faced: { label: "Balls Faced", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      strike_rate: { label: "Strike Rate", ...getQualificationFilterObject(true) },
      centuries: { label: "Centuries", ...getQualificationFilterObject(true) },
      half_centuries: { label: "Half Centuries", ...getQualificationFilterObject(true) },
      fifty_plus_scores: { label: "50+ scores", ...getQualificationFilterObject(true) },
      ducks: { label: "Ducks", ...getQualificationFilterObject(true) },
      fours_scored: { label: "Fours scored", ...getQualificationFilterObject(true) },
      sixes_scored: { label: "Sixes scored", ...getQualificationFilterObject(true) },
      player_name: { label: "Player Name", ...getQualificationFilterObject(false, false, true, ["batters"]) },
      innings_number: {
        label: "Innings Number",
        ...getQualificationFilterObject(false, false, false, [], ["innings"]),
      },
    },
    bowling: {
      wickets_taken: { label: "Wickets Taken", ...getQualificationFilterObject(true) },
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      innings_bowled: { label: "Innings Bowled", ...getQualificationFilterObject(true) },
      overs_bowled: { label: "Overs Bowled", ...getQualificationFilterObject(true) },
      maiden_overs: { label: "Maiden Overs", ...getQualificationFilterObject(true) },
      runs_conceded: { label: "Runs Conceded", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      strike_rate: { label: "Strike Rate", ...getQualificationFilterObject(true) },
      economy: { label: "Economy", ...getQualificationFilterObject(true) },
      four_wkt_hauls: {
        label: "4-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      five_wkt_hauls: {
        label: "5-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      ten_wkt_hauls: {
        label: "10-wkt Hauls",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      best_bowling_match: {
        label: "Best Bowling Match",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      best_bowling_innings: {
        label: "Best Bowling Innings",
        ...getQualificationFilterObject(
          false,
          true,
          false,
          [],
          ["series", "tournaments", "grounds", "host-nations", "oppositions", "years", "seasons"]
        ),
      },
      player_name: { label: "Player Name", ...getQualificationFilterObject(false, false, true, ["bowlers"]) },
    },
    team: {
      matches_won: { label: "Matches Won", ...getQualificationFilterObject(true) },
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      matches_lost: { label: "Matches Lost", ...getQualificationFilterObject(true) },
      matches_tied: { label: "Matches Tied", ...getQualificationFilterObject(true) },
      matches_drawn: { label: "Matches Drawn", ...getQualificationFilterObject(true) },
      matches_with_no_result: { label: "N/R Matches", ...getQualificationFilterObject(true) },
      win_loss_ratio: { label: "W/L Ratio", ...getQualificationFilterObject(true) },
      innings_count: { label: "Innings Count", ...getQualificationFilterObject(true) },
      total_runs: { label: "Total Runs", ...getQualificationFilterObject(true) },
      total_balls: { label: "Total Balls", ...getQualificationFilterObject(true) },
      total_wickets: { label: "Total Wickets", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      scoring_rate: { label: "Scoring Rate", ...getQualificationFilterObject(true) },
      highest_score: { label: "Highest Score", ...getQualificationFilterObject(true) },
      lowest_score: { label: "Lowest Score", ...getQualificationFilterObject(true) },
      start_date: { label: "Start Date", ...getQualificationFilterObject(true) },
      innings_number: {
        label: "Innings Number",
        ...getQualificationFilterObject(false, false, false, [], ["innings"]),
      },
      team_name: {
        label: "Team Name",
        ...getQualificationFilterObject(false, false, true, ["teams"]),
      },
      player_name: {
        label: "Player Name",
        ...getQualificationFilterObject(false, false, false, ["players"]),
      },
    },
  };

  const options: { label: string; value: string }[] = [];

  for (const sortKey in sortOptionsMap[statsType]) {
    const sortKeyInfo = sortOptionsMap[statsType][sortKey];
    if (
      sortKeyInfo.commonAll ||
      (viewValue === "overall" && sortKeyInfo.commonOverall) ||
      (viewValue === "individual" && sortKeyInfo.commonIndividual) ||
      (viewValue === "overall" && sortKeyInfo.overallGroups.includes(groupValue)) ||
      (viewValue === "individual" && sortKeyInfo.individualGroups.includes(groupValue))
    ) {
      options.push({ label: sortKeyInfo.label, value: sortKey });
    }
  }

  return options;
}

function getQualificationFilterObject(
  commonAll = false,
  commonOverall = false,
  commonIndividual = false,
  overallGroups: string[] = [],
  individualGroups: string[] = []
) {
  const obj = {
    commonAll,
    commonOverall,
    commonIndividual,
    overallGroups,
    individualGroups,
  };

  if (commonAll) {
    obj.commonOverall = obj.commonIndividual = false;
    obj.overallGroups = [];
    obj.individualGroups = [];
  }

  if (obj.commonOverall) obj.overallGroups = [];
  if (obj.commonIndividual) obj.individualGroups = [];

  return obj;
}
