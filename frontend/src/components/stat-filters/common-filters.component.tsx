import { IStatsFilters, IStatsFiltersMap } from "@/lib/types/filters-stats.types";
import { SelectCheckboxInput } from "./select-checkbox-input.component";
import { CheckboxInput } from "./checkbox-input.component";
import { RadioInput } from "./radio-input.component";
import { DateInput } from "./date-input.component";

export function CommonFilters({
  filtersData,
  filtersMap,
}: {
  filtersData: IStatsFilters;
  filtersMap: IStatsFiltersMap;
}) {
  return (
    <div className="flex flex-col gap-4">
      {filtersData?.teams?.length && (
        <SelectCheckboxInput
          name="primary_team"
          label="Primary Team"
          options={filtersData?.teams?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
          defaultValues={filtersMap?.primary_team}
        />
      )}

      {filtersData?.teams?.length && (
        <SelectCheckboxInput
          name="opposition_team"
          label="Opposition Team"
          options={filtersData?.teams?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
          defaultValues={filtersMap?.opposition_team}
        />
      )}

      <CheckboxInput
        name="home_or_away"
        label="Home or Away"
        options={[
          { value: "home", label: "home" },
          { value: "away", label: "away" },
          { value: "neutral", label: "neutral" },
        ]}
        defaultValues={filtersMap?.home_or_away}
      />

      {filtersData?.host_nations?.length && (
        <SelectCheckboxInput
          name="host_nation"
          label="Host Nation"
          options={filtersData?.host_nations?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
          defaultValues={filtersMap?.host_nation}
        />
      )}

      {filtersData?.continents?.length && (
        <SelectCheckboxInput
          name="continent"
          label="Continent"
          options={filtersData?.continents?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
          defaultValues={filtersMap?.continent}
        />
      )}

      {filtersData?.grounds?.length && (
        <SelectCheckboxInput
          name="ground"
          label="Ground"
          options={filtersData?.grounds?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
          defaultValues={filtersMap?.ground}
        />
      )}

      <DateInput
        name="min_start_date"
        label="Min. Start Date"
        defaultDate={filtersMap?.min_start_date}
        minDate={filtersData?.min_date}
        maxDate={filtersData?.max_date}
      />

      <DateInput
        name="max_start_date"
        label="Max. Start Date"
        defaultDate={filtersMap?.max_start_date}
        minDate={filtersData?.min_date}
        maxDate={filtersData?.max_date}
      />

      {filtersData?.seasons?.length && (
        <SelectCheckboxInput
          name="season"
          label="Season"
          options={filtersData?.seasons?.map((item) => {
            return { label: item, value: item };
          })}
          defaultValues={filtersMap?.season}
        />
      )}

      {filtersData?.series?.length && (
        <SelectCheckboxInput
          name="series"
          label="Series"
          options={filtersData?.series?.map((item) => {
            return { label: `${item.name}, ${item.season}`, value: item.id.toString() };
          })}
          defaultValues={filtersMap?.series}
        />
      )}

      {filtersData?.tournaments?.length && (
        <SelectCheckboxInput
          name="tournament"
          label="Tournament"
          options={filtersData?.tournaments?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
          defaultValues={filtersMap?.tournament}
        />
      )}

      <CheckboxInput
        name="match_result"
        label="Match Result"
        options={[
          { value: "won", label: "won" },
          { value: "lost", label: "lost" },
          { value: "tied", label: "tied" },
          { value: "drawn", label: "drawn" },
          { value: "no result", label: "no result" },
        ]}
        defaultValues={filtersMap?.match_result}
      />

      <RadioInput
        name="toss_result"
        label="Toss Result"
        defaultValue={filtersMap?.toss_result || ""}
        options={[
          { value: "won", label: "won" },
          { value: "lost", label: "lost" },
          { value: "", label: "either" },
        ]}
      />

      <RadioInput
        name="bat_field_first"
        label="Batting or Fielding First"
        defaultValue={filtersMap?.bat_field_first || ""}
        options={[
          { value: "bat", label: "bat" },
          { value: "field", label: "field" },
          { value: "", label: "either" },
        ]}
      />

      <CheckboxInput
        name="innings_number"
        label="Innings Number"
        options={[
          { value: "1", label: "1st Innings" },
          { value: "2", label: "2nd Innings" },
          { value: "3", label: "3rd Innings" },
          { value: "4", label: "4th Innings" },
        ]}
        defaultValues={filtersMap?.innings_number}
      />
    </div>
  );
}
