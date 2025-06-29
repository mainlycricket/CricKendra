import Image from "next/image";
import Link from "next/link";
import { ChevronsUpDownIcon } from "lucide-react";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

import { ISingleSeriesSquadEntry, ISingleSeriesSquadPlayer } from "@/lib/types/single-series.types";

export function SingleSquadComponent({
  seriesId,
  activeSquadId,
  squad_list,
  players,
}: {
  seriesId: number;
  activeSquadId: number;
  squad_list: ISingleSeriesSquadEntry[];
  players: ISingleSeriesSquadPlayer[];
}) {
  return (
    <div className="w-full md:w-3/4 h-fit flex gap-4">
      <DesktopSquadList squad_list={squad_list} seriesId={seriesId} activeSquadId={activeSquadId} />
      <Card className="md:w-3/4 gap-2">
        <CardHeader>
          <CardTitle>
            {/* Current Squad Label is a part of Mobile Squad List */}
            <MobileSquadList squad_list={squad_list} seriesId={seriesId} activeSquadId={activeSquadId} />
          </CardTitle>
        </CardHeader>
        <CardContent className="p-0 flex flex-wrap">
          <SquadPlayersList players={players} />
        </CardContent>
      </Card>
    </div>
  );
}

function DesktopSquadList({
  squad_list,
  seriesId,
  activeSquadId,
}: {
  squad_list: ISingleSeriesSquadEntry[];
  seriesId: number;
  activeSquadId: number;
}) {
  return (
    <Card className="hidden md:flex gap-2 h-fit">
      <CardHeader>
        <CardTitle>Squads</CardTitle>
      </CardHeader>
      <CardContent className="p-0">
        {squad_list.map((item) => {
          const bgClass = item.squad_id == activeSquadId ? "bg-sky-500 text-white" : `hover:text-sky-500`;
          return (
            <Link
              key={item.squad_id}
              className={`block py-2 px-4 ${bgClass}`}
              href={`/series/${seriesId}/squads/${item.squad_id}`}
            >
              {item.squad_label}
            </Link>
          );
        })}
      </CardContent>
    </Card>
  );
}

function MobileSquadList({
  squad_list,
  seriesId,
  activeSquadId,
}: {
  squad_list: ISingleSeriesSquadEntry[];
  seriesId: number;
  activeSquadId: number;
}) {
  const activeSquadLabel = squad_list?.filter((item) => item.squad_id == activeSquadId)?.[0]?.squad_label;

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="block w-full text-xl flex items-center justify-between">
        <p>{activeSquadLabel || ""}</p>
        <ChevronsUpDownIcon className="md:hidden" />
      </DropdownMenuTrigger>

      <DropdownMenuContent align="end" className="p-0 md-hidden">
        <DropdownMenuLabel className="p-0 px-4 pt-2">Squads</DropdownMenuLabel>
        <DropdownMenuSeparator />
        {squad_list.map((item) => {
          const bgClass = item.squad_id == activeSquadId ? "bg-sky-500 text-white" : `hover:text-sky-500`;
          return (
            <DropdownMenuItem key={item.squad_id} className={`rounded-none py-2 px-4 ${bgClass}`}>
              <Link href={`/series/${seriesId}/squads/${item.squad_id}`}>{item.squad_label}</Link>
            </DropdownMenuItem>
          );
        })}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

function SquadPlayersList({ players }: { players: ISingleSeriesSquadPlayer[] }) {
  return players.map((player) => {
    return (
      <div key={player.player_id} className="flex gap-4 w-full md:w-[50%] px-8 py-4 border-1">
        <div>
          <Link href={`/players/${player.player_id}`}>
            <Image
              src={"/file.svg"}
              width={40}
              height={40}
              alt={player.player_name}
              className="rounded-full"
            />
          </Link>
        </div>
        <div>
          <p>
            <Link href={`/players/${player.player_id}`} className="hover:text-sky-500">
              {player.player_name}
            </Link>
          </p>
        </div>
      </div>
    );
  });
}
