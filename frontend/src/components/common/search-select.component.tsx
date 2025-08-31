"use client";

import * as React from "react";
import {
  ColumnDef,
  ColumnFiltersState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  useReactTable,
  VisibilityState,
} from "@tanstack/react-table";

import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Input } from "@/components/ui/input";
import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table";
import { Label } from "../ui/label";
import { Select, SelectContent, SelectTrigger, SelectValue } from "../ui/select";

export function SearchSelect({
  options,
  name,
  label,
  defaultChecked,
}: {
  options: { label: string; value: string }[];
  name: string;
  label: string;
  defaultChecked?: string[];
}) {
  options.sort((a, b) => {
    const isASelected = defaultChecked?.includes(a.value);
    const isBSelected = defaultChecked?.includes(b.value);
    if (isASelected && !isBSelected) return -1;
    if (isBSelected && !isASelected) return 1;
    if (a.label.toLowerCase() > b.label.toLowerCase()) return 1;
    if (a.label.toLowerCase() < b.label.toLowerCase()) return -1;
    return 0;
  });

  const [selectedOptions, setSelectedOptions] = React.useState(defaultChecked || []);
  const [sorting, setSorting] = React.useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = React.useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = React.useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = React.useState({});

  const columns = getColumns({ label, name, selectedOptions, setSelectedOptions });

  const table = useReactTable({
    data: options,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
    initialState: {
      pagination: {
        pageSize: 5,
      },
    },
  });

  return (
    <div>
      <div className="hidden">
        {options?.map((option) => (
          <input
            key={option.value}
            type="checkbox"
            name={name}
            defaultValue={option.value}
            checked={selectedOptions.includes(option.value)}
            onChange={() => {}}
          />
        ))}
      </div>
      <Select>
        <SelectTrigger className="w-full">
          <SelectValue placeholder={`${label} - ${selectedOptions.length} selected`} />
        </SelectTrigger>
        <SelectContent>
          <div className="p-2 flex flex-col gap-4">
            <div className="flex items-center">
              <Input
                placeholder={`Search...`}
                value={(table.getColumn("label")?.getFilterValue() as string) ?? ""}
                onChange={(event) => table.getColumn("label")?.setFilterValue(event.target.value)}
              />
            </div>
            <div className="overflow-hidden rounded-md border">
              <Table>
                <TableBody>
                  {table.getRowModel().rows?.length ? (
                    table.getRowModel().rows.map((row) => (
                      <TableRow key={row.id} data-state={row.getIsSelected() && "selected"}>
                        {row.getVisibleCells().map((cell) => (
                          <TableCell key={cell.id}>
                            {flexRender(cell.column.columnDef.cell, cell.getContext())}
                          </TableCell>
                        ))}
                      </TableRow>
                    ))
                  ) : (
                    <TableRow>
                      <TableCell colSpan={columns.length} className="h-24 text-center">
                        No results.
                      </TableCell>
                    </TableRow>
                  )}
                </TableBody>
              </Table>
            </div>
            <div className="flex items-center justify-between">
              <Button
                type="button"
                variant="outline"
                size="sm"
                onClick={() => table.previousPage()}
                disabled={!table.getCanPreviousPage()}
              >
                Previous
              </Button>
              <Button
                type="button"
                variant="outline"
                size="sm"
                onClick={() => table.nextPage()}
                disabled={!table.getCanNextPage()}
              >
                Next
              </Button>
            </div>
          </div>
        </SelectContent>
      </Select>
    </div>
  );
}

function getColumns({
  label,
  name,
  selectedOptions,
  setSelectedOptions,
}: {
  label: string;
  name: string;
  selectedOptions: string[];
  setSelectedOptions: React.Dispatch<React.SetStateAction<string[]>>;
}) {
  const columns: ColumnDef<{ value: string; label: string }>[] = [
    {
      id: "select",
      header: ({ table }) => (
        <Checkbox
          checked={table.getIsAllPageRowsSelected()}
          onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
          aria-label="Select all"
        />
      ),
      cell: ({ row }) => (
        <Checkbox
          checked={selectedOptions.includes(row.original.value)}
          onCheckedChange={(value) => {
            if (value) {
              selectedOptions.push(row.original.value);
              setSelectedOptions(selectedOptions.slice());
            } else {
              setSelectedOptions(selectedOptions.filter((val) => val !== row.original.value));
            }
            row.toggleSelected(!!value);
          }}
          aria-label="Select row"
          name={name}
          id={`${name}_${row.original.value}`}
          value={row.original.value}
        />
      ),
      enableSorting: false,
      enableHiding: false,
    },
    {
      accessorKey: "label",
      header: label,
      cell: ({ row }) => {
        return <Label htmlFor={`${name}_${row.original.value}`}>{row.original.label}</Label>;
      },
    },
  ];

  return columns;
}
