import { IStatsFilters, IStatsFiltersData } from "@/lib/types/filters-stats.types";
import { capitalizeFirstLetter, getDismissalTypeDisplayText, getDisplayDate } from "@/lib/utils";

export function FiltersList({
  filtersMap,
  filtersIdData,
}: {
  filtersMap: IStatsFiltersData;
  filtersIdData: IStatsFilters;
}) {
  return (
    <div className="flex flex-col gap-2">
      <div className="flex gap-2">
        <p className="w-[180px] font-medium">Stats Type</p>
        <p className="capitalize">{filtersMap["statsType"]}</p>
      </div>

      <div className="flex gap-2">
        <p className="w-[180px] font-medium capitalize">Playing Format</p>
        <p className="capitalize">{filtersMap["playing_format"] || "all"}</p>
      </div>

      <div className="flex gap-2">
        <p className="w-[180px] font-medium">Gender</p>
        <p>{filtersMap["is_male"] === "false" ? "Female" : "Male"}</p>
      </div>

      {filtersMap?.primary_team?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Primary Team</p>
          <p>
            {filtersMap?.primary_team
              ?.map((id) => {
                const name =
                  filtersIdData?.primary_teams?.find((item) => item.id.toString() === id)?.name || "";
                return name;
              })
              .join(" or ")}
          </p>
        </div>
      )}

      {filtersMap?.opposition_team?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Opposition Team</p>
          <p>
            {filtersMap?.opposition_team
              ?.map((id) => {
                const name = filtersIdData?.opposition_teams?.find((item) => item.id.toString() === id)?.name;
                if (name) return name;
              })
              .join(" or ")}
          </p>
        </div>
      )}

      {(filtersMap?.min_start_date || filtersMap?.max_start_date) && (
        <div className="hidden md:flex gap-2">
          <p className="w-[180px] font-medium">Start Date</p>
          <p>
            {filtersMap?.min_start_date && <span>Since {getDisplayDate(filtersMap?.min_start_date)}</span>}
            {filtersMap?.min_start_date && filtersMap?.max_start_date && <span> and </span>}
            {filtersMap?.max_start_date && <span>Till {getDisplayDate(filtersMap?.max_start_date)}</span>}
          </p>
        </div>
      )}

      {filtersMap?.min_start_date && (
        <div className="flex gap-2 md:hidden">
          <p className="w-[180px] font-medium">Min. Start Date</p>
          <p>{getDisplayDate(filtersMap?.min_start_date)}</p>
        </div>
      )}

      {filtersMap?.max_start_date && (
        <div className="flex gap-2 md:hidden">
          <p className="w-[180px] font-medium">Max. Start Date</p>
          <p>{getDisplayDate(filtersMap?.max_start_date)}</p>
        </div>
      )}

      {filtersMap?.season?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Seasons</p>
          <p>{filtersMap?.season?.join(" or ")}</p>
        </div>
      )}

      {filtersMap?.home_or_away?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Home or Away</p>
          <p>{filtersMap?.home_or_away?.map((value) => capitalizeFirstLetter(value))?.join(" or ")}</p>
        </div>
      )}

      {filtersMap?.continent?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Continents</p>
          <p>
            {filtersMap?.continent
              ?.map((id) => {
                const name = filtersIdData?.continents?.find((item) => item.id.toString() === id)?.name;
                if (name) return name;
              })
              .join(" or ")}
          </p>
        </div>
      )}

      {filtersMap?.host_nation?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] capitalize">host nations</p>
          <p>
            {filtersMap?.host_nation
              ?.map((id) => {
                const name = filtersIdData?.host_nations?.find((item) => item.id.toString() === id)?.name;
                if (name) return name;
              })
              .join(" or ")}
          </p>
        </div>
      )}

      {filtersMap?.ground?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] capitalize">grounds</p>
          <p>
            {filtersMap?.ground
              ?.map((id) => {
                const name = filtersIdData?.grounds?.find((item) => item.id.toString() === id)?.name;
                if (name) return name;
              })
              .join(" or ")}
          </p>
        </div>
      )}

      {filtersMap?.tournament?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] capitalize">tournaments</p>
          <p>
            {filtersMap?.tournament
              ?.map((id) => {
                const name = filtersIdData?.tournaments?.find((item) => item.id.toString() === id)?.name;
                if (name) return name;
              })
              .join(" or ")}
          </p>
        </div>
      )}

      {filtersMap?.series?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] capitalize">series</p>
          <p>
            {filtersMap?.series
              ?.map((id) => {
                const item = filtersIdData?.series?.find((item) => item.id.toString() === id);
                if (item) {
                  return `${item.name}, ${item.season}`;
                }
              })
              .join(" or ")}
          </p>
        </div>
      )}

      {filtersMap?.match_result && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Match Result</p>
          <p>{filtersMap?.match_result?.map((value) => capitalizeFirstLetter(value))?.join(" or ")}</p>
        </div>
      )}

      {filtersMap?.toss_result && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Toss Result</p>
          <p className="capitalize">{filtersMap?.toss_result}</p>
        </div>
      )}

      {filtersMap?.bat_field_first && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Bat or Field First</p>
          <p className="capitalize">{filtersMap?.bat_field_first}</p>
        </div>
      )}

      {filtersMap?.innings_number?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Innings Number</p>
          <p>{filtersMap?.innings_number?.sort()?.join(" or ")}</p>
        </div>
      )}

      {/* Batting Specific filters */}

      {(filtersMap?.min__innings_runs_scored || filtersMap?.max__innings_runs_scored) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Runs scored in an inns</p>
          <p>
            {filtersMap?.min__innings_runs_scored && <span>Min. {filtersMap?.min__innings_runs_scored}</span>}
            {filtersMap?.min__innings_runs_scored && filtersMap?.max__innings_runs_scored && (
              <span> and </span>
            )}
            {filtersMap?.max__innings_runs_scored && <span>Max. {filtersMap?.max__innings_runs_scored}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__innings_batting_position || filtersMap?.max__innings_batting_position) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Batting Position</p>
          <p>
            {filtersMap?.min__innings_batting_position && (
              <span>Min. {filtersMap?.min__innings_batting_position}</span>
            )}
            {filtersMap?.min__innings_batting_position && filtersMap?.max__innings_batting_position && (
              <span> and </span>
            )}
            {filtersMap?.max__innings_batting_position && (
              <span>Max. {filtersMap?.max__innings_batting_position}</span>
            )}
          </p>
        </div>
      )}

      {filtersMap?.innings_is_batter_dismissed && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Dismissed</p>
          <p className="capitalize">{filtersMap?.innings_is_batter_dismissed?.replaceAll("_", " ")}</p>
        </div>
      )}

      {filtersMap?.innings_batter_dismissal_type?.length && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Dismissal Type</p>
          <p>
            {filtersMap?.innings_batter_dismissal_type
              ?.map((value) => getDismissalTypeDisplayText(value))
              .join(" or ")}
          </p>
        </div>
      )}

      {/* Bowling Specific filters */}
      {(filtersMap?.min__innings_balls_bowled || filtersMap?.max__innings_balls_bowled) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Balls Bowled in an inns</p>
          <p>
            {filtersMap?.min__innings_balls_bowled && (
              <span>Min. {filtersMap?.min__innings_balls_bowled}</span>
            )}
            {filtersMap?.min__innings_balls_bowled && filtersMap?.max__innings_balls_bowled && (
              <span> and </span>
            )}
            {filtersMap?.max__innings_balls_bowled && (
              <span>Max. {filtersMap?.max__innings_balls_bowled}</span>
            )}
          </p>
        </div>
      )}

      {(filtersMap?.min__innings_runs_conceded || filtersMap?.max__innings_runs_conceded) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Runs conceded in an inns</p>
          <p>
            {filtersMap?.min__innings_runs_conceded && (
              <span>Min. {filtersMap?.min__innings_runs_conceded}</span>
            )}
            {filtersMap?.min__innings_runs_conceded && filtersMap?.max__innings_runs_conceded && (
              <span> and </span>
            )}
            {filtersMap?.max__innings_runs_conceded && (
              <span>Max. {filtersMap?.max__innings_runs_conceded}</span>
            )}
          </p>
        </div>
      )}

      {(filtersMap?.min__innings_wickets_taken || filtersMap?.max__innings_wickets_taken) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Wkts Taken in an inns</p>
          <p>
            {filtersMap?.min__innings_wickets_taken && (
              <span>Min. {filtersMap?.min__innings_wickets_taken}</span>
            )}
            {filtersMap?.min__innings_wickets_taken && filtersMap?.max__innings_wickets_taken && (
              <span> and </span>
            )}
            {filtersMap?.max__innings_wickets_taken && (
              <span>Max. {filtersMap?.max__innings_wickets_taken}</span>
            )}
          </p>
        </div>
      )}

      {(filtersMap?.min__innings_bowling_position || filtersMap?.max__innings_bowling_position) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Bowling Position</p>
          <p>
            {filtersMap?.min__innings_bowling_position && (
              <span>Min. {filtersMap?.min__innings_bowling_position}</span>
            )}
            {filtersMap?.min__innings_bowling_position && filtersMap?.max__innings_bowling_position && (
              <span> and </span>
            )}
            {filtersMap?.max__innings_bowling_position && (
              <span>Max. {filtersMap?.max__innings_bowling_position}</span>
            )}
          </p>
        </div>
      )}

      {/* Common Qualification Filters */}
      {(filtersMap?.min__matches_played || filtersMap?.max__matches_played) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Matches Played</p>
          <p>
            {filtersMap?.min__matches_played && <span>Min. {filtersMap?.min__matches_played}</span>}
            {filtersMap?.min__matches_played && filtersMap?.max__matches_played && <span> and </span>}
            {filtersMap?.max__matches_played && <span>Max. {filtersMap?.max__matches_played}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__average || filtersMap?.max__average) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Average</p>
          <p>
            {filtersMap?.min__average && <span>Min. {filtersMap?.min__average}</span>}
            {filtersMap?.min__average && filtersMap?.max__average && <span> and </span>}
            {filtersMap?.max__average && <span>Max. {filtersMap?.max__average}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__strike_rate || filtersMap?.max__strike_rate) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Strike Rate</p>
          <p>
            {filtersMap?.min__strike_rate && <span>Min. {filtersMap?.min__strike_rate}</span>}
            {filtersMap?.min__strike_rate && filtersMap?.max__strike_rate && <span> and </span>}
            {filtersMap?.max__strike_rate && <span>Max. {filtersMap?.max__strike_rate}</span>}
          </p>
        </div>
      )}

      {/* Batting Specific Qualification Filters */}
      {(filtersMap?.min__innings_batted || filtersMap?.max__innings_batted) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Innings Batted</p>
          <p>
            {filtersMap?.min__innings_batted && <span>Min. {filtersMap?.min__innings_batted}</span>}
            {filtersMap?.min__innings_batted && filtersMap?.max__innings_batted && <span> and </span>}
            {filtersMap?.max__innings_batted && <span>Max. {filtersMap?.max__innings_batted}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__not_outs || filtersMap?.max__not_outs) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Not Outs</p>
          <p>
            {filtersMap?.min__not_outs && <span>Min. {filtersMap?.min__not_outs}</span>}
            {filtersMap?.min__not_outs && filtersMap?.max__not_outs && <span> and </span>}
            {filtersMap?.max__not_outs && <span>Max. {filtersMap?.max__not_outs}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__runs_scored || filtersMap?.max__runs_scored) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Runs Scored</p>
          <p>
            {filtersMap?.min__runs_scored && <span>Min. {filtersMap?.min__runs_scored}</span>}
            {filtersMap?.min__runs_scored && filtersMap?.max__runs_scored && <span> and </span>}
            {filtersMap?.max__runs_scored && <span>Max. {filtersMap?.max__runs_scored}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__balls_faced || filtersMap?.max__balls_faced) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Balls Faced</p>
          <p>
            {filtersMap?.min__balls_faced && <span>Min. {filtersMap?.min__balls_faced}</span>}
            {filtersMap?.min__balls_faced && filtersMap?.max__balls_faced && <span> and </span>}
            {filtersMap?.max__balls_faced && <span>Max. {filtersMap?.max__balls_faced}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__centuries || filtersMap?.max__centuries) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Centuries</p>
          <p>
            {filtersMap?.min__centuries && <span>Min. {filtersMap?.min__centuries}</span>}
            {filtersMap?.min__centuries && filtersMap?.max__centuries && <span> and </span>}
            {filtersMap?.max__centuries && <span>Max. {filtersMap?.max__centuries}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__half_centuries || filtersMap?.max__half_centuries) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Half Centuries</p>
          <p>
            {filtersMap?.min__half_centuries && <span>Min. {filtersMap?.min__half_centuries}</span>}
            {filtersMap?.min__half_centuries && filtersMap?.max__half_centuries && <span> and </span>}
            {filtersMap?.max__half_centuries && <span>Max. {filtersMap?.max__half_centuries}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__fifty_plus_scores || filtersMap?.max__fifty_plus_scores) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">50+ scores</p>
          <p>
            {filtersMap?.min__fifty_plus_scores && <span>Min. {filtersMap?.min__fifty_plus_scores}</span>}
            {filtersMap?.min__fifty_plus_scores && filtersMap?.max__fifty_plus_scores && <span> and </span>}
            {filtersMap?.max__fifty_plus_scores && <span>Max. {filtersMap?.max__fifty_plus_scores}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__ducks || filtersMap?.max__ducks) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Ducks</p>
          <p>
            {filtersMap?.min__ducks && <span>Min. {filtersMap?.min__ducks}</span>}
            {filtersMap?.min__ducks && filtersMap?.max__ducks && <span> and </span>}
            {filtersMap?.max__ducks && <span>Max. {filtersMap?.max__ducks}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__fours_scored || filtersMap?.max__fours_scored) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">4s scored</p>
          <p>
            {filtersMap?.min__fours_scored && <span>Min. {filtersMap?.min__fours_scored}</span>}
            {filtersMap?.min__fours_scored && filtersMap?.max__fours_scored && <span> and </span>}
            {filtersMap?.max__fours_scored && <span>Max. {filtersMap?.max__fours_scored}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__sixes_scored || filtersMap?.max__sixes_scored) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">6s scored</p>
          <p>
            {filtersMap?.min__sixes_scored && <span>Min. {filtersMap?.min__sixes_scored}</span>}
            {filtersMap?.min__sixes_scored && filtersMap?.max__sixes_scored && <span> and </span>}
            {filtersMap?.max__sixes_scored && <span>Max. {filtersMap?.max__sixes_scored}</span>}
          </p>
        </div>
      )}

      {/* Bowling Specific Qualification Filters */}
      {(filtersMap?.min__innings_bowled || filtersMap?.max__innings_bowled) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Innings Bowled</p>
          <p>
            {filtersMap?.min__innings_bowled && <span>Min. {filtersMap?.min__innings_bowled}</span>}
            {filtersMap?.min__innings_bowled && filtersMap?.max__innings_bowled && <span> and </span>}
            {filtersMap?.max__innings_bowled && <span>Max. {filtersMap?.max__innings_bowled}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__overs_bowled || filtersMap?.max__overs_bowled) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Overs Bowled</p>
          <p>
            {filtersMap?.min__overs_bowled && <span>Min. {filtersMap?.min__overs_bowled}</span>}
            {filtersMap?.min__overs_bowled && filtersMap?.max__overs_bowled && <span> and </span>}
            {filtersMap?.max__overs_bowled && <span>Max. {filtersMap?.max__overs_bowled}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__maiden_overs || filtersMap?.max__maiden_overs) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Maiden Overs</p>
          <p>
            {filtersMap?.min__maiden_overs && <span>Min. {filtersMap?.min__maiden_overs}</span>}
            {filtersMap?.min__maiden_overs && filtersMap?.max__maiden_overs && <span> and </span>}
            {filtersMap?.max__maiden_overs && <span>Max. {filtersMap?.max__maiden_overs}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__runs_conceded || filtersMap?.max__runs_conceded) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Runs Conceded</p>
          <p>
            {filtersMap?.min__runs_conceded && <span>Min. {filtersMap?.min__runs_conceded}</span>}
            {filtersMap?.min__runs_conceded && filtersMap?.max__runs_conceded && <span> and </span>}
            {filtersMap?.max__runs_conceded && <span>Max. {filtersMap?.max__runs_conceded}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__wickets_taken || filtersMap?.max__wickets_taken) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Wickets Taken</p>
          <p>
            {filtersMap?.min__wickets_taken && <span>Min. {filtersMap?.min__wickets_taken}</span>}
            {filtersMap?.min__wickets_taken && filtersMap?.max__wickets_taken && <span> and </span>}
            {filtersMap?.max__wickets_taken && <span>Max. {filtersMap?.max__wickets_taken}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__economy || filtersMap?.max__economy) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Economy</p>
          <p>
            {filtersMap?.min__economy && <span>Min. {filtersMap?.min__economy}</span>}
            {filtersMap?.min__economy && filtersMap?.max__economy && <span> and </span>}
            {filtersMap?.max__economy && <span>Max. {filtersMap?.max__economy}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__fours_conceded || filtersMap?.max__fours_conceded) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Fours Conceded</p>
          <p>
            {filtersMap?.min__fours_conceded && <span>Min. {filtersMap?.min__fours_conceded}</span>}
            {filtersMap?.min__fours_conceded && filtersMap?.max__fours_conceded && <span> and </span>}
            {filtersMap?.max__fours_conceded && <span>Max. {filtersMap?.max__fours_conceded}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__sixes_conceded || filtersMap?.max__sixes_conceded) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Sixes Conceded</p>
          <p>
            {filtersMap?.min__sixes_conceded && <span>Min. {filtersMap?.min__sixes_conceded}</span>}
            {filtersMap?.min__sixes_conceded && filtersMap?.max__sixes_conceded && <span> and </span>}
            {filtersMap?.max__sixes_conceded && <span>Max. {filtersMap?.max__sixes_conceded}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__four_wkt_hauls || filtersMap?.max__four_wkt_hauls) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">4-wkt hauls</p>
          <p>
            {filtersMap?.min__four_wkt_hauls && <span>Min. {filtersMap?.min__four_wkt_hauls}</span>}
            {filtersMap?.min__four_wkt_hauls && filtersMap?.max__four_wkt_hauls && <span> and </span>}
            {filtersMap?.max__four_wkt_hauls && <span>Max. {filtersMap?.max__four_wkt_hauls}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__five_wkt_hauls || filtersMap?.max__five_wkt_hauls) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">5-wkt hauls</p>
          <p>
            {filtersMap?.min__five_wkt_hauls && <span>Min. {filtersMap?.min__five_wkt_hauls}</span>}
            {filtersMap?.min__five_wkt_hauls && filtersMap?.max__five_wkt_hauls && <span> and </span>}
            {filtersMap?.max__five_wkt_hauls && <span>Max. {filtersMap?.max__five_wkt_hauls}</span>}
          </p>
        </div>
      )}

      {(filtersMap?.min__ten_wkt_hauls || filtersMap?.max__ten_wkt_hauls) && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">10-wkt hauls</p>
          <p>
            {filtersMap?.min__ten_wkt_hauls && <span>Min. {filtersMap?.min__ten_wkt_hauls}</span>}
            {filtersMap?.min__ten_wkt_hauls && filtersMap?.max__ten_wkt_hauls && <span> and </span>}
            {filtersMap?.max__ten_wkt_hauls && <span>Max. {filtersMap?.max__ten_wkt_hauls}</span>}
          </p>
        </div>
      )}

      <div className="flex gap-2">
        <p className="w-[180px] font-medium capitalize">View</p>
        <p className="capitalize">{filtersMap["view"]}</p>
      </div>

      <div className="flex gap-2">
        <p className="w-[180px] font-medium">Group By</p>
        <p className="capitalize">{filtersMap["group"]}</p>
      </div>

      {filtersMap["sort_by"] && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Sort By</p>
          <p className="capitalize">{filtersMap["sort_by"].replaceAll("_", " ")}</p>
        </div>
      )}

      {filtersMap["sort_order"] && (
        <div className="flex gap-2">
          <p className="w-[180px] font-medium">Sort Order</p>
          <p className="capitalize">{filtersMap["sort_order"]}</p>
        </div>
      )}
    </div>
  );
}
