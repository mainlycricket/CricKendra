import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Factory } from "lucide-react";
import { useState } from "react";
import { Checkbox } from "../ui/checkbox";
import { Label } from "../ui/label";

export function StatOption({
  name,
  label,
  optionLabels,
  optionValues,
}: {
  name: string;
  label: string;
  optionLabels: string[];
  optionValues: string[];
}) {
  const [isMultiple, setIsMultiple] = useState(false);
  return (
    <div className="flex justify-between">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>{label}</p>
        <div>
          {isMultiple ? (
            <div className="flex flex-wrap gap-4">
              {optionValues?.map((value, idx) => {
                return (
                  <div className="flex gap-2" key={value}>
                    <Checkbox id={name + value} name={name} value={value} />
                    <Label htmlFor={name + value}>{optionLabels[idx]}</Label>
                  </div>
                );
              })}
            </div>
          ) : (
            <Select name={name}>
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="All" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectLabel className="text-sm">{label}</SelectLabel>
                  {optionValues.map((value, idx) => (
                    <SelectItem key={value} value={value}>
                      {optionLabels[idx] || ""}
                    </SelectItem>
                  ))}
                </SelectGroup>
              </SelectContent>
            </Select>
          )}
        </div>
      </div>
      <div className="hidden md:block" style={{ minWidth: "60px" }}>
        <button type="button" className="text-sm btn" onClick={(e) => setIsMultiple(!isMultiple)}>
          {isMultiple ? "Single" : "Multiple"}
        </button>
      </div>
    </div>
  );
}
