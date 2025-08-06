import { EnumPlayingFormat } from "@/lib/types/enums.types";
import Link from "next/link";

export function PlayingFormatOptions({
  selectedPlayingFormat,
}: {
  selectedPlayingFormat: EnumPlayingFormat;
}) {
  return (
    <div className="w-full flex flex-row bg-secondary justify-center gap-2">
      <Link
        href={`/stats/filters?playing_format=Test`}
        className="bg-secondary px-2 py-1 rounded"
        style={selectedPlayingFormat === "Test" ? { color: "var(--color-sky-500)" } : {}}
      >
        Tests
      </Link>
      <Link
        href={`/stats/filters?playing_format=ODI`}
        className="bg-secondary px-2 py-1 rounded"
        style={selectedPlayingFormat === "ODI" ? { color: "var(--color-sky-500)" } : {}}
      >
        ODIs
      </Link>
      <Link
        href={`/stats/filters?playing_format=T20I`}
        className="bg-secondary px-2 py-1 rounded"
        style={selectedPlayingFormat === "T20I" ? { color: "var(--color-sky-500)" } : {}}
      >
        T20Is
      </Link>
      <Link
        href={`/stats/filters?playing_format=first_class`}
        className="bg-secondary px-2 py-1 rounded"
        style={selectedPlayingFormat === "first_class" ? { color: "var(--color-sky-500)" } : {}}
      >
        FC
      </Link>
      <Link
        href={`/stats/filters?playing_format=list_a`}
        className="bg-secondary px-2 py-1 rounded"
        style={selectedPlayingFormat === "list_a" ? { color: "var(--color-sky-500)" } : {}}
      >
        List A
      </Link>
      <Link
        href={`/stats/filters?playing_format=T20`}
        className="bg-secondary px-2 py-1 rounded"
        style={selectedPlayingFormat === "T20" ? { color: "var(--color-sky-500)" } : {}}
      >
        T20
      </Link>
    </div>
  );
}
