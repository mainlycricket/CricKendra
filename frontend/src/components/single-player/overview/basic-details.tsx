import { ISinglePlayer } from "@/lib/types/single-player.types";
import { getDisplayDate } from "@/lib/utils";
import Image from "next/image";
import Link from "next/link";

export function PlayerBasicDetails({ player }: { player: ISinglePlayer }) {
  return (
    <div className="flex flex-col justify-between gap-2">
      <hr className="hidden md:block" />
      <DetailRow label="full name" value={player.full_name} />
      <hr />
      <DetailRow label="date of birth" value={getDisplayDate(player.date_of_birth || new Date())} />
      <hr />
      <DetailRow label="playing role" value={player?.playing_role || "Top order Batter"} />
      <hr />
      <DetailRow label="batting style" value={player.is_rhb ? "Right hand bat" : "Left hand bat"} />
      <hr />
      <DetailRow label="bowling style" value={player?.bowling_styles?.join(", ") || "Right-arm fast"} />
      <hr />
      <div className="flex flex-col gap-2">
        <p className="font-thin uppercase">Teams</p>
        <div className="mt-1 flex flex-wrap gap-4 justify-between">
          {player.teams_represented?.map((team) => (
            <div key={team.id} className="flex flex-row gap-2 hover:text-sky-500">
              <Image src="/file.svg" alt={team.name} width={20} height={20} />
              <Link href={`/teams/${team.id}`}>{team.name}</Link>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

function DetailRow({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex justify-between">
      <p className="font-thin uppercase ">{label}</p>
      <p className="uppercase tracking-wide text-right font-medium">{value}</p>
    </div>
  );
}
