import { IPlayerSquadEntry } from "@/lib/types/single-match.types";
import Link from "next/link";

export function TeamSquad({ players }: { players: IPlayerSquadEntry[] }) {
  return (
    <div>
      {players.map((player) => {
        return (
          <div key={player.player_id}>
            <Link
              href={`/players/${player.player_id}`}
              className="block py-2 tracking-wide text-lg border-1 underline"
            >
              {player.player_name} {player.is_captain && <span>(c)</span>}{" "}
              {/* <span className="">{player.playing_status}</span> */}
              {player.is_captain && <span>(wk)</span>}
            </Link>
          </div>
        );
      })}
    </div>
  );
}
