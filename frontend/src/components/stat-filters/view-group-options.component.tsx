import { EnumStatsType, EnumStatsView } from "@/lib/types/enums.types";
import { useEffect, useState } from "react";
import { getSortOptions, SortFilters } from "./sort-filters.component";
import { RadioInput } from "./radio-input.component";
import { getQualificationOptions, QualificationFilters } from "./qualification-filters.component";
import { SelectInput } from "./select-input.component";
import { IStatsFiltersMap } from "@/lib/types/filters-stats.types";

export function ViewGroupOptions({
  statsType,
  filtersMap,
}: {
  statsType: EnumStatsType;
  filtersMap: IStatsFiltersMap;
}) {
  const [viewValue, setViewValue] = useState(filtersMap.view);
  const [groupOptions, setGroupOptions] = useState(getGroupOptions({ statsType, viewValue }));
  const [groupValue, setGroupValue] = useState(filtersMap?.group || groupOptions?.[0]?.value);

  const [qualificationOptions, setQualificationOptions] = useState(
    getQualificationOptions({ statsType, viewValue, groupValue })
  );
  const [sortOptions, setSortOptions] = useState(getSortOptions({ statsType, viewValue, groupValue }));

  useEffect(() => {
    const options = getGroupOptions({ statsType, viewValue });
    setGroupOptions([...options]);
  }, [statsType, viewValue]);

  useEffect(() => {
    setGroupValue(groupOptions?.[0]?.value);
  }, [groupOptions]);

  useEffect(() => {
    const qualificicationOptions = getQualificationOptions({ statsType, viewValue, groupValue });
    setQualificationOptions(qualificicationOptions);

    const sortOptions = getSortOptions({ statsType, viewValue, groupValue });
    setSortOptions(sortOptions);
  }, [statsType, viewValue, groupValue]);

  return (
    <div className="flex flex-col gap-4">
      <RadioInput
        name="view"
        label="View Format"
        defaultValue={viewValue}
        options={[
          { label: "overall", value: "overall" },
          { label: "individual", value: "individual" },
        ]}
        onValueChange={setViewValue as (value: string) => void}
      />

      <SelectInput name="group" label="group" options={groupOptions} onValueChange={setGroupValue} />

      <QualificationFilters options={qualificationOptions} filtersMap={filtersMap} />
      <SortFilters
        sortOptions={sortOptions}
        defaultSortKey={filtersMap?.sort_by}
        defaultSortOrder={filtersMap?.sort_order || "default"}
      />
    </div>
  );
}

function getGroupOptions({ statsType, viewValue }: { statsType: EnumStatsType; viewValue: EnumStatsView }) {
  const commonOptions = {
    overall: [
      { value: "matches", label: "Matches" },
      { value: "teams", label: "Teams" },
      { value: "grounds", label: "Grounds" },
      { value: "host-nations", label: "Host Nations" },
      { value: "continents", label: "Continents" },
      { value: "series", label: "Series" },
      { value: "tournaments", label: "Tournaments" },
      { value: "years", label: "Years" },
      { value: "seasons", label: "Seasons" },
      { value: "decades", label: "Decades" },
      { value: "aggregate", label: "Aggregate" },
    ],
    individual: [
      { value: "innings", label: "Innings" },
      { value: "match-totals", label: "Match Totals" },
      { value: "series", label: "Series" },
      { value: "tournaments", label: "Tournaments" },
      { value: "grounds", label: "Grounds" },
      { value: "host-nations", label: "Host Nations" },
      { value: "years", label: "Years" },
      { value: "seasons", label: "Seasons" },
    ],
  };

  const finalOptions = {
    batting: {
      overall: [
        { value: "batters", label: "Batters" },
        { value: "team-innings", label: "Team Innings" },
        { value: "oppositions", label: "Oppositions" },
        ...commonOptions["overall"],
      ],
      individual: [...commonOptions["individual"], { value: "oppositions", label: "Oppositions" }],
    },
    bowling: {
      overall: [
        { value: "bowlers", label: "Bowlers" },
        { value: "team-innings", label: "Team Innings" },
        { value: "oppositions", label: "Oppositions" },
        ...commonOptions["overall"],
      ],
      individual: [...commonOptions["individual"], { value: "oppositions", label: "Oppositions" }],
    },
    team: {
      overall: [{ value: "players", label: "Players" }, ...commonOptions["overall"]],
      individual: [...commonOptions["individual"], { value: "match-results", label: "Match Results" }],
    },
  };

  return finalOptions[statsType][viewValue];
}
