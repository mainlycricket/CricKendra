import { EnumStatsType } from "@/lib/types/enums.types";
import { Dispatch, SetStateAction } from "react";
import { SelectInput } from "./select-input.component";

export function StatsTypeGenderDropdowns({
  defaultIsMale,
  defaultStatsType,
  setStatsTypeValue,
}: {
  defaultIsMale: "true" | "false";
  defaultStatsType: EnumStatsType;
  setStatsTypeValue: Dispatch<SetStateAction<EnumStatsType>>;
}) {
  return (
    <div className="flex flex-col gap-4">
      {/* Gender */}
      <SelectInput
        name="is_male"
        label="Gender"
        options={[
          { label: "Male", value: "true" },
          { label: "Female", value: "false" },
        ]}
        defaultValue={defaultIsMale || "true"}
      />

      {/* Stats Type */}
      <SelectInput
        name="statsType"
        label="Stats Type"
        options={[
          { label: "Batting", value: "batting" },
          { label: "Bowling", value: "bowling" },
          { label: "Team", value: "team" },
        ]}
        defaultValue={defaultStatsType}
        onValueChange={(value) => {
          setStatsTypeValue(value as EnumStatsType);
        }}
      />
    </div>
  );
}
