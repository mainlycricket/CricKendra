import { Input } from "../ui/input";

export function BowlingFilters() {
  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Balls bowled in an inns</p>
        <div className="flex gap-2">
          <Input
            type="number"
            name="min__innings_balls_bowled"
            min={0}
            step={1}
            className="w-24"
            placeholder="Min"
          />
          <Input
            type="number"
            name="max__innings_balls_bowled"
            min={0}
            step={1}
            className="w-24"
            placeholder="Max"
          />
        </div>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Runs conc. in an inns</p>
        <div className="flex gap-2">
          <Input
            type="number"
            name="min__innings_runs_conceded"
            min={0}
            step={1}
            className="w-24"
            placeholder="Min"
          />
          <Input
            type="number"
            name="max__innings_runs_conceded"
            min={0}
            step={1}
            className="w-24"
            placeholder="Max"
          />
        </div>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Wkts Taken in an inns</p>
        <div className="flex gap-2">
          <Input
            type="number"
            name="min__innings_wickets_taken"
            min={0}
            step={1}
            className="w-24"
            placeholder="Min"
          />
          <Input
            type="number"
            name="max__innings_wickets_taken"
            min={0}
            step={1}
            className="w-24"
            placeholder="Max"
          />
        </div>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Bowling Position</p>
        <div className="flex gap-2">
          <Input
            type="number"
            name="min__innings_bowling_position"
            min={1}
            step={1}
            max={12}
            className="w-24"
            placeholder="Min"
          />
          <Input
            type="number"
            name="max__innings_bowling_position"
            min={1}
            step={1}
            max={12}
            className="w-24"
            placeholder="Max"
          />
        </div>
      </div>
    </div>
  );
}
