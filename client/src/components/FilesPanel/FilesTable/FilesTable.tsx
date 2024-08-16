import { Box, Table } from '@radix-ui/themes'
import { ColumnFiltersState, flexRender, getCoreRowModel, getFilteredRowModel, getPaginationRowModel, getSortedRowModel, SortingState, useReactTable, VisibilityState } from '@tanstack/react-table'
import React from 'react'
import FilesTablePageing from './FilesTablePageing'
import FilesManuHeader from './FilesManuHeader'

type Props = { data: any[], columns: any[] }

const FilesTable = ({ data, columns }: Props) => {

    const [sorting, setSorting] = React.useState<SortingState>([])
    const [columnFilters, setColumnFilters] = React.useState<ColumnFiltersState>(
        []
    )
    const [columnVisibility, setColumnVisibility] =
        React.useState<VisibilityState>({})
    const [rowSelection, setRowSelection] = React.useState({})

    const table = useReactTable({
        data,
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


    })

    return (
        <div className="rounded-sm pt-5">
            <Box className='w-full'>
                <FilesManuHeader />
            </Box>
            <Table.Root variant="surface"  >
                <Table.Header className='h-1'>
                    {table.getHeaderGroups().map((headerGroup) => (
                        <Table.Row key={headerGroup.id}>
                            {headerGroup.headers.map((header) => {
                                return (
                                    <Table.ColumnHeaderCell key={header.id} className='text-sm '>

                                        <div className="flex items-center text-xs  ">
                                            {header.isPlaceholder
                                                ? null
                                                : flexRender(
                                                    header.column.columnDef.header,
                                                    header.getContext(),
                                                )}
                                        </div>
                                    </Table.ColumnHeaderCell>
                                )
                            })}
                        </Table.Row>
                    ))}
                </Table.Header>
                <Table.Body>
                    {table.getRowModel().rows?.length ? (
                        table.getRowModel().rows.map((row) => (
                            <Table.Row
                                key={row.id}
                                data-state={row.getIsSelected() && "selected"}
                            >
                                {row.getVisibleCells().map((cell) => (
                                    <Table.Cell key={cell.id}>
                                        {flexRender(
                                            cell.column.columnDef.cell,
                                            cell.getContext()
                                        )}
                                    </Table.Cell>
                                ))}
                            </Table.Row>
                        ))
                    ) : (
                        <Table.Row>
                            <Table.Cell
                                colSpan={columns.length}
                                className="h-24 text-center"
                            >
                                No results.
                            </Table.Cell>
                        </Table.Row>
                    )}
                </Table.Body>
            </Table.Root>

            <FilesTablePageing table={table} />
        </div>
    )
}

export default FilesTable