import { IBatterScorecardEntry, IInningsExtrasData, IInningsTotalData } from "@/lib/types/single-match.types";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../../ui/table";
import Link from "next/link";
import { isBowlerDismissal } from "@/lib/utils";

export function ScorecardBatters({
  batters,
  extrasData,
  totalData,
}: {
  batters: IBatterScorecardEntry[];
  extrasData: IInningsExtrasData;
  totalData: IInningsTotalData;
}) {
  batters.sort((a, b) => a.batting_position - b.batting_position);
  const lastDidNotBatIdx = batters.findLastIndex((entry) => !entry.has_batted);

  return (
    <div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Batting</TableHead>
            <TableHead className="hidden md:table-cell"></TableHead>
            <TableHead>R</TableHead>
            <TableHead>B</TableHead>
            <TableHead>4s</TableHead>
            <TableHead>6s</TableHead>
            <TableHead>SR</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {batters.map((batter) => {
            if (batter.has_batted) {
              return (
                <TableRow key={batter.batter_id}>
                  <TableCell>
                    <Link href={`/players/${batter.batter_id}`} className="underline">
                      {batter.batter_name}
                    </Link>
                    <div className="inline md:hidden">
                      <br />
                      {batter.dismissal_type
                        ? `${batter.dismissal_type} ${
                            isBowlerDismissal(batter.dismissal_type) ? `- b ${batter.dismissed_by_name}` : ""
                          }`
                        : "not out"}
                    </div>
                  </TableCell>
                  <TableCell className="hidden md:table-cell">
                    {batter.dismissal_type
                      ? `${batter.dismissal_type} ${
                          isBowlerDismissal(batter.dismissal_type) ? `- b ${batter.dismissed_by_name}` : ""
                        }`
                      : "not out"}
                  </TableCell>
                  <TableCell className="font-bold">{batter.runs_scored}</TableCell>
                  <TableCell>{batter.balls_faced}</TableCell>
                  <TableCell>{batter.fours_scored}</TableCell>
                  <TableCell>{batter.sixes_scored}</TableCell>
                  <TableCell>
                    {batter.balls_faced ? ((batter.runs_scored * 100) / batter.balls_faced).toFixed(2) : "-"}
                  </TableCell>
                </TableRow>
              );
            }
          })}
        </TableBody>
      </Table>

      <div className="mt-4 px-2 flex justify-between">
        <div className="font-bold">Extras</div>
        <div>
          <span className="font-bold">
            {extrasData.byes +
              extrasData.leg_byes +
              extrasData.wides +
              extrasData.noballs +
              extrasData.penalty}
          </span>{" "}
          <span>
            (b {extrasData.byes}, lb {extrasData.leg_byes}, w {extrasData.wides}, nb {extrasData.noballs}, p{" "}
            {extrasData.penalty})
          </span>
        </div>
      </div>

      <div className="mt-4 px-2 flex justify-between">
        <div className="font-bold">Total</div>
        <div>
          <span className="font-bold">
            {totalData.total_runs}-{totalData.total_wickets}
          </span>{" "}
          <span>
            ({totalData.total_overs} Overs, RR: {(totalData.total_runs / totalData.total_overs).toFixed(2)})
          </span>
        </div>
      </div>

      <div className="mt-4">
        <span className="font-bold">Did not bat:</span> {"  "}
        {batters.map((batter, idx) => {
          if (!batter.has_batted) {
            return (
              <span key={batter.batter_id}>
                <Link href={`/players/${batter.batter_id}`} className="underline">
                  {batter.batter_name}
                </Link>
                {idx < lastDidNotBatIdx ? ", " : ""}
              </span>
            );
          }
        })}
      </div>
    </div>
  );
}
