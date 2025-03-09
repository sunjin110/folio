import * as Dialog from "@radix-ui/react-dialog";
import { Textarea } from "../ui/textarea";
import { LanguageCodeCombobox } from "../melecules/languageCodeCombobox";
import { Dispatch, SetStateAction, useEffect, useRef } from "react";
import { ArrowRight } from "lucide-react";
import { Button } from "../ui/button";
import { translateLanguages } from "@/domain/model/translateLanguages";

export interface TranslateModalProps {
  isOpen: boolean;
  onRequestClose?(): void;

  text: string;
  setText: Dispatch<SetStateAction<string>>;

  translatedText: string;

  sourceLanguageCode: string;
  setSourceLanguageCode: Dispatch<SetStateAction<string>>;

  targetLanguageCode: string;
  setTargetLanguageCode: Dispatch<SetStateAction<string>>;
}

export function TranslateModal(props: TranslateModalProps) {
  const {
    isOpen,
    onRequestClose,
    text,
    setText,
    translatedText,
    sourceLanguageCode,
    setSourceLanguageCode,
    targetLanguageCode,
    setTargetLanguageCode,
  } = props;

  const textareaRef = useRef<HTMLTextAreaElement | null>(null);

  const onExchangeLanguages = () => {
    const tmpCode = sourceLanguageCode;
    setSourceLanguageCode(targetLanguageCode);
    setTargetLanguageCode(tmpCode);
  };

  useEffect(() => {
    if (isOpen) {
      setTimeout(() => {
        // 遅延させないとフォーカスできない
        if (textareaRef.current) {
          textareaRef.current.focus();
        }
      }, 100);
    }
  }, [isOpen]);

  return (
    <Dialog.Root
      open={isOpen}
      onOpenChange={(open) => {
        if (!open && onRequestClose) {
          onRequestClose();
        }
      }}
    >
      <Dialog.Portal>
        <Dialog.Overlay
          className="fixed inset-0"
          style={{ background: "rgba(0,0,0,0.7)", zIndex: 15 }}
        />
        <Dialog.Content
          className="fixed bg-white p-4"
          style={{
            width: "70vw",
            top: "25%",
            left: "50%",
            right: "auto",
            bottom: "auto",
            marginRight: "-50%",
            transform: "translate(-50%, -50%)",
            background: "black",
            border: "2px solid white",
            zIndex: 20,
          }}
        >
          <div className="m-1">
            <h1 className="text-2xl pb-3">🐼 Translate Panel 🐼</h1>
            <Textarea
              className="bg-black text-white p-3 mb-4"
              placeholder="翻訳前"
              value={text}
              ref={textareaRef}
              onChange={(event) => setText(event.target.value)}
            />
            <div className="w-full pb-3 flex justify-between">
              <LanguageCodeCombobox
                languageCode={sourceLanguageCode}
                setLanguageCode={setSourceLanguageCode}
                emptyText="Source Language..."
                className=""
                languages={translateLanguages}
              />
              <Button className="" onClick={onExchangeLanguages}>
                <ArrowRight />
              </Button>
              <LanguageCodeCombobox
                languageCode={targetLanguageCode}
                setLanguageCode={setTargetLanguageCode}
                emptyText="Target Language..."
                className=""
                languages={translateLanguages}
              />
            </div>
            <Textarea
              className="bg-black text-white p-3"
              readOnly
              placeholder="翻訳後"
              value={translatedText}
            />
            <Dialog.Close asChild>
              <Button className="mt-4" onClick={onRequestClose}>
                Close
              </Button>
            </Dialog.Close>
          </div>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
}
