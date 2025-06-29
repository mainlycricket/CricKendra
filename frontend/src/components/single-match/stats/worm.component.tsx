"use client";

import { CartesianGrid, Line, LineChart, XAxis, YAxis } from "recharts";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { IOverStats } from "@/lib/types/single-match.types";

export function StatsWorm({
  data,
}: {
  data: {
    inningsId: number;
    inningsNumber: number;
    battingTeamName: string;
    overs: IOverStats[];
  }[];
}) {
  const maxOvers = Math.max(...data.map((item) => item.overs.length));
  const chartData: { overNumber: number; team1Runs: number; team2Runs: number }[] = [];

  let team1Runs = 0,
    team2Runs = 0;

  for (let i = 0; i < maxOvers; i++) {
    team1Runs += data?.[0]?.overs?.[i]?.runs || 0;
    team2Runs += data?.[1]?.overs?.[i]?.runs || 0;
    chartData.push({ overNumber: i + 1, team1Runs, team2Runs });
  }

  const chartConfig: ChartConfig = {};

  if (data.length > 0) {
    chartConfig.team1Runs = {
      label: data?.[0]?.battingTeamName,
      color: "var(--chart-1)",
    };
  }

  if (data.length > 1) {
    chartConfig.team2Runs = {
      label: data?.[1]?.battingTeamName,
      color: "var(--chart-2)",
    };
  }

  return (
    <Card className="gap-0">
      <CardHeader>
        <CardTitle className="px-4 md:px-12 text-xl">Worm</CardTitle>
      </CardHeader>
      <CardContent className="pl-0 pr-2 py-0 md:px-8">
        <ChartContainer config={chartConfig} className="min-h-[100px] max-h-[500px] w-full">
          <LineChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <XAxis
              label={{ value: "Over Number", position: "insideBottom", offset: -5 }}
              dataKey="overNumber"
            />
            <YAxis label={{ value: "Runs", angle: -90, position: "left", offset: -20 }} />
            <ChartTooltip cursor={true} content={<ChartTooltipContent hideLabel={true} />} />
            <ChartLegend verticalAlign="top" content={<ChartLegendContent />} />
            <Line dataKey="team1Runs" stroke="var(--color-team1Runs)" strokeWidth={2} dot={false} />
            <Line dataKey="team2Runs" stroke="var(--color-team2Runs)" strokeWidth={2} dot={false} />
          </LineChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
