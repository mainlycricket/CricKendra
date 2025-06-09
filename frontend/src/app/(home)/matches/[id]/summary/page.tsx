import { InningsCommentary } from "@/components/single-match/commentary/innings-commentary";
import { MatchHeader } from "@/components/single-match/common/match-header";
import { MatchTabs } from "@/components/single-match/common/match-tabs";
import { SummaryScorecard } from "@/components/single-match/summary/summary-scorecard";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchSummary } from "@/lib/types/single-match";
import Image from "next/image";
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

        {matchHeader.player_awards.length && (
          <div className="block md:hidden border-y-2">
            {matchHeader.player_awards.map((entry) => {
              return (
                <div key={`${entry.player_id}_${entry.award_type}`} className="flex flex-col gap-2">
                  <p className="capitalize text-lg font-thin px-4 pt-2">{entry.award_type.split("_").join(" ")}</p>
                  <div className="flex gap-4 px-4 pb-4">
                    <Image
                      src={`/file.svg`}
                      width={50}
                      height={50}
                      alt={entry.player_name}
                      className="rounded-full"
                    />
                    <p className="text-lg py-2 w-full">{entry.player_name}</p>
                  </div>
                </div>
              );
            })}
          </div>
        )}

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
