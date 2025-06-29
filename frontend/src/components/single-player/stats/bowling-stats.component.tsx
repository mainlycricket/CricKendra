import { IOverall_Bowling_Summary_Group } from "@/lib/types/bowling-stats.types";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";

import {
  IOverall_Bowling_Continent_Group,
  IOverall_Bowling_HostNation_Group,
  IOverall_Bowling_Opposition_Group,
  IOverall_Bowling_Season_Group,
  IOverall_Bowling_Summary_BatBowlFirst_Group,
  IOverall_Bowling_Summary_BowlingPosition_Group,
  IOverall_Bowling_Summary_HomeAway_Group,
  IOverall_Bowling_Summary_InningsNumber_Group,
  IOverall_Bowling_Summary_MatchResult_Group,
  IOverall_Bowling_Summary_MatchResultBatBowlFirst_Group,
  IOverall_Bowling_Summary_SeriesMatchNumber_Group,
  IOverall_Bowling_Summary_SeriesTeamsCount_Group,
  IOverall_Bowling_Team_Group,
  IOverall_Bowling_Summary_TossDecision_Group,
  IOverall_Bowling_Summary_TossResult_Group,
  IOverall_Bowling_Tournament_Group,
  IOverall_Bowling_Year_Group,
  isT,
} from "@/lib/types/bowling-stats.types";
import { capitalizeFirstLetter } from "@/lib/utils";

type CombinedBowlingStatsType =
  | IOverall_Bowling_Team_Group
  | IOverall_Bowling_Opposition_Group
  | IOverall_Bowling_HostNation_Group
  | IOverall_Bowling_Continent_Group
  | IOverall_Bowling_Year_Group
  | IOverall_Bowling_Season_Group
  | IOverall_Bowling_Summary_HomeAway_Group
  | IOverall_Bowling_Summary_TossResult_Group
  | IOverall_Bowling_Summary_TossDecision_Group
  | IOverall_Bowling_Summary_BatBowlFirst_Group
  | IOverall_Bowling_Summary_InningsNumber_Group
  | IOverall_Bowling_Summary_MatchResult_Group
  | IOverall_Bowling_Summary_MatchResultBatBowlFirst_Group
  | IOverall_Bowling_Summary_SeriesTeamsCount_Group
  | IOverall_Bowling_Summary_SeriesMatchNumber_Group
  | IOverall_Bowling_Tournament_Group
  | IOverall_Bowling_Summary_BowlingPosition_Group;

export function BowlingStats({ stats }: { stats: IOverall_Bowling_Summary_Group }) {
  const data: {
    triggerValue: string;
    triggerLabel: string;
    stats: CombinedBowlingStatsType[];
  }[] = [
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
      triggerLabel: "in match number per series",
      stats: stats.series_event_match_number || [],
    },
    { triggerValue: "tournaments", triggerLabel: "by tournament", stats: stats.tournaments || [] },
    {
      triggerValue: "batting_positions",
      triggerLabel: "by batting position",
      stats: stats.bowling_positions || [],
    },
  ];

  return (
    <div className="flex flex-col gap-4">
      <Accordion type="multiple" defaultValue={["teams"]}>
        {data.map((item) => (
          <AccordionItem value={item.triggerValue} key={item.triggerValue}>
            <AccordionTrigger className="uppercase tracking-wider">{item.triggerLabel}</AccordionTrigger>
            <AccordionContent>
              <BowlingStatsTable stats={item.stats} />
            </AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>
    </div>
  );
}

function BowlingStatsTable({ stats }: { stats: CombinedBowlingStatsType[] }) {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead></TableHead>
          {!isT<IOverall_Bowling_Year_Group>(stats?.[0], ["year"]) &&
            !isT<IOverall_Bowling_Season_Group>(stats?.[0], ["season"]) && (
              <TableHead className="hidden md:table-cell">Span</TableHead>
            )}
          <TableHead className="hidden md:table-cell" title="Matches">
            Mat
          </TableHead>
          <TableHead title="Innings">Inns</TableHead>
          <TableHead>Overs</TableHead>
          <TableHead className="hidden md:table-cell" title="Maiden Overs">
            Mdns
          </TableHead>
          <TableHead className="hidden md:table-cell">Runs</TableHead>
          <TableHead title="Wickets">Wkts</TableHead>
          <TableHead className="hidden md:table-cell" title="Best Bowling Innings">
            BBI
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Best Bowling Match">
            BBM
          </TableHead>
          <TableHead title="Average">Ave</TableHead>
          <TableHead title="Economy">Econ</TableHead>
          <TableHead title="Strike Rate">SR</TableHead>
          <TableHead className="hidden md:table-cell" title="Four-wicket hauls">
            4w
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Five-wicket hauls">
            5w
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Ten-wicket hauls">
            10w
          </TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {stats.map((row) => {
          const { key, label, minSpanYear, maxSpanYear } = getRowMetaData(row);

          return (
            <TableRow key={key}>
              <TableCell className="font-medium">{label}</TableCell>
              {minSpanYear && maxSpanYear && (
                <TableCell className="hidden md:table-cell">
                  {minSpanYear}-{maxSpanYear}
                </TableCell>
              )}
              <TableCell className="hidden md:table-cell">{row.matches_played}</TableCell>
              <TableCell>{row.innings_bowled}</TableCell>
              <TableCell>{row.overs_bowled}</TableCell>
              <TableCell className="hidden md:table-cell">{row.maiden_overs}</TableCell>
              <TableCell className="hidden md:table-cell">{row.runs_conceded}</TableCell>
              <TableCell>{row.wickets_taken}</TableCell>
              <TableCell className="hidden md:table-cell">
                {row.wickets_taken ? `${row.best_innings_wickets}/${row.best_innings_runs}` : "-"}
              </TableCell>
              <TableCell className="hidden md:table-cell">
                {row.wickets_taken ? `${row.best_match_wickets}/${row.best_match_runs}` : "-"}
              </TableCell>
              <TableCell>{row?.average?.toFixed(2) || "-"}</TableCell>
              <TableCell>{row?.economy?.toFixed(2) || "-"}</TableCell>
              <TableCell>{row?.strike_rate?.toFixed(2) || "-"}</TableCell>
              <TableCell className="hidden md:table-cell">{row.four_wicket_hauls}</TableCell>
              <TableCell className="hidden md:table-cell">{row.five_wicket_hauls}</TableCell>
              <TableCell className="hidden md:table-cell">{row.ten_wicket_hauls || 0}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}

function getRowMetaData(row: CombinedBowlingStatsType): {
  key: number | string;
  label: string;
  minSpanYear?: number;
  maxSpanYear?: number;
} {
  if (isT<IOverall_Bowling_Team_Group>(row, ["team_id", "team_name"])) {
    return {
      key: row.team_id,
      label: row.team_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Opposition_Group>(row, ["opposition_team_id", "opposition_team_name"])) {
    return {
      key: row.opposition_team_id,
      label: row.opposition_team_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_HostNation_Group>(row, ["host_nation_id", "host_nation_name"])) {
    return {
      key: row.host_nation_id,
      label: row.host_nation_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Continent_Group>(row, ["continent_id", "continent_name"])) {
    return {
      key: row.continent_id,
      label: row.continent_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Year_Group>(row, ["year"])) {
    return {
      key: row.year,
      label: row.year.toString(),
    };
  }

  if (isT<IOverall_Bowling_Season_Group>(row, ["season"])) {
    return {
      key: row.season,
      label: row.season,
    };
  }

  if (isT<IOverall_Bowling_Summary_HomeAway_Group>(row, ["home_away_label"])) {
    return {
      key: row.home_away_label,
      label: capitalizeFirstLetter(row.home_away_label),
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_TossDecision_Group>(row, ["toss_result", "is_toss_decision_bat"])) {
    return {
      key: row.toss_result + row.is_toss_decision_bat ? "batted" : "fielded",
      label: `${capitalizeFirstLetter(row.toss_result)} Toss & ${
        row.is_toss_decision_bat ? "Batted" : "Fielded"
      }`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_TossResult_Group>(row, ["toss_result"])) {
    return {
      key: row.toss_result,
      label: `${capitalizeFirstLetter(row.toss_result)} Toss`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_MatchResultBatBowlFirst_Group>(row, ["match_result", "bat_bowl_first"])) {
    return {
      key: row.match_result + row.bat_bowl_first,
      label: `${capitalizeFirstLetter(row.bat_bowl_first)} First & ${row.match_result}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_MatchResult_Group>(row, ["match_result"])) {
    return {
      key: row.match_result,
      label: `${capitalizeFirstLetter(row.match_result)}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_BatBowlFirst_Group>(row, ["bat_bowl_first"])) {
    return {
      key: row.bat_bowl_first,
      label: `${capitalizeFirstLetter(row.bat_bowl_first)} First`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_InningsNumber_Group>(row, ["innings_number"])) {
    return {
      key: row.innings_number,
      label: `Innings No. ${row.innings_number}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_SeriesTeamsCount_Group>(row, ["teams_count"])) {
    return {
      key: row.teams_count,
      label: `${row.teams_count} Teams Series`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_SeriesMatchNumber_Group>(row, ["event_match_number"])) {
    return {
      key: row.event_match_number,
      label: `Match No. ${row.event_match_number}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Tournament_Group>(row, ["tournament_id", "tournament_name"])) {
    return {
      key: row.tournament_id,
      label: row.tournament_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Tournament_Group>(row, ["tournament_id", "tournament_name"])) {
    return {
      key: row.tournament_id,
      label: row.tournament_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Bowling_Summary_BowlingPosition_Group>(row, ["bowling_position"])) {
    return {
      key: row.bowling_position,
      label: `Bowl Pos. ${row.bowling_position}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  return { key: "", label: "" };
}
