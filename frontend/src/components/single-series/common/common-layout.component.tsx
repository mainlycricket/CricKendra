import { SeriesTabs } from "./series-tabs.component";
import { TopBatters } from "./top-batters.component";
import { TopBowlers } from "./top-bowlers.component";

import { ISeriesHeader } from "@/lib/types/single-series.types";

export default function SingleSeriesLayout({
  children,
  seriesHeader,
  active,
}: Readonly<{ children: React.ReactNode; seriesHeader: ISeriesHeader; active: string }>) {
  return (
    <div className="flex flex-col">
      <p className="md:hidden px-2 font-bold tracking-wide">
        {seriesHeader.series_name} {seriesHeader.season}
      </p>

      <SeriesTabs
        seriesId={seriesHeader.series_id}
        seriesName={seriesHeader.series_name}
        seriesSeason={seriesHeader.season}
        active={active}
      />

      <div className="flex justify-between gap-4 mt-4">
        {children}

        <div className="hidden w-1/4 md:flex md:flex-col md:gap-4">
          {seriesHeader?.top_batters?.length && <TopBatters topBatters={seriesHeader.top_batters} />}
          {seriesHeader?.top_bowlers?.length && <TopBowlers topBowlers={seriesHeader.top_bowlers} />}
        </div>
      </div>
    </div>
  );
}
