import { Checkbox } from "../ui/checkbox";
import { Label } from "../ui/label";

export function CheckboxOption({
  label,
  name,
  options,
}: {
  name: string;
  label: string;
  options: { value: string; label: string }[];
}) {
  return (
    <div className="flex gap-4">
      <p style={{ minWidth: "175px" }}>{label}</p>
      <div className="flex flex-wrap gap-4">
        {options.map((option) => (
          <div className="flex items-center gap-2" key={`${name}_${option.value}`}>
            <Checkbox id={`${name}_${option.value}`} name={name} value={option.value} />
            <Label htmlFor={`${name}_${option.value}`} className="font-normal capitalize">
              {option.label}
            </Label>
          </div>
        ))}
      </div>
    </div>
  );
}
