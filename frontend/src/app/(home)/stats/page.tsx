"use client";

import { CommonStatsLayout } from "@/components/stats/common-layout.component";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IStats } from "@/lib/types/common-stats.types";
import { IStatsFilters, prepareStatsFiltersMap, stringifyFiltersMap } from "@/lib/types/filters-stats.types";
import { useSearchParams } from "next/navigation";
import { Suspense, useEffect, useState } from "react";

export default function Stats() {
  return (
    <Suspense>
      <Component />
    </Suspense>
  );
}

function Component() {
  const searchParams = useSearchParams();
  const filterMap = prepareStatsFiltersMap(searchParams);
  const { playing_format, is_male, statsType, view, group } = filterMap;
  const query = stringifyFiltersMap(filterMap);

  const [filtersIdData, setFiltersIdData] = useState<IStatsFilters>();
  const [statsData, setStatsData] = useState<IStats>();

  useEffect(() => {
    doBackendRequest<null, IStatsFilters>({
      url: `/stats/filter-options?playing_format=${playing_format}&is_male=${is_male}`,
      method: "GET",
    }).then((data) => {
      setFiltersIdData(data.data!);
    });

    doBackendRequest<null, IStats>({
      url: `/stats/${statsType}/${view}/${group}?${query}`,
      method: "GET",
    }).then((data) => {
      setStatsData(data.data!);
    });
  }, [playing_format, is_male, statsType, view, group, query]);

  return filtersIdData && statsData ? (
    <CommonStatsLayout stats={statsData} filtersMap={filterMap} filtersIdData={filtersIdData} />
  ) : (
    <></>
  );
}
