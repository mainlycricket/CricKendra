export type EnumPlayingFormat = "Test" | "ODI" | "T20I" | "first_class" | "list_a" | "T20";
export type EnumStatsType = "batting" | "bowling" | "team";
export type EnumStatsView = "overall" | "individual";
export type EnumHomeAway = "home" | "away" | "neutral";
export type EnumMatchResult = "won" | "lost" | "drawn" | "tied" | "no result";
export type EnumTossResult = "won" | "lost";
export type EnumBatField = "bat" | "field";
export type EnumInningsNumber = "1" | "2" | "3" | "4";
export type EnumDismissalType =
  | "caught"
  | "bowled"
  | "lbw"
  | "run out"
  | "stumped"
  | "hit wicket"
  | "handled the ball"
  | "obstructing the field"
  | "timed out"
  | "retired hurt"
  | "hit the ball twice"
  | "caught and bowled"
  | "retired out"
  | "retired not out";

export function isEnumDismissalType(value: string): value is EnumDismissalType {
  const validValues = [
    "caught",
    "bowled",
    "lbw",
    "run out",
    "stumped",
    "hit wicket",
    "handled the ball",
    "obstructing the field",
    "timed out",
    "retired hurt",
    "hit the ball twice",
    "caught and bowled",
    "retired out",
    "retired not out",
  ];

  return validValues.includes(value);
}
