import Image from "next/image";
import Link from "next/link";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

import { ISeriesTopBatter } from "@/lib/types/single-series.types";

export function TopBatters({ topBatters }: { topBatters: ISeriesTopBatter[] }) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-center text-xl">Top Batters</CardTitle>
      </CardHeader>

      <CardContent className="flex flex-col gap-6">
        {topBatters?.length &&
          topBatters.map((batter) => {
            return (
              <div key={batter.batter_id} className="flex flex-col gap-4">
                <hr />
                <div className="flex px-4 gap-8 items-center">
                  <div className="flex flex-col">
                    <Link href={`/players/${batter.batter_id}`}>
                      <Image
                        src={`${batter.batter_image_url || "/file.svg"}`}
                        width={50}
                        height={50}
                        alt={batter.batter_name}
                        className="rounded-full"
                      />
                    </Link>
                  </div>
                  <div className="flex flex-col gap-1">
                    <Link
                      href={`/players/${batter.batter_id}`}
                      className="font-medium text-lg hover:text-sky-500"
                    >
                      {batter.batter_name}
                    </Link>
                    <p className="text-2xl font-bold">{batter.runs_scored}</p>
                    <p className="flex gap-2 font-light">
                      <span>Innings: {batter.innings_batted}</span>
                      <span>Average: {batter?.average?.toFixed(2) || "-"}</span>
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
