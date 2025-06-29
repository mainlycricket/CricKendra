import { doBackendRequest } from "@/lib/axiosFetch";
import { ISingleSeriesSquadsList } from "@/lib/types/single-series.types";
import SingleSeriesLayout from "@/components/single-series/common/common-layout.component";
import { SquadsComponent } from "@/components/single-series/squads/squads.component";

export default async function SeriesSquads({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, ISingleSeriesSquadsList>({
      url: `/series/${id}/squads-list`,
      method: "GET",
    });
    const { series_header, squad_list } = response.data!;

    return (
      <SingleSeriesLayout active="squads" seriesHeader={series_header}>
        <SquadsComponent
          seriesId={series_header.series_id}
          squad_list={squad_list || []}
        />
      </SingleSeriesLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
