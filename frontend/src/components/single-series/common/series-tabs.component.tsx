import Link from "next/link";

export function SeriesTabs({
  seriesId,
  seriesName,
  seriesSeason,
  active,
}: {
  seriesId: number;
  seriesName: string;
  seriesSeason: string;
  active: string;
}) {
  const normalClass = "hover:text-blue-500";
  const activeClass = "border-b-2 border-blue-400";

  return (
    <div className="mt-2 flex justify-between px-6 md:px-12 overflow-auto bg-secondary py-2 rounded-lg">
      <p className="hidden md:block font-bold tracking-wide">
        {seriesName} {seriesSeason}
      </p>
      <Link
        href={`/series/${seriesId}`}
        className={active === "overview" ? activeClass : normalClass}
      >
        Overview
      </Link>
      <Link href={`/series/${seriesId}/matches`} className={active === "matches" ? activeClass : normalClass}>
        Matches
      </Link>
      <Link href={`/series/${seriesId}/teams`} className={active === "teams" ? activeClass : normalClass}>
        Teams
      </Link>
      <Link href={`/series/${seriesId}/squads`} className={active === "squads" ? activeClass : normalClass}>
        Squads
      </Link>
    </div>
  );
}
