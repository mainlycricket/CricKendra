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
import { Checkbox } from "../ui/checkbox";
import { Label } from "../ui/label";

export function SelectCheckboxInput({
  name,
  label,
  options,
  defaultValues,
}: {
  name: string;
  label: string;
  options: { value: string; label: string }[];
  defaultValues?: string[];
}) {
  const [isMultiple, setIsMultiple] = useState((defaultValues?.length || 0) > 1 ? true : false);
  return (
    <div className="flex justify-between">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>{label}</p>
        <div>
          {isMultiple ? (
            <div className="flex flex-wrap gap-4">
              {options?.map((option) => {
                return (
                  <div className="flex gap-2" key={`${name}_${option.value}`}>
                    <Checkbox
                      id={`${name}_${option.value}`}
                      name={name}
                      value={option.value}
                      defaultChecked={defaultValues?.includes(option.value)}
                    />
                    <Label htmlFor={`${name}_${option.value}`}>{option.label}</Label>
                  </div>
                );
              })}
            </div>
          ) : (
            <Select name={name} defaultValue={defaultValues?.[0]}>
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="All" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectLabel className="text-sm">{label}</SelectLabel>
                  {options?.map((option) => (
                    <SelectItem key={`${name}_${option.value}`} value={option.value}>
                      {option.label}
                    </SelectItem>
                  ))}
                </SelectGroup>
              </SelectContent>
            </Select>
          )}
        </div>
      </div>
      <div className="hidden md:block" style={{ minWidth: "60px" }}>
        <button type="button" className="text-sm btn" onClick={() => setIsMultiple(!isMultiple)}>
          {isMultiple ? "Single" : "Multiple"}
        </button>
      </div>
    </div>
  );
}
