import Image from "next/image";
import Link from "next/link";

import { Card, CardContent, CardFooter, CardHeader } from "../../ui/card";

import { IMatchHeader, ITeamInningsShortInfo } from "@/lib/types/single-match.types";
import { getDisplayDate } from "@/lib/utils";

export function MatchHeader({ matchHeader }: { matchHeader: IMatchHeader }) {
  let { team1_id, team1_name, team1_image_url, team2_id, team2_name, team2_image_url } = matchHeader;

  if (matchHeader?.innings_scores?.[0]?.batting_team_id === team2_id) {
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
          {matchHeader.match_state === "break"
            ? matchHeader.match_state_description
            : matchHeader.match_state}
        </p>
        <p className="line-clamp-1">
          {matchHeader.ground_id && (
            <Link href={`/grounds/${matchHeader.ground_id}`}>{matchHeader.ground_name}</Link>
          )}

          {matchHeader.start_date && <span>, {getDisplayDate(matchHeader.start_date)}</span>}
          {matchHeader.main_series_id && (
            <span>
              ,{" "}
              <Link href={`/series/${matchHeader.main_series_id}`} className="underline">
                {matchHeader.main_series_name}
              </Link>
            </span>
          )}
        </p>
      </CardHeader>

      <hr />

      <CardContent className="flex flex-row justify-between gap-4 h-[100%]">
        <div className="w-full flex flex-col gap-2">
          <div
            className={`flex justify-between text-lg ${
              matchHeader.match_winner_team_id == team1_id ? "font-bold" : ""
            }`}
          >
            <div>
              <Link href={team1_id ? `/teams/${team1_id}` : ""} className="flex gap-2">
                <Image src={team1_image_url || "/file.svg"} alt="team1_image" width={20} height={20} />
                {team1_name}
              </Link>
            </div>
            <div>{getTeamScores(matchHeader?.innings_scores || [], team1_id)}</div>
          </div>

          <div
            className={`flex justify-between text-lg ${
              matchHeader.match_winner_team_id == team2_id ? "font-bold" : ""
            }`}
          >
            <div>
              <Link href={team2_id ? `/teams/${matchHeader.team2_id}` : ""} className="flex gap-2">
                <Image src={team2_image_url || "/file.svg"} alt="team2_image" width={20} height={20} />
                {team2_name}
              </Link>
            </div>
            <div>{getTeamScores(matchHeader?.innings_scores || [], team2_id)}</div>
          </div>
        </div>

        {matchHeader?.player_awards?.length && (
          <div className="hidden md:block w-1/5 px-2 border-l-1">
            {matchHeader?.player_awards?.map((entry) => {
              return (
                <div key={`${entry.player_id}_${entry.award_type}`} className="flex flex-col gap-2">
                  <p className="capitalize font-thin px-4">{entry.award_type.split("_").join(" ")}</p>
                  <div>
                    <div className="flex gap-4 px-4 py-2">
                      <Image
                        src={`/file.svg`}
                        width={40}
                        height={40}
                        alt={entry.player_name}
                        className="rounded-full"
                      />
                      <p className="text-lg py-2 w-full">{entry.player_name}</p>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </CardContent>

      <CardFooter className="py-0 mt-[-2]">{getMatchResult(matchHeader)}</CardFooter>
    </Card>
  );
}

function getMatchResult(matchHeader: IMatchHeader): string {
  switch (matchHeader.match_state) {
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
        matchHeader.match_winner_team_id ||
        matchHeader.super_over_winner_id ||
        matchHeader.bowl_out_winner_id;
      const winningTeamName =
        winningTeamId === matchHeader.team1_id
          ? matchHeader.team1_name
          : winningTeamId === matchHeader.team2_id
          ? matchHeader.team2_name
          : "";

      switch (matchHeader.final_result) {
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
                  matchHeader.super_over_winner_id
                    ? "Super Over"
                    : matchHeader.bowl_out_winner_id
                    ? "Bowl Out"
                    : "Tie Breaker"
                })`
              : `(${matchHeader.outcome_special_method ? `${matchHeader.outcome_special_method}` : ""})`
          }`;
        case "winner decided":
          return `${winningTeamName} won by ${matchHeader.win_margin} ${
            matchHeader.is_won_by_runs
              ? "runs"
              : `wickets ${
                  matchHeader.balls_margin ? `(with ${matchHeader.balls_margin} balls remaining)` : ""
                }`
          } ${matchHeader.is_won_by_innings ? "and an innings" : ""}`;
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
