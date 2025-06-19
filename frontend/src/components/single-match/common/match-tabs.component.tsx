import Link from "next/link";

export function MatchTabs({
  matchId,
  inningsId,
  active,
}: {
  matchId: number;
  inningsId: number;
  active: string;
}) {
  const normalClass = "hover:text-blue-500";
  const activeClass = "border-b-2 border-blue-400";

  return (
    <div className="px-4 md:px-0 py-3 flex justify-evenly gap-4 rounded-xl overflow-auto bg-secondary">
      <Link href={`/matches/${matchId}`} className={active === "summary" ? activeClass : normalClass}>
        Summary
      </Link>
      <Link
        href={`/matches/${matchId}/scorecard`}
        className={active === "scorecard" ? activeClass : normalClass}
      >
        Scorecard
      </Link>
      {inningsId !== -1 && (
        <Link
          href={`/matches/${matchId}/commentary/${inningsId}`}
          className={active === "commentary" ? activeClass : normalClass}
        >
          Commentary
        </Link>
      )}
      <Link href={`/matches/${matchId}/stats`} className={active === "stats" ? activeClass : normalClass}>
        Stats
      </Link>
      <Link href={`/matches/${matchId}/squads`} className={active === "squads" ? activeClass : normalClass}>
        Squads
      </Link>
    </div>
  );
}
