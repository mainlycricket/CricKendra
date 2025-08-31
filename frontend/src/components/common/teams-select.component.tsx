"use client";

import { doBackendRequest } from "@/lib/axiosFetch";
import { IAllTeams, IAllTeamsResponse } from "@/lib/types/teams.types";
import { useEffect, useState } from "react";
import { SearchSelect } from "./search-select.component";

export function SelectTeams({ name, defaultSelected }: { name: string; defaultSelected: string[] }) {
  const [isLoading, setIsLoading] = useState(false);
  const [teams, setTeams] = useState<IAllTeams[]>([]);
  try {
    useEffect(() => {
      doBackendRequest<null, IAllTeamsResponse>({
        url: "/teams?__limit=1000",
        method: "GET",
      })
        .then((data) => {
          setTeams(data?.data?.teams || []);
        })
        .finally(() => {
          setIsLoading(false);
        });
    }, []);

    if (isLoading) return <div>Loading...</div>;

    return (
      <SearchSelect
        label="Teams"
        name={name}
        options={teams.map((team) => {
          return { value: team.id.toString(), label: `${team.name} ${team.is_male ? "" : "- W"}` };
        })}
        defaultChecked={defaultSelected}
      />
    );
  } catch (error) {
    console.error(error);
    return <div>failed to fetch teams...</div>;
  }
}
