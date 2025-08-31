import { ICombinedBattingStatsType, ICombinedBattingStatsType2 } from "./batting-stats.types";
import { ICombinedBowlingStatsType, ICombinedBowlingStatsType2 } from "./bowling-stats.types";
import { ICombinedTeamStatsType, ICombinedTeamStatsType2 } from "./team-stats.types";

export interface IStats {
  stats: (
    | ICombinedBattingStatsType
    | ICombinedBattingStatsType2
    | ICombinedBowlingStatsType
    | ICombinedBowlingStatsType2
    | ICombinedTeamStatsType
    | ICombinedTeamStatsType2
  )[];
  next: boolean;
}
