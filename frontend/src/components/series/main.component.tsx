import { IAllSeriesResponse } from "@/lib/types/series.types";
import { SingleSeries } from "./single-series.component";
import { SeriesFilters } from "./filters.component";
import { ReadonlyURLSearchParams } from "next/navigation";
import { Pagination } from "../common/pagination.component";

export function MainSeriesLayout({
  data,
  searchParams,
}: {
  data: IAllSeriesResponse;
  searchParams: ReadonlyURLSearchParams;
}) {
  return (
    <div>
      <div className="flex flex-col gap-6">
        <SeriesFilters searchParams={searchParams} />
        {data?.series?.length ? (
          <div className="flex flex-col gap-2">
            {data?.series?.map((entry) => (
              <SingleSeries series={entry} key={entry.id} />
            ))}
            <Pagination
              mainPageLink="series"
              recordsCount={data?.series?.length || 0}
              currentPage={Number(searchParams.get("__page")) || 1}
              disableNext={!data?.next}
            />
          </div>
        ) : (
          <div>No series found...</div>
        )}
      </div>
    </div>
  );
}
