import { Dispatch, SetStateAction, useState } from "react";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { Button } from "../ui/button";
import { Check, ChevronsUpDown } from "lucide-react";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "../ui/command";
import { ArticleTag } from "@/domain/model/article";
import { cn } from "@/lib/utils";

export interface ArticleTagSearchComboboxProps {
  className?: string;
  searchText: string;
  setSearchText: Dispatch<SetStateAction<string>>;
  candidateTags: ArticleTag[];
  selectedTagMap: Map<string, ArticleTag>;
  setSelectedTagMap: Dispatch<SetStateAction<Map<string, ArticleTag>>>;
}

export function ArticleTagSearchCombobox(props: ArticleTagSearchComboboxProps) {
  const {
    className,
    searchText,
    setSearchText,
    candidateTags,
    selectedTagMap,
    setSelectedTagMap,
  } = props;

  const [open, setOpen] = useState(false);

  const toggleTagSelection = (tag: ArticleTag) => {
    setSelectedTagMap((prevMap) => {
      const newMap = new Map(prevMap);
      if (newMap.has(tag.id)) {
        newMap.delete(tag.id);
      } else {
        newMap.set(tag.id, tag);
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
                {candidateTags.map((tag) => (
                  <CommandItem
                    key={tag.id}
                    value={tag.name}
                    onSelect={() => {
                      toggleTagSelection(tag);
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
