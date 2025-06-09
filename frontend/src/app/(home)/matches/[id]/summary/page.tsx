import { InningsCommentary } from "@/components/single-match/commentary/innings-commentary";
import { MatchHeader } from "@/components/single-match/common/match-header";
import { MatchTabs } from "@/components/single-match/common/match-tabs";
import { SummaryScorecard } from "@/components/single-match/summary/summary-scorecard";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchSummary } from "@/lib/types/single-match";
import Link from "next/link";

export default async function Summary({ params }: { params: Promise<{ id: string }> }) {
  try {
    const { id } = await params;

    const response = await doBackendRequest<null, IMatchSummary>({
      url: `/matches/${id}/summary`,
      method: "GET",
    });
    const {
      match_header: matchHeader,
      scorecard_summary: scorecardSummary,
      latest_commentary: commentary,
    } = response.data!;

    matchHeader.innings_scores.sort((a, b) => a.innings_number - b.innings_number);

    return (
      <div className="flex flex-col gap-4">
        <MatchHeader matchHeader={matchHeader} />
        <MatchTabs
          matchId={matchHeader.match_id}
          inningsId={matchHeader?.innings_scores?.[matchHeader?.innings_scores?.length - 1]?.innings_id || -1}
          active="summary"
        />

        <div className="w-full">
          <h3 className="bg-secondary px-4 py-2 tracking-wider">Scorecard Summary</h3>
          {scorecardSummary.map((innings) => {
            return <SummaryScorecard key={innings.innings_id} entry={innings} />;
          })}
          <div className="mt-4 text-blue-400 text-center hover:underline">
            <Link href={`/matches/${id}/scorecard`}>View Full Scorecard</Link>
          </div>
        </div>

        {matchHeader?.innings_scores?.toReversed().map((innings) => {
          const inningsCommentary = commentary.filter((item) => item.innings_id === innings.innings_id);
          return (
            <InningsCommentary
              key={innings.innings_id}
              commentary={inningsCommentary}
              battingTeamName={innings.batting_team_name}
              displayOverSummary={false}
            />
          );
        })}
      </div>
    );
  } catch (error) {
    console.error(error);
  }
}
