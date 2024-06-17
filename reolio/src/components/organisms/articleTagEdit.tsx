import { ArticleTag } from "@/domain/model/article";
import { Dispatch, SetStateAction } from "react";
import { ArticleTagCombobox } from "../melecules/articleTagCombobox";
import { ArticleTagPreview } from "../melecules/articleTagPreview";

export interface ArticleTagEditProps {
  className?: string;
  searchText: string;
  setSearchText: Dispatch<SetStateAction<string>>;
  candidateTagMap: Map<string, ArticleTag>;
  selectedTagMap: Map<string, ArticleTag>;
  setSelectedTagMap: Dispatch<SetStateAction<Map<string, ArticleTag>>>;
  selectedNewTagNameMap: Map<string, boolean>;
  setSelectedNewTagNameMap: Dispatch<SetStateAction<Map<string, boolean>>>;
}

export function ArticleTagEdit(props: ArticleTagEditProps) {
  const {
    className,
    searchText,
    setSearchText,
    candidateTagMap,
    selectedTagMap,
    setSelectedTagMap,
    selectedNewTagNameMap,
    setSelectedNewTagNameMap,
  } = props;

  return (
    <div className={className}>
      <ArticleTagCombobox
        className="mb-2"
        candidateTagMap={candidateTagMap}
        searchText={searchText}
        setSearchText={setSearchText}
        selectedTagMap={selectedTagMap}
        setSelectedTagMap={setSelectedTagMap}
        selectedNewTagNameMap={selectedNewTagNameMap}
        setSelectedNewTagNameMap={setSelectedNewTagNameMap}
      />

      <ArticleTagPreview
        className="mb-1"
        selectedTagMap={selectedTagMap}
        setSelectedTagMap={setSelectedTagMap}
        selectedNewTagNameMap={selectedNewTagNameMap}
        setSelectedNewTagNameMap={setSelectedNewTagNameMap}
      />
    </div>
  );
}
