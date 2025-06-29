"use client";

import { Card, CardContent } from "@/components/ui/card";
import { StatsDropdown } from "./dropdown-options.component";
import { BattingStats } from "./batting-stats.component";
import { BowlingStats } from "./bowling-stats.component";
import { ISinglePlayer } from "@/lib/types/single-player.types";
import { useEffect, useState } from "react";
import { IOverall_Batting_Summary_Group } from "@/lib/types/batting-stats.types";
import { IOverall_Bowling_Summary_Group } from "@/lib/types/bowling-stats.types";
import { doBackendRequest } from "@/lib/axiosFetch";

export function MainStats({ player }: { player: ISinglePlayer }) {
  const formatOptions = getFormatOptions(player);
  const defaultFormat = getDefaultFormat(player);
  const battingBowlingOptions = getBattingBowlingOptions(player);
  const isDefaultBatting = isDefaultBattingStats(player);
  const [format, setFormat] = useState(defaultFormat);
  const [isBattingStats, setIsBattingStats] = useState(isDefaultBatting);
  const [battingStats, setBattingStats] = useState<IOverall_Batting_Summary_Group | null>(null);
  const [bowlingStats, setBowlingStats] = useState<IOverall_Bowling_Summary_Group | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (format && isBattingStats) {
      setIsLoading(true);
      doBackendRequest<null, IOverall_Batting_Summary_Group>({
        url: `/stats/batting/overall/summary?batter=${player.id}&playing_format=${format}`,
        method: "GET",
      })
        .then((data) => {
          setBattingStats(data.data!);
        })
        .finally(() => setIsLoading(false));
    } else if (format && !isBattingStats) {
      setIsLoading(true);
      doBackendRequest<null, IOverall_Bowling_Summary_Group>({
        url: `/stats/bowling/overall/summary?bowler=${player.id}&playing_format=${format}`,
        method: "GET",
      })
        .then((data) => {
          setBowlingStats(data.data!);
        })
        .finally(() => setIsLoading(false));
    }
  }, [player.id, format, isBattingStats]);

  return (
    <Card>
      <CardContent>
        {format === "" && <p>No stats found</p>}
        {format !== "" && (
          <div className="flex flex-col gap-4">
            <StatsDropdown
              formatOptions={formatOptions}
              defaultFormat={format}
              setFormat={setFormat}
              battingBowlingOptions={battingBowlingOptions}
              isDefaultBatting={isBattingStats}
              setIsBattingStats={setIsBattingStats}
            />
            <hr />
          </div>
        )}
        {isLoading && <p className="p-4 text-center">Loading...</p>}
        {!isLoading && isBattingStats && battingStats && <BattingStats stats={battingStats!} />}
        {!isLoading && !isBattingStats && bowlingStats && <BowlingStats stats={bowlingStats!} />}
      </CardContent>
    </Card>
  );
}

function getDefaultFormat(player: ISinglePlayer) {
  const format = player?.test_stats
    ? "Test"
    : player?.odi_stats
    ? "ODI"
    : player?.t20i_stats
    ? "T20I"
    : player?.fc_stats
    ? "first_class"
    : player?.lista_stats
    ? "list_a"
    : player?.t20_stats
    ? "T20"
    : "";

  return format;
}

function isDefaultBattingStats(player: ISinglePlayer) {
  const battingInnings =
    (player?.test_stats?.innings_batted || 0) +
    (player?.odi_stats?.innings_batted || 0) +
    (player?.t20i_stats?.innings_batted || 0) +
    (player?.fc_stats?.innings_batted || 0) +
    (player?.lista_stats?.innings_batted || 0) +
    (player?.t20_stats?.innings_batted || 0);

  const bowlingInnings =
    (player?.test_stats?.innings_bowled || 0) +
    (player?.odi_stats?.innings_bowled || 0) +
    (player?.t20i_stats?.innings_bowled || 0) +
    (player?.fc_stats?.innings_bowled || 0) +
    (player?.lista_stats?.innings_bowled || 0) +
    (player?.t20_stats?.innings_bowled || 0);

  return battingInnings > bowlingInnings;
}

function getBattingBowlingOptions(player: ISinglePlayer) {
  const options: { value: string; label: string }[] = [];

  const battingInnings =
    (player?.test_stats?.innings_batted || 0) +
    (player?.odi_stats?.innings_batted || 0) +
    (player?.t20i_stats?.innings_batted || 0) +
    (player?.fc_stats?.innings_batted || 0) +
    (player?.lista_stats?.innings_batted || 0) +
    (player?.t20_stats?.innings_batted || 0);

  const bowlingInnings =
    (player?.test_stats?.innings_bowled || 0) +
    (player?.odi_stats?.innings_bowled || 0) +
    (player?.t20i_stats?.innings_bowled || 0) +
    (player?.fc_stats?.innings_bowled || 0) +
    (player?.lista_stats?.innings_bowled || 0) +
    (player?.t20_stats?.innings_bowled || 0);

  if (battingInnings) options.push({ value: "batting", label: "Batting" });
  if (bowlingInnings) options.push({ value: "bowling", label: "Bowling" });

  return options;
}

function getFormatOptions(player: ISinglePlayer) {
  const options: { value: string; label: string }[] = [];

  if (player?.test_stats) {
    options.push({ value: "Test", label: "Test" });
  }
  if (player?.odi_stats) {
    options.push({ value: "ODI", label: "ODI" });
  }
  if (player?.t20i_stats) {
    options.push({ value: "T20I", label: "T20I" });
  }
  if (player?.fc_stats) {
    options.push({ value: "first_class", label: "First Class" });
  }
  if (player?.lista_stats) {
    options.push({ value: "list_a", label: "List A" });
  }
  if (player?.t20_stats) {
    options.push({ value: "T20", label: "T20" });
  }

  return options;
}
