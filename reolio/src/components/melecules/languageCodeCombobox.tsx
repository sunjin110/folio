import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";
import { Dispatch, SetStateAction, useState } from "react";
import { Button } from "../ui/button";
import {
  Command,
  CommandEmpty,
  CommandInput,
  CommandItem,
  CommandList,
} from "../ui/command";
import { CommandGroup } from "cmdk";
import { Check, ChevronsUpDown } from "lucide-react";
import { cn } from "@/lib/utils";
import { TranslateLanguage } from "@/domain/model/translateLanguages";

export interface LanguageCodeComboboxProps {
  languageCode: string;
  setLanguageCode: Dispatch<SetStateAction<string>>;
  emptyText: string;
  className?: string;
  languages: TranslateLanguage[];
}

export function LanguageCodeCombobox(props: LanguageCodeComboboxProps) {
  const { languageCode, setLanguageCode, emptyText, className, languages } =
    props;

  const [open, setOpen] = useState(false);

  return (
    <div className={className}>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant={"outline"}
            role="combobox"
            aria-expanded={open}
            className="w-[200px] justify-between bg-black hover:bg-zinc-900 hover:text-white"
          >
            {languageCode
              ? languages.find((language) => language.code === languageCode)
                  ?.name
              : emptyText}
            <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-[200px] p-0">
          <Command className="bg-black text-white">
            <CommandInput placeholder={emptyText} />
            <CommandList>
              <CommandEmpty className="text-white">
                No language found.
              </CommandEmpty>
              <CommandGroup>
                {languages.map((language) => (
                  <CommandItem
                    key={language.code}
                    value={language.code}
                    onSelect={(currentCode: string) => {
                      setLanguageCode(
                        currentCode === languageCode ? "" : currentCode,
                      );
                      setOpen(false);
                    }}
                    className="hover:bg-zinc-900"
                  >
                    <Check
                      className={cn(
                        "mr-2 h-4 w-4",
                        languageCode === language.code
                          ? "opacity-100"
                          : "opacity-0",
                      )}
                    />
                    {language.name}
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
