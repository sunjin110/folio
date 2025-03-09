import { TranslateModal } from "@/components/organisms/translateModal";
import { useToast } from "@/components/ui/use-toast";
import { TranslationRepository } from "@/domain/repository/translation";
import { AuthError, InternalError } from "@/error/error";
import { useDelayState } from "@/hooks/useDelayState";
import { useEffect, useState } from "react";

interface TranslateModalPageProps {
  translationRepository: TranslationRepository;
}

export function TranslateModalPage(props: TranslateModalPageProps) {
  const { translationRepository } = props;

  const [text, setText] = useState("");
  const [translatedText, setTranslatedText] = useState("");
  const [delayText] = useDelayState(text, 1000); // 変更終了後1秒後にリクエスト
  const [sourceLanguageCode, setSourceLanguageCode] = useState("ja");
  const [targetLanguageCode, setTargetLanguageCode] = useState("en");

  const { toast } = useToast();

  // translate modal setting
  const [isTranslateModalOpen, setIsTranslateModalOpen] = useState(false);
  const onTranslateModalRequestClose = () => {
    setIsTranslateModalOpen(false);
  };

  // cmd + u で modal を開閉できるようにする
  useEffect(() => {
    const handleTranslateModalOpen = (event: KeyboardEvent) => {
      if ((event.metaKey || event.ctrlKey) && event.key === "u") {
        event.preventDefault();
        setIsTranslateModalOpen((prev) => !prev);
      }
    };
    window.addEventListener("keydown", handleTranslateModalOpen);
    return () => {
      window.removeEventListener("keydown", handleTranslateModalOpen);
    };
  }, []);

  useEffect(() => {
    if (!delayText) {
      setTranslatedText("");
      return;
    }

    const fetchTranslation = async (text: string) => {
      try {
        const output = await translationRepository.TranslationText(
          text,
          sourceLanguageCode,
          targetLanguageCode,
        );
        setTranslatedText(output);
      } catch (err) {
        console.error(err);
        if (err instanceof AuthError) {
          toast({
            title: "AuthError",
            description: err.message,
          });
          setIsTranslateModalOpen(false);
          return;
        } else if (err instanceof InternalError) {
          toast({
            title: "InternalError",
            description: err.message,
          });
          return;
        }
        toast({
          title: "Error",
          description: `${err}`,
        });
      }
    };
    fetchTranslation(delayText);
  }, [
    delayText,
    toast,
    translationRepository,
    sourceLanguageCode,
    targetLanguageCode,
  ]);

  return (
    <TranslateModal
      isOpen={isTranslateModalOpen}
      onRequestClose={onTranslateModalRequestClose}
      text={text}
      setText={setText}
      translatedText={translatedText}
      sourceLanguageCode={sourceLanguageCode}
      setSourceLanguageCode={setSourceLanguageCode}
      targetLanguageCode={targetLanguageCode}
      setTargetLanguageCode={setTargetLanguageCode}
    />
  );
}
