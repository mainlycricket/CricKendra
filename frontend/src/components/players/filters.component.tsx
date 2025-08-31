import { ReadonlyURLSearchParams } from "next/navigation";
import { SelectInput } from "../common/select-input.component";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Card, CardContent } from "../ui/card";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion";

export function PlayerFilters({ searchParams }: { searchParams: ReadonlyURLSearchParams }) {
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
      <SelectInput
        label="Gender"
        name="is_male"
        options={[
          { value: "true", label: "Male" },
          { value: "false", label: "Female" },
        ]}
        defaultValue={searchParams.get("is_male") || undefined}
      />
      <Input
        type="text"
        name="name__like"
        defaultValue={searchParams.get("name__like") || ""}
        placeholder="Player Name (type last name for better results)"
      />
      <Button type="submit">Search Players</Button>
    </form>
  );
}
