import { IBowlerScorecardEntry } from "@/lib/types/single-match.types";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../../ui/table";
import Link from "next/link";

export function ScorecardBowlers({ bowlers }: { bowlers: IBowlerScorecardEntry[] }) {
  bowlers.sort((a, b) => a.bowling_position - b.bowling_position);

  return (
    <div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Bowling</TableHead>
            <TableHead>O</TableHead>
            <TableHead>M</TableHead>
            <TableHead>R</TableHead>
            <TableHead>W</TableHead>
            <TableHead>ECON</TableHead>
            <TableHead className="hidden md:table-cell">4s</TableHead>
            <TableHead className="hidden md:table-cell">6s</TableHead>
            <TableHead>WD</TableHead>
            <TableHead>NB</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {bowlers.map((bowler) => {
            if (bowler.bowling_position) {
              return (
                <TableRow key={bowler.bowler_id}>
                  <TableCell className="underline" width={"30%"}>
                    <Link href={`/players/${bowler.bowler_id}`}>{bowler.bowler_name}</Link>
                  </TableCell>
                  <TableCell>{bowler.overs_bowled}</TableCell>
                  <TableCell>{bowler.maiden_overs}</TableCell>
                  <TableCell>{bowler.runs_conceded}</TableCell>
                  <TableCell className="font-bold">{bowler.wickets_taken}</TableCell>
                  <TableCell>
                    {bowler.overs_bowled ? (bowler.runs_conceded / bowler.overs_bowled).toFixed(2) : "-"}
                  </TableCell>
                  <TableCell className="hidden md:table-cell">{bowler.fours_conceded}</TableCell>
                  <TableCell className="hidden md:table-cell">{bowler.sixes_conceded}</TableCell>
                  <TableCell>{bowler.wides_conceded}</TableCell>
                  <TableCell>{bowler.noballs_conceded}</TableCell>
                </TableRow>
              );
            }
          })}
        </TableBody>
      </Table>
    </div>
  );
}
