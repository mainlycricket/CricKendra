"use client";

import { doBackendRequest } from "@/lib/axiosFetch";
import { useEffect, useState } from "react";
import { SearchSelect } from "./search-select.component";
import { IAllSeasonsResponse } from "@/lib/types/seasons.types";

export function SelectSeasons({ defaultSelected }: { defaultSelected: string[] }) {
  const [isLoading, setIsLoading] = useState(false);
  const [seasons, setSeasons] = useState<string[]>([]);
  try {
    useEffect(() => {
      doBackendRequest<null, IAllSeasonsResponse>({
        url: "/seasons?__limit=10000&__sort=-season",
        method: "GET",
      })
        .then((data) => {
          setSeasons(data?.data?.seasons || []);
        })
        .finally(() => {
          setIsLoading(false);
        });
    }, []);

    if (isLoading) return <div>Loading...</div>;

    return (
      <SearchSelect
        label="Seasons"
        name="season"
        options={seasons.map((season) => {
          return { value: season, label: season };
        })}
        defaultChecked={defaultSelected}
      />
    );
  } catch (error) {
    console.error(error);
    return <div>failed to fetch seasons...</div>;
  }
}
