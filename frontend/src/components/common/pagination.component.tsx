"use client";

import Link from "next/link";
import { Button } from "../ui/button";
import { useSearchParams } from "next/navigation";

export function Pagination({
  mainPageLink,
  currentPage,
  recordsCount,
  disableNext,
}: {
  mainPageLink: string;
  currentPage: number;
  recordsCount: number;
  disableNext: boolean;
}) {
  const searchParams = useSearchParams();

  return (
    <div className="mt-2 flex justify-between">
      <Button type="button" disabled={currentPage === 1}>
        <Link href={`/${mainPageLink}?${getUpdatePageQuery(searchParams.toString(), currentPage - 1)}`}>
          Prev
        </Link>
      </Button>
      <p style={{ fontStyle: "italic" }}>
        Current Page: {currentPage}, Records: {recordsCount}
      </p>
      <Button type="button" disabled={disableNext}>
        <Link href={`/${mainPageLink}/?${getUpdatePageQuery(searchParams.toString(), currentPage + 1)}`}>
          Next
        </Link>
      </Button>
    </div>
  );
}

function getUpdatePageQuery(searchParams: string, page: number): string {
  const searchParamsObj = new URLSearchParams(searchParams);
  searchParamsObj.set("__page", page.toString());
  return searchParamsObj.toString();
}
