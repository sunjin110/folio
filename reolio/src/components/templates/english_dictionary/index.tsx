import { Navigation } from "@/components/organisms/navigation";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { WordDetailWithTranslation } from "@/domain/model/word_detail_with_translation";
import { Dispatch, SetStateAction, useState } from "react";

interface EnglishDictionaryProps {
  word: string;
  setWord: Dispatch<SetStateAction<string>>;
  wordDetailWithTranslation?: WordDetailWithTranslation;
}

export default function EnglishDictionaryTemplate(
  props: EnglishDictionaryProps,
) {
  const { word, setWord, wordDetailWithTranslation } = props;

  const [viewWord, setViewWord] = useState(word);

  const handleWordKeyDown = (
    e: React.KeyboardEvent<HTMLInputElement>,
  ) => {
    if (e.nativeEvent.isComposing || e.key !== "Enter") return;
    setWord(viewWord);
  };

  return (
    <Navigation title="English Dictionary" sidebarPosition="english_dictionary">
      <div className="flex flex-col h-full">
        <div>
          <Input
            type="text"
            placeholder="word"
            onKeyDown={handleWordKeyDown}
            value={viewWord}
            onChange={(event) => setViewWord(event.target.value)}
          />
        </div>
        {wordDetailWithTranslation && (
          <div className="mt-2">
            <h1 className="text-xl">
              {wordDetailWithTranslation.origin.word}:{" "}
              {wordDetailWithTranslation.translated.word}
            </h1>

            {wordDetailWithTranslation.origin.definitions.map(
              (wordDefinition, index) => {
                const translatedDefinition =
                  wordDetailWithTranslation.translated.definitions[index];

                return (
                  <Card className="mt-4">
                    <CardHeader>
                      <CardTitle>{wordDefinition.partOfSpeech}: {wordDefinition.definition}</CardTitle>
                      <CardDescription>
                        {translatedDefinition.definition}
                      </CardDescription>
                    </CardHeader>
                    <CardContent className="">
                      {wordDefinition.synonyms && (
                        <div>
                          <h2 className="text-base">Synonyms</h2>
                          <ol className="list-disc pl-10">
                            {wordDefinition.synonyms.map((synonym, index) => {
                              return (
                                <li>
                                  {synonym}:{" "}
                                  {translatedDefinition.synonyms[index]}
                                </li>
                              );
                            })}
                          </ol>
                        </div>
                      )}

                      {wordDefinition.antonyms && (
                        <div>
                          <h2 className="text-base">Antonyms</h2>
                          <ol className="list-disc pl-10">
                            {wordDefinition.antonyms.map((antonyms, index) => {
                              return (
                                <li>
                                  {antonyms}:{" "}
                                  {translatedDefinition.antonyms[index]}
                                </li>
                              );
                            })}
                          </ol>
                        </div>
                      )}

                      {wordDefinition.examples && (
                        <div>
                          <h2 className="text-base">Examples</h2>
                          <ol className="list-disc pl-10">
                            {wordDefinition.examples.map((example, index) => {
                              return (
                                <li>
                                  {example}
                                  <ol className="list-disc text-sm text-slate-400">
                                    {translatedDefinition.examples[index]}
                                  </ol>
                                </li>
                              );
                            })}
                          </ol>
                        </div>
                      )}
                    </CardContent>
                  </Card>
                );
              },
            )}
          </div>
        )}
      </div>
    </Navigation>
  );
}
