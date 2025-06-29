"use client";

import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";
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

export function StatsManhattan({
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

  for (let i = 0; i < maxOvers; i++) {
    const team1Runs = data?.[0]?.overs?.[i]?.runs || 0;
    const team2Runs = data?.[1]?.overs?.[i]?.runs || 0;
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
        <CardTitle className="px-4 md:px-12 text-xl">Manhattan</CardTitle>
      </CardHeader>
      <CardContent className="p-0 pr-2 md:p-8">
        <ChartContainer config={chartConfig} className="min-h-[150px] max-h-[500px] w-full">
          <BarChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <XAxis
              label={{ value: "Over Number", position: "insideBottom", offset: -5 }}
              dataKey="overNumber"
            />
            <YAxis label={{ value: "Runs", angle: -90, position: "left", offset: -20 }} />
            <ChartTooltip cursor={true} content={<ChartTooltipContent hideLabel={true} />} />
            <ChartLegend verticalAlign="top" content={<ChartLegendContent />} />
            <Bar dataKey="team1Runs" fill="var(--color-team1Runs)" radius={4} />
            <Bar dataKey="team2Runs" fill="var(--color-team2Runs)" radius={4} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
