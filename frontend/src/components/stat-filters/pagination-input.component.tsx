import { Input } from "../ui/input";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export function PaginationInputs({
  defaultPage,
  defaultLimit,
}: {
  defaultPage: number;
  defaultLimit: number;
}) {
  return (
    <div className="flex flex-col gap-4">
      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Page No.</p>
        <Input
          type="number"
          name="__page"
          min={1}
          step={1}
          className="w-24"
          defaultValue={defaultPage || 1}
        />
      </div>

      <div className="flex gap-4">
        <p style={{ minWidth: "175px" }}>Records per Page</p>
        <div>
          <Select name="__limit" defaultValue={defaultLimit?.toString() || "50"}>
            <SelectTrigger className="w-24">
              <SelectValue placeholder="50" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel className="text-sm">Records per Page</SelectLabel>
                <SelectItem value="50">50</SelectItem>
                <SelectItem value="100">100</SelectItem>
                <SelectItem value="150">150</SelectItem>
                <SelectItem value="200">200</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>
      </div>
    </div>
  );
}
