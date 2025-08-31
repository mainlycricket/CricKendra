import { ReadonlyURLSearchParams } from "next/navigation";
import { SelectSeasons } from "../common/season-select.component";
import { SelectInput } from "../common/select-input.component";
import { SelectTeams } from "../common/teams-select.component";
import { Button } from "../ui/button";
import { Card, CardContent } from "../ui/card";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion";
import { DateInput } from "../common/date-input.component";

export function MatchFilters({ searchParams }: { searchParams: ReadonlyURLSearchParams }) {
  return (
    <Card className="p-0">
      <CardContent>
        <Accordion type="single" className="p-0" defaultValue="item-1" collapsible>
          <AccordionItem value="item-1">
            <AccordionTrigger className="text-xl font-bold">Filters</AccordionTrigger>
            <AccordionContent>
              <FilterForm searchParams={searchParams} />
            </AccordionContent>
          </AccordionItem>
        </Accordion>
      </CardContent>
    </Card>
  );
}

function FilterForm({ searchParams }: { searchParams: ReadonlyURLSearchParams }) {
  return (
    <form className="flex flex-col gap-4">
      <div className="flex flex-col md:flex-row gap-4">
        <SelectInput
          label="Gender"
          name="is_male"
          options={[
            { value: "true", label: "Male" },
            { value: "false", label: "Female" },
          ]}
          defaultValue={searchParams.get("is_male") || undefined}
        />
        <SelectInput
          label="Playing Format"
          name="playing_format"
          options={[
            { value: "Test", label: "Test" },
            { value: "ODI", label: "ODI" },
            { value: "T20I", label: "T20I" },
            { value: "first_class", label: "First Class" },
            { value: "list_a", label: "List A" },
            { value: "T20", label: "T20" },
          ]}
          defaultValue={searchParams.get("playing_format") || undefined}
        />
      </div>
      <div className="flex flex-col md:flex-row gap-4">
        <DateInput
          label="Min. Start Date"
          name="start_date__min"
          defaultValue={searchParams.get("start_date__min") || undefined}
        />
        <DateInput
          label="Max. Start Date"
          name="start_date__max"
          defaultValue={searchParams.get("start_date__max") || undefined}
        />
      </div>
      <SelectTeams name="teams_id" defaultSelected={searchParams.getAll("teams_id") || []} />
      <SelectSeasons defaultSelected={searchParams.getAll("season") || []} />
      <Button type="submit">Search Matches</Button>
    </form>
  );
}
