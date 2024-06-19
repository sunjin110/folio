import { ArticleTag } from "@/domain/model/article";
import { ArticleTagSearchCombobox } from "../melecules/articleTagSearchCombobox";
import { Dispatch, SetStateAction } from "react";
import { ArticleTagPreview } from "../melecules/articleTagPreview";

export interface ArticleTagSearchProps {
  className?: string;
  searchText: string;
  setSearchText: Dispatch<SetStateAction<string>>;
  candidateTags: ArticleTag[];
  selectedTagMap: Map<string, ArticleTag>;
  setSelectedTagMap: Dispatch<SetStateAction<Map<string, ArticleTag>>>;
}

export function ArticleTagSearch(props: ArticleTagSearchProps) {
  const {
    className,
    searchText,
    setSearchText,
    candidateTags,
    selectedTagMap,
    setSelectedTagMap,
  } = props;
  return (
    <div className={className}>
      <ArticleTagSearchCombobox
        className="mb-2"
        candidateTags={candidateTags}
        selectedTagMap={selectedTagMap}
        setSelectedTagMap={setSelectedTagMap}
        searchText={searchText}
        setSearchText={setSearchText}
      />
      <ArticleTagPreview
        className="mb-1"
        selectedTagMap={selectedTagMap}
        setSelectedTagMap={setSelectedTagMap}
      />
    </div>
  );
}
