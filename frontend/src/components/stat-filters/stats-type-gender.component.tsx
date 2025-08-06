import { EnumStatsType } from "@/lib/types/enums.types";
import { Dispatch, SetStateAction } from "react";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

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
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Gender</p>
        <div>
          <Select name="is_male" defaultValue={defaultIsMale || "true"}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Gender" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel className="text-sm">Gender</SelectLabel>
                <SelectItem value="true">Male</SelectItem>
                <SelectItem value="false">Female</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>
      </div>

      {/* Stats Type */}
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Stats Type</p>
        <div>
          <Select
            name="type"
            value={defaultStatsType}
            onValueChange={(value: EnumStatsType) => {
              setStatsTypeValue(value);
            }}
          >
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Stats Type" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel className="text-sm">Type</SelectLabel>
                <SelectItem value="batting">Batting</SelectItem>
                <SelectItem value="bowling">Bowling</SelectItem>
                <SelectItem value="team">Team</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>
      </div>
    </div>
  );
}
