"use client";

import { doBackendRequest } from "@/lib/axiosFetch";
import { useSearchParams } from "next/navigation";
import { Suspense, useEffect, useState } from "react";
import { IAllPlayersResponse } from "@/lib/types/player.types";
import { MainPlayersLayout } from "@/components/players/main.component";

export default function Players() {
  return (
    <Suspense>
      <Component />
    </Suspense>
  );
}

function Component() {
  const searchParams = useSearchParams();

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [data, setData] = useState<IAllPlayersResponse>();

  useEffect(() => {
    setLoading(true);
    doBackendRequest<null, IAllPlayersResponse>({
      url: `/players?${searchParams.toString()}`,
      method: "GET",
    })
      .then((data) => {
        if (data.success) setData(data.data!);
        else setError(data.message);
      })
      .catch((error) => {
        setError(error);
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  if (loading) return <div>Loading...</div>;

  if (error) return <div>{error}</div>;

  return <MainPlayersLayout data={data!} searchParams={searchParams} />;
}
