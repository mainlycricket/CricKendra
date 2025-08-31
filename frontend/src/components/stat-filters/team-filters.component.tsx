import { IStatsFiltersMap } from "@/lib/types/filters-stats.types";
import { RangeInput } from "./range-input.component";
import { RadioInput } from "./radio-input.component";

export function TeamFilters({ filtersMap }: { filtersMap: IStatsFiltersMap }) {
  type keyEnum = "team_innings_runs" | "team_innings_wickets" | "team_innings_balls";

  const filters: Record<keyEnum, string> = {
    team_innings_runs: "Team runs in an inns",
    team_innings_wickets: "Team wkts in an inns",
    team_innings_balls: "Team balls in an inns",
  };

  return (
    <div className="flex flex-col gap-4">
      <RadioInput
        label="team total for"
        name="team_total_for"
        options={[
          { label: "Batting", value: "batting" },
          { label: "Bowling", value: "bowling" },
        ]}
        defaultValue="batting"
      />

      {(Object.keys(filters) as keyEnum[]).map((key) => {
        return (
          <RangeInput
            key={key}
            label={filters[key]}
            minName={`min__${key}`}
            maxName={`max__${key}`}
            defaultMin={filtersMap?.[`min__${key}`]}
            defaultMax={filtersMap?.[`max__${key}`]}
          />
        );
      })}
    </div>
  );
}
