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
import { IStatsFiltersMap } from "@/lib/types/filters-stats.types";

export function QualificationFilters({
  options,
  filtersMap,
}: {
  options: { label: string; value: string }[];
  filtersMap: IStatsFiltersMap;
}) {
  const inititalUsed = options?.filter((option) => {
    const minKey = `min__${option.value}`,
      maxKey = `max__${option.value}`;
    if (filtersMap[minKey] || filtersMap[maxKey]) return { ...option };
  });
  const [usedFilters, setUsedFilters] = useState(inititalUsed?.length ? inititalUsed : [options?.[0]]);

  useEffect(() => {
    if (!inititalUsed?.length) {
      setUsedFilters([options?.[0]]);
    }
  }, [options]);

  function addNewUsedFilter() {
    const arr = [...usedFilters];
    for (const option of options) {
      const isUsed = usedFilters.find((item) => item.value === option.value);
      if (!isUsed) {
        arr.push({ ...option });
        break;
      }
    }
    setUsedFilters(arr);
  }

  function removeUsedFilter(value: string) {
    const updatedFilters = usedFilters.filter((option) => option.value !== value);
    setUsedFilters(updatedFilters);
  }

  function modifyUsedFiltersIdx(value: string, idx: number) {
    const updatedUsedOptions = [...usedFilters];

    const newOption = options.find((option) => option.value === value);
    if (newOption) {
      updatedUsedOptions[idx] = { ...newOption };
    }

    setUsedFilters(updatedUsedOptions);
  }

  return (
    <div className="hidden md:flex justify-between">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Result Qualifications</p>
        <div className="flex flex-col gap-2">
          {usedFilters?.length ? (
            usedFilters?.map((usedOption, idx) => {
              return (
                <div key={usedOption.value} className="flex gap-2">
                  <Select
                    defaultValue={usedOption.value}
                    onValueChange={(value) => {
                      modifyUsedFiltersIdx(value, idx);
                    }}
                  >
                    <SelectTrigger className="w-[180px]">
                      <SelectValue placeholder="Select" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectLabel className="text-sm">Qualification Filters</SelectLabel>
                        {options.map((option) => {
                          const isUsed = usedFilters.find((filter) => filter.value === option.value);
                          if (!isUsed || option.value === usedOption.value) {
                            return (
                              <SelectItem key={option.value} value={option.value}>
                                {option.label}
                              </SelectItem>
                            );
                          }
                        })}
                      </SelectGroup>
                    </SelectContent>
                  </Select>

                  <Input
                    type="number"
                    name={`min__${usedOption.value}`}
                    placeholder="Min"
                    min={1}
                    step={1}
                    className="w-24"
                    defaultValue={filtersMap[`min__${usedOption.value}`]}
                  />
                  <Input
                    type="number"
                    name={`max__${usedOption.value}`}
                    placeholder="Max"
                    min={1}
                    step={1}
                    className="w-24"
                    defaultValue={filtersMap[`max__${usedOption.value}`]}
                  />
                  <Button type="button" onClick={() => removeUsedFilter(usedOption.value)}>
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
          disabled={options?.length === usedFilters?.length}
          onClick={addNewUsedFilter}
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
    team: {
      matches_played: { label: "Matches Played", ...getQualificationFilterObject(true) },
      matches_won: { label: "Matches Won", ...getQualificationFilterObject(true) },
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
    },
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
