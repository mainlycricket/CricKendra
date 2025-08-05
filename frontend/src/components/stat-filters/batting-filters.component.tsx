import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import { StatOption } from "./options.component";

export function BattingFilters() {
  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Runs scored in an inns</p>
        <div className="flex gap-2">
          <Input
            type="number"
            name="min__innings_runs_scored"
            min={0}
            step={1}
            className="w-24"
            placeholder="Min"
          />
          <Input
            type="number"
            name="max__innings_runs_scored"
            min={0}
            step={1}
            className="w-24"
            placeholder="Max"
          />
        </div>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Batting Position</p>
        <div className="flex gap-2">
          <Input
            type="number"
            name="min__innings_batting_position"
            min={1}
            step={1}
            max={12}
            className="w-24"
            placeholder="Min"
          />
          <Input
            type="number"
            name="max__innings_batting_position"
            min={1}
            step={1}
            max={12}
            className="w-24"
            placeholder="Max"
          />
        </div>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Dismissed</p>
        <RadioGroup defaultValue="" name="innings_is_batter_dismissed" className="flex flex-wrap gap-4">
          <div className="flex items-center gap-1">
            <RadioGroupItem value="dismissed" id="dismissed" />
            <Label htmlFor="dismissed" className="capitalize font-normal">
              out
            </Label>
          </div>
          <div className="flex items-center gap-1">
            <RadioGroupItem value="not_out" id="not_out" />
            <Label htmlFor="not_out" className="capitalize font-normal">
              not out
            </Label>
          </div>
          <div className="flex items-center gap-1">
            <RadioGroupItem value="" id="either" />
            <Label htmlFor="either" className="capitalize font-normal">
              either
            </Label>
          </div>
        </RadioGroup>
      </div>

      <StatOption
        name="innings_batter_dismissal_type"
        label="Dismissal Type"
        optionValues={[
          "caught",
          "bowled",
          "lbw",
          "run out",
          "stumped",
          "hit wicket",
          "handled the ball",
          "obstructing the field",
          "timed out",
          "retired hurt",
          "hit the ball twice",
          "caught and bowled",
          "retired out",
          "retired not out",
        ]}
        optionLabels={[
          "Caught",
          "Bowled",
          "LBW",
          "Run Out",
          "Stumped",
          "Hit Wicket",
          "Handled the Ball",
          "Obstructing the Field",
          "Timed Out",
          "Retired Hurt",
          "Hit the Ball Twice",
          "Caught and Bowled",
          "Retired Out",
          "Retired Not Out",
        ]}
      />
    </div>
  );
}
