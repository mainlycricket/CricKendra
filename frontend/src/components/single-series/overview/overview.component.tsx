import Link from "next/link";

import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";

import { MatchInfo } from "@/components/single-match/common/match-info.component";
import { TopBowlers } from "../common/top-bowlers.component";
import { TopBatters } from "../common/top-batters.component";

import { ISeriesHeader } from "@/lib/types/single-series.types";
import { IMatchInfo } from "@/lib/types/single-match.types";

export function OverviewComponent({
  series_header,
  fixture_matches,
  result_matches,
}: {
  series_header: ISeriesHeader;
  fixture_matches: IMatchInfo[];
  result_matches: IMatchInfo[];
}) {
  return (
    <div className="w-full md:w-3/4">
      <Tabs defaultValue={fixture_matches?.length ? "fixtures" : "results"}>
        <TabsList className="w-full">
          <TabsTrigger value="fixtures">Fixtures</TabsTrigger>
          <TabsTrigger value="results">Results</TabsTrigger>
          <TabsTrigger value="top-performers" className="md:hidden">
            Top Performers
          </TabsTrigger>
        </TabsList>
        <TabsContent value="fixtures" className="mt-2 flex flex-col gap-4">
          {fixture_matches?.length ? (
            fixture_matches.map((match) => {
              return (
                <Link key={match.match_id} href={`/matches/${match.match_id}`}>
                  <MatchInfo matchInfo={match} />
                </Link>
              );
            })
          ) : (
            <Card>
              <CardHeader>
                <CardTitle>No upcoming fixtures</CardTitle>
              </CardHeader>
            </Card>
          )}
        </TabsContent>
        <TabsContent value="results" className="mt-2 flex flex-col gap-4">
          {result_matches ? (
            result_matches.map((match) => {
              return (
                <Link key={match.match_id} href={`/matches/${match.match_id}`}>
                  <MatchInfo matchInfo={match} />
                </Link>
              );
            })
          ) : (
            <Card>
              <CardHeader>
                <CardTitle>No recent results</CardTitle>
              </CardHeader>
            </Card>
          )}
        </TabsContent>

        <TabsContent value="top-performers" className="mt-2 flex flex-col gap-4 md:flex-row">
          {series_header?.top_batters?.length && <TopBatters topBatters={series_header.top_batters} />}
          {series_header?.top_bowlers?.length && <TopBowlers topBowlers={series_header.top_bowlers} />}
        </TabsContent>
      </Tabs>
    </div>
  );
}
