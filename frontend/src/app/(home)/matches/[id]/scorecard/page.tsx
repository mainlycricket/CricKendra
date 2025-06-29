import { CommonMatchLayout } from "@/components/single-match/common/common-layout.component";
import { InningsScorecard } from "@/components/single-match/scorecard/scorecard-innings";
import { Accordion } from "@/components/ui/accordion";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchScorecard } from "@/lib/types/single-match.types";

export default async function Scorecard({ params }: { params: Promise<{ id: string }> }) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, IMatchScorecard>({
      url: `/matches/${id}/full-scorecard`,
      method: "GET",
    });
    const { match_header, innings_scorecards: teamInnings } = response.data!;
    match_header?.innings_scores?.sort((a, b) => a.innings_number - b.innings_number);
    teamInnings?.sort((a, b) => a.innings_number - b.innings_number);

    return (
      <CommonMatchLayout matchHeader={match_header} active="scorecard">
        <Accordion
          type="multiple"
          defaultValue={[teamInnings?.[teamInnings?.length - 1]?.innings_id.toString() || ""]}
          className="flex flex-col gap-8"
        >
          {teamInnings?.length &&
            teamInnings.map((innings) => <InningsScorecard innings={innings} key={innings.innings_id} />)}
        </Accordion>
      </CommonMatchLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
