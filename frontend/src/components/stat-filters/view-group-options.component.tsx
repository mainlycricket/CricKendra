import { EnumStatsType, EnumStatsView } from "@/lib/types/enums.types";
import { useEffect, useState } from "react";
import { getSortOptions, SortFilters } from "./sort-filters.component";
import { RadioOption } from "./radio-option.component";
import { getQualificationOptions, QualificationFilters } from "./qualification-filters.component";
import { SelectInput } from "./select-input.component";

export function ViewGroupOptions({
  statsType,
  defaultView,
}: {
  statsType: EnumStatsType;
  defaultView: EnumStatsView;
}) {
  const [viewValue, setViewValue] = useState(defaultView);
  const [groupOptions, setGroupOptions] = useState(getGroupOptions({ statsType, viewValue }));
  const [groupValue, setGroupValue] = useState(groupOptions?.[0]?.value);

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
  }, [groupValue]);

  return (
    <div className="flex flex-col gap-4">
      <RadioOption
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

      <QualificationFilters options={qualificationOptions} />
      <SortFilters sortOptions={sortOptions} />
    </div>
  );
}

function getGroupOptions({ statsType, viewValue }: { statsType: EnumStatsType; viewValue: EnumStatsView }) {
  const commonOptions = {
    overall: [
      { value: "team-innings", label: "Team Innings" },
      { value: "matches", label: "Matches" },
      { value: "teams", label: "Teams" },
      { value: "oppositions", label: "Oppositions" },
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
      { value: "oppositions", label: "Oppositions" },
      { value: "years", label: "Years" },
      { value: "seasons", label: "Seasons" },
    ],
  };

  const finalOptions = {
    batting: {
      overall: [{ value: "batters", label: "Batters" }, ...commonOptions["overall"]],
      individual: [...commonOptions["individual"]],
    },
    bowling: {
      overall: [{ value: "bowlers", label: "Bowlers" }, ...commonOptions["overall"]],
      individual: [...commonOptions["individual"]],
    },
    team: {
      overall: [],
      individual: [],
    },
  };

  return finalOptions[statsType][viewValue];
}
