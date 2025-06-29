import { IBbbCommentary, IOverSummary } from "@/lib/types/single-match.types";
import { SingleBallCommentary } from "./single-commentary";
import { OverSummary } from "./over-summary";
import { isBowlerDismissal } from "@/lib/utils";

export function InningsCommentary({
  battingTeamName,
  commentary,
  displayOverSummary,
}: {
  battingTeamName: string;
  commentary: IBbbCommentary[];
  displayOverSummary: boolean;
}) {
  if (!commentary?.length) return <div className="py-2">No commentary yet</div>;

  commentary.sort((a, b) => a.innings_delivery_number - b.innings_delivery_number);
  let previousBowlerId = 0,
    currentBowlerRuns = 0,
    currentBowlerBalls = 0;

  const overSummary: IOverSummary = {
    battingTeamName,
    overNumber: commentary?.[0]?.over_number,
    overRuns: 0,
    overWickets: 0,
    totalRuns: 0,
    totalWickets: 0,
    batters: [],
    bowlers: [],
  };

  const rows = commentary.map((item, idx) => {
    let striker = overSummary.batters.find((batter) => batter.id === item.batter_id!);
    if (!striker) {
      striker = {
        id: item.batter_id!,
        name: item.batter_name!,
        runs: 0,
        balls: 0,
        fours: 0,
        sixes: 0,
        isActive: true,
      };

      overSummary.batters.push(striker);
    }
    striker.isActive = true;

    let bowler = overSummary.bowlers.find((bowler) => bowler.id === item.bowler_id!);
    if (!bowler) {
      bowler = {
        id: item.bowler_id!,
        name: item.bowler_name!,
        runs: 0,
        balls: 0,
        maidens: 0,
        wickets: 0,
      };
    }
    overSummary.bowlers = overSummary.bowlers.filter((bowler) => bowler.id !== item.bowler_id);
    overSummary.bowlers.push(bowler);

    if (bowler.id !== previousBowlerId) {
      currentBowlerRuns = currentBowlerBalls = 0;
    }
    previousBowlerId = bowler.id;

    if (!item.wides) striker.balls++;
    if (!item.noballs && !item.wides) {
      bowler.balls++;
      currentBowlerBalls++;
    }
    if (item.is_four) striker.fours++;
    if (item.is_six) striker.sixes++;
    striker.runs += item.total_runs - (item.wides + item.noballs + item.legbyes + item.byes);
    bowler.runs += item.total_runs - (item.legbyes + item.byes);
    currentBowlerRuns += item.total_runs - (item.legbyes + item.byes);

    if (currentBowlerBalls === 6 && currentBowlerRuns === 0) {
      bowler.maidens++;
    }

    overSummary.overRuns += item.total_runs;
    overSummary.totalRuns += item.total_runs;

    if (item.player1_dismissed_id) {
      if (isBowlerDismissal(item.player1_dismissal_type!)) bowler.wickets++;

      const batter = overSummary.batters.find((batter) => batter.id === item.player1_dismissed_id);
      if (batter) {
        batter.isActive = false;
      }

      overSummary.overWickets++;
      overSummary.totalWickets++;
    }

    if (item.player2_dismissed_id) {
      const batter = overSummary.batters.find((batter) => batter.id === item.player2_dismissed_id);
      if (batter) {
        batter.isActive = false;
      }

      if (isBowlerDismissal(item.player2_dismissal_type!)) bowler.wickets++;
      overSummary.overWickets++;
      overSummary.totalWickets++;
    }

    const overSummary2 = Object.assign({}, overSummary);
    overSummary2.batters = overSummary2.batters
      .filter((batter) => batter.isActive)
      .map((batter) => {
        return { ...batter };
      });
    overSummary2.bowlers = overSummary2.bowlers
      .map((bowler) => {
        return { ...bowler };
      })
      .reverse()
      .slice(0, 2);

    const flag = idx == commentary.length - 1 ? true : commentary[idx + 1].over_number > item.over_number;

    const component = (
      <div key={item.innings_delivery_number}>
        {displayOverSummary && flag && <OverSummary overSummary={overSummary2} />}
        <SingleBallCommentary data={item} />
      </div>
    );

    if (flag) {
      overSummary.overNumber = item.over_number + 1;
      overSummary.overRuns = 0;
      overSummary.overWickets = 0;
    }

    return component;
  });

  return <div className="py-2">{rows.reverse()}</div>;
}
