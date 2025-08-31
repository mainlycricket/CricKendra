import Link from "next/link";

import { SingleMatch } from "@/components/matches/single-match-info.component";

import { IMatchInfo } from "@/lib/types/match.types";

export function MatchesComponent({ matches }: { matches: IMatchInfo[] }) {
  return (
    <div className="w-full md:w-3/4 flex flex-col gap-4">
      {matches.map((match) => {
        return (
          <Link key={match.match_id} href={`/matches/${match.match_id}`}>
            <SingleMatch matchInfo={match} />
          </Link>
        );
      })}
    </div>
  );
}
