import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";

import {
  ICombinedTeamStatsType2,
  IIndividual_Team_Innings_Group,
  IIndividual_Team_MatchResults_Group,
  IIndividual_Team_MatchTotals_Group,
  isT2,
} from "@/lib/types/team-stats.types";
import { getDisplayDate } from "@/lib/utils";
import Link from "next/link";

export function TeamListTable({ stats }: { stats: ICombinedTeamStatsType2[] }) {
  const group = isT2<IIndividual_Team_Innings_Group>(stats?.[0] || {}, ["innings_id"])
    ? "innings"
    : isT2<IIndividual_Team_MatchTotals_Group>(stats?.[0] || {}, ["total_balls"])
    ? "match_totals"
    : "match_results";

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Team</TableHead>
          <TableHead>Result</TableHead>
          {group === "match_results" && <TableHead title="Win Margin">Margin</TableHead>}
          {group === "innings" && <TableHead title="Innings">Inns</TableHead>}
          {group !== "match_results" && <TableHead>Runs</TableHead>}
          {group === "innings" && <TableHead title="Overs">Overs</TableHead>}
          {group === "match_totals" && <TableHead>Balls</TableHead>}
          {group !== "match_results" && <TableHead title="Wickets">Wkts</TableHead>}
          {group === "match_totals" && <TableHead title="Average">Ave</TableHead>}
          {group !== "match_results" && <TableHead title="Scoring Rate (RPO)">RPO</TableHead>}
          {group === "match_results" && <TableHead title="Toss Result">Toss</TableHead>}
          <TableHead>Opposition</TableHead>
          <TableHead className="hidden md:table-cell">City</TableHead>
          <TableHead title="Match Start Date">Date</TableHead>
          <TableHead></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {stats.map((row) => {
          const key = isT2<IIndividual_Team_Innings_Group>(row, ["innings_id"])
            ? `${row.match_id}_${row.team_id}_${row.innings_id}`
            : `${row.match_id}_${row.team_id}`;
          return (
            <TableRow key={key}>
              <TableCell className="font-medium">{row.team_name}</TableCell>
              <TableCell className="capitalize">
                {row.final_result === "winner decided"
                  ? row.match_winner_id === row.team_id
                    ? "won"
                    : "lost"
                  : row.final_result}
              </TableCell>
              {isT2<IIndividual_Team_MatchResults_Group>(row, ["win_margin"]) && (
                <TableCell>
                  {row.is_won_by_innings && "inns &"} {row.win_margin}{" "}
                  {row.is_won_by_runs ? "runs" : "wickets"}
                </TableCell>
              )}
              {isT2<IIndividual_Team_Innings_Group>(row, ["innings_id"]) && (
                <TableCell>{row.innings_number}</TableCell>
              )}
              {!isT2<IIndividual_Team_MatchResults_Group>(row, ["win_margin"]) && (
                <TableCell>{row.total_runs}</TableCell>
              )}
              {isT2<IIndividual_Team_Innings_Group>(row, ["innings_id"]) && (
                <TableCell>{row.total_overs}</TableCell>
              )}
              {isT2<IIndividual_Team_MatchTotals_Group>(row, ["total_balls"]) && (
                <TableCell>{row.total_balls}</TableCell>
              )}
              {!isT2<IIndividual_Team_MatchResults_Group>(row, ["win_margin"]) && (
                <TableCell>{row.total_wickets}</TableCell>
              )}
              {isT2<IIndividual_Team_MatchTotals_Group>(row, ["average"]) && (
                <TableCell>{row?.average?.toFixed(2) || "-"}</TableCell>
              )}
              {!isT2<IIndividual_Team_MatchResults_Group>(row, ["win_margin"]) && (
                <TableCell>{row?.scoring_rate?.toFixed(2) || "-"}</TableCell>
              )}
              {isT2<IIndividual_Team_MatchResults_Group>(row, ["win_margin"]) && (
                <TableCell className="capitalize">
                  {row.toss_winner_id === row.team_id ? "won" : "lost"}
                </TableCell>
              )}
              <TableCell>{row.opposition_name}</TableCell>
              <TableCell className="hidden md:table-cell">{row.city_name}</TableCell>
              <TableCell>{getDisplayDate(new Date(row.start_date))}</TableCell>
              <TableCell>
                <Link href={`/matches/${row.match_id}`} className="underline" target="_blank">
                  View Match
                </Link>
              </TableCell>
            </TableRow>
          );
        })}
      </TableBody>
    </Table>
  );
}
