import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useEffect, useState } from "react";

export function SelectInput({
  label,
  name,
  options,
  onValueChange,
}: {
  label: string;
  name: string;
  options: { label: string; value: string }[];
  onValueChange?: (value: string) => void;
}) {
  const [value, setValue] = useState(options?.[0]?.value);

  useEffect(() => {
    setValue(options?.[0]?.value);
  }, [options]);

  return (
    <div className="flex gap-4">
      <p style={{ minWidth: "175px" }} className="capitalize">
        {label}
      </p>
      <Select
        name={name}
        value={value}
        onValueChange={(value: string) => {
          setValue(value);
          if (onValueChange) onValueChange(value);
        }}
      >
        <SelectTrigger className="w-[180px]">
          <SelectValue />
        </SelectTrigger>
        <SelectContent>
          <SelectGroup>
            <SelectLabel className="text-sm">{label}</SelectLabel>
            {options.map((item) => (
              <SelectItem
                key={item.value}
                value={item.value}
                onSelect={() => {
                  setValue(value);
                  if (onValueChange) onValueChange(value);
                }}
              >
                {item.label}
              </SelectItem>
            ))}
          </SelectGroup>
        </SelectContent>
      </Select>
    </div>
  );
}
