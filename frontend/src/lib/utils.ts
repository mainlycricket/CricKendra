import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
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
