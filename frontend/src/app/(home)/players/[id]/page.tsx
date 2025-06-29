import { SinglePlayerLayout } from "@/components/single-player/common/common-layout.component";
import { PlayerBasicDetails } from "@/components/single-player/overview/basic-details";
import { PlayerCareerStats } from "@/components/single-player/overview/career-stats";
import { Card, CardContent } from "@/components/ui/card";
import { doBackendRequest } from "@/lib/axiosFetch";
import { ISinglePlayer } from "@/lib/types/single-player.types";

export default async function SinlgePlayerOverview({ params }: { params: Promise<{ id: number }> }) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, ISinglePlayer>({
      url: `/players/${id}`,
      method: "GET",
    });

    const { data: player } = response;

    return (
      <div>
        <SinglePlayerLayout player={player!} active="overview">
          <div className="flex flex-col md:flex-col-reverse gap-4">
            <Card className="md:hidden">
              <CardContent>
                <PlayerBasicDetails player={player!} />
              </CardContent>
            </Card>
            <PlayerCareerStats player={player!} />
          </div>
        </SinglePlayerLayout>
      </div>
    );
  } catch (error) {
    console.error(error);
  }
}
