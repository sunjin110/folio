import { useCallback, useEffect, useState } from "react";
import { DataTable } from "@/components/data-table";
import { ColumnDef } from "@tanstack/react-table";
import { ArticleSummary, ArticleTag } from "@/domain/model/article";
import { formatDateFromRFC } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Link, useNavigate } from "react-router-dom";
import { useToast } from "@/components/ui/use-toast";
import { Navigation } from "@/components/organisms/navigation";
import { Input } from "@/components/ui/input";
import { ArticleUsecase } from "@/usecase/article";
import { AuthError, InternalError } from "@/error/error";
import { Badge } from "@/components/ui/badge";
import { ArticleTagSearch } from "@/components/organisms/articleTagSearch";
import { useDelayState } from "@/hooks/useDelayState";
import { handleError } from "@/error/pageErrorHandle";

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
    accessorKey: "tags",
    header: "Tag",
    cell: ({ row }) => {
      const tags: ArticleTag[] | undefined = row.getValue("tags");
      if (!tags) {
        return <div></div>;
      }

      return (
        <div>
          {tags.map((tag) => {
            return (
              <Badge id={tag.id} variant="secondary">
                {tag.name}
              </Badge>
            );
          })}
        </div>
      );
    },
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

  // tag
  const [tagSearchText, setTagSearchText] = useState("");
  const [selectedTagMap, setSelectedTagMap] = useState<Map<string, ArticleTag>>(
    new Map(),
  );
  const [candidateTags, setCandidateTags] = useState<ArticleTag[]>([]);
  const [delayTagSearchText] = useDelayState(tagSearchText, 700);

  const { toast } = useToast();
  const navigate = useNavigate();

  const fetchArticles = useCallback(
    async (searchTitleText: string) => {
      const offset = pageIndex * pageSize;
      const limit = pageSize;

      const tagIDs: string[] = Array.from(selectedTagMap).map(
        ([tagID]) => tagID,
      );
      try {
        const resp = await articleUsecase.FindSummaries(
          offset,
          limit,
          searchTitleText,
          tagIDs,
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
    [articleUsecase, navigate, pageIndex, pageSize, toast, selectedTagMap],
  );

  useEffect(() => {
    fetchArticles(searchTitleText);
  }, [
    pageIndex,
    pageSize,
    navigate,
    toast,
    fetchArticles,
    searchTitleText,
    selectedTagMap,
  ]);

  useEffect(() => {
    if (!delayTagSearchText) {
      setCandidateTags([]);
      return;
    }

    const fetch = async (tagSearchText: string) => {
      try {
        const articleTags = await articleUsecase.FindTags(tagSearchText, 0, 10);
        setCandidateTags(articleTags);
      } catch (err) {
        const resp = handleError(err);
        toast(resp.toast);
        if (resp.navigationPath) {
          navigate(resp.navigationPath);
        }
      }
    };
    fetch(delayTagSearchText);
  }, [delayTagSearchText, articleUsecase, navigate, toast]);

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
    <Navigation
      title="Articles"
      sidebarPosition="articles"
      headerContent={
        <div className="grid justify-items-end">
          <Link to={`/articles/create`}>
            <Button>Create Article</Button>
          </Link>
        </div>
      }
    >
      <div className="container mx-auto py-3">
        <div className="pb-2 pt-2">
          <Input
            id="search_title_text"
            type="text"
            placeholder="Search Title"
            onKeyDown={handleSearchTitleKeyDown}
            value={viewSearchTitleText}
            onChange={(event) => setViewSearchTitleText(event.target.value)}
          ></Input>
        </div>
        <div className="pb-2 pt-2">
          <ArticleTagSearch
            className=""
            searchText={tagSearchText}
            setSearchText={setTagSearchText}
            candidateTags={candidateTags}
            selectedTagMap={selectedTagMap}
            setSelectedTagMap={setSelectedTagMap}
          />
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
