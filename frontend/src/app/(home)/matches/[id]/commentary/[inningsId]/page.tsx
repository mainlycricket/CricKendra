import { InningsCommentary } from "@/components/single-match/commentary/innings-commentary";
import { InningsDropdown } from "@/components/single-match/commentary/innings-dropdown";
import { CommonMatchLayout } from "@/components/single-match/common/common-layout.component";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchCommentary } from "@/lib/types/single-match.types";

export default async function Commentary({ params }: { params: Promise<{ id: string; inningsId: string }> }) {
  try {
    const { id: matchId, inningsId } = await params;

    const response = await doBackendRequest<null, IMatchCommentary>({
      url: `/matches/${matchId}/innings/${inningsId}/commentary`,
      method: "GET",
    });

    const { match_header, commentary } = response.data!;
    match_header?.innings_scores?.sort((a, b) => a.innings_number - b.innings_number);

    const battingTeamName =
      match_header?.innings_scores?.find((innings) => innings.innings_id.toString() === inningsId)
        ?.batting_team_name || "";

    return (
      <CommonMatchLayout matchHeader={match_header} active="commentary">
        <InningsDropdown
          matchId={matchId}
          inningsId={inningsId}
          inningsList={match_header?.innings_scores || []}
        />
        <InningsCommentary
          battingTeamName={battingTeamName}
          commentary={commentary || []}
          displayOverSummary={true}
        />
      </CommonMatchLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
