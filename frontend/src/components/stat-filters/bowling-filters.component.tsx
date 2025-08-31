import { IStatsFiltersMap } from "@/lib/types/filters-stats.types";
import { RangeInput } from "./range-input.component";

export function BowlingFilters({ filtersMap }: { filtersMap: IStatsFiltersMap }) {
  type keyEnum =
    | "innings_balls_bowled"
    | "innings_runs_conceded"
    | "innings_wickets_taken"
    | "innings_bowling_position";

  const filters: Record<keyEnum, string> = {
    innings_balls_bowled: "Balls bowled in an inns",
    innings_runs_conceded: "Runs conc. in an inns",
    innings_wickets_taken: "Wkts Taken in an inns",
    innings_bowling_position: "Bowling Position",
  };

  return (
    <div className="flex flex-col gap-4">
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
