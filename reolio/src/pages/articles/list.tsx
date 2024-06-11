import { useCallback, useEffect, useState } from "react";
import { DataTable } from "@/components/data-table";
import { ColumnDef } from "@tanstack/react-table";
import { ArticleSummary } from "@/domain/model/article";
import { formatDateFromRFC } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Link, useNavigate } from "react-router-dom";
import { useToast } from "@/components/ui/use-toast";
import { Navigation } from "@/components/organisms/navigation";
import { Input } from "@/components/ui/input";
import { Label } from "@radix-ui/react-label";
import { ArticleUsecase } from "@/usecase/article";
import { AuthError, InternalError } from "@/error/error";

const columns: ColumnDef<ArticleSummary>[] = [
  {
    accessorKey: "id",
    header: "",
    cell: ({ row }) => {
      return (
        <Link to={`/articles/${row.getValue("id")}`}>
          <Button>Detial</Button>
        </Link>
      );
    },
  },
  {
    accessorKey: "title",
    header: "Title",
  },
  {
    accessorKey: "created_at",
    header: () => <div className="">CreatedTime</div>,
    cell: ({ row }) => {
      return formatDateFromRFC(row.getValue("created_at"));
    },
  },
  {
    accessorKey: "id",
    header: "Edit",
    cell: ({ row }) => {
      return (
        <Link to={`/articles/edit/${row.getValue("id")}`}>
          <Button>Edit</Button>
        </Link>
      );
    },
  },
];

export interface ArticlesProps {
  articleUsecase: ArticleUsecase;
}

export default function Articles(props: ArticlesProps) {
  const { articleUsecase } = props;

  const [data, setData] = useState<ArticleSummary[]>([]);

  const [pageSize] = useState(10);
  const [pageIndex, setPageIndex] = useState(0);
  const [pageCount, setPageCount] = useState(0);
  const [searchTitleText, setSearchTitleText] = useState("");
  const [viewSearchTitleText, setViewSearchTitleText] = useState("");

  const { toast } = useToast();
  const navigate = useNavigate();

  const fetch = useCallback(
    async (searchTitleText: string) => {
      const offset = pageIndex * pageSize;
      const limit = pageSize;

      try {
        const resp = await articleUsecase.FindSummaries(
          offset,
          limit,
          searchTitleText,
        );
        setData(resp.summaries);
        setPageCount(Math.ceil(resp.totalCount / pageSize));
      } catch (err) {
        if (err instanceof AuthError) {
          toast({
            title: "Please login again",
            description: err.message,
          });
          navigate("/login");
          return;
        } else if (err instanceof InternalError) {
          toast({
            title: "Error",
            description: err.message,
          });
          return;
        }

        toast({
          title: "Error",
          description: `${err}`,
        });

        console.error(err);
      }
    },
    [articleUsecase, navigate, pageIndex, pageSize, toast],
  );

  useEffect(() => {
    fetch(searchTitleText);
  }, [pageIndex, pageSize, navigate, toast, fetch, searchTitleText]);

  const handleSearchTitleKeyDown = (
    e: React.KeyboardEvent<HTMLInputElement>,
  ) => {
    if (e.nativeEvent.isComposing || e.key !== "Enter") return;
    setSearchTitleText(viewSearchTitleText);
  };

  const onPageChange = (newPageIndex: number) => {
    setPageIndex(newPageIndex);
  };

  return (
    <Navigation title="Articles" sidebarPosition="articles">
      <div className="container mx-auto py-10">
        <Link to={`/articles/create`}>
          <Button>Create Article</Button>
        </Link>

        <div className="pb-2 pt-2">
          <Label id="search_title_text_label" htmlFor="search_title_text">
            Search Title: {searchTitleText}
          </Label>
          <Input
            id="search_title_text"
            type="text"
            placeholder="Title"
            onKeyDown={handleSearchTitleKeyDown}
            value={viewSearchTitleText}
            onChange={(event) => setViewSearchTitleText(event.target.value)}
          ></Input>
        </div>

        <DataTable
          columns={columns}
          data={data}
          pageIndex={pageIndex}
          pageSize={pageSize}
          onPageChange={onPageChange}
          pageCount={pageCount}
        ></DataTable>
      </div>
    </Navigation>
  );
}
