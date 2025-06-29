import { doBackendRequest } from "@/lib/axiosFetch";
import { ISingleSeriesMatches } from "@/lib/types/single-series.types";
import SingleSeriesLayout from "@/components/single-series/common/common-layout.component";
import { MatchesComponent } from "@/components/single-series/matches/matches.component";

export default async function SeriesMatches({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, ISingleSeriesMatches>({
      url: `/series/${id}/matches`,
      method: "GET",
    });
    const { series_header, matches } = response.data!;

    return (
      <SingleSeriesLayout active="matches" seriesHeader={series_header}>
        <MatchesComponent matches={matches || []} />
      </SingleSeriesLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
