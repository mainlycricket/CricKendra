import { CommonMatchLayout } from "@/components/single-match/common/common-layout.component";
import { StatsManhattan } from "@/components/single-match/stats/manhattan.component";
import { Partnerships } from "@/components/single-match/stats/partnership.component";
import { StatsRunRate } from "@/components/single-match/stats/runrate.component";
import { StatsWorm } from "@/components/single-match/stats/worm.component";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchStats } from "@/lib/types/single-match.types";

export default async function Stats({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, IMatchStats>({
      url: `/matches/${id}/stats`,
      method: "GET",
    });

    const { match_header, innings: inningsStats } = response.data!;
    match_header?.innings_scores?.sort(
      (a, b) => a.innings_number - b.innings_number,
    );
    inningsStats?.sort((a, b) => a.innings_number - b.innings_number);

    const overStatsData = inningsStats?.map((innings) => {
      const {
        innings_id: inningsId,
        innings_number: inningsNumber,
        batting_team_name: battingTeamName,
        overs,
      } = innings;

      overs?.sort((a, b) => a.over_number - b.over_number);

      return {
        inningsId,
        inningsNumber,
        battingTeamName,
        overs: overs || [],
      };
    });

    return (
      <CommonMatchLayout matchHeader={match_header} active="stats">
        {/* Partnerships */}
        {inningsStats?.length && (
          <Card>
            <CardHeader>
              <CardTitle className="px-2 md:px-12 text-xl">
                Partnerships
              </CardTitle>
            </CardHeader>
            <CardContent className="flex flex-col md:flex-row md:justify-between md:px-16">
              {inningsStats.map((innings) => (
                <div key={innings.innings_id} className="w-full md:max-w-[45%]">
                  <h3 className="text-lg underline text-center">
                    {innings.batting_team_name}
                  </h3>
                  <Partnerships partnerships={innings.partnerships || []} />
                </div>
              ))}
            </CardContent>
          </Card>
        )}

        {/* Manhattan */}
        {inningsStats?.length && inningsStats?.length <= 2 && (
          <StatsManhattan data={overStatsData || []} />
        )}

        {/* Run Rate */}
        {inningsStats?.length && inningsStats?.length <= 2 && (
          <StatsRunRate data={overStatsData || []} />
        )}

        {/* Worm */}
        {inningsStats?.length && inningsStats?.length <= 2 && (
          <StatsWorm data={overStatsData || []} />
        )}
      </CommonMatchLayout>
    );
  } catch (error) {
    console.error(error);
  }
}
