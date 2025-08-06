import { StatsFiltersComponent } from "@/components/stat-filters/main.component";
import { doBackendRequest } from "@/lib/axiosFetch";
import { EnumPlayingFormat, EnumStatsType, EnumStatsView } from "@/lib/types/enums.types";
import { IStatsFilters } from "@/lib/types/filters-stats.types";

export default async function StatsFilters({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string }>;
}) {
  try {
    let { is_male, playing_format, type, view } = await searchParams;
    if (!["Test", "ODI", "T20I", "first_class", "list_a", "t20"].includes(playing_format))
      playing_format = "ODI";
    if (is_male !== "true" && is_male !== "false") is_male = "true";

    const response = await doBackendRequest<null, IStatsFilters>({
      url: `/stats/filter-options?is_male=${is_male}&playing_format=${playing_format}`,
      method: "GET",
    });

    if (type !== "batting" && type !== "bowling" && type !== "team") type = "batting";
    if (view !== "overall" && type !== "individual") view = "overall";
  
    return (
      <StatsFiltersComponent
        commonFilterOptions={response.data!}
        defaultPlayingFormat={playing_format as EnumPlayingFormat}
        defaultStatsType={type as EnumStatsType}
        defaultIsMale={is_male as "true" | "false"}
        defaultView={view as EnumStatsView}
      />
    );
  } catch (error) {
    console.error(error);
  }
}
