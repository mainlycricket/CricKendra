import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";

import {
  ICombinedBattingStatsType2,
  IIndividual_Batting_Innings_Group,
  IIndividual_Batting_MatchTotals_Group,
  isT2,
} from "@/lib/types/batting-stats.types";
import { getDisplayDate } from "@/lib/utils";
import Link from "next/link";

export function BattingListTable({ stats }: { stats: ICombinedBattingStatsType2[] }) {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Batter</TableHead>
          <TableHead>Runs</TableHead>
          {/* {isT2<IIndividual_Batting_MatchTotals_Group>(stats?.[0], ["innings"]) && (
            <TableHead>Innings Wise</TableHead>
          )} */}
          <TableHead title="Balls Faced">BF</TableHead>
          <TableHead title="Strike Rate">SR</TableHead>
          <TableHead className="hidden md:table-cell" title="Fours">
            4s
          </TableHead>
          <TableHead className="hidden md:table-cell" title="Sixes">
            6s
          </TableHead>
          {isT2<IIndividual_Batting_Innings_Group>(stats?.[0], ["innings_number"]) && (
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
            <TableRow key={`${row.batter_id}_${row.match_id}`}>
              <TableCell className="font-medium">
                {row.batter_name} ({row.batting_team_name})
              </TableCell>
              <TableCell>
                {row.runs_scored}
                {isT2<IIndividual_Batting_Innings_Group>(row, ["innings_number"]) && row.is_not_out && "*"}
              </TableCell>

              {/* {isT2<IIndividual_Batting_MatchTotals_Group>(row, ["innings"]) && (
                <TableHead>
                  {row?.innings
                    ?.map((item) => `${item.runs_scored}${item.is_not_out ? "*" : ""}`)
                    ?.join(", ")}
                </TableHead>
              )} */}
              <TableCell>{row.balls_faced}</TableCell>
              <TableCell>{row.strike_rate.toFixed(2)}</TableCell>
              <TableCell className="hidden md:table-cell">{row.fours_scored}</TableCell>
              <TableCell className="hidden md:table-cell">{row.sixes_scored}</TableCell>
              {isT2<IIndividual_Batting_Innings_Group>(row, ["innings_number"]) && (
                <TableCell>{row.innings_number}</TableCell>
              )}
              <TableCell>{row.bowling_team_name}</TableCell>
              <TableCell>{getDisplayDate(new Date(row.start_date))}</TableCell>
              <TableCell className="hidden md:table-cell">{row.city_name}</TableCell>
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
