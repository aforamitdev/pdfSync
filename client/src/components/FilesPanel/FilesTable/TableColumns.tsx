import { Avatar, Checkbox } from "@radix-ui/themes";
import { ColumnDef } from "@tanstack/react-table";
import { TFilesList } from "../types";


export const columns: ColumnDef<TFilesList>[] = [
    {
        id: "id",
        header: ({ table }) => (
            <div className="flex">

                <Checkbox
                    size="1"
                    checked={
                        table.getIsAllPageRowsSelected() ||
                        (table.getIsSomePageRowsSelected() && "indeterminate")
                    }
                    onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
                    aria-label="Select all"
                />
                <div className="px-2">Files & Folders</div>
            </div>
        ),
        cell: ({ row }) => {

            console.log(row)

            return <div className="flex items-center gap-2">
                <Checkbox
                    size="1"
                    checked={row.getIsSelected()}
                    onCheckedChange={(value) => row.toggleSelected(!!value)}
                    aria-label="Select row"
                />
                <div className="font-medium">

                    {row.original.file}
                </div>
            </div>
        }
        ,

        enableSorting: false,
        enableHiding: false,
    },

    {
        accessorKey: "creator",
        header: "Created By",
        cell: ({ row }) => (
            <div className="flex items-center gap-1">
                <Avatar size="1" radius="full" fallback={row.getValue("creator")[0] || "A"} />
                <div className="capitalize">{row.getValue("creator")}</div>
            </div>
        ),
    },

    {
        accessorKey: "createdat",
        header: "Created On",
        cell: ({ row }) => (
            <div className="flex items-center gap-1">
                <div className="capitalize">{new Date(row.getValue("createdat")).getFullYear()}</div>
            </div>
        ),
    },
    {
        accessorKey: "access",
        header: "Access",
        cell: ({ row }) => (
            <div className="flex items-center gap-1">
                <div className="capitalize">{row.getValue("access")}</div>
            </div>
        ),
    },


]