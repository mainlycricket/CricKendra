import { PopulatedTable } from "@/components/common/table.component";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { ICareerStats, ISinglePlayer } from "@/lib/types/single-player.types";
import { rotate2DArray } from "@/lib/utils";

export function PlayerCareerStats({ player }: { player: ISinglePlayer }) {
  const battingStats = getBattingStats(player);
  const bowlingStats = getBowlingStats(player);

  const battingInnings =
    (player?.fc_stats?.innings_batted || 0) +
    (player?.lista_stats?.innings_batted || 0) +
    (player?.t20_stats?.innings_batted || 0);

  const bowlingInnings =
    (player?.fc_stats?.innings_bowled || 0) +
    (player?.lista_stats?.innings_bowled || 0) +
    (player?.t20_stats?.innings_bowled || 0);

  return (
    <Card className="gap-4">
      <CardHeader>
        <CardTitle className="text-lg px-2">{player.name} Career Stats</CardTitle>
      </CardHeader>
      <CardContent>
        <MobileStats
          battingStats={rotate2DArray(battingStats)}
          bowlingStats={rotate2DArray(bowlingStats)}
          isDefaultBatting={battingInnings > bowlingInnings}
        />
        <DesktopStats
          battingStats={battingStats}
          bowlingStats={bowlingStats}
          isDefaultBatting={battingInnings > bowlingInnings}
        />
      </CardContent>
    </Card>
  );
}

function DesktopStats({
  battingStats,
  bowlingStats,
  isDefaultBatting,
}: {
  battingStats: string[][];
  bowlingStats: string[][];
  isDefaultBatting: boolean;
}) {
  return (
    <div className={`hidden md:flex md:gap-4 ${isDefaultBatting ? "md:flex-col" : "md:flex-col-reverse"}`}>
      <div>
        <p className="px-2 pb-2 font-medium uppercase tracking-wider">Batting</p>
        <hr />
        <PopulatedTable data={battingStats} isFirstColHead={true} />
      </div>
      <div>
        <p className="px-2 pb-2 font-medium uppercase tracking-wider">Bowling</p>
        <hr />
        <PopulatedTable data={bowlingStats} isFirstColHead={true} />
      </div>
    </div>
  );
}

function MobileStats({
  battingStats,
  bowlingStats,
  isDefaultBatting,
}: {
  battingStats: string[][];
  bowlingStats: string[][];
  isDefaultBatting: boolean;
}) {
  return (
    <div className="md:hidden">
      <Tabs defaultValue={isDefaultBatting ? "batting" : "bowling"}>
        <TabsList className="w-full">
          <TabsTrigger value="batting">Batting</TabsTrigger>
          <TabsTrigger value="bowling">Bowling</TabsTrigger>
        </TabsList>
        <TabsContent value="batting">
          <PopulatedTable data={battingStats} isFirstColHead={true} />
        </TabsContent>
        <TabsContent value="bowling">
          <PopulatedTable data={bowlingStats} isFirstColHead={true} />
        </TabsContent>
      </Tabs>
    </div>
  );
}

function getBattingStats(player: ISinglePlayer): string[][] {
  const rows: string[][] = [
    ["Format", "Mat", "Inn", "NO", "Runs", "HS", "Ave", "BF", "SR", "100s", "50s", "4s", "6s"],
  ];

  if (player.test_stats) {
    const testRow = getBattingStatsRow("Test", player.test_stats);
    rows.push(testRow);
  }

  if (player.odi_stats) {
    const odiRow = getBattingStatsRow("ODI", player.odi_stats);
    rows.push(odiRow);
  }

  if (player.t20i_stats) {
    const t20iRow = getBattingStatsRow("T20I", player.t20i_stats);
    rows.push(t20iRow);
  }

  if (player.fc_stats) {
    const fcRow = getBattingStatsRow("First Class", player.fc_stats);
    rows.push(fcRow);
  }

  if (player.lista_stats) {
    const listaRow = getBattingStatsRow("List A", player.lista_stats);
    rows.push(listaRow);
  }

  if (player.t20_stats) {
    const t20Row = getBattingStatsRow("T20", player.t20_stats);
    rows.push(t20Row);
  }

  return rows;
}

function getBattingStatsRow(playingFormat: string, stats: ICareerStats) {
  const matchesPlayed = stats?.matches_played?.toString() || "-";

  if (!stats.innings_batted)
    return [playingFormat, matchesPlayed, "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"];

  const dismissals = stats.innings_batted - stats.not_outs;
  const average = dismissals ? (stats.runs_scored / (stats.innings_batted - stats.not_outs)).toFixed(2) : "-";
  const strikeRate = stats.balls_faced ? ((stats.runs_scored * 100) / stats.balls_faced).toFixed(2) : "-";

  const row: string[] = [
    playingFormat,
    matchesPlayed,
    stats.innings_batted.toString(),
    stats.not_outs.toString(),
    stats.runs_scored.toString(),
    `${stats.highest_score}${stats.is_highest_not_out ? "*" : ""}`,
    average,
    stats.balls_faced.toString(),
    strikeRate,
    stats.centuries_scored.toString(),
    stats.fifties_scored.toString(),
    stats.fours_scored.toString(),
    stats.sixes_scored.toString(),
  ];

  return row;
}

function getBowlingStats(player: ISinglePlayer): string[][] {
  const rows: string[][] = [
    ["Format", "Mat", "Inn", "Balls", "Runs", "Wkts", "BBI", "BBM", "Ave", "Econ", "SR", "4w", "5w", "10w"],
  ];

  if (player.test_stats) {
    const testRow = getBowlingStatsRow("Test", player.test_stats);
    rows.push(testRow);
  }

  if (player.odi_stats) {
    const odiRow = getBowlingStatsRow("ODI", player.odi_stats);
    rows.push(odiRow);
  }

  if (player.t20i_stats) {
    const t20iRow = getBowlingStatsRow("T20I", player.t20i_stats);
    rows.push(t20iRow);
  }

  if (player.fc_stats) {
    const fcRow = getBowlingStatsRow("First Class", player.fc_stats);
    rows.push(fcRow);
  }

  if (player.lista_stats) {
    const listaRow = getBowlingStatsRow("List A", player.lista_stats);
    rows.push(listaRow);
  }

  if (player.t20_stats) {
    const t20Row = getBowlingStatsRow("T20", player.t20_stats);
    rows.push(t20Row);
  }

  return rows;
}

function getBowlingStatsRow(playingFormat: string, stats: ICareerStats) {
  const matchesPlayed = stats?.matches_played?.toString() || "-";

  if (!stats.innings_bowled)
    return [playingFormat, matchesPlayed, "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"];

  const bestMatchFigures = stats.best_match_fig_wkts
    ? `${stats.best_match_fig_wkts}/${stats.best_match_fig_runs}`
    : "-";
  const bestInningsFigures = stats.best_inn_fig_wkts
    ? `${stats.best_inn_fig_wkts}/${stats.best_inn_fig_runs}`
    : "-";

  const average = stats.wickets_taken ? (stats.runs_conceded / stats.wickets_taken).toFixed(2) : "-";
  const economy = stats.balls_bowled ? ((stats.runs_conceded / stats.balls_bowled) * 6).toFixed(2) : "-";
  const strikeRate = stats.wickets_taken ? (stats.balls_bowled / stats.wickets_taken).toFixed(2) : "-";

  const row: string[] = [
    playingFormat,
    matchesPlayed,
    stats.innings_bowled.toString(),
    stats.balls_bowled.toString(),
    stats.runs_conceded.toString(),
    stats.wickets_taken.toString(),
    bestInningsFigures,
    bestMatchFigures,
    average,
    economy,
    strikeRate,
    stats.four_wkt_hauls.toString(),
    stats.five_wkt_hauls.toString(),
    stats?.ten_wkt_hauls?.toString() || "0",
  ];

  return row;
}
