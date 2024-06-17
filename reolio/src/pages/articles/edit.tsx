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
import { handleError } from "@/error/pageErrorHandle";
import {
  ArticleTagEdit,
  ArticleTagEditProps,
} from "@/components/organisms/articleTagEdit";
import { ArticleTag } from "@/domain/model/article";
import { useDelayState } from "@/hooks/useDelayState";

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

  const [candidateTagMap, setCandidateTagMap] = useState<
    Map<string, ArticleTag>
  >(new Map());

  const [selectedTagMap, setSelectedTagMap] = useState<Map<string, ArticleTag>>(
    new Map(),
  );
  const [selectedNewTagNameMap, setSelectedNewTagNameMap] = useState<
    Map<string, boolean>
  >(new Map());

  const [tagSearchText, setTagSearchText] = useState("");

  const [deplayTagSearchText] = useDelayState(tagSearchText, 700);

  const articleTagEditProps: ArticleTagEditProps = {
    candidateTagMap,
    searchText: tagSearchText,
    setSearchText: setTagSearchText,
    selectedTagMap,
    setSelectedTagMap,
    selectedNewTagNameMap,
    setSelectedNewTagNameMap,
  };

  useEffect(() => {
    const fetch = async (id: string) => {
      try {
        const article = await articleUsecase.Get(id);
        setTitle(article.title);
        setBody(article.body);
        setSelectedTagMap((prevMap) => {
          const newMap = new Map(prevMap);
          for (let tag of article.tags) {
            newMap.set(tag.id, tag);
          }
          return newMap;
        });
      } catch (err) {
        const resp = handleError(err);
        toast(resp.toast);
        if (resp.navigationPath) {
          navigate(resp.navigationPath);
        }
      }
    };

    if (articleId) {
      fetch(articleId);
    }
  }, [articleId, navigate, toast, articleUsecase]);

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

      // 先に存在しないtagを追加する
      const tagIDs: string[] = [];
      try {
        const itr = selectedNewTagNameMap.entries();
        while (true) {
          const e = itr.next();
          if (e.done) {
            break;
          }
          const [tagName, isSelected] = e.value;
          if (!isSelected) {
            continue;
          }
          const tagID = await articleUsecase.InsertTag(tagName);
          tagIDs.push(tagID);
        }
      } catch (err) {
        const resp = handleError(err);
        toast(resp.toast);
        if (resp.navigationPath) {
          navigate(resp.navigationPath);
        }
      }

      Array.from(selectedTagMap).forEach((keyValue) => {
        tagIDs.push(keyValue[1].id);
      });

      await articleUsecase.UpdateArticle(
        articleId,
        title,
        body === undefined ? "" : body,
        tagIDs,
      );

      const emoji = getRandomEmoji();
      toast({
        title: `${emoji} Success ${emoji}`,
        description: "updated your article!",
      });
    } catch (err) {
      const resp = handleError(err);
      toast(resp.toast);
      if (resp.navigationPath) {
        navigate(resp.navigationPath);
      }
    }
  }, [
    articleId,
    title,
    body,
    toast,
    navigate,
    articleUsecase,
    selectedNewTagNameMap,
    selectedTagMap,
  ]);

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

  useEffect(() => {
    if (!deplayTagSearchText) {
      setCandidateTagMap(new Map());
      return;
    }

    const fetch = async (tagSearchText: string) => {
      try {
        const articleTags = await articleUsecase.FindTags(tagSearchText, 0, 10);

        const newCandidateTagMap = new Map<string, ArticleTag>();
        for (let articleTag of articleTags) {
          newCandidateTagMap.set(articleTag.id, articleTag);
        }
        setCandidateTagMap(newCandidateTagMap);
      } catch (err) {
        const resp = handleError(err);
        toast(resp.toast);
        if (resp.navigationPath) {
          navigate(resp.navigationPath);
        }
      }
    };

    fetch(deplayTagSearchText);
  }, [
    deplayTagSearchText,
    toast,
    navigate,
    articleUsecase,
    setCandidateTagMap,
  ]);

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
          <div className="pb-2">
            <Label htmlFor="">Tags</Label>
            <ArticleTagEdit
              candidateTagMap={articleTagEditProps.candidateTagMap}
              searchText={articleTagEditProps.searchText}
              setSearchText={articleTagEditProps.setSearchText}
              selectedTagMap={articleTagEditProps.selectedTagMap}
              setSelectedTagMap={articleTagEditProps.setSelectedTagMap}
              selectedNewTagNameMap={articleTagEditProps.selectedNewTagNameMap}
              setSelectedNewTagNameMap={
                articleTagEditProps.setSelectedNewTagNameMap
              }
            />
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
