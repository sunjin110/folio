import React from "react"; // Reactのインポートが必要
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
  getPaginationRowModel,
} from "@tanstack/react-table";
import { Button } from "./ui/button";

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data: TData[];
  pageIndex: number;
  pageSize: number;
  onPageChange: (newPageIndex: number) => void;
  pageCount: number;
}

export function DataTable<TData, TValue>({
  columns,
  data,
  pageIndex,
  pageSize,
  onPageChange,
  pageCount,
}: DataTableProps<TData, TValue>): React.JSX.Element {
  const table = useReactTable({
    data,
    columns,
    // state: { // データが固定のときはこれを使う
    //     pagination: {
    //         pageIndex: pageIndex,
    //         pageSize: pageSize,
    //     }
    // },
    pageCount: pageCount,
    onStateChange: (state) => {
      if (typeof state !== "function") {
        const newState = state as unknown as {
          pagination: { pageIndex: number };
        };
        const newPageIndex = newState.pagination.pageIndex;
        if (newPageIndex !== pageIndex) {
          onPageChange(newPageIndex);
        }
      }
    },
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
  });

  return (
    <div>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext(),
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-sate={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext(),
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <Button
          variant="outline"
          size="sm"
          onClick={() => onPageChange(Math.max(pageIndex - 1, 0))}
          disabled={pageIndex <= 0}
        >
          Previous
        </Button>
        <Button
          variant="outline"
          size="sm"
          onClick={() => onPageChange(pageIndex + 1)}
          disabled={pageIndex >= pageCount - 1}
        >
          Next
        </Button>
      </div>
    </div>
  );
}
