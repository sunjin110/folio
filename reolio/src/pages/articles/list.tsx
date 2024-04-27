import { useEffect, useState } from "react";
import { DataTable } from "@/components/data-table";
import { ColumnDef } from "@tanstack/react-table";
import { getArticles } from "@/api/api";
import { ArticleSummary } from "@/domain/model/article";
import { formatDateFromRFC } from "@/lib/utils";

const columns: ColumnDef<ArticleSummary>[] = [
    {
        accessorKey: "id",
        header: "ID",
    },
    {
        accessorKey: "title",
        header: "Title"
    },
    {
        accessorKey: "created_at",
        // header: "CreateTime",
        header: () => <div className="">CreatedTime</div>,
        cell: ({row}) => {
            return formatDateFromRFC(row.getValue("created_at"));
        }
    },
];

export default function Articles() {
    const [data, setData] = useState<ArticleSummary[]>([]);
    useEffect(() => {
        const fetch = async () => {
            try {
                setData(await getArticles());
            } catch (err) {
                console.error("failed get article summaries", err);
            }
        };
        fetch();
    },[]);
    return (<div className="container mx-auto py-10">
        <DataTable columns={columns} data={data}></DataTable>
    </div>);
}
