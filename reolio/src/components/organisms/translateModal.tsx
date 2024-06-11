import ReactModal from "react-modal";
import { Textarea } from "../ui/textarea";
import { LanguageCodeCombobox } from "../melecules/languageCodeCombobox";
import { Dispatch, SetStateAction, useEffect, useRef } from "react";
import { ArrowRight } from "lucide-react";
import { Button } from "../ui/button";
import { translateLanguages } from "@/domain/model/translateLanguages";

export interface TranslateModalProps {
  isOpen: boolean;
  onRequestClose?(event: React.MouseEvent | React.KeyboardEvent): void;

  text: string;
  setText: Dispatch<SetStateAction<string>>;

  translatedText: string;

  sourceLanguageCode: string;
  setSourceLanguageCode: Dispatch<SetStateAction<string>>;

  targetLanguageCode: string;
  setTargetLanguageCode: Dispatch<SetStateAction<string>>;
}

const defaultStyle = {
  overlay: {
    background: "rgba(0,0,0, 0.7)",
    zIndex: 15, // headerより上
  },
  content: {
    width: "70vw",
    top: "20%",
    left: "50%",
    right: "auto",
    bottom: "auto",
    marginRight: "-50%",
    transform: "translate(-50%, -50%)",
    background: "black",
  },
};

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
      }, 100); // 100ミリ秒の遅延を入れる
    }
  }, [isOpen]);

  return (
    <ReactModal
      isOpen={isOpen}
      style={defaultStyle}
      onRequestClose={onRequestClose}
      contentLabel="Translate Panel"
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
          ></LanguageCodeCombobox>
          <Button className="" onClick={onExchangeLanguages}>
            <ArrowRight></ArrowRight>
          </Button>
          <LanguageCodeCombobox
            languageCode={targetLanguageCode}
            setLanguageCode={setTargetLanguageCode}
            emptyText="Target Language..."
            className=""
            languages={translateLanguages}
          ></LanguageCodeCombobox>
        </div>
        <Textarea
          className="bg-black text-white p-3"
          readOnly
          placeholder="翻訳後"
          value={translatedText}
        />
      </div>
    </ReactModal>
  );
}
