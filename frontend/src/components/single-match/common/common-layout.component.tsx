import { MatchHeader } from "./match-header.component";
import { MatchTabs } from "./match-tabs.component";

import { IMatchHeader } from "@/lib/types/single-match.types";

export function CommonMatchLayout({
  matchHeader,
  children,
  active,
}: {
  matchHeader: IMatchHeader;
  children: React.ReactNode;
  active: string;
}) {
  return (
    <div className="flex flex-col gap-4">
      <MatchHeader matchHeader={matchHeader} />
      <MatchTabs
        matchId={matchHeader.match_id}
        inningsId={matchHeader?.innings_scores?.[matchHeader?.innings_scores?.length - 1]?.innings_id || -1}
        active={active}
      />
      {children}
    </div>
  );
}
