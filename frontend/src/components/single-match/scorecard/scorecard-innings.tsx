import { IInningsExtrasData, IInningsTotalData, ITeamInningsScorecard } from "@/lib/types/single-match.types";
import { ScorecardBatters } from "./scorecard-batters";
import { ScorecardBowlers } from "./scorecard-bowlers";
import { ScorecardFoWs } from "./scorecard-fow";
import { Card, CardContent, CardHeader, CardTitle } from "../../ui/card";
import { AccordionContent, AccordionTrigger } from "../../ui/accordion";
import { AccordionItem } from "@radix-ui/react-accordion";

export function InningsScorecard({ innings }: { innings: ITeamInningsScorecard }) {
  const extrasData: IInningsExtrasData = {
    byes: innings.byes,
    leg_byes: innings.leg_byes,
    wides: innings.wides,
    noballs: innings.noballs,
    penalty: innings.penalty,
  };

  const totalData: IInningsTotalData = {
    total_runs: innings.total_runs,
    total_overs: innings.total_overs,
    total_wickets: innings.total_wickets,
    innings_end: innings.innings_end,
    target_runs: innings.target_runs,
    max_overs: innings.max_overs,
  };

  return (
    <Card className="p-0">
      <AccordionItem value={innings.innings_id.toString()} className="w-full py-4">
        <AccordionTrigger className="bg-cyan pl-2 pe-8 py-2 text-xl md:text-2xl ">
          <CardHeader className="w-full p-0 pl-6">
            <CardTitle className="flex justify-between">
              <p>{innings.batting_team_name}</p>
              <p>
                {innings.total_runs}-{innings.total_wickets}{" "}
                {innings.max_overs ? `${innings.max_overs} ovs maximum` : ""}
              </p>
            </CardTitle>
          </CardHeader>
        </AccordionTrigger>

        <AccordionContent className="p-0">
          <CardContent className="flex flex-col gap-4">
            <ScorecardBatters
              batters={innings.batter_scorecard_entries || []}
              extrasData={extrasData}
              totalData={totalData}
            />
            <ScorecardFoWs entries={innings.fall_of_wickets || []} />
            <ScorecardBowlers bowlers={innings.bowler_scorecard_entries || []} />
          </CardContent>
        </AccordionContent>
      </AccordionItem>
    </Card>
  );
}
