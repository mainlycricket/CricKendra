import { CommonMatchLayout } from "@/components/single-match/common/common-layout.component";
import { TeamSquad } from "@/components/single-match/squads/team-squad";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchSquad } from "@/lib/types/single-match.types";

export default async function Squads({ params }: { params: Promise<{ id: string }> }) {
  try {
    const { id } = await params;

    const response = await doBackendRequest<null, IMatchSquad>({
      url: `/matches/${id}/squads`,
      method: "GET",
    });
    const { match_header, team_squads } = response.data!;

    match_header?.innings_scores?.sort((a, b) => a.innings_number - b.innings_number);

    return (
      <CommonMatchLayout matchHeader={match_header} active="squads">
        {team_squads?.length && (
          <div className="flex justify-center px-4">
            {team_squads?.map((squad) => (
              <div className="w-full md:w-[40%] text-center" key={squad.team_id}>
                <h3 className="py-2 border-1 text-xl font-medium">{squad.team_name}</h3>
                <TeamSquad
                  players={squad?.players?.filter((player) => player.playing_status === "playing_xi") || []}
                />
              </div>
            ))}
          </div>
        )}
      </CommonMatchLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
