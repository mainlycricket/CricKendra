import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchSummary } from "@/lib/types/single-match.types";
import { CommonMatchLayout } from "@/components/single-match/common/common-layout.component";
import { MatchSummary } from "@/components/single-match/summary/summary.component";

export default async function Summary({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  try {
    const { id } = await params;

    const response = await doBackendRequest<null, IMatchSummary>({
      url: `/matches/${id}/summary`,
      method: "GET",
    });
    const { match_header, scorecard_summary, latest_commentary } =
      response.data!;

    match_header?.innings_scores?.sort(
      (a, b) => a.innings_number - b.innings_number,
    );

    return (
      <CommonMatchLayout matchHeader={match_header} active="summary">
        <MatchSummary
          matchHeader={match_header}
          latestCommentary={latest_commentary || []}
          scorecardSummary={scorecard_summary || []}
        />
      </CommonMatchLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
