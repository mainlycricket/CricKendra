import { SingleMatch } from "./single-match-info.component";
import { MatchFilters } from "./filters.component";
import { ReadonlyURLSearchParams } from "next/navigation";
import { Pagination } from "../common/pagination.component";
import { AllMatchesResponse } from "@/lib/types/match.types";
import Link from "next/link";

export function MainMatchesLayout({
  data,
  searchParams,
}: {
  data: AllMatchesResponse;
  searchParams: ReadonlyURLSearchParams;
}) {
  return (
    <div>
      <div className="flex flex-col gap-6">
        <MatchFilters searchParams={searchParams} />
        {data?.matches?.length ? (
          <div className="flex flex-col gap-2">
            {data?.matches?.map((entry) => (
              <Link key={entry.match_id} href={`/matches/${entry.match_id}`}>
                <SingleMatch matchInfo={entry} />
              </Link>
            ))}
            <Pagination
              mainPageLink="matches"
              recordsCount={data?.matches?.length || 0}
              currentPage={Number(searchParams.get("__page")) || 1}
              disableNext={!data?.next}
            />
          </div>
        ) : (
          <div>No matches found...</div>
        )}
      </div>
    </div>
  );
}
