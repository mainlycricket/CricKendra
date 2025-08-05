import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";

import {
  ICombinedBowlingStatsType2,
  IIndividual_Bowling_Innings_Group,
  isT2,
} from "@/lib/types/bowling-stats.types";
import { getDisplayDate } from "@/lib/utils";
import Link from "next/link";

export function BowlingListTable({ stats }: { stats: ICombinedBowlingStatsType2[] }) {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Bowler</TableHead>
          <TableHead>Overs</TableHead>
          <TableHead title="Wickets">Wkts</TableHead>
          <TableHead className="hidden md:table-cell" title="Maiden Overs">
            Mdns
          </TableHead>
          <TableHead title="Runs Conceded">Runs</TableHead>
          <TableHead className="hidden md:table-cell" title="Economy">
            Econ
          </TableHead>
          <TableHead className="hidden md:table-cell">4s</TableHead>
          <TableHead className="hidden md:table-cell">6s</TableHead>
          {isT2<IIndividual_Bowling_Innings_Group>(stats?.[0], ["innings_number"]) && (
            <TableHead title="Innings Number">Inns</TableHead>
          )}
          <TableHead title="Opposition">Opposition</TableHead>
          <TableHead title="Match Date">Match Date</TableHead>
          <TableHead className="hidden md:table-cell" title="City">
            City
          </TableHead>
          <TableHead></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {stats.map((row) => {
          return (
            <TableRow key={`${row.bowler_id}_${row.match_id}`}>
              <TableCell>{row.bowler_name}</TableCell>
              <TableCell>{row.overs_bowled}</TableCell>
              <TableCell>{row.wickets_taken}</TableCell>
              <TableCell className="hidden md:table-cell">{row.maiden_overs}</TableCell>
              <TableCell>{row.runs_conceded}</TableCell>
              <TableCell className="hidden md:table-cell">{row?.economy?.toFixed(2) || "-"}</TableCell>
              <TableCell className="hidden md:table-cell">{row.fours_conceded}</TableCell>
              <TableCell className="hidden md:table-cell">{row.sixes_conceded}</TableCell>
              {isT2<IIndividual_Bowling_Innings_Group>(row, ["innings_number"]) && (
                <TableCell>{row.innings_number}</TableCell>
              )}
              <TableCell>{row.batting_team_name}</TableCell>
              <TableCell>{getDisplayDate(new Date(row.start_date))}</TableCell>
              <TableCell className="hidden md:table-cell">{row.city_name}</TableCell>
              <TableCell>
                <Link href={`/matches/${row.match_id}`} target="_blank" className="underline">
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
