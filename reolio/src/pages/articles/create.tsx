import CreateArticleTemplate from "@/components/templates/articles/create";
import { useToast } from "@/components/ui/use-toast";
import { handleError } from "@/error/pageErrorHandle";
import { ArticleUsecase } from "@/usecase/article";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

interface CreateArticleProps {
  articleUsecase: ArticleUsecase;
}

export default function CreateArticle(props: CreateArticleProps) {
  const { articleUsecase } = props;

  const [title, setTitle] = useState("");
  const [body, setBody] = useState<string | undefined>("");

  const [aiPrompt, setAiPrompt] = useState("");

  const { toast } = useToast();
  const navigate = useNavigate();

  const handlePost = async () => {
    try {
      const articleId = await articleUsecase.InsertArticle(
        title,
        body === undefined ? "" : body,
        [],
      ); // TODO tagIDs
      navigate(`/articles/edit/${articleId}`);
    } catch (err) {
      const resp = handleError(err);
      toast(resp.toast);
      if (resp.navigationPath) {
        navigate(resp.navigationPath);
      }
    }
  };

  const handleGenerateAI = async () => {
    if (!aiPrompt) {
      return;
    }
    toast({
      title: "🧠 start generating article 🧠",
      description: `prompt: ${aiPrompt}`,
    });

    try {
      const output = await articleUsecase.GenerateArtixleByAI(aiPrompt);
      toast({
        title: "🐨 finish generating article 🐨",
      });
      navigate(`/articles/${output}`);
      return;
    } catch (err) {
      const resp = handleError(err);
      toast(resp.toast);
      if (resp.navigationPath) {
        navigate(resp.navigationPath);
      }
    }
  };

  return (
    <CreateArticleTemplate
      aiPrompt={aiPrompt}
      setAiPrompt={setAiPrompt}
      title={title}
      setTitle={setTitle}
      body={body}
      setBody={setBody}
      handlePost={handlePost}
      handleGenerateAI={handleGenerateAI}
    />
  );
}
