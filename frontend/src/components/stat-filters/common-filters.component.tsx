import { IStatsFilters } from "@/lib/types/filters-stats.types";
import { SelectCheckboxInput } from "./select-checkbox-input.component";
import { CheckboxOption } from "./checkbox-option.component";
import { RadioOption } from "./radio-option.component";
import { DateInput } from "./date-input.component";

export function CommonFilters({ data }: { data: IStatsFilters }) {
  const {
    primary_teams,
    opposition_teams,
    host_nations,
    continents,
    grounds,
    min_date,
    max_date,
    seasons,
    series,
    tournaments,
  } = data;

  return (
    <div className="flex flex-col gap-4">
      {primary_teams?.length && (
        <SelectCheckboxInput
          name="primary_team"
          label="Primary Team"
          options={primary_teams?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
        />
      )}

      {opposition_teams?.length && (
        <SelectCheckboxInput
          name="opposition_team"
          label="Opposition Team"
          options={opposition_teams?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
        />
      )}

      <CheckboxOption
        name="home_or_away"
        label="Home or Away"
        options={[
          { value: "home", label: "home" },
          { value: "away", label: "away" },
          { value: "neutral", label: "neutral" },
        ]}
      />

      {host_nations?.length && (
        <SelectCheckboxInput
          name="host_nation"
          label="Host Nation"
          options={host_nations?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
        />
      )}

      {continents?.length && (
        <SelectCheckboxInput
          name="continent"
          label="Continent"
          options={continents?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
        />
      )}

      {grounds?.length && (
        <SelectCheckboxInput
          name="ground"
          label="Ground"
          options={grounds?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
        />
      )}

      <DateInput
        name="min_start_date"
        label="Min. Start Date"
        defaultDate={min_date}
        minDate={min_date}
        maxDate={max_date}
      />

      <DateInput
        name="max_start_date"
        label="Max. Start Date"
        defaultDate={max_date}
        minDate={min_date}
        maxDate={max_date}
      />

      {seasons?.length && (
        <SelectCheckboxInput
          name="season"
          label="Season"
          options={seasons?.map((item) => {
            return { label: item, value: item };
          })}
        />
      )}

      {series?.length && (
        <SelectCheckboxInput
          name="series"
          label="Series"
          options={series?.map((item) => {
            return { label: `${item.name}, ${item.season}`, value: item.id.toString() };
          })}
        />
      )}

      {tournaments?.length && (
        <SelectCheckboxInput
          name="tournament"
          label="Tournament"
          options={tournaments?.map((item) => {
            return { label: item.name, value: item.id.toString() };
          })}
        />
      )}

      <CheckboxOption
        name="match_result"
        label="Match Result"
        options={[
          { value: "won", label: "won" },
          { value: "lost", label: "lost" },
          { value: "tied", label: "tied" },
          { value: "drawn", label: "drawn" },
          { value: "no result", label: "no result" },
        ]}
      />

      <RadioOption
        name="toss_result"
        label="Toss Result"
        defaultValue=""
        options={[
          { value: "won", label: "won" },
          { value: "lost", label: "lost" },
          { value: "", label: "either" },
        ]}
      />

      <RadioOption
        name="bat_field_first"
        label="Batting or Fielding First"
        defaultValue=""
        options={[
          { value: "bat", label: "bat" },
          { value: "field", label: "field" },
          { value: "", label: "either" },
        ]}
      />

      <CheckboxOption
        name="innings_number"
        label="Innings Number"
        options={[
          { value: "1", label: "1st Innings" },
          { value: "2", label: "2nd Innings" },
          { value: "3", label: "3rd Innings" },
          { value: "4", label: "4th Innings" },
        ]}
      />
    </div>
  );
}
