import { RadioOption } from "./radio-option.component";
import { RangeInput } from "./range-input.component";
import { SelectCheckboxInput } from "./select-checkbox-input.component";

export function BattingFilters() {
  return (
    <div className="flex flex-col gap-4">
      <RangeInput
        label="Runs scored in an inns"
        minName="min__innings_runs_scored"
        maxName="max__innings_runs_scored"
      />

      <RangeInput
        label="Batting Position"
        minName="min__innings_batting_position"
        maxName="max__innings_batting_position"
      />

      <RadioOption
        name="innings_is_batter_dismissed"
        label="Dismissed"
        defaultValue=""
        options={[
          { label: "out", value: "dismissed" },
          { label: "not out", value: "not_out" },
          { label: "either", value: "" },
        ]}
      />

      <SelectCheckboxInput
        name="innings_batter_dismissal_type"
        label="Dismissal Type"
        options={[
          { label: "Caught", value: "caught" },
          { label: "Bowled", value: "bowled" },
          { label: "LBW", value: "lbw" },
          { label: "Run Out", value: "run out" },
          { label: "Stumped", value: "stumped" },
          { label: "Hit Wicket", value: "hit wicket" },
          { label: "Handled the Ball", value: "handled the ball" },
          { label: "Obstructing the Field", value: "obstructing the field" },
          { label: "Timed Out", value: "timed out" },
          { label: "Retired Hurt", value: "retired hurt" },
          { label: "Hit the Ball Twice", value: "hit the ball twice" },
          { label: "Caught and Bowled", value: "caught and bowled" },
          { label: "Retired Out", value: "retired out" },
          { label: "Retired Not Out", value: "retired not out" },
        ]}
      />
    </div>
  );
}
