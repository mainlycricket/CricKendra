import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useState } from "react";

export function SelectInput({
  label,
  name,
  options,
  defaultValue,
}: {
  label: string;
  name: string;
  options: { label: string; value: string }[];
  defaultValue?: string;
}) {
  const [value, setValue] = useState(defaultValue);

  return (
    <Select
      name={name}
      value={value}
      onValueChange={(value: string) => {
        setValue(value);
      }}
    >
      <SelectTrigger className="w-full">
        <SelectValue placeholder={label} />
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
              }}
            >
              {item.label}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
