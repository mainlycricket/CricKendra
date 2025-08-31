"use client";

import { doBackendRequest } from "@/lib/axiosFetch";
import { useEffect, useState } from "react";
import { SearchSelect } from "./search-select.component";
import { IAllTournaments, IAllTournamentsResponse } from "@/lib/types/tournaments.types";

export function SelectTournaments({ defaultSelected }: { defaultSelected: string[] }) {
  const [isLoading, setIsLoading] = useState(false);
  const [tournaments, setTournaments] = useState<IAllTournaments[]>([]);
  try {
    useEffect(() => {
      doBackendRequest<null, IAllTournamentsResponse>({
        url: "/tournaments?__limit=1000",
        method: "GET",
      })
        .then((data) => {
          setTournaments(data?.data?.tournaments || []);
        })
        .finally(() => {
          setIsLoading(false);
        });
    }, []);

    if (isLoading) return <div>Loading...</div>;

    return (
      <SearchSelect
        label="Tournaments"
        name="tournament_id"
        options={tournaments.map((tournament) => {
          return { value: tournament.id.toString(), label: tournament.name };
        })}
        defaultChecked={defaultSelected}
      />
    );
  } catch (error) {
    console.error(error);
    return <div>failed to fetch tournaments...</div>;
  }
}
