import { doBackendRequest } from "@/lib/axiosFetch";
import { ISingleSeriesSingleSquad } from "@/lib/types/single-series.types";
import SingleSeriesLayout from "@/components/single-series/common/common-layout.component";
import { SingleSquadComponent } from "@/components/single-series/squads/single-squad.component";

export default async function SeriesSingleSquad({
  params,
}: {
  params: Promise<{ id: string; squadId: number }>;
}) {
  try {
    const { id, squadId } = await params;
    const response = await doBackendRequest<null, ISingleSeriesSingleSquad>({
      url: `/series/${id}/squads/${squadId}`,
      method: "GET",
    });
    const { series_header, squad_list, players } = response.data!;

    return (
      <SingleSeriesLayout active="squads" seriesHeader={series_header}>
        <SingleSquadComponent
          seriesId={series_header.series_id}
          activeSquadId={squadId}
          squad_list={squad_list || []}
          players={players || []}
        />
      </SingleSeriesLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
