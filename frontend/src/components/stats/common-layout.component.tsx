import { IStats } from "@/lib/types/common-stats.types";
import { Card, CardContent } from "../ui/card";
import { BattingStats } from "./batting-stats.component";
import { ICombinedBattingStatsType } from "@/lib/types/batting-stats.types";
import { BowlingStats } from "./bowling-stats.component";
import { ICombinedBowlingStatsType } from "@/lib/types/bowling-stats.types";
import { EnumStatsType } from "@/lib/types/enums.types";
import { IStatsFilters, IStatsFiltersData } from "@/lib/types/filters-stats.types";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "../ui/accordion";
import { FiltersList } from "./filters-data.component";
import { Pagination } from "./pagination.component";

export function CommonStatsLayout({
  statsType,
  stats,
  filtersMap,
  filtersIdData,
}: {
  statsType: EnumStatsType;
  stats: IStats;
  filtersMap: IStatsFiltersData;
  filtersIdData: IStatsFilters;
}) {
  return (
    <div className="flex flex-col gap-4">
      <Card className="p-0">
        <CardContent>
          <Accordion type="single" className="p-0" defaultValue="item-1" collapsible>
            <AccordionItem value="item-1">
              <AccordionTrigger className="text-xl font-bold">Filters</AccordionTrigger>
              <AccordionContent>
                <FiltersList filtersMap={filtersMap} filtersIdData={filtersIdData} />
              </AccordionContent>
            </AccordionItem>
          </Accordion>
        </CardContent>
      </Card>
      <Card>
        <CardContent>
          {statsType === "batting" ? (
            <BattingStats
              stats={stats.stats as ICombinedBattingStatsType[]}
              viewType={filtersMap["view"]}
              groupType={filtersMap["group"]}
            />
          ) : statsType === "bowling" ? (
            <BowlingStats
              stats={stats.stats as ICombinedBowlingStatsType[]}
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
}
