import { IOverSummary } from "@/lib/types/single-match.types";

export function OverSummary({ overSummary }: { overSummary: IOverSummary }) {
  return (
    <div className="text-black">
      <div className="py-2 px-4 flex justify-between bg-over-summary-20 text-primary border-l-5 border-sky-500">
        <div>
          <span className="font-medium">End of Over {overSummary.overNumber}: </span>
          <span className="md:hidden">
            <br />
          </span>
          {overSummary.overRuns} runs{" "}
          {overSummary.overWickets > 1
            ? `& ${overSummary.overWickets} wickets`
            : overSummary.overWickets
            ? `& ${overSummary.overWickets} wicket`
            : ""}
        </div>
        <div>
          <span className="font-medium">
            {overSummary.battingTeamName}: {overSummary.totalRuns}-{overSummary.totalWickets}{" "}
          </span>
          <span className="md:hidden">
            <br />
          </span>
          <span className="text-sm">CRR: {(overSummary.totalRuns / overSummary.overNumber).toFixed(2)}</span>
        </div>
      </div>

      <div className="px-4 flex bg-over-summary-10 text-primary justify-between gap-4">
        <div className="w-[50%] py-2">
          {overSummary.batters.map((batter) => (
            <div key={batter.id} className="flex justify-between">
              <p>{batter.name}</p>
              <p>
                {batter.runs} ({batter.balls}b{batter.fours ? <span> {batter.fours}✕4</span> : ""}
                {batter.sixes ? <span> {batter.sixes}✕6</span> : ""})
              </p>
            </div>
          ))}
        </div>

        <div className="border-white border-l-1"></div>

        <div className="w-[50%] py-2">
          {overSummary.bowlers.map((bowler) => (
            <div key={bowler.id} className="flex justify-between">
              <p>{bowler.name}</p>
              <p>
                {parseInt((bowler.balls / 6).toFixed(2))}-{bowler.maidens}-{bowler.runs}-{bowler.wickets}
              </p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
