import { ArticleTagSearchCombobox } from "@/components/melecules/articleTagSearchCombobox";
import { ArticleTag } from "@/domain/model/article";
import { useState } from "react";

export default function About() {
  const [searchText, setSearchText] = useState("");

  const [selectedTagMap, setSelectedTagMap] = useState<Map<string, ArticleTag>>(
    new Map(),
  );

  return (
    <div>
      <ArticleTagSearchCombobox
        searchText={searchText}
        setSearchText={setSearchText}
        candidateTags={[]}
        selectedTagMap={selectedTagMap}
        setSelectedTagMap={setSelectedTagMap}
      />
    </div>
  );
}
