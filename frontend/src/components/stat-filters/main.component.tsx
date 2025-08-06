"use client";

import { IStatsFilters } from "@/lib/types/filters-stats.types";
import { CommonFilters } from "./common-filters.component";
import { Button } from "../ui/button";
import { FormEvent, useState } from "react";

import { useRouter } from "next/navigation";
import { EnumPlayingFormat, EnumStatsType, EnumStatsView } from "@/lib/types/enums.types";
import { BattingFilters } from "./batting-filters.component";
import { BowlingFilters } from "./bowling-filters.component";
import { PlayingFormatOptions } from "./playing-formats.component";
import { StatsTypeGenderDropdowns } from "./stats-type-gender.component";
import { PaginationInputs } from "./pagination-input.component";
import { ViewGroupOptions } from "./view-group-options.component";

export function StatsFiltersComponent({
  commonFilterOptions,
  defaultPlayingFormat,
  defaultStatsType,
  defaultIsMale,
  defaultView,
}: {
  commonFilterOptions: IStatsFilters;
  defaultPlayingFormat: EnumPlayingFormat;
  defaultStatsType: EnumStatsType;
  defaultIsMale: "true" | "false";
  defaultView: EnumStatsView;
}) {
  const [statsTypeValue, setStatsTypeValue] = useState(defaultStatsType || "batting");
  const router = useRouter();

  function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const keys = formData.keys().toArray();
    const filters = [`playing_format=${defaultPlayingFormat}`];
    const done = new Set<string>();
    for (const key of keys) {
      if (done.has(key)) continue;
      const values = formData.getAll(key);
      for (const value of values) {
        if (value) {
          filters.push(`${key}=${value}`);
        }
      }
      done.add(key);
    }

    const url = `/stats?${filters.join("&")}`;
    router.push(url);
  }

  return (
    <div className="flex flex-col gap-4">
      <PlayingFormatOptions selectedPlayingFormat={defaultPlayingFormat} />
      <form onSubmit={(e) => handleSubmit(e)} className="px-2 flex flex-col gap-4">
        <StatsTypeGenderDropdowns
          defaultIsMale={defaultIsMale}
          defaultStatsType={statsTypeValue}
          setStatsTypeValue={setStatsTypeValue}
        />
        <CommonFilters data={commonFilterOptions} />
        {statsTypeValue === "batting" ? (
          <BattingFilters />
        ) : statsTypeValue === "bowling" ? (
          <BowlingFilters />
        ) : (
          <></>
        )}
        <ViewGroupOptions statsType={statsTypeValue} defaultView={defaultView} />
        <PaginationInputs defaultPage={1} defaultLimit={50} />
        <Button type="submit" className="w-24">
          Submit
        </Button>
      </form>
    </div>
  );
}
