import { IInningsScorecardSummary } from "@/lib/types/single-match";
import Link from "next/link";

export function SummaryScorecard({ entry }: { entry: IInningsScorecardSummary }) {
  return (
    <div className="border-1">
      <div className="px-4 py-2 border-1 border-t-0">
        {entry.batting_team_name} • {entry.total_runs}-{entry.total_wickets} ({entry.total_overs} overs)
      </div>
      <div className="flex justify-between">
        <div className="w-[50%]">
          {entry.top_batters.map((batter) => {
            return (
              <div key={batter.batter_id} className="border-1 px-4 py-2 flex justify-between">
                <Link href={`/players/${batter.batter_id}`} className="underline">
                  {batter.batter_name}
                </Link>
                <span>
                  {batter.runs_scored} ({batter.balls_faced})
                </span>
              </div>
            );
          })}
        </div>
        <div className="w-[50%]">
          {entry.top_bowlers.map((bowler) => {
            return (
              <div key={bowler.bowler_id} className="border-1 px-4 py-2 flex justify-between">
                <Link href={`/players/${bowler.bowler_id}`} className="underline">
                  {bowler.bowler_name}
                </Link>
                <span>
                  {bowler.wickets_taken}-{bowler.runs_conceded} ({bowler.overs_bowled})
                </span>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
}
