import { PlayerFilters } from "./filters.component";
import { ReadonlyURLSearchParams } from "next/navigation";
import { Pagination } from "../common/pagination.component";
import { IAllPlayersResponse } from "@/lib/types/player.types";
import { SinglePlayer } from "./single-player.component";

export function MainPlayersLayout({
  data,
  searchParams,
}: {
  data: IAllPlayersResponse;
  searchParams: ReadonlyURLSearchParams;
}) {
  return (
    <div>
      <div className="flex flex-col gap-6">
        <PlayerFilters searchParams={searchParams} />
        {data?.players?.length ? (
          <div className="flex flex-col gap-2">
            {data?.players?.map((entry) => (
              <SinglePlayer key={entry.id} player={entry} />
            ))}
            <Pagination
              mainPageLink="players"
              recordsCount={data?.players?.length || 0}
              currentPage={Number(searchParams.get("__page")) || 1}
              disableNext={!data?.next}
            />
          </div>
        ) : (
          <div>No players found...</div>
        )}
      </div>
    </div>
  );
}
