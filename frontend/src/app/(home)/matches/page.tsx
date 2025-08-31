"use client";

import { doBackendRequest } from "@/lib/axiosFetch";
import { useSearchParams } from "next/navigation";
import { Suspense, useEffect, useState } from "react";
import { MainMatchesLayout } from "@/components/matches/main.component";
import { AllMatchesResponse } from "@/lib/types/match.types";

export default function Matches() {
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
  const [data, setData] = useState<AllMatchesResponse>();

  useEffect(() => {
    setLoading(true);
    doBackendRequest<null, AllMatchesResponse>({
      url: `/matches?${searchParams.toString()}`,
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

  return <MainMatchesLayout data={data!} searchParams={searchParams} />;
}
