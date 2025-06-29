import Link from "next/link";

import { MatchInfo } from "@/components/single-match/common/match-info.component";

import { IMatchInfo } from "@/lib/types/single-match.types";

export function MatchesComponent({ matches }: { matches: IMatchInfo[] }) {
  return (
    <div className="w-full md:w-3/4 flex flex-col gap-4">
      {matches.map((match) => {
        return (
          <Link key={match.match_id} href={`/matches/${match.match_id}`}>
            <MatchInfo matchInfo={match} />
          </Link>
        );
      })}
    </div>
  );
}
