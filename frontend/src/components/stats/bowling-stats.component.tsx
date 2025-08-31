import { EnumStatsView } from "@/lib/types/enums.types";
import { BowlingCommonTable } from "./tables/bowling-common-table.component";
import { ICombinedBowlingStatsType, ICombinedBowlingStatsType2 } from "@/lib/types/bowling-stats.types";
import { BowlingListTable } from "./tables/bowling-list-table.component";

export function BowlingStats({
  stats,
  viewType,
  groupType,
}: {
  stats: (ICombinedBowlingStatsType | ICombinedBowlingStatsType2)[];
  viewType: EnumStatsView;
  groupType: string;
}) {
  return (
    <div>
      {viewType === "individual" && (groupType === "innings" || groupType === "match-totals") ? (
        <BowlingListTable stats={stats as ICombinedBowlingStatsType2[]} />
      ) : (
        <BowlingCommonTable stats={stats as ICombinedBowlingStatsType[]} />
      )}
    </div>
  );
}
