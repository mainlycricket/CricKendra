"use client";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { ITeamInningsShortInfo } from "@/lib/types/single-match.types";
import { useRouter } from "next/navigation";

export function InningsDropdown({
  inningsList,
  matchId,
  inningsId,
}: {
  inningsList: ITeamInningsShortInfo[];
  matchId: string;
  inningsId: string;
}) {
  const router = useRouter();

  return (
    <Select
      defaultValue={inningsId}
      onValueChange={(value) => {
        router.push(`/matches/${matchId}/commentary/${value}`);
      }}
    >
      <SelectTrigger className="w-[180px]">
        <SelectValue placeholder="Innings" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Innings</SelectLabel>
          {inningsList?.length &&
            inningsList.map((innings) => (
              <SelectItem value={innings.innings_id.toString()} key={innings.innings_id}>
                {innings.batting_team_name} - {innings.innings_number}
              </SelectItem>
            ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
