import { IFallOfWickets } from "@/lib/types/single-match.types";
import Link from "next/link";

export function ScorecardFoWs({ entries }: { entries: IFallOfWickets[] }) {
  entries.sort((a, b) => a.wicket_number - b.wicket_number);

  return (
    <div>
      <p className="font-semibold">Fall of Wickets:</p>
      <div>
        {entries.map((entry, idx) => {
          return (
            <span key={entry.batter_id}>
              <span>
                {entry.wicket_number}-{entry.team_runs}
              </span>{" "}
              (<Link href={`/players/${entry.batter_id}`}>{entry.batter_name}</Link>, {entry.ball_number} ov)
              {idx < entries.length - 1 ? ", " : ""}
            </span>
          );
        })}
      </div>
    </div>
  );
}
