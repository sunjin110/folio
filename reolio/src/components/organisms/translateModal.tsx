import ReactModal from "react-modal";
import { Textarea } from "../ui/textarea";
import { LanguageCodeCombobox } from "../melecules/languageCodeCombobox";
import { useState } from "react";
import { ArrowRight } from "lucide-react";
import { Button } from "../ui/button";
import { translateLanguages } from "@/domain/model/translateLanguages";

export interface TranslateModalProps {
    isOpen: boolean;
    onRequestClose?(event: React.MouseEvent | React.KeyboardEvent): void;
};

const defaultStyle = {
    overlay: {
      background: 'rgba(0,0,0, 0.7)',
      zIndex: 15, // headerより上
    },
    content: {
      width: '70vw',
      top: '20%',
      left: '50%',
      right: 'auto',
      bottom: 'auto',
      marginRight: '-50%',
      transform: 'translate(-50%, -50%)',
      background: 'black',
    }
  };

export function TranslateModal(props: TranslateModalProps) {
    const { isOpen, onRequestClose } = props;

    const [sourceLanguageCode, setSourceLanguageCode] = useState("ja");
    const [targetLanguageCode, setTargetLanguageCode] = useState("en");

    const onExchangeLanguages = () => {
      const tmpCode = sourceLanguageCode;
      setSourceLanguageCode(targetLanguageCode);
      setTargetLanguageCode(tmpCode);
    };

    return <ReactModal isOpen={isOpen} style={defaultStyle} onRequestClose={onRequestClose} contentLabel="Translate Panel">
        <div className="m-1">
            <h1 className="text-2xl pb-3">🐼 Translate Panel 🐼</h1>
            <Textarea className="bg-black text-white p-3 mb-4" placeholder="翻訳前" />
            <div className="w-full pb-3 flex justify-between">
              <LanguageCodeCombobox 
                languageCode={sourceLanguageCode} 
                setLanguageCode={setSourceLanguageCode} 
                emptyText="Source Language..."
                className=""
                languages={translateLanguages}
                ></LanguageCodeCombobox>
              <Button className="" onClick={onExchangeLanguages}><ArrowRight></ArrowRight></Button>
              <LanguageCodeCombobox
                languageCode={targetLanguageCode}
                setLanguageCode={setTargetLanguageCode}
                emptyText="Target Language..."
                className=""
                languages={translateLanguages}
              ></LanguageCodeCombobox>
            </div>
            <Textarea className="bg-black text-white p-3" placeholder="翻訳後" />
        </div>
    </ReactModal>;
}
