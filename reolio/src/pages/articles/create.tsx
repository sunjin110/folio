import { createArticle } from "@/api/api";
import CreateArticleTemplate from "@/components/templates/articles/create";
import { useToast } from "@/components/ui/use-toast";
import { AuthError } from "@/error/error";
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
      const output = await createArticle(title, body === undefined ? "" : body);
      if (output.type === "error") {
        toast({
          title: "ログインし直してください",
          description: output.message,
        });
        navigate("/login");
        return;
      }
      toast({
        title: "Success",
        description: "posted your article!",
      });
      navigate(`/articles/edit/${output.id}`);
    } catch (err) {
      console.error("failed create article", err);
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
      if (err instanceof AuthError) {
        console.error(err);
        toast({
          title: "please login again",
          description: err.message,
        });
        navigate("/login");
        return;
      }
      toast({
        title: "internal error",
        description: `${err}`,
      });
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
