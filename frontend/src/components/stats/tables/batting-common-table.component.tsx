import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";

import {
  IOverall_Batting_Batter_Group,
  IOverall_Batting_Continent_Group,
  IOverall_Batting_HostNation_Group,
  IOverall_Batting_Opposition_Group,
  IOverall_Batting_Season_Group,
  IOverall_Batting_Summary_BatBowlFirst_Group,
  IOverall_Batting_Summary_BattingPosition_Group,
  IOverall_Batting_Summary_HomeAway_Group,
  IOverall_Batting_Summary_InningsNumber_Group,
  IOverall_Batting_Summary_MatchResult_Group,
  IOverall_Batting_Summary_MatchResultBatBowlFirst_Group,
  IOverall_Batting_Summary_SeriesMatchNumber_Group,
  IOverall_Batting_Summary_SeriesTeamsCount_Group,
  IOverall_Batting_Team_Group,
  IOverall_Batting_Summary_TossDecision_Group,
  IOverall_Batting_Summary_TossResult_Group,
  IOverall_Batting_Tournament_Group,
  IOverall_Batting_Year_Group,
  ICombinedBattingStatsType,
  isT,
  IOverall_Batting_Ground_Group,
  IOverall_Batting_Series_Group,
  IOverall_Batting_Decade_Group,
  IOverall_Batting_Aggregate_Group,
  IOverall_Batting_Match_Group,
  IOverall_Batting_TeamInnings_Group,
  IIndividual_Batting_Series_Group,
  IIndividual_Batting_Season_Group,
  IIndividual_Batting_Year_Group,
  IIndividual_Batting_Opposition_Group,
  IIndividual_Batting_HostNation_Group,
  IIndividual_Batting_Ground_Group,
  IIndividual_Batting_Tournaments_Group,
} from "@/lib/types/batting-stats.types";
import { capitalizeFirstLetter } from "@/lib/utils";

export function BattingCommonTable({ stats }: { stats: ICombinedBattingStatsType[] }) {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead></TableHead>
          {!isT<IOverall_Batting_Year_Group>(stats?.[0] || {}, ["year"]) &&
            !isT<IOverall_Batting_Season_Group>(stats?.[0] || {}, ["season"]) &&
            !isT<IOverall_Batting_Decade_Group>(stats?.[0] || {}, ["decade"]) && (
              <TableHead className="hidden md:table-cell">Span</TableHead>
            )}
          <TableHead className="hidden md:table-cell" title="Matches">
            Mat
          </TableHead>
          <TableHead title="Innings">Inns</TableHead>
          <TableHead className="hidden md:table-cell" title="Not Outs">
            NO
          </TableHead>
          <TableHead>Runs</TableHead>
          <TableHead className="hidden md:table-cell" title="Highest Score">
            HS
          </TableHead>
          <TableHead title="Average">Ave</TableHead>
          <TableHead className="hidden md:table-cell" title="Balls Faced">
            BF
          </TableHead>
          <TableHead title="Strike Rate">SR</TableHead>
          <TableHead title="Centuries">100s</TableHead>
          <TableHead title="Half-Centuries">50s</TableHead>
          <TableHead className="hidden md:table-cell" title="Ducks">
            0s
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Fours">
            4s
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Sixes">
            6s
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
                  {minSpanYear} - {maxSpanYear}
                </TableCell>
              )}
              <TableCell className="hidden md:table-cell">{row.matches_played}</TableCell>
              <TableCell>{row.innings_batted}</TableCell>
              <TableCell className="hidden md:table-cell">{row.not_outs}</TableCell>
              <TableCell>{row.runs_scored}</TableCell>
              <TableCell className="hidden md:table-cell">
                {row.highest_score}
                {row.highest_not_out_score === row.highest_score ? "*" : ""}
              </TableCell>
              <TableCell>{row?.average?.toFixed(2) || "-"}</TableCell>
              <TableCell className="hidden md:table-cell">{row.balls_faced}</TableCell>
              <TableCell>{row.strike_rate.toFixed(2)}</TableCell>
              <TableCell>{row.centuries}</TableCell>
              <TableCell>{row.half_centuries}</TableCell>
              <TableCell className="hidden md:table-cell">{row.ducks}</TableCell>
              <TableCell className="hidden md:table-cell">{row.fours_scored}</TableCell>
              <TableCell className="hidden md:table-cell">{row.sixes_scored}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}

function getRowMetaData(row: ICombinedBattingStatsType): {
  key: number | string;
  label: string;
  minSpanYear?: number;
  maxSpanYear?: number;
} {
  /* Individual */
  if (isT<IIndividual_Batting_Series_Group>(row, ["batter_id", "series_id"])) {
    return {
      key: `${row.batter_id}_${row.series_id}`,
      label: `${row.batter_name} in ${row.series_name}, ${row.series_season}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Batting_Tournaments_Group>(row, ["batter_id", "tournament_id"])) {
    return {
      key: `${row.batter_id}_${row.tournament_id}`,
      label: `${row.batter_name} in ${row.tournament_name}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Batting_Ground_Group>(row, ["batter_id", "ground_id"])) {
    return {
      key: `${row.batter_id}_${row.ground_id}`,
      label: `${row.batter_name} at ${row.ground_name}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Batting_HostNation_Group>(row, ["batter_id", "host_nation_id"])) {
    return {
      key: `${row.batter_id}_${row.host_nation_id}`,
      label: `${row.batter_name} in ${row.host_nation_name}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Batting_Opposition_Group>(row, ["batter_id", "opposition_team_id"])) {
    return {
      key: `${row.batter_id}_${row.opposition_team_id}`,
      label: `${row.batter_name} v ${row.opposition_team_name}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Batting_Year_Group>(row, ["batter_id", "year"])) {
    return {
      key: `${row.batter_id}_${row.year}`,
      label: `${row.batter_name}, ${row.year}`,
    };
  }

  if (isT<IIndividual_Batting_Season_Group>(row, ["batter_id", "season"])) {
    return {
      key: `${row.batter_id}_${row.season}`,
      label: `${row.batter_name}, ${row.season}`,
    };
  }

  /* Overall */
  if (isT<IOverall_Batting_Batter_Group>(row, ["batter_id", "batter_name"])) {
    return {
      key: row.batter_id,
      label: row.batter_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_TeamInnings_Group>(row, ["match_id", "innings_number"])) {
    return {
      key: `${row.match_id}-${row.innings_number}`,
      label: `${row.batting_team_name} (innings ${row.innings_number}) v ${row.bowling_team_name} at ${row.city_name}, ${row.season}`,
    };
  }

  if (isT<IOverall_Batting_Match_Group>(row, ["match_id"])) {
    return {
      key: row.match_id,
      label: `${row.team1_name} v ${row.team2_name} at ${row.city_name}, ${row.season}`,
    };
  }

  if (isT<IOverall_Batting_Team_Group>(row, ["team_id", "team_name"])) {
    return {
      key: row.team_id,
      label: row.team_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Opposition_Group>(row, ["opposition_team_id", "opposition_team_name"])) {
    return {
      key: row.opposition_team_id,
      label: row.opposition_team_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Ground_Group>(row, ["ground_id", "ground_name"])) {
    return {
      key: row.ground_id,
      label: row.ground_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_HostNation_Group>(row, ["host_nation_id", "host_nation_name"])) {
    return {
      key: row.host_nation_id,
      label: row.host_nation_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Continent_Group>(row, ["continent_id", "continent_name"])) {
    return {
      key: row.continent_id,
      label: row.continent_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Year_Group>(row, ["year"])) {
    return {
      key: row.year,
      label: row.year.toString(),
    };
  }

  if (isT<IOverall_Batting_Season_Group>(row, ["season"])) {
    return {
      key: row.season,
      label: row.season,
    };
  }

  if (isT<IOverall_Batting_Summary_HomeAway_Group>(row, ["home_away_label"])) {
    return {
      key: row.home_away_label,
      label: capitalizeFirstLetter(row.home_away_label),
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Summary_TossDecision_Group>(row, ["toss_result", "is_toss_decision_bat"])) {
    return {
      key: row.toss_result + row.is_toss_decision_bat ? "batted" : "fielded",
      label: `${capitalizeFirstLetter(row.toss_result)} Toss & ${
        row.is_toss_decision_bat ? "Batted" : "Fielded"
      }`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Summary_TossResult_Group>(row, ["toss_result"])) {
    return {
      key: row.toss_result,
      label: `${capitalizeFirstLetter(row.toss_result)} Toss`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Summary_MatchResultBatBowlFirst_Group>(row, ["match_result", "bat_bowl_first"])) {
    return {
      key: row.match_result + row.bat_bowl_first,
      label: `${capitalizeFirstLetter(row.bat_bowl_first)} First & ${row.match_result}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Summary_MatchResult_Group>(row, ["match_result"])) {
    return {
      key: row.match_result,
      label: `${capitalizeFirstLetter(row.match_result)}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Summary_BatBowlFirst_Group>(row, ["bat_bowl_first"])) {
    return {
      key: row.bat_bowl_first,
      label: `${capitalizeFirstLetter(row.bat_bowl_first)} First`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Summary_InningsNumber_Group>(row, ["innings_number"])) {
    return {
      key: row.innings_number,
      label: `Innings No. ${row.innings_number}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Summary_SeriesTeamsCount_Group>(row, ["teams_count"])) {
    return {
      key: row.teams_count,
      label: `${row.teams_count} Teams Series`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Summary_SeriesMatchNumber_Group>(row, ["event_match_number"])) {
    return {
      key: row.event_match_number,
      label: `Match No. ${row.event_match_number}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Series_Group>(row, ["series_id", "series_name"])) {
    return {
      key: row.series_id,
      label: row.series_name + ", " + row.series_season,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Tournament_Group>(row, ["tournament_id", "tournament_name"])) {
    return {
      key: row.tournament_id,
      label: row.tournament_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Decade_Group>(row, ["decade"])) {
    return {
      key: row.decade,
      label: row.decade.toString(),
    };
  }

  if (isT<IOverall_Batting_Summary_BattingPosition_Group>(row, ["batting_position"])) {
    return {
      key: row.batting_position,
      label: `Bat Pos. ${row.batting_position}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Batting_Aggregate_Group>(row, ["players_count"])) {
    return {
      key: row.players_count,
      label: "overall",
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  return { key: "", label: "" };
}
