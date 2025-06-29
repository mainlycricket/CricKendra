"use client";

import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Dispatch, SetStateAction } from "react";

export function StatsDropdown({
  formatOptions,
  defaultFormat,
  setFormat,
  battingBowlingOptions,
  isDefaultBatting,
  setIsBattingStats,
}: {
  formatOptions: { value: string; label: string }[];
  defaultFormat: string;
  setFormat: Dispatch<SetStateAction<string>>;
  battingBowlingOptions: { value: string; label: string }[];
  isDefaultBatting: boolean;
  setIsBattingStats: Dispatch<SetStateAction<boolean>>;
}) {
  return (
    <div className="flex justify-between items-center">
      <p className="font-bold">Career Statistics</p>
      <div className="flex gap-4">
        <Select value={defaultFormat} onValueChange={(value) => setFormat(value)}>
          <SelectTrigger>
            <SelectValue placeholder="Format" />
          </SelectTrigger>
          <SelectContent>
            {formatOptions.map((item) => (
              <SelectItem key={item.value} value={item.value}>
                {item.label}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>

        <Select
          value={isDefaultBatting ? "batting" : "bowling"}
          onValueChange={(value) => {
            if (value === "batting") {
              setIsBattingStats(true);
            } else if (value === "bowling") {
              setIsBattingStats(false);
            }
          }}
        >
          <SelectTrigger>
            <SelectValue placeholder="Format" />
          </SelectTrigger>
          <SelectContent>
            {battingBowlingOptions.map((item) => (
              <SelectItem key={item.value} value={item.value}>
                {item.label}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>
    </div>
  );
}
