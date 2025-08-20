import { BattingCommonTable } from "@/components/stats/tables/batting-common-table.component";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";

import { IOverall_Batting_Summary_Group, ICombinedBattingStatsType } from "@/lib/types/batting-stats.types";

export function BattingStats({ stats }: { stats: IOverall_Batting_Summary_Group }) {
  const data: { triggerValue: string; triggerLabel: string; stats: ICombinedBattingStatsType[] }[] = [
    { triggerValue: "teams", triggerLabel: "for teams", stats: stats.teams || [] },
    { triggerValue: "oppositions", triggerLabel: "vs teams", stats: stats.oppositions || [] },
    { triggerValue: "host_nations", triggerLabel: "in host country", stats: stats.host_nations || [] },
    { triggerValue: "continents", triggerLabel: "in continent", stats: stats.continents || [] },
    { triggerValue: "home_away", triggerLabel: "home vs away", stats: stats.home_away || [] },
    { triggerValue: "years", triggerLabel: "by year", stats: stats.years || [] },
    { triggerValue: "seasons", triggerLabel: "by season", stats: stats.seasons || [] },
    { triggerValue: "toss_won_lost", triggerLabel: "by toss result", stats: stats.toss_won_lost || [] },
    { triggerValue: "toss_decision", triggerLabel: "by toss decision", stats: stats.toss_decision || [] },
    { triggerValue: "bat_bowl_first", triggerLabel: "by bat/bowl first", stats: stats.bat_bowl_first || [] },
    { triggerValue: "innings_number", triggerLabel: "by innings number", stats: stats.innings_number || [] },
    { triggerValue: "match_result", triggerLabel: "by match result", stats: stats.match_result || [] },
    {
      triggerValue: "match_result_bat_bowl_first",
      triggerLabel: "by match result & toss decision",
      stats: stats.match_result_bat_bowl_first || [],
    },
    {
      triggerValue: "series_teams_count",
      triggerLabel: "by tournament type",
      stats: stats.series_teams_count || [],
    },
    {
      triggerValue: "series_event_match_number",
      triggerLabel: "by match number per series",
      stats: stats.series_event_match_number || [],
    },
    { triggerValue: "tournaments", triggerLabel: "by tournament", stats: stats.tournaments || [] },
    {
      triggerValue: "batting_positions",
      triggerLabel: "by batting position",
      stats: stats.batting_positions || [],
    },
  ];

  return (
    <div className="flex flex-col gap-4">
      <Accordion type="multiple" defaultValue={["teams"]}>
        {data.map((item) => (
          <AccordionItem value={item.triggerValue} key={item.triggerValue}>
            <AccordionTrigger className="uppercase tracking-wider">{item.triggerLabel}</AccordionTrigger>
            <AccordionContent>
              <BattingCommonTable stats={item.stats} />
            </AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>
    </div>
  );
}
