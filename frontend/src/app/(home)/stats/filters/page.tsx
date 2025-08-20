"use client";

import { StatsFiltersComponent } from "@/components/stat-filters/main.component";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IStatsFilters, prepareStatsFiltersMap } from "@/lib/types/filters-stats.types";
import { useSearchParams } from "next/navigation";
import { Suspense, useEffect, useState } from "react";

export default function StatsFilters() {
  return (
    <Suspense>
      <Component />
    </Suspense>
  );
}

function Component() {
  const searchParams = useSearchParams();
  const filtersMap = prepareStatsFiltersMap(searchParams);
  const { playing_format, is_male } = filtersMap;

  const [filtersData, setFiltersData] = useState<IStatsFilters>();

  useEffect(() => {
    doBackendRequest<null, IStatsFilters>({
      url: `/stats/filter-options?is_male=${is_male}&playing_format=${playing_format}`,
      method: "GET",
    })
      .then((data) => {
        setFiltersData(data.data!);
      })
      .catch((error) => {
        console.error(error);
      });
  }, [is_male, playing_format]);

  return filtersData ? <StatsFiltersComponent filtersMap={filtersMap} filtersData={filtersData} /> : <></>;
}
