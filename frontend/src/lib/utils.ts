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
