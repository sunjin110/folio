import { ArticleTagEditProps } from "@/components/organisms/articleTagEdit";
import CreateArticleTemplate from "@/components/templates/articles/create";
import { useToast } from "@/components/ui/use-toast";
import { ArticleTag } from "@/domain/model/article";
import { handleError } from "@/error/pageErrorHandle";
import { useDelayState } from "@/hooks/useDelayState";
import { ArticleUsecase } from "@/usecase/article";
import { useEffect, useState } from "react";
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

  const handlePost = async () => {
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

    try {
      const articleId = await articleUsecase.InsertArticle(
        title,
        body === undefined ? "" : body,
        tagIDs,
      );
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
    <CreateArticleTemplate
      aiPrompt={aiPrompt}
      setAiPrompt={setAiPrompt}
      title={title}
      setTitle={setTitle}
      body={body}
      setBody={setBody}
      handlePost={handlePost}
      handleGenerateAI={handleGenerateAI}
      articleTagEditProps={articleTagEditProps}
    />
  );
}
