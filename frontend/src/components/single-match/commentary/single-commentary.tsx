import { IBbbCommentary } from "@/lib/types/single-match.types";
import { isBowlerDismissal } from "@/lib/utils";

export function SingleBallCommentary({ data }: { data: IBbbCommentary }) {
  return (
    <div className="flex gap-2 border-1 p-4 items-center">
      <div className="font-thin p-2">{data.ball_number}</div>
      <ScoreBox data={data} />
      <div className="ml-2">
        <TextCommentary data={data} />
      </div>
    </div>
  );
}

function TextCommentary({ data }: { data: IBbbCommentary }) {
  if (!data.bowler_id) {
    return <div>{data.commentary}</div>;
  }

  let text = "";

  if (data.wides > 0) {
    text += "wide ball";
  }

  if (data.noballs > 0) {
    text += "no ball";
  }

  if (data.legbyes > 0) {
    text += `${data.legbyes} legbyes`;
  }

  if (data.byes > 0) {
    text += `${data.byes} byes`;
  }

  if (data.is_four) {
    text += `FOUR runs`;
  }

  if (data.is_six) {
    text += `SIX runs`;
  }

  if (data.player1_dismissed_id || data.player2_dismissed_id) {
    text += `OUT`;
  }

  if (text == "") {
    text = data.total_runs === 0 ? `no run` : data.total_runs === 1 ? `1 run` : `${data.total_runs} runs`;
  }

  return (
    <div>
      <p>
        {data.bowler_name} to {data.batter_name}, {text}
      </p>
      <p>{data.commentary}</p>
      <p className="font-bold">
        {data.player1_dismissed_id
          ? `${data.player1_dismissed_name} - ${data.player1_dismissal_type} ${
              isBowlerDismissal(data.player1_dismissal_type!) ? `- b ${data.bowler_name}` : ""
            } - ${data.player1_dismissed_runs} (${data.player1_dismissed_balls}b ${
              data.player1_dismissed_fours
            }✕4 ${data.player1_dismissed_sixes}✕6)`
          : ""}
      </p>
    </div>
  );
}

function ScoreBox({ data }: { data: IBbbCommentary }) {
  const texts: string[] = [];
  const classes = ["p-2", "text-center", "min-w-8", "h-10", "font-medium"];

  if (data.wides > 0) {
    classes.push("bg-secondary");
    classes.push("text-primary");
    texts.push(`${data.wides.toString()}w`);
  }

  if (data.noballs > 0) {
    classes.push("bg-secondary");
    classes.push("text-primary");
    texts.push(`${data.noballs.toString()}nb`);
  }

  if (data.legbyes > 0) {
    classes.push("bg-secondary");
    classes.push("text-primary");
    texts.push(`${data.legbyes.toString()}lb`);
  }

  if (data.byes > 0) {
    classes.push("bg-secondary");
    classes.push("text-primary");
    texts.push(`${data.byes.toString()}b`);
  }

  if (data.is_four) {
    classes.push("bg-green-700");
    classes.push("text-white");
    texts.push("4");
  }

  if (data.is_six) {
    classes.push("bg-purple-500");
    classes.push("text-white");
    texts.push("6");
  }

  if (data.player1_dismissed_id || data.player2_dismissed_id) {
    classes.push("bg-red-500");
    classes.push("text-white");
    texts.push("W");
  }

  if (texts.length == 0) {
    classes.push("bg-secondary");
    classes.push("text-primary");
    texts.push(data.total_runs ? data.total_runs.toString() : ".");
  }

  return <div className={classes.join(" ")}>{texts.join(" ")}</div>;
}
