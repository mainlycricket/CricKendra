import Link from "next/link";

export function PlayerTabs({ playerId, active }: { playerId: number; active: string }) {
  const normalClass = "hover:text-blue-500";
  const activeClass = "border-b-2 border-blue-400";

  return (
    <div className="px-4 md:px-0 py-3 flex justify-evenly gap-4 rounded-xl overflow-auto bg-secondary">
      <Link href={`/players/${playerId}`} className={active === "overview" ? activeClass : normalClass}>
        Overview
      </Link>
      <Link href={`/players/${playerId}/stats`} className={active === "stats" ? activeClass : normalClass}>
        Stats
      </Link>
    </div>
  );
}
