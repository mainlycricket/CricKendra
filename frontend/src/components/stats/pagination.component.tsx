"use client";

import Link from "next/link";
import { Button } from "../ui/button";
import { useSearchParams } from "next/navigation";

export function Pagination({
  currentPage,
  recordsCount,
  disableNext,
}: {
  currentPage: number;
  recordsCount: number;
  disableNext: boolean;
}) {
  const searchParams = useSearchParams();

  return (
    <div className="mt-2 flex justify-between">
      <Link href={`/stats?${getUpdatePageQuery(searchParams.toString(), currentPage - 1)}`}>
        <Button type="button" disabled={currentPage === 1}>
          Prev
        </Button>
      </Link>
      <p style={{ fontStyle: "italic" }}>
        Current Page: {currentPage}, Records: {recordsCount}
      </p>
      <Link href={`/stats?${getUpdatePageQuery(searchParams.toString(), currentPage + 1)}`}>
        <Button type="button" disabled={disableNext}>
          Next
        </Button>
      </Link>
    </div>
  );
}

function getUpdatePageQuery(searchParams: string, page: number): string {
  const searchParamsObj = new URLSearchParams(searchParams);
  searchParamsObj.set("__page", page.toString());
  return searchParamsObj.toString();
}
