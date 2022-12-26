import { useMemo, useState } from "react";
import z from "zod";
import dayjs from "@utils/date";
import {
  useReactTable,
  getCoreRowModel,
  getSortedRowModel,
  flexRender,
} from "@tanstack/react-table";
import type { ColumnDef, SortingState } from "@tanstack/react-table";
import type { Complaint } from "@api/complaints";

const ComplaintsTable: React.FC<{ tableData: Complaint[] }> = ({
  tableData,
}) => {
  const [sorting, setSorting] = useState<SortingState>([]);

  const columns = useMemo<ColumnDef<Complaint>[]>(
    () => [
      {
        header: "ID",
        accessorKey: "id",
        cell: (info) => info.getValue(),
      },
      {
        header: "Datum",
        accessorKey: "date",
        cell: (info) =>
          dayjs(z.string().parse(info.getValue())).format("DD.MM.YYYY"),
      },
      {
        header: "Auftragsnummer",
        accessorKey: "order",
        cell: (info) => info.getValue(),
      },
      {
        header: "Maschine",
        accessorKey: "machine",
        cell: (info) => info.getValue(),
      },
      {
        header: "Menge",
        accessorKey: "quantity",
        cell: (info) => info.getValue(),
      },
      {
        header: "Kosten",
        accessorKey: "cost",
        cell: (info) => info.getValue(),
      },
      {
        header: "Grund",
        accessorKey: "reason",
        cell: (info) => info.getValue(),
      },
      {
        header: "Material",
        accessorKey: "material",
        cell: (info) => info.getValue(),
      },
    ],
    [tableData]
  );

  const table = useReactTable({
    data: tableData,
    columns,
    state: {
      sorting,
    },
    onSortingChange: setSorting,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
  });

  return (
    <div className="w-full">
      <table>
        <thead>
          {table.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <th key={header.id} colSpan={header.colSpan}>
                    {header.isPlaceholder ? null : (
                      <div
                        {...{
                          className: header.column.getCanSort()
                            ? "cursor-pointer select-none"
                            : "",
                          onClick: header.column.getToggleSortingHandler(),
                        }}
                      >
                        {flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                        {{
                          asc: " 🔼",
                          desc: " 🔽",
                        }[header.column.getIsSorted() as string] ?? null}
                      </div>
                    )}
                  </th>
                );
              })}
            </tr>
          ))}
        </thead>
        <tbody>
          {table
            .getRowModel()
            .rows.slice(0, 10)
            .map((row) => {
              return (
                <tr key={row.id}>
                  {row.getVisibleCells().map((cell) => {
                    return (
                      <td key={cell.id}>
                        {flexRender(
                          cell.column.columnDef.cell,
                          cell.getContext()
                        )}
                      </td>
                    );
                  })}
                </tr>
              );
            })}
        </tbody>
      </table>
    </div>
  );
};

export default ComplaintsTable;
