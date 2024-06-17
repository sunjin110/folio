import { ArticleTag } from "@/domain/model/article";
import { Dispatch, SetStateAction } from "react";
import { Badge } from "../ui/badge";
import { Button } from "../ui/button";
import { CircleX } from "lucide-react";

export interface ArticleTagPreview {
  className?: string;
  selectedTagMap: Map<string, ArticleTag>;
  setSelectedTagMap: Dispatch<SetStateAction<Map<string, ArticleTag>>>;
  selectedNewTagNameMap: Map<string, boolean>;
  setSelectedNewTagNameMap: Dispatch<SetStateAction<Map<string, boolean>>>;
}

export function ArticleTagPreview(props: ArticleTagPreview) {
  const {
    className,
    selectedTagMap,
    setSelectedTagMap,
    selectedNewTagNameMap,
    setSelectedNewTagNameMap,
  } = props;

  const presseDeleteTag = (tagID: string) => {
    setSelectedTagMap((prevMap) => {
      const newMap = new Map(prevMap);
      newMap.delete(tagID);
      return newMap;
    });
  };

  const pressedDeleteNewTagName = (tagName: string) => {
    setSelectedNewTagNameMap((prevMap) => {
      const newMap = new Map(prevMap);
      newMap.delete(tagName);
      return newMap;
    });
  };

  return (
    <div className={className}>
      {selectedTagMap &&
        Array.from(selectedTagMap).map(([key, value]) => {
          return (
            <Badge className="m-1" id={key} variant={"secondary"}>
              {value.name}
              <Button
                variant="ghost"
                size="icon"
                className="size-5 ml-1"
                onClick={() => {
                  presseDeleteTag(key);
                }}
              >
                <CircleX className="size-5" />
              </Button>
            </Badge>
          );
        })}
      {selectedNewTagNameMap &&
        Array.from(selectedNewTagNameMap)
          .filter(([_, value]) => value)
          .map(([key]) => {
            return (
              <Badge className="m-1" id={key} variant={"secondary"}>
                {key}
                <Button
                  variant="ghost"
                  size="icon"
                  className="size-5 ml-1"
                  onClick={() => {
                    pressedDeleteNewTagName(key);
                  }}
                >
                  <CircleX className="size-5" />
                </Button>
              </Badge>
            );
          })}
    </div>
  );
}
