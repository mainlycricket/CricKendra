import Image from "next/image";
import Link from "next/link";

import { ISeriesTeam } from "@/lib/types/single-series.types";

export function TeamsComponent({ teams }: { teams: ISeriesTeam[] }) {
  return (
    <div className="w-full md:w-3/4 h-fit flex flex-col md:flex-row md:flex-wrap">
      {teams.map((team) => {
        return (
          <div key={team.team_id} className="p-8 h-20 md:w-[50%] flex items-center gap-6 border-1">
            <Link href={`/teams/${team.team_id}`}>
              <Image
                src={`${team.team_image_url || "/file.svg"}`}
                width={40}
                height={40}
                alt={team.team_name}
                className="rounded-full"
              />
            </Link>

            <Link href={`/teams/${team.team_id}`} className="hover:text-sky-500">
              {team.team_name}
            </Link>
          </div>
        );
      })}
    </div>
  );
}
