import { EnumStatsView } from "@/lib/types/enums.types";
import { ICombinedTeamStatsType, ICombinedTeamStatsType2 } from "@/lib/types/team-stats.types";
import { TeamCommonTable } from "./tables/team-common-table.component";
import { TeamListTable } from "./tables/team-list-table.component";

export function TeamStats({
  stats,
  viewType,
  groupType,
}: {
  stats: (ICombinedTeamStatsType | ICombinedTeamStatsType2)[];
  viewType: EnumStatsView;
  groupType: string;
}) {
  return (
    <div>
      {viewType === "individual" &&
      (groupType === "innings" || groupType === "match-totals" || groupType === "match-results") ? (
        <TeamListTable stats={stats as ICombinedTeamStatsType2[]} />
      ) : (
        <TeamCommonTable stats={stats as ICombinedTeamStatsType[]} />
      )}
    </div>
  );
}
