import { InningsCommentary } from "@/components/single-match/commentary/innings-commentary";
import { InningsDropdown } from "@/components/single-match/commentary/innings-dropdown";
import { MatchHeader } from "@/components/single-match/common/match-header";
import { MatchTabs } from "@/components/single-match/common/match-tabs";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchCommentary } from "@/lib/types/single-match";

export default async function Commentary({ params }: { params: Promise<{ id: string; inningsId: string }> }) {
  try {
    const { id: matchId, inningsId } = await params;

    const response = await doBackendRequest<null, IMatchCommentary>({
      url: `/matches/${matchId}/innings/${inningsId}/commentary`,
      method: "GET",
    });

    const { match_header: matchHeader, commentary } = response.data!;
    matchHeader.innings_scores.sort((a, b) => a.innings_number - b.innings_number);

    return (
      <div className="flex flex-col gap-4">
        <MatchHeader matchHeader={matchHeader} />
        <MatchTabs
          matchId={matchHeader.match_id}
          inningsId={matchHeader?.innings_scores?.[matchHeader?.innings_scores?.length - 1]?.innings_id || -1}
          active="commentary"
        />
        <InningsDropdown matchId={matchId} inningsId={inningsId} inningsList={matchHeader.innings_scores} />
        <InningsCommentary
          battingTeamName={
            matchHeader?.innings_scores?.[matchHeader?.innings_scores?.length - 1]?.batting_team_name || ""
          }
          commentary={commentary}
          displayOverSummary={true}
        />
      </div>
    );
  } catch (error) {
    console.error(error);
  }
}
