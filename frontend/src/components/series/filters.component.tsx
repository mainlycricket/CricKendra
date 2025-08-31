import { ReadonlyURLSearchParams } from "next/navigation";
import { SelectSeasons } from "../common/season-select.component";
import { SelectInput } from "../common/select-input.component";
import { SelectTeams } from "../common/teams-select.component";
import { SelectTournaments } from "../common/tournament-select-component";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Card, CardContent } from "../ui/card";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion";
import { DateInput } from "../common/date-input.component";

export function SeriesFilters({ searchParams }: { searchParams: ReadonlyURLSearchParams }) {
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
          label="Playing Level"
          name="playing_level"
          options={[
            { value: "international", label: "International" },
            { value: "domestic", label: "Domestic" },
          ]}
          defaultValue={searchParams.get("playing_level") || undefined}
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
      <Input
        type="text"
        name="name__like"
        placeholder="Series Name"
        defaultValue={searchParams.get("name__like") || ""}
      />
      <SelectTeams name="teams_id__all" defaultSelected={searchParams.getAll("teams_id__all") || []} />
      <SelectSeasons defaultSelected={searchParams.getAll("season") || []} />
      <SelectTournaments defaultSelected={searchParams.getAll("tournament_id") || []} />
      <Button type="submit">Search Series</Button>
    </form>
  );
}
