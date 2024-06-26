import EnglishDictionaryTemplate from "@/components/templates/english_dictionary";
import { useToast } from "@/components/ui/use-toast";
import { WordDetailWithTranslation } from "@/domain/model/word_detail_with_translation";
import { handleError } from "@/error/pageErrorHandle";
import { EnglishDictionaryUsecase } from "@/usecase/english_dictionary";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

interface EnglishDictionaryProps {
    englishDctionaryUsecase: EnglishDictionaryUsecase;
}

export default function EnglishDictionary(props: EnglishDictionaryProps) {

  const { englishDctionaryUsecase } = props;
  const [word, setWord] = useState("");

  const { toast } = useToast();
  const navigate = useNavigate();

  const [wordDetailWithTranslation, setWordDetailWithTranslation] = useState<WordDetailWithTranslation | undefined>(undefined);

  useEffect(() => {
    if (!word) {
        setWordDetailWithTranslation(undefined);
        return;
    }

    const fetch = async (word: string) => {
        try {
            const wordDetailWithTranslation = await englishDctionaryUsecase.GetWordDetialWithTranslation(word);
            setWordDetailWithTranslation(wordDetailWithTranslation);
        } catch (err) {
            const resp = handleError(err);
            toast(resp.toast);
            if (resp.navigationPath) {
                navigate(resp.navigationPath);
            }
        }
    };

    fetch(word);
  }, [word, englishDctionaryUsecase, navigate, toast]);

  return (
    <EnglishDictionaryTemplate
      word={word}
      setWord={setWord}
      wordDetailWithTranslation={wordDetailWithTranslation}
    />
  );
}
