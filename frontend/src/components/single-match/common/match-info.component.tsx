import Image from "next/image";

import { Card, CardContent, CardFooter, CardHeader } from "../../ui/card";

import { IMatchInfo, ITeamInningsShortInfo } from "@/lib/types/single-match.types";
import { getDisplayDate } from "@/lib/utils";

export function MatchInfo({ matchInfo }: { matchInfo: IMatchInfo }) {
  matchInfo?.innings_scores?.sort((a, b) => a.innings_number - b.innings_number);

  let { team1_id, team1_name, team1_image_url, team2_id, team2_name, team2_image_url } = matchInfo;

  if (matchInfo?.innings_scores?.[0]?.batting_team_id === team2_id) {
    const team_id = team2_id,
      team_name = team2_name,
      team_image_url = team2_image_url;
    team2_id = team1_id;
    team2_name = team1_name;
    team2_image_url = team1_image_url;
    team1_id = team_id;
    team1_name = team_name;
    team1_image_url = team_image_url;
  }

  return (
    <Card className="gap-4 py-4">
      <CardHeader>
        <p className="font-bold uppercase">
          {matchInfo.match_state === "break" ? matchInfo.match_state_description : matchInfo.match_state}
        </p>
        <p className="line-clamp-1">
          {matchInfo.ground_id && <span>{matchInfo.ground_name}</span>}

          {matchInfo.start_date && <span>, {getDisplayDate(matchInfo.start_date)}</span>}
          {matchInfo.main_series_id && <span>, {matchInfo.main_series_name}</span>}
        </p>
      </CardHeader>

      <hr />

      <CardContent className="flex flex-row justify-between gap-4 h-[100%]">
        <div className="w-full flex flex-col gap-2">
          <div
            className={`flex justify-between text-lg ${
              matchInfo.match_winner_team_id == team1_id ? "font-bold" : ""
            }`}
          >
            <div className="flex gap-2">
              <Image src={team1_image_url || "/file.svg"} alt="team1_image" width={20} height={20} />
              {team1_name}
            </div>
            <div>{getTeamScores(matchInfo.innings_scores || [], team1_id)}</div>
          </div>

          <div
            className={`flex justify-between text-lg ${
              matchInfo.match_winner_team_id == team2_id ? "font-bold" : ""
            }`}
          >
            <div className="flex gap-2">
              <Image src={team2_image_url || "/file.svg"} alt="team2_image" width={20} height={20} />
              {team2_name}
            </div>
            <div>{getTeamScores(matchInfo.innings_scores || [], team2_id)}</div>
          </div>
        </div>
      </CardContent>

      <CardFooter className="py-0 mt-[-2]">{getMatchResult(matchInfo)}</CardFooter>
    </Card>
  );
}

function getMatchResult(matchInfo: IMatchInfo): string {
  switch (matchInfo.match_state) {
    case "upcoming": {
      return "Match will start...";
    }
    case "live": {
      return "";
    }
    case "break": {
      return "";
    }
    case "completed": {
      const winningTeamId =
        matchInfo.match_winner_team_id || matchInfo.super_over_winner_id || matchInfo.bowl_out_winner_id;
      const winningTeamName =
        winningTeamId === matchInfo.team1_id
          ? matchInfo.team1_name
          : winningTeamId === matchInfo.team2_id
          ? matchInfo.team2_name
          : "";

      switch (matchInfo.final_result) {
        case "abandoned":
          return "Match Abandoned";
        case "draw":
          return "Match Drawn";
        case "no result":
          return "No Result";
        case "tie":
          return `Match Tied ${
            winningTeamName
              ? `(${winningTeamName} won the ${
                  matchInfo.super_over_winner_id
                    ? "Super Over"
                    : matchInfo.bowl_out_winner_id
                    ? "Bowl Out"
                    : "Tie Breaker"
                })`
              : `(${matchInfo.outcome_special_method ? `${matchInfo.outcome_special_method}` : ""})`
          }`;
        case "winner decided":
          return `${winningTeamName} won by ${matchInfo.win_margin} ${
            matchInfo.is_won_by_runs
              ? "runs"
              : `wickets ${matchInfo.balls_margin ? `(with ${matchInfo.balls_margin} balls remaining)` : ""}`
          } ${matchInfo.is_won_by_innings ? "and an innings" : ""}`;
      }
    }
  }

  return "";
}

function getTeamScores(inningsScores: ITeamInningsShortInfo[], teamId: number) {
  const teamInnings = inningsScores.filter((entry) => entry.batting_team_id === teamId);
  const entries = teamInnings.map(
    (entry) => `${entry.total_runs}-${entry.total_wickets} (${entry.total_overs})`
  );
  return entries.join("&");
}
