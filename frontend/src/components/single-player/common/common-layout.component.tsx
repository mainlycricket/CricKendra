import { ISinglePlayer } from "@/lib/types/single-player.types";
import { PlayerTabs } from "./player-tabs.component";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import Image from "next/image";
import { PlayerBasicDetails } from "../overview/basic-details";

export function SinglePlayerLayout({
  player,
  active,
  children,
}: {
  player: ISinglePlayer;
  active: string;
  children: React.ReactNode;
}) {
  return (
    <div className="flex flex-col md:flex-row gap-4">
      <Card className="md:w-1/3 h-fit">
        <CardContent className="flex md:flex-col-reverse md:gap-8 items-center justify-between">
          <div className="font-semibold flex flex-col gap-1 md:text-center">
            <h3 className="text-2xl">{player.name}</h3>
            <p>
              {player.nationality || "India"} | {player.playing_role || "Top order Batter"}
            </p>
          </div>
          <div className="">
            <Image
              src={player.image_url || "/file.svg"}
              alt={player.name}
              width={100}
              height={40}
              className="rounded-full md:rounded-none"
            />
          </div>
        </CardContent>
        <CardFooter className="hidden md:block">
          <PlayerBasicDetails player={player} />
        </CardFooter>
      </Card>

      <div className="w-full flex flex-col gap-4">
        <PlayerTabs playerId={player.id} active={active} />
        <div>{children}</div>
      </div>
    </div>
  );
}
