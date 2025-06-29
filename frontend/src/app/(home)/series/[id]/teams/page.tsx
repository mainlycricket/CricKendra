import { doBackendRequest } from "@/lib/axiosFetch";
import { ISingleSeriesTeams } from "@/lib/types/single-series.types";
import SingleSeriesLayout from "@/components/single-series/common/common-layout.component";
import { TeamsComponent } from "@/components/single-series/teams/teams.component";

export default async function SeriesTeams({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, ISingleSeriesTeams>({
      url: `/series/${id}/teams`,
      method: "GET",
    });
    const { series_header, teams } = response.data!;

    return (
      <SingleSeriesLayout active="teams" seriesHeader={series_header}>
        <TeamsComponent teams={teams || []} />
      </SingleSeriesLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
