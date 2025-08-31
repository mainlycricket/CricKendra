import Image from "next/image";

import { Card, CardHeader, CardTitle } from "@/components/ui/card";
import { IAllPlayers } from "@/lib/types/player.types";
import Link from "next/link";
import { Badge } from "../ui/badge";

export function SinglePlayer({ player }: { player: IAllPlayers }) {
  return (
    <Link href={`/players/${player.id}`}>
      <Card className="gap-4 py-4">
        <CardHeader>
          <CardTitle>
            <div className="flex gap-2 items-center">
              <Image src={player.image_url || `/file.svg`} alt={player.name} width={30} height={30} />{" "}
              <span>{player.name}</span>
              {!player.is_male && <Badge variant="secondary"> Women</Badge>}
            </div>
          </CardTitle>
        </CardHeader>
      </Card>
    </Link>
  );
}
