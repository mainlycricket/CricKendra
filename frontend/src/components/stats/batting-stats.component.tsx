import { ICombinedBattingStatsType, ICombinedBattingStatsType2 } from "@/lib/types/batting-stats.types";
import { BattingCommonTable } from "./tables/batting-common-table.component";
import { EnumStatsView } from "@/lib/types/enums.types";
import { BattingListTable } from "./tables/batting-list-table.component";

export function BattingStats({
  stats,
  viewType,
  groupType,
}: {
  stats: (ICombinedBattingStatsType | ICombinedBattingStatsType2)[];
  viewType: EnumStatsView;
  groupType: string;
}) {
  return (
    <div>
      {viewType === "individual" && (groupType === "innings" || groupType === "match-totals") ? (
        <BattingListTable stats={stats as ICombinedBattingStatsType2[]} />
      ) : (
        <BattingCommonTable stats={stats as ICombinedBattingStatsType[]} />
      )}
    </div>
  );
}
