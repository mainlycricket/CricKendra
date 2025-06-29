import Image from "next/image";
import Link from "next/link";

import { ISingleSeriesSquadEntry } from "@/lib/types/single-series.types";

export function SquadsComponent({
  squad_list,
  seriesId,
}: {
  squad_list: ISingleSeriesSquadEntry[];
  seriesId: number;
}) {
  return (
    <div className="w-full md:w-3/4 h-fit flex flex-col md:flex-row md:flex-wrap">
      {squad_list.map((item) => {
        return (
          <div key={item.squad_id} className="p-8 h-20 md:w-[50%] flex items-center gap-6 border-1">
            <Image src={"/file.svg"} width={40} height={40} alt={item.squad_label} className="rounded-full" />
            <Link href={`/series/${seriesId}/squads/${item.squad_id}`} className="hover:text-sky-500">
              {item.squad_label}
            </Link>
          </div>
        );
      })}
    </div>
  );
}
