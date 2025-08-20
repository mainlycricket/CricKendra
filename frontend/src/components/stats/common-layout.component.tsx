import { IStats } from "@/lib/types/common-stats.types";
import { Card, CardContent } from "../ui/card";
import { BattingStats } from "./batting-stats.component";
import { ICombinedBattingStatsType, ICombinedBattingStatsType2 } from "@/lib/types/batting-stats.types";
import { BowlingStats } from "./bowling-stats.component";
import { ICombinedBowlingStatsType, ICombinedBowlingStatsType2 } from "@/lib/types/bowling-stats.types";
import { IStatsFilters, IStatsFiltersMap, stringifyFiltersMap } from "@/lib/types/filters-stats.types";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion";
import { FiltersList } from "./filters-data.component";
import { Pagination } from "./pagination.component";
import Link from "next/link";
import { TeamStats } from "./team-stats.component";
import { ICombinedTeamStatsType, ICombinedTeamStatsType2 } from "@/lib/types/team-stats.types";

export function CommonStatsLayout({
  stats,
  filtersMap,
  filtersIdData,
}: {
  stats: IStats;
  filtersMap: IStatsFiltersMap;
  filtersIdData: IStatsFilters;
}) {
  try {
    return (
      <div className="flex flex-col gap-4">
        <Card className="p-0">
          <CardContent>
            <Accordion type="single" className="p-0" defaultValue="item-1" collapsible>
              <AccordionItem value="item-1">
                <AccordionTrigger>
                  <div className="flex justify-between w-full items-center">
                    <p className="text-xl font-bold">Filters</p>
                    <Link
                      href={`/stats/filters?${stringifyFiltersMap(filtersMap)}`}
                      className="underline"
                      style={{ color: "var(--color-sky-500)" }}
                    >
                      View Form
                    </Link>
                  </div>
                </AccordionTrigger>
                <AccordionContent>
                  <FiltersList filtersMap={filtersMap} filtersIdData={filtersIdData} />
                </AccordionContent>
              </AccordionItem>
            </Accordion>
          </CardContent>
        </Card>
        <Card>
          <CardContent>
            {filtersMap.statsType === "batting" ? (
              <BattingStats
                stats={stats.stats as (ICombinedBattingStatsType | ICombinedBattingStatsType2)[]}
                viewType={filtersMap["view"]}
                groupType={filtersMap["group"]}
              />
            ) : filtersMap.statsType === "bowling" ? (
              <BowlingStats
                stats={stats.stats as (ICombinedBowlingStatsType | ICombinedBowlingStatsType2)[]}
                viewType={filtersMap["view"]}
                groupType={filtersMap["group"]}
              />
            ) : filtersMap.statsType === "team" ? (
              <TeamStats
                stats={stats.stats as (ICombinedTeamStatsType | ICombinedTeamStatsType2)[]}
                viewType={filtersMap["view"]}
                groupType={filtersMap["group"]}
              />
            ) : (
              <></>
            )}
          </CardContent>
        </Card>

        <Pagination
          currentPage={parseInt(filtersMap["__page"] || "1")}
          recordsCount={stats?.stats?.length || 0}
          disableNext={!stats?.next}
        />
      </div>
    );
  } catch (error) {
    console.error(error);
  }
}
