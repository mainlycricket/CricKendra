import { MatchHeader } from "@/components/single-match/common/match-header";
import { MatchTabs } from "@/components/single-match/common/match-tabs";
import { StatsManhattan } from "@/components/single-match/stats/stats-manhattan";
import { Partnerships } from "@/components/single-match/stats/stats-partnership";
import { StatsRunRate } from "@/components/single-match/stats/stats-runrate";
import { StatsWorm } from "@/components/single-match/stats/stats-worm";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { doBackendRequest } from "@/lib/axiosFetch";
import { IMatchStats } from "@/lib/types/single-match";

export default async function Stats({ params }: { params: Promise<{ id: string }> }) {
  try {
    const { id } = await params;
    const response = await doBackendRequest<null, IMatchStats>({
      url: `/matches/${id}/stats`,
      method: "GET",
    });

    const { match_header: matchHeader, innings: inningsStats } = response.data!;
    matchHeader.innings_scores.sort((a, b) => a.innings_number - b.innings_number);
    inningsStats.sort((a, b) => a.innings_number - b.innings_number);

    const overStatsData = inningsStats.map((innings) => {
      const {
        innings_id: inningsId,
        innings_number: inningsNumber,
        batting_team_name: battingTeamName,
        overs,
      } = innings;

      overs.sort((a, b) => a.over_number - b.over_number);

      return {
        inningsId,
        inningsNumber,
        battingTeamName,
        overs,
      };
    });

    return (
      <div className="flex flex-col gap-4">
        <MatchHeader matchHeader={matchHeader} />
        <MatchTabs
          matchId={matchHeader.match_id}
          inningsId={matchHeader?.innings_scores?.[matchHeader?.innings_scores?.length - 1]?.innings_id || -1}
          active="stats"
        />

        {/* Partnerships */}
        {inningsStats?.length && (
          <Card>
            <CardHeader>
              <CardTitle className="px-2 md:px-12 text-xl">Partnerships</CardTitle>
            </CardHeader>
            <CardContent className="flex flex-col md:flex-row md:justify-between md:px-16">
              {inningsStats.map((innings) => (
                <div key={innings.innings_id} className="w-full md:max-w-[45%]">
                  <h3 className="text-lg underline text-center">{innings.batting_team_name}</h3>
                  <Partnerships partnerships={innings.partnerships} />
                </div>
              ))}
            </CardContent>
          </Card>
        )}

        {/* Manhattan */}
        {inningsStats?.length && inningsStats?.length <= 2 && <StatsManhattan data={overStatsData} />}

        {/* Run Rate */}
        {inningsStats?.length && inningsStats?.length <= 2 && <StatsRunRate data={overStatsData} />}

        {/* Worm */}
        {inningsStats?.length && inningsStats?.length <= 2 && <StatsWorm data={overStatsData} />}
      </div>
    );
  } catch (error) {
    console.error(error);
  }
}
