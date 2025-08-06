import { RangeInput } from "./range-input.component";

export function BowlingFilters() {
  return (
    <div className="flex flex-col gap-4">
      <RangeInput
        label="Balls bowled in an inns"
        minName="min__innings_balls_bowled"
        maxName="max__innings_balls_bowled"
      />

      <RangeInput
        label="Runs conc. in an inns"
        minName="min__innings_runs_conceded"
        maxName="max__innings_runs_conceded"
      />

      <RangeInput
        label="Wkts Taken in an inns"
        minName="min__innings_wickets_taken"
        maxName="max__innings_wickets_taken"
      />

      <RangeInput
        label="Bowling Position"
        minName="min__innings_bowling_position"
        maxName="max__innings_bowling_position"
      />
    </div>
  );
}
