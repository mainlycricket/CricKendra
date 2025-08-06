import { Input } from "../ui/input";

export function RangeInput({ label, minName, maxName }: { label: string; minName: string; maxName: string }) {
  return (
    <div className="flex gap-4">
      <p style={{ minWidth: "175px" }}>{label}</p>
      <div className="flex gap-2">
        <Input type="number" name={minName} min={0} step={1} className="w-24" placeholder="Min" />
        <Input type="number" name={maxName} min={0} step={1} className="w-24" placeholder="Max" />
      </div>
    </div>
  );
}
