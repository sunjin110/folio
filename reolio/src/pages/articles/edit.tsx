import { getArticleById, updateArticle } from "@/api/api";
import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/components/ui/use-toast";
import { useCallback, useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import MDEditor from "@uiw/react-md-editor";
import { getRandomEmoji } from "@/domain/service/joke";
import { ArticleUsecase } from "@/usecase/article";
import { AuthError, InternalError } from "@/error/error";

export interface ArticleEditProps {
  articleUsecase: ArticleUsecase;
}

export default function EditArticle(props: ArticleEditProps) {
  const { articleUsecase } = props;

  const { articleId } = useParams();
  const { toast } = useToast();
  const navigate = useNavigate();

  const [title, setTitle] = useState<string>("");
  const [body, setBody] = useState<string | undefined>("");

  const [prompt, setPrompt] = useState<string>("");

  useEffect(() => {
    const fetch = async (id: string) => {
      try {
        const output = await getArticleById(id);
        if (output.type === "error") {
          toast({ title: "please login", description: output.message });
          navigate("/login");
          return;
        }
        setTitle(output.article.title);
        setBody(output.article.body);
      } catch (err) {
        console.error("Error fetching article", err);
      }
    };
    if (articleId) {
      fetch(articleId);
    }
  }, [articleId, navigate, toast]);

  const handleEdit = useCallback(async () => {
    try {
      if (!articleId) {
        console.log("articleId is empty");
        return;
      }

      if (!title) {
        toast({
          title: "タイトルを入力してください",
        });
        return;
      }

      if (!body) {
        toast({
          title: "本文を入力してください",
        });
        return;
      }

      const output = await updateArticle(
        articleId,
        title,
        body === undefined ? "" : body,
      );
      if (output.type === "error") {
        toast({
          title: "ログインし直してください",
          description: output.message,
        });
        navigate("/login");
        return;
      }
      const emoji = getRandomEmoji();
      toast({
        title: `${emoji} Success ${emoji}`,
        description: "updated your article!",
      });
    } catch (err) {
      console.error("failed cedit article", err);
    }
  }, [articleId, title, body, toast, navigate]);

  const handleGenerateBody = useCallback(async () => {
    if (!articleId) {
      return;
    }

    const beforeBody = body;

    toast({
      title: "🧠🧠🧠🧠 AI generate start!!! 🧠🧠🧠🧠",
      description: `prompt: ${prompt}`,
    });

    try {
      const generatedBody = await articleUsecase.AsistantBodyByAI(
        articleId,
        prompt,
      );
      setBody(
        `${generatedBody}\n\n---\n\n # article before change\n\n${beforeBody}`,
      );
    } catch (err) {
      if (err instanceof AuthError) {
        toast({
          title: "Please login again",
          description: err.message,
        });
        navigate("/login");
        return;
      } else if (err instanceof InternalError) {
        toast({
          title: "Error",
          description: err.message,
        });
        return;
      }
      toast({
        title: "Error",
        description: `${err}`,
      });
      console.error(err);
    }

    toast({
      title: "🐎🐎🐎🐎 AI generate finished!!! 🐎🐎🐎🐎",
      description: `prompt: ${prompt}`,
    });
  }, [prompt, articleId, articleUsecase, navigate, toast, body]);

  useEffect(() => {
    const handleSaveShortcut = (event: KeyboardEvent) => {
      if ((event.metaKey || event.ctrlKey) && event.key === "s") {
        event.preventDefault();
        handleEdit();
      }
    };

    window.addEventListener("keydown", handleSaveShortcut);

    return () => {
      // コンポーネントがアンマウントされるときにリスナーをクリーンアップ
      window.removeEventListener("keydown", handleSaveShortcut);
    };
  }, [handleEdit]);

  return (
    <Navigation title="Articles" sidebarPosition="articles">
      <div className="flex flex-col h-full p-2">
        <div className="flex flex-col flex-grow">
          <div className="pb-4">
            <Label htmlFor="title">Title</Label>
            <Input
              id="title"
              type="text"
              placeholder="article title"
              required
              value={title}
              onChange={(event) => setTitle(event.target.value)}
            />
          </div>
          <div className="flex pb-4">
            <div className="w-1/2">
              <Label>AI: </Label>
              <Input
                type="text"
                placeholder="Prompt"
                className="p-2"
                value={prompt}
                onChange={(event) => setPrompt(event.target.value)}
              />
            </div>
            <Button
              type="submit"
              className="flex-none"
              onClick={handleGenerateBody}
            >
              Generate!
            </Button>
          </div>
          <div className="flex flex-col flex-grow">
            <Label htmlFor="body">Body</Label>
            <div className="flex-grow">
              <MDEditor value={body} onChange={setBody} height={"100%"} />
            </div>
          </div>
          <div className="flex items-center justify-between p-5">
            <Link to={"/articles"}>
              <Button>Cancel</Button>
            </Link>
            <Button onClick={handleEdit}>Edit</Button>
          </div>
        </div>
      </div>
    </Navigation>
  );
}
