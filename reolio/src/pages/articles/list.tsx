import { useEffect, useState } from "react";
import { DataTable } from "@/components/data-table";
import { ColumnDef } from "@tanstack/react-table";
import { getArticles } from "@/api/api";
import { ArticleSummary } from "@/domain/model/article";
import { formatDateFromRFC } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Link, useNavigate } from "react-router-dom";
import { useToast } from "@/components/ui/use-toast";
import { Navigation } from "@/components/organisms/navigation";

const columns: ColumnDef<ArticleSummary>[] = [
    {
        accessorKey: "id",
        header: "",
        cell: ({row}) => {
            return (
                <Link to={`/articles/${row.getValue("id")}`}>
                    <Button>Detial</Button>
                </Link>
            )
        }
    },
    {
        accessorKey: "title",
        header: "Title"
    },
    {
        accessorKey: "created_at",
        header: () => <div className="">CreatedTime</div>,
        cell: ({row}) => {
            return formatDateFromRFC(row.getValue("created_at"));
        }
    },
    {
        accessorKey: "id",
        header: "Edit",
        cell: ({row}) => {
            return (
                <Link to={`/articles/${row.getValue("id")}`}>
                    <Button>TODO Edit</Button>
                </Link>
            )
        }
    },
];

export default function Articles() {
    const [data, setData] = useState<ArticleSummary[]>([]);

    const [pageSize, setPageSize] = useState(10);
    const [pageIndex, setPageIndex] = useState(0);
    const [pageCount, setPageCount] = useState(0);
    const { toast } = useToast();
    const navigate = useNavigate();

    useEffect(() => {
        const fetch = async () => {

            const offset = pageIndex * pageSize;
            const limit = pageSize;
            try {
                const output = await getArticles(offset, limit);
                if (output.type === 'error') {
                    toast({title: "ログインし直してください", description: output.message});
                    navigate("/login");
                    return;
                }
                setData(output.articles);
                const total = output.total;
                setPageCount(Math.ceil(total / pageSize));
            } catch (err) {
                console.error("failed get article summaries", err);
            }
        };
        fetch();
    },[pageIndex, pageSize]);

    const onPageChange = (newPageIndex: number) => {
        setPageIndex(newPageIndex);
    };

    return (
        <Navigation title="Articles" sidebarPosition='articles'>
            <div className="container mx-auto py-10">
                <Link to={`/articles/create`}>
                    <Button>Create Article</Button>
                </Link>
                <DataTable columns={columns} data={data} pageIndex={pageIndex} pageSize={pageSize} onPageChange={onPageChange} pageCount={pageCount}></DataTable>
            </div>
        </Navigation>
    )
}
