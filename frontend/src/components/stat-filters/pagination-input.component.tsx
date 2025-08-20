import { Input } from "../ui/input";
import { SelectInput } from "./select-input.component";

export function PaginationInputs({
  defaultPage,
  defaultLimit,
}: {
  defaultPage: string;
  defaultLimit: string;
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

      <SelectInput
        name="__limit"
        label="Records per Page"
        options={[
          { label: "50", value: "50" },
          { label: "100", value: "100" },
          { label: "150", value: "150" },
          { label: "200", value: "200" },
        ]}
        defaultValue={defaultLimit || "50"}
      />
    </div>
  );
}
