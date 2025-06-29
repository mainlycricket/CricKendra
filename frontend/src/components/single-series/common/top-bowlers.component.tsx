import Image from "next/image";
import Link from "next/link";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

import { ISeriesTopBowler } from "@/lib/types/single-series.types";

export function TopBowlers({ topBowlers }: { topBowlers: ISeriesTopBowler[] }) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-center text-xl">Top Bowlers</CardTitle>
      </CardHeader>

      <CardContent className="flex flex-col gap-6">
        {topBowlers?.length &&
          topBowlers.map((bowler) => {
            return (
              <div key={bowler.bowler_id} className="flex flex-col gap-4">
                <hr />

                <div className="flex px-4 gap-8 items-center">
                  <div className="flex flex-col">
                    <Link href={`/players/${bowler.bowler_id}`}>
                      <Image
                        src={`${bowler.bowler_image_url || "/file.svg"}`}
                        width={50}
                        height={50}
                        alt={bowler.bowler_name}
                        className="rounded-full"
                      />
                    </Link>
                  </div>
                  <div className="flex flex-col gap-1">
                    <Link
                      href={`/players/${bowler.bowler_id}`}
                      className="font-medium text-lg hover:text-sky-500"
                    >
                      {bowler.bowler_name}
                    </Link>
                    <p className="text-2xl font-bold">{bowler.wickets_taken}</p>
                    <p className="flex gap-2 font-light">
                      <span>Innings: {bowler.innings_bowled}</span>
                      <span>Average: {bowler?.average?.toFixed(2) || "-"}</span>
                    </p>
                  </div>
                </div>
              </div>
            );
          })}
      </CardContent>
    </Card>
  );
}
