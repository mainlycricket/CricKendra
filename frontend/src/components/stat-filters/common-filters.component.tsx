import { IStatsFilters } from "@/lib/types/filters-stats.types";
import { StatOption } from "./options.component";
import { DateInput } from "../common/date.component";
import { Checkbox } from "../ui/checkbox";
import { Label } from "../ui/label";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";

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
        <StatOption
          name="primary_team"
          label="Primary Team"
          optionValues={primary_teams.map((item) => item.id.toString())}
          optionLabels={primary_teams.map((item) => item.name)}
        />
      )}

      {opposition_teams?.length && (
        <StatOption
          name="opposition_team"
          label="Opposition Team"
          optionValues={opposition_teams.map((item) => item.id.toString())}
          optionLabels={opposition_teams.map((item) => item.name)}
        />
      )}

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Home or Away</p>
        <div className="flex flex-wrap gap-4">
          <div className="flex items-center gap-2">
            <Checkbox id="home" name="home_or_away" value="home" />
            <Label htmlFor="home" className="font-normal">
              Home
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="away" name="home_or_away" value="away" />
            <Label htmlFor="away" className="font-normal">
              Away
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="neutral" name="home_or_away" value="neutral" />
            <Label htmlFor="neutral" className="font-normal">
              Neutral
            </Label>
          </div>
        </div>
      </div>

      {host_nations?.length && (
        <StatOption
          name="host_nation"
          label="Host Nation"
          optionValues={host_nations.map((item) => item.id.toString())}
          optionLabels={host_nations.map((item) => item.name)}
        />
      )}

      {continents?.length && (
        <StatOption
          name="continent"
          label="Continent"
          optionValues={continents.map((item) => item.id.toString())}
          optionLabels={continents.map((item) => item.name)}
        />
      )}

      {grounds?.length && (
        <StatOption
          name="ground"
          label="Ground"
          optionValues={grounds.map((item) => item.id.toString())}
          optionLabels={grounds.map((item) => `${item.name}${item.city_name ? `, ${item.city_name}` : ""}`)}
        />
      )}

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Min. Start Date</p>
        <input
          type="date"
          defaultValue={min_date}
          min={min_date}
          max={max_date}
          name="min_start_date"
          className="dark:bg-input/30 dark:hover:bg-input/50  border-1 rounded-lg w-[180px] px-2 h-9"
        />
        {/* <DateInput defaultValue={min_date} name="min_start_date" /> */}
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Max. Start Date</p>
        <input
          type="date"
          defaultValue={max_date}
          min={min_date}
          max={max_date}
          name="max_start_date"
          className="dark:bg-input/30 dark:hover:bg-input/50  border-1 rounded-lg w-[180px] px-2 h-9"
        />
        {/* <DateInput defaultValue={max_date} name="max_start_date" /> */}
      </div>

      {seasons?.length && (
        <StatOption name="season" label="Season" optionValues={seasons} optionLabels={seasons} />
      )}

      {series?.length && (
        <StatOption
          name="series"
          label="Series"
          optionValues={series.map((item) => item.id.toString())}
          optionLabels={series.map((item) => `${item.name} ${item.season}`)}
        />
      )}

      {tournaments?.length && (
        <StatOption
          name="tournament"
          label="Tournament"
          optionValues={tournaments.map((item) => item.id.toString())}
          optionLabels={tournaments.map((item) => item.name)}
        />
      )}

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Match Result</p>
        <div className="flex flex-wrap gap-4">
          <div className="flex items-center gap-2">
            <Checkbox id="won" name="match_result" value="won" />
            <Label className="font-normal" htmlFor="won">
              Won
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="lost" name="match_result" value="lost" />
            <Label className="font-normal" htmlFor="lost">
              Lost
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="tied" name="match_result" value="tied" />
            <Label className="font-normal" htmlFor="tied">
              Tied
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="drawn" name="match_result" value="drawn" />
            <Label className="font-normal" htmlFor="drawn">
              Drawn
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="no result" name="match_result" value="no result" />
            <Label className="font-normal" htmlFor="no result">
              No Result
            </Label>
          </div>
        </div>
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Toss Result</p>
        <RadioGroup defaultValue="" name="toss_result" className="flex flex-wrap gap-4">
          <div className="flex items-center gap-1">
            <RadioGroupItem value="won" id="won" />
            <Label htmlFor="won" className="capitalize font-normal">
              won
            </Label>
          </div>
          <div className="flex items-center gap-1">
            <RadioGroupItem value="lost" id="lost" />
            <Label htmlFor="lost" className="capitalize font-normal">
              lost
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

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Batting or Fielding First</p>
        <RadioGroup defaultValue="" name="bat_field_first" className="flex flex-wrap gap-4">
          <div className="flex items-center gap-1">
            <RadioGroupItem value="bat" id="bat" />
            <Label htmlFor="bat" className="capitalize font-normal">
              bat
            </Label>
          </div>
          <div className="flex items-center gap-1">
            <RadioGroupItem value="field" id="field" />
            <Label htmlFor="field" className="capitalize font-normal">
              field
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

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Innings Number</p>
        <div className="flex flex-wrap gap-4">
          <div className="flex items-center gap-2">
            <Checkbox id="1" name="innings_number" value="1" />
            <Label htmlFor="1" className="font-normal">
              1st Innings
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="2" name="innings_number" value="2" />
            <Label htmlFor="2" className="font-normal">
              2nd Innings
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="3" name="innings_number" value="3" />
            <Label htmlFor="3" className="font-normal">
              3rd Innings
            </Label>
          </div>
          <div className="flex items-center gap-2">
            <Checkbox id="4" name="innings_number" value="4" />
            <Label htmlFor="4" className="font-normal">
              4th Innings
            </Label>
          </div>
        </div>
      </div>
    </div>
  );
}
