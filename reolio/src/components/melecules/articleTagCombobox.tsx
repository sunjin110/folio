import { Dispatch, SetStateAction, useState } from "react";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { Button } from "../ui/button";
import { ArticleTag } from "@/domain/model/article";
import { Check, ChevronsUpDown, Plus } from "lucide-react";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "../ui/command";
import { cn } from "@/lib/utils";

export interface ArticleTagComboboxProps {
  className?: string;
  searchText: string;
  setSearchText: Dispatch<SetStateAction<string>>;
  candidateTagMap: Map<string, ArticleTag>;
  selectedTagMap: Map<string, ArticleTag>;
  setSelectedTagMap: Dispatch<SetStateAction<Map<string, ArticleTag>>>;
  selectedNewTagNameMap: Map<string, boolean>;
  setSelectedNewTagNameMap: Dispatch<SetStateAction<Map<string, boolean>>>;
}

export function ArticleTagCombobox(props: ArticleTagComboboxProps) {
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

  const [open, setOpen] = useState(false);

  // タグの選択状態をトグルする関数
  const toggleTagSelection = (tag: ArticleTag) => {
    setSelectedTagMap((prevMap) => {
      const newMap = new Map(prevMap); // 新しいMapオブジェクトを作成
      if (newMap.has(tag.id)) {
        newMap.delete(tag.id); // タグが選択済みの場合は削除
      } else {
        newMap.set(tag.id, tag); // タグが未選択の場合は追加
      }
      return newMap;
    });
  };

  const toggleTagNameSelection = (tagName: string) => {
    setSelectedNewTagNameMap((prevMap) => {
      const newMap = new Map(prevMap);
      if (newMap.get(tagName)) {
        newMap.set(tagName, false);
      } else {
        newMap.set(tagName, true);
      }
      return newMap;
    });
  };

  return (
    <div className={className}>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            className="w-[200px] justify-between bg-black hover:bg-zinc-900 hover:text-white"
          >
            select tags...
            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-[200px] p-0">
          <Command className="bg-black text-white">
            <CommandInput
              value={searchText}
              onValueChange={(searchText) => setSearchText(searchText)}
              placeholder="Search tag..."
            />
            <CommandList>
              <CommandEmpty className="text-white">No tag found.</CommandEmpty>

              <CommandGroup>
                {/* create */}
                {searchText && !selectedNewTagNameMap.has(searchText) && (
                  <CommandItem
                    key={"new_tag"}
                    value={searchText}
                    onSelect={(currentValue) => {
                      toggleTagNameSelection(currentValue);
                    }}
                    className="hover:bg-zinc-900 text-white"
                  >
                    <Plus className="mr-2 h-4 w-4" />
                    {searchText}
                  </CommandItem>
                )}

                {selectedNewTagNameMap &&
                  Array.from(selectedNewTagNameMap).map(([key, value]) => {
                    return (
                      <CommandItem
                        key={key}
                        value={key}
                        onSelect={(currentKey) => {
                          toggleTagNameSelection(currentKey);
                        }}
                        className="hover:bg-zinc-900 text-white"
                      >
                        <Check
                          className={cn(
                            "mr-2 h-4 w-4",
                            value ? "opacity-100" : "opacity-0",
                          )}
                        />
                        {key}
                      </CommandItem>
                    );
                  })}

                {/* candidate tags */}
                {Array.from(candidateTagMap).map(([tagID, tag]) => (
                  <CommandItem
                    key={tag.id}
                    value={tag.name} // 検索した時のここの値で検索されるため
                    onSelect={() => {
                      const tag = candidateTagMap.get(tagID);
                      if (tag) {
                        toggleTagSelection(tag);
                      }
                    }}
                    className="hover:bg-zinc-900 text-white"
                  >
                    <Check
                      className={cn(
                        "mr-2 h-4 w-4",
                        selectedTagMap.has(tag.id)
                          ? "opacity-100"
                          : "opacity-0",
                      )}
                    />
                    {tag.name}
                  </CommandItem>
                ))}
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
    </div>
  );
}
