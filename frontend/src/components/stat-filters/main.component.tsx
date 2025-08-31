"use client";

import { IStatsFilters, IStatsFiltersMap } from "@/lib/types/filters-stats.types";
import { CommonFilters } from "./common-filters.component";
import { Button } from "../ui/button";
import { FormEvent, useState } from "react";

import { useRouter } from "next/navigation";
import { BattingFilters } from "./batting-filters.component";
import { BowlingFilters } from "./bowling-filters.component";
import { PlayingFormatOptions } from "./playing-formats.component";
import { StatsTypeGenderDropdowns } from "./stats-type-gender.component";
import { PaginationInputs } from "./pagination-input.component";
import { ViewGroupOptions } from "./view-group-options.component";
import { TeamFilters } from "./team-filters.component";

export function StatsFiltersComponent({
  filtersMap,
  filtersData,
}: {
  filtersMap: IStatsFiltersMap;
  filtersData: IStatsFilters;
}) {
  const [statsTypeValue, setStatsTypeValue] = useState(filtersMap.statsType);
  const router = useRouter();

  function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const keys = formData.keys().toArray();
    const filters = [`playing_format=${filtersMap.playing_format}`];
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
      <PlayingFormatOptions
        selectedPlayingFormat={filtersMap.playing_format}
        isMale={filtersMap.is_male || "true"}
      />
      <form onSubmit={(e) => handleSubmit(e)} className="px-2 flex flex-col gap-4">
        <StatsTypeGenderDropdowns
          playingFormat={filtersMap.playing_format || "ODI"}
          defaultIsMale={filtersMap.is_male}
          defaultStatsType={statsTypeValue}
          setStatsTypeValue={setStatsTypeValue}
        />
        <CommonFilters filtersData={filtersData} filtersMap={filtersMap} />
        {statsTypeValue === "batting" ? (
          <BattingFilters filtersMap={filtersMap} />
        ) : statsTypeValue === "bowling" ? (
          <BowlingFilters filtersMap={filtersMap} />
        ) : statsTypeValue === "team" ? (
          <TeamFilters filtersMap={filtersMap} />
        ) : (
          <></>
        )}
        <ViewGroupOptions statsType={statsTypeValue} filtersMap={filtersMap} />
        <PaginationInputs
          defaultPage={filtersMap?.__page || "1"}
          defaultLimit={filtersMap?.__limit || "50"}
        />
        <Button type="submit" className="w-24">
          Submit
        </Button>
      </form>
    </div>
  );
}
