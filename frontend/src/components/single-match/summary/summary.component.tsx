import Image from "next/image";

import { SingleScorecardSummary } from "./scorecard-summary.component";

import {
  IBbbCommentary,
  IInningsScorecardSummary,
  IMatchHeader,
  IPlayerAwardInfo,
} from "@/lib/types/single-match.types";
import Link from "next/link";
import { InningsCommentary } from "../commentary/innings-commentary";

export function MatchSummary({
  matchHeader,
  scorecardSummary,
  latestCommentary,
}: {
  matchHeader: IMatchHeader;
  scorecardSummary: IInningsScorecardSummary[];
  latestCommentary: IBbbCommentary[];
}) {
  return (
    <>
      {/* Mobile POM */}
      {matchHeader?.player_awards?.length && <MobilePoM playerAwards={matchHeader.player_awards} />}

      {/* Scorecard Summary */}
      <div className="w-full">
        <h3 className="bg-secondary px-4 py-2 tracking-wider">Scorecard Summary</h3>
        {scorecardSummary.map((innings) => {
          return <SingleScorecardSummary key={innings.innings_id} entry={innings} />;
        })}
        <div className="mt-4 text-blue-400 text-center hover:underline">
          <Link href={`/matches/${matchHeader.match_id}/scorecard`}>View Full Scorecard</Link>
        </div>
      </div>

      {/* Latest Commentary */}
      {matchHeader?.innings_scores?.toReversed().map((innings) => {
        const inningsCommentary = latestCommentary.filter((item) => item.innings_id === innings.innings_id);
        return (
          <InningsCommentary
            key={innings.innings_id}
            commentary={inningsCommentary}
            battingTeamName={innings.batting_team_name}
            displayOverSummary={false}
          />
        );
      })}
    </>
  );
}

function MobilePoM({ playerAwards }: { playerAwards: IPlayerAwardInfo[] }) {
  return (
    <div className="block md:hidden border-y-2">
      {playerAwards.map((entry) => {
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
  );
}
