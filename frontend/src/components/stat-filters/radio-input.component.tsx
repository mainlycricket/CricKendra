import { useState } from "react";
import { Label } from "../ui/label";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";

export function RadioInput({
  name,
  label,
  options,
  defaultValue,
  onValueChange,
}: {
  name: string;
  label: string;
  options: { value: string; label: string }[];
  defaultValue: string;
  onValueChange?: (value: string) => void;
}) {
  const [value, setValue] = useState(defaultValue);

  return (
    <div className="flex gap-4">
      <p style={{ minWidth: "175px" }} className="capitalize">{label}</p>
      <RadioGroup
        value={value}
        onValueChange={(value) => {
          setValue(value);
          if (onValueChange) {
            onValueChange(value);
          }
        }}
        name={name}
        className="flex flex-wrap gap-4"
      >
        {options.map((option) => (
          <div className="flex items-center gap-1" key={`${name}_${option.value}`}>
            <RadioGroupItem value={option.value} id={`${name}_${option.value}`} />
            <Label htmlFor={`${name}_${option.value}`} className="capitalize font-normal">
              {option.label}
            </Label>
          </div>
        ))}
      </RadioGroup>
    </div>
  );
}
