import { FC, useEffect, useMemo, useState } from "react";
import z from "zod";

import {
  Column,
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";

import { jsonSchema } from "@utils/validators";

import type { JsonSchema, Json } from "@utils/validators";
import type { ColumnDef, PaginationState } from "@tanstack/react-table";

const PropsValidator = z.object({
  data: z.array(jsonSchema),
  columnHeaders: z.array(z.string()),
  rowsPerPage: z.number().min(1),
  onPageChange: z.function().args(z.number()).returns(z.void()).optional(),
});

type Props = z.infer<typeof PropsValidator>;

const Table: FC<Props> = ({
  data,
  columnHeaders,
  rowsPerPage,
  onPageChange,
}) => {
  const columnHelper = createColumnHelper<any>();

  // const [columns, setColumns] = useState<ColumnDef<any>[]>([]);

  const columns = useMemo<ColumnDef<any>[]>(() => [], []);

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  // useEffect(() => {
  //   const columns = columnHeaders.map((col) =>
  //     columnHelper.accessor(col, {
  //       id: col,
  //       cell: (info) => info.getValue(),
  //       footer: (info) => info.column.id,
  //     })
  //   );
  //   setColumns(columns);
  // }, [data]);

  return (
    <table>
      <thead>
        {table.getHeaderGroups().map((headerGroup) => (
          <tr key={headerGroup.id}>
            {headerGroup.headers.map((header) => (
              <th key={header.id}>
                {header.isPlaceholder
                  ? null
                  : flexRender(
                      header.column.columnDef.header,
                      header.getContext()
                    )}
              </th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody>
        {table.getRowModel().rows.map((row) => (
          <tr key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <td key={cell.id}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default Table;
