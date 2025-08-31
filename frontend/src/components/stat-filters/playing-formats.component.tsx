import { EnumPlayingFormat } from "@/lib/types/enums.types";
import Link from "next/link";

export function PlayingFormatOptions({
  selectedPlayingFormat,
  isMale,
}: {
  selectedPlayingFormat: EnumPlayingFormat;
  isMale: "true" | "false";
}) {
  const formats: { label: string; value: EnumPlayingFormat }[] = [
    { label: "Tests", value: "Test" },
    { label: "ODIs", value: "ODI" },
    { label: "T20Is", value: "T20I" },
    { label: "FC", value: "first_class" },
    { label: "List A", value: "list_a" },
    { label: "T20s", value: "T20" },
  ];

  return (
    <div className="w-full flex flex-row bg-secondary justify-center gap-2">
      {formats.map((option) => (
        <Link
          key={option.value}
          href={`/stats/filters?playing_format=${option.value}&is_male=${isMale}`}
          className="bg-secondary px-2 py-1 rounded"
          style={selectedPlayingFormat === option.value ? { color: "var(--color-sky-500)" } : {}}
        >
          {option.label}
        </Link>
      ))}
    </div>
  );
}
