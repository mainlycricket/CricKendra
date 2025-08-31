"use client";

import * as React from "react";
import { ChevronDownIcon } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { formatDateForApi } from "@/lib/utils";

export function DateInput({
  label,
  name,
  defaultValue,
}: {
  label: string;
  name: string;
  defaultValue?: string;
}) {
  const [open, setOpen] = React.useState(false);
  const [date, setDate] = React.useState<Date | undefined>(defaultValue ? new Date(defaultValue) : undefined);
  const [strDate, setStrDate] = React.useState("");

  return (
    <div className="w-full flex flex-col gap-3">
      <input type="text" name={name} value={strDate} onChange={() => {}} className="hidden" />
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button variant="outline" id="date" className="justify-between font-normal">
            {label}
            {date ? `: ${date?.toLocaleDateString()}` : ""}
            <ChevronDownIcon />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-auto overflow-hidden p-0" align="start">
          <Calendar
            mode="single"
            selected={date}
            captionLayout="dropdown"
            className="p-4 w-full rounded-md border shadow-sm"
            onSelect={(date) => {
              setDate(date);
              setOpen(false);
              setStrDate(date ? formatDateForApi(date) : "");
            }}
          />
        </PopoverContent>
      </Popover>
    </div>
  );
}
