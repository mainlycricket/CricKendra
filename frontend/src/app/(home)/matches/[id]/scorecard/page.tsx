import { MatchHeader } from "@/components/single-match/common/match-header";
import { MatchTabs } from "@/components/single-match/common/match-tabs";
import { InningsScorecard } from "@/components/single-match/scorecard/scorecard-innings";
import { Accordion } from "@/components/ui/accordion";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchScorecard } from "@/lib/types/single-match";

export default async function Scorecard({ params }: { params: Promise<{ id: string }> }) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, IMatchScorecard>({
      url: `/matches/${id}/full-scorecard`,
      method: "GET",
    });
    const { match_header: matchHeader, innings_scorecards: teamInnings } = response.data!;
    matchHeader.innings_scores.sort((a, b) => a.innings_number - b.innings_number);
    teamInnings.sort((a, b) => a.innings_number - b.innings_number);

    return (
      <div className="flex flex-col gap-4">
        <MatchHeader matchHeader={matchHeader} />
        <MatchTabs
          matchId={matchHeader.match_id}
          inningsId={matchHeader?.innings_scores?.[matchHeader?.innings_scores?.length - 1]?.innings_id || -1}
          active="scorecard"
        />

        <Accordion
          type="multiple"
          defaultValue={[teamInnings?.[teamInnings.length - 1]?.innings_id.toString() || ""]}
          className="flex flex-col gap-8"
        >
          {teamInnings.map((innings) => (
            <InningsScorecard innings={innings} key={innings.innings_id} />
          ))}
        </Accordion>
      </div>
    );
  } catch (error) {
    console.error(error);
  }
}
