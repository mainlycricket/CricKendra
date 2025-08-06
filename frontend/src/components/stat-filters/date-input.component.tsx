export function DateInput({
  label,
  name,
  minDate,
  maxDate,
  defaultDate,
}: {
  label: string;
  name: string;
  minDate?: string;
  maxDate?: string;
  defaultDate?: string;
}) {
  return (
    <div className="flex gap-4">
      <p style={{ minWidth: "175px" }}>{label}</p>
      <input
        type="date"
        defaultValue={defaultDate}
        min={minDate}
        max={maxDate}
        name={name}
        className="dark:bg-input/30 dark:hover:bg-input/50  border-1 rounded-lg w-[180px] px-2 h-9"
      />
    </div>
  );
}
