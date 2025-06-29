import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../ui/table";

export function PopulatedTable({ data, isFirstColHead }: { data: string[][]; isFirstColHead: boolean }) {
  if (!data?.length) return <></>;

  const headerRow = data?.[0];
  const valueRows = data.slice(1);

  return (
    <Table>
      <TableHeader>
        <TableRow>
          {headerRow.map((header) => (
            <TableHead key={header}>{header}</TableHead>
          ))}
        </TableRow>
      </TableHeader>
      <TableBody>
        {valueRows?.length === 0 ? (
          <TableRow>
            <TableCell colSpan={headerRow.length} className="text-center">No records</TableCell>
          </TableRow>
        ) : (
          valueRows.map((row, idx) => {
            return (
              <TableRow key={idx}>
                {row.map((value, idx) => (
                  <TableCell key={idx} className={idx == 0 && isFirstColHead ? "font-medium" : ""}>
                    {value}
                  </TableCell>
                ))}
              </TableRow>
            );
          })
        )}
      </TableBody>
    </Table>
  );
}
