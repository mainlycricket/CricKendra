import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { Input } from "../ui/input";
import { Button } from "../ui/button";

import { PlusSquareIcon, Trash2 } from "lucide-react";
import { useEffect, useState } from "react";
import { EnumStatsType, EnumStatsView } from "@/lib/types/enums.types";

export function QualificationFilters({ options }: { options: { label: string; value: string }[] }) {
  const [usedOptions, setUsedOptions] = useState([options?.[0]]);

  useEffect(() => {
    setUsedOptions([options?.[0]]);
  }, [options]);

  return (
    <div className="hidden md:flex justify-between">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Result Qualifications</p>
        <div className="flex flex-col gap-2">
          {usedOptions?.length ? (
            usedOptions?.map((item, idx) => {
              return (
                <div key={item.value} className="flex gap-2">
                  <Select
                    defaultValue={item.value}
                    onValueChange={(value) => {
                      const updatedUsedOptions = [...usedOptions];
                      const newOption = options.find((x) => x.value === value);
                      if (newOption) {
                        updatedUsedOptions[idx] = { ...newOption };
                      }
                      setUsedOptions(updatedUsedOptions);
                    }}
                  >
                    <SelectTrigger className="w-[180px]">
                      <SelectValue placeholder="Select" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectLabel className="text-sm">Qualification Filters</SelectLabel>
                        {options.map((x) => {
                          const isUsed = usedOptions.find((y) => y.value === x.value);
                          if (!isUsed || x.value === item.value) {
                            return (
                              <SelectItem key={x.value} value={x.value}>
                                {x.label}
                              </SelectItem>
                            );
                          }
                        })}
                      </SelectGroup>
                    </SelectContent>
                  </Select>

                  <Input
                    type="number"
                    name={`min__${item.value}`}
                    placeholder="Min"
                    min={1}
                    step={1}
                    className="w-24"
                  />
                  <Input
                    type="number"
                    name={`max__${item.value}`}
                    placeholder="Max"
                    min={1}
                    step={1}
                    className="w-24"
                  />
                  <Button
                    type="button"
                    onClick={() => {
                      const updatedFilters = usedOptions.filter((x) => x.value !== item.value);
                      setUsedOptions(updatedFilters);
                    }}
                  >
                    <Trash2 />
                  </Button>
                </div>
              );
            })
          ) : (
            <p>Add a qualification Filter</p>
          )}
        </div>
      </div>
      <div>
        <Button
          type="button"
          className="flex gap-2"
          disabled={options?.length === usedOptions?.length}
          onClick={() => {
            const arr = [...usedOptions];
            for (const option of options) {
              const isUsed = usedOptions.find((item) => item.value === option.value);
              if (!isUsed) {
                arr.push({ ...option });
                break;
              }
            }
            setUsedOptions(arr);
          }}
        >
          <PlusSquareIcon /> Add
        </Button>
      </div>
    </div>
  );
}

export function getQualificationOptions({
  statsType,
  viewValue,
  groupValue,
}: {
  statsType: EnumStatsType;
  viewValue: EnumStatsView;
  groupValue: string;
}) {
  const qualificationFilters: Record<
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
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      innings_batted: { label: "Innings Batted", ...getQualificationFilterObject(true) },
      not_outs: { label: "Not Outs", ...getQualificationFilterObject(true) },
      runs_scored: { label: "Runs Scored", ...getQualificationFilterObject(true) },
      balls_faced: { label: "Balls Faced", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      strike_rate: { label: "Strike Rate", ...getQualificationFilterObject(true) },
      centuries: { label: "Centuries", ...getQualificationFilterObject(true) },
      half_centuries: { label: "Half Centuries", ...getQualificationFilterObject(true) },
      fifty_plus_scores: { label: "50+ scores", ...getQualificationFilterObject(true) },
      ducks: { label: "Ducks", ...getQualificationFilterObject(true) },
      fours_scored: { label: "Fours scored", ...getQualificationFilterObject(true) },
      sixes_scored: { label: "Sixes scored", ...getQualificationFilterObject(true) },
    },
    bowling: {
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      innings_bowled: { label: "Innings Bowled", ...getQualificationFilterObject(true) },
      overs_bowled: { label: "Overs Bowled", ...getQualificationFilterObject(true) },
      maiden_overs: { label: "Maiden Overs", ...getQualificationFilterObject(true) },
      runs_conceded: { label: "Runs Conceded", ...getQualificationFilterObject(true) },
      wickets_taken: { label: "Wickets Taken", ...getQualificationFilterObject(true) },
      average: { label: "Average", ...getQualificationFilterObject(true) },
      strike_rate: { label: "Strike Rate", ...getQualificationFilterObject(true) },
      economy: { label: "Economy", ...getQualificationFilterObject(true) },
      fours_conceded: { label: "Fours Conceded", ...getQualificationFilterObject(true) },
      sixes_conceded: { label: "Sixes Conceded", ...getQualificationFilterObject(true) },
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
    },
    team: {},
  };

  const options: { label: string; value: string }[] = [];

  for (const sortKey in qualificationFilters[statsType]) {
    const sortKeyInfo = qualificationFilters[statsType][sortKey];
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
