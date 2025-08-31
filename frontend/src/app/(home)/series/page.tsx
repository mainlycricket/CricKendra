"use client";

import { doBackendRequest } from "@/lib/axiosFetch";
import { IAllSeriesResponse } from "@/lib/types/series.types";
import { MainSeriesLayout } from "@/components/series/main.component";
import { useSearchParams } from "next/navigation";
import { Suspense, useEffect, useState } from "react";

export default function Series() {
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
  const [data, setData] = useState<IAllSeriesResponse>();

  useEffect(() => {
    setLoading(true);

    doBackendRequest<null, IAllSeriesResponse>({
      url: `/series?${searchParams.toString()}`,
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

  return <MainSeriesLayout data={data!} searchParams={searchParams} />;
}
