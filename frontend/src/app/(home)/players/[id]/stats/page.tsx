import { SinglePlayerLayout } from "@/components/single-player/common/common-layout.component";
import { MainStats } from "@/components/single-player/stats/main-stats.component";
import { doBackendRequest } from "@/lib/axiosFetch";
import { ISinglePlayer } from "@/lib/types/single-player.types";

export default async function SinlgePlayerStats({ params }: { params: Promise<{ id: number }> }) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, ISinglePlayer>({
      url: `/players/${id}`,
      method: "GET",
    });

    const { data: player } = response;

    return (
      <div>
        <SinglePlayerLayout player={player!} active="stats">
          <MainStats player={player!} />
        </SinglePlayerLayout>
      </div>
    );
  } catch (error) {
    console.error(error);
  }
}
