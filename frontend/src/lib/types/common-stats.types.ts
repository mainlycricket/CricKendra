import { ICombinedBattingStatsType } from "./batting-stats.types";
import { ICombinedBowlingStatsType } from "./bowling-stats.types";

export interface IStats {
  stats: (ICombinedBattingStatsType | ICombinedBowlingStatsType)[];
  next: boolean;
}


