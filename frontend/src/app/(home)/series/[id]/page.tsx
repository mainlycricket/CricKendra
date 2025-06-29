import { doBackendRequest } from "@/lib/axiosFetch";
import { ISingleSeriesOverview } from "@/lib/types/single-series.types";
import SingleSeriesLayout from "@/components/single-series/common/common-layout.component";
import { OverviewComponent } from "@/components/single-series/overview/overview.component";

export default async function SeriesOverview({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, ISingleSeriesOverview>({
      url: `/series/${id}`,
      method: "GET",
    });
    const { series_header, fixture_matches, result_matches } = response.data!;

    return (
      <SingleSeriesLayout active="overview" seriesHeader={series_header}>
        <OverviewComponent
          series_header={series_header}
          fixture_matches={fixture_matches || []}
          result_matches={result_matches || []}
        />
      </SingleSeriesLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
