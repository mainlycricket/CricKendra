import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";

import {
  ICombinedTeamStatsType,
  IIndividual_Team_Grounds_Group,
  IIndividual_Team_HostNations_Group,
  IIndividual_Team_Seasons_Group,
  IIndividual_Team_Series_Group,
  IIndividual_Team_Tournaments_Group,
  IIndividual_Team_Years_Group,
  IOverall_Team_Aggregate_Group,
  IOverall_Team_Continents_Group,
  IOverall_Team_Decades_Group,
  IOverall_Team_Grounds_Group,
  IOverall_Team_HostNations_Group,
  IOverall_Team_Matches_Group,
  IOverall_Team_Players_Group,
  IOverall_Team_Seasons_Group,
  IOverall_Team_Series_Group,
  IOverall_Team_Teams_Group,
  IOverall_Team_Tournament_Group,
  IOverall_Team_Years_Group,
  isT,
} from "@/lib/types/team-stats.types";

export function TeamCommonTable({ stats }: { stats: ICombinedTeamStatsType[] }) {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead></TableHead>
          {!isT<IOverall_Team_Years_Group>(stats?.[0], ["year"]) &&
            !isT<IOverall_Team_Seasons_Group>(stats?.[0], ["season"]) &&
            !isT<IOverall_Team_Decades_Group>(stats?.[0], ["decade"]) && (
              <TableHead className="hidden md:table-cell">Span</TableHead>
            )}
          <TableHead title="Matches Played">Mat</TableHead>
          <TableHead title="Matches Won">Won</TableHead>
          <TableHead title="Matches Lost">Lost</TableHead>
          <TableHead title="Win/Loss Ratio">W/L</TableHead>
          <TableHead className="hidden md:table-cell" title="Matches Drawn">
            Drawn
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Matches Tied">
            Tied
          </TableHead>
          <TableHead className="hidden md:table-cell" title="N/R Matches">
            NR
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Innings">
            Inns
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Runs">
            Runs
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Balls">
            Balls
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Wickets">
            Wkts
          </TableHead>
          <TableHead title="Average">Ave</TableHead>
          <TableHead title="Scoring Rate (RPO)">RPO</TableHead>
          <TableHead className="hidden md:table-cell" title="Highest Score">
            HS
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Lowest Score">
            LS
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
              <TableCell>{row.matches_played}</TableCell>
              <TableCell>{row.matches_won}</TableCell>
              <TableCell>{row.matches_lost}</TableCell>
              <TableCell>{row?.win_loss_ratio?.toFixed(2) || "-"}</TableCell>
              <TableCell className="hidden md:table-cell">{row.matches_drawn}</TableCell>
              <TableCell className="hidden md:table-cell">{row.matches_tied}</TableCell>
              <TableCell className="hidden md:table-cell">{row.matches_no_result}</TableCell>
              <TableCell className="hidden md:table-cell">{row.innings_count}</TableCell>
              <TableCell className="hidden md:table-cell">{row.total_runs}</TableCell>
              <TableCell className="hidden md:table-cell">{row.total_balls}</TableCell>
              <TableCell className="hidden md:table-cell">{row.total_wickets}</TableCell>
              <TableCell>{row?.average?.toFixed(2) || "-"}</TableCell>
              <TableCell>{row?.scoring_rate?.toFixed(2) || "-"}</TableCell>
              <TableCell className="hidden md:table-cell">{row.highest_score}</TableCell>
              <TableCell className="hidden md:table-cell">{row.lowest_score || "-"}</TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}

function getRowMetaData(row: ICombinedTeamStatsType): {
  key: number | string;
  label: string;
  minSpanYear?: number;
  maxSpanYear?: number;
} {
  /* Individual */
  if (isT<IIndividual_Team_Series_Group>(row, ["team_id", "series_id"])) {
    return {
      key: `${row.team_id}_${row.series_id}`,
      label: `${row.team_name} in ${row.series_name}, ${row.series_season}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Team_Tournaments_Group>(row, ["team_id", "tournament_id"])) {
    return {
      key: `${row.team_id}_${row.tournament_id}`,
      label: `${row.team_name} in ${row.tournament_name}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Team_Grounds_Group>(row, ["team_id", "ground_id"])) {
    return {
      key: `${row.team_id}_${row.ground_id}`,
      label: `${row.team_name} at ${row.ground_name}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Team_HostNations_Group>(row, ["team_id", "host_nation_id"])) {
    return {
      key: `${row.team_id}_${row.host_nation_id}`,
      label: `${row.team_name} in ${row.host_nation_name}`,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IIndividual_Team_Years_Group>(row, ["team_id", "year"])) {
    return {
      key: `${row.team_id}_${row.year}`,
      label: `${row.team_name}, ${row.year}`,
    };
  }

  if (isT<IIndividual_Team_Seasons_Group>(row, ["team_id", "season"])) {
    return {
      key: `${row.team_id}_${row.season}`,
      label: `${row.team_name}, ${row.season}`,
    };
  }

  /* Overall */
  if (isT<IOverall_Team_Players_Group>(row, ["player_id", "player_name"])) {
    return {
      key: row.player_id,
      label: row.player_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Team_Matches_Group>(row, ["match_id"])) {
    return {
      key: row.match_id,
      label: `${row.team1_name} v ${row.team2_name} at ${row.city_name}, ${row.season}`,
    };
  }

  if (isT<IOverall_Team_Teams_Group>(row, ["team_id", "team_name"])) {
    return {
      key: row.team_id,
      label: row.team_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Team_Grounds_Group>(row, ["ground_id", "ground_name"])) {
    return {
      key: row.ground_id,
      label: row.ground_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Team_HostNations_Group>(row, ["host_nation_id", "host_nation_name"])) {
    return {
      key: row.host_nation_id,
      label: row.host_nation_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Team_Continents_Group>(row, ["continent_id", "continent_name"])) {
    return {
      key: row.continent_id,
      label: row.continent_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Team_Years_Group>(row, ["year"])) {
    return {
      key: row.year,
      label: row.year.toString(),
    };
  }

  if (isT<IOverall_Team_Seasons_Group>(row, ["season"])) {
    return {
      key: row.season,
      label: row.season,
    };
  }

  if (isT<IOverall_Team_Series_Group>(row, ["series_id", "series_name"])) {
    return {
      key: row.series_id,
      label: row.series_name + ", " + row.series_season,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Team_Tournament_Group>(row, ["tournament_id", "tournament_name"])) {
    return {
      key: row.tournament_id,
      label: row.tournament_name,
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  if (isT<IOverall_Team_Decades_Group>(row, ["decade"])) {
    return {
      key: row.decade,
      label: row.decade.toString(),
    };
  }

  if (isT<IOverall_Team_Aggregate_Group>(row, ["teams_count"])) {
    return {
      key: row.teams_count,
      label: "overall",
      minSpanYear: new Date(row.min_date).getFullYear(),
      maxSpanYear: new Date(row.max_date).getFullYear(),
    };
  }

  return { key: "", label: "" };
}
