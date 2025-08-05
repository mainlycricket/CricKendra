import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { EnumDismissalType } from "./types/enums.types";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function isValidIsoDate(date: string) {
  const regex = /^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/;
  if (!regex.test(date)) return false;
  const parsedDate = new Date(date);
  return !isNaN(parsedDate?.getDate());
}

export function getDisplayDate(date: Date | string) {
  if (typeof date === "string") {
    date = new Date(date);
  }

  const months = ["Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sep", "Oct", "Nov", "Dec"];

  return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
}

export function isBowlerDismissal(dismissalType: string): boolean {
  const bowlerWickets = ["caught", "bowled", "lbw", "stumped", "hit wicket", "caught and bowled"];
  return bowlerWickets.includes(dismissalType);
}

export function capitalizeFirstLetter(word: string): string {
  if (!word?.length) return word;

  return word?.[0]?.toUpperCase() + word.slice(1);
}

export function rotate2DArray<T>(data: T[][]): T[][] {
  const res: T[][] = [];

  const rows = data?.length || 0;
  const cols = data?.[0]?.length || 0;

  // ensure rows
  for (let i = 0; i < cols; i++) {
    res.push([]);
  }

  // populate data
  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      res[j].push(data[i][j]);
    }
  }

  return res;
}

export function getDismissalTypeDisplayText(value: EnumDismissalType): string {
  switch (value) {
    case "caught":
      return "Caught";
    case "bowled":
      return "Bowled";
    case "lbw":
      return "LBW";
    case "run out":
      return "Run Out";
    case "stumped":
      return "Stumped";
    case "hit wicket":
      return "Hit Wicket";
    case "handled the ball":
      return "Handled the Ball";
    case "obstructing the field":
      return "Obstructing the Field";
    case "timed out":
      return "Timed Out";
    case "retired hurt":
      return "Retired Hurt";
    case "hit the ball twice":
      return "Hit the ball Twice";
    case "caught and bowled":
      return "Caught & Bowled";
    case "retired out":
      return "Retired Out";
    case "retired not out":
      return "Retired Not Out";
  }
}
