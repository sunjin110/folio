import {
  ArticleTagEdit,
  ArticleTagEditProps,
} from "@/components/organisms/articleTagEdit";
import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import MDEditor from "@uiw/react-md-editor";
import { BrainCircuit } from "lucide-react";
import { Dispatch, SetStateAction } from "react";
import { Link } from "react-router-dom";

interface CreateArticleProps {
  aiPrompt: string;
  setAiPrompt: Dispatch<SetStateAction<string>>;
  title: string;
  setTitle: Dispatch<SetStateAction<string>>;
  body: string | undefined;
  setBody: Dispatch<SetStateAction<string | undefined>>;
  handlePost: () => Promise<void>;
  handleGenerateAI: () => Promise<void>;
  articleTagEditProps: ArticleTagEditProps;
}

export default function CreateArticleTemplate(props: CreateArticleProps) {
  const {
    title,
    setTitle,
    body,
    setBody,
    handlePost,
    aiPrompt,
    setAiPrompt,
    handleGenerateAI,
    articleTagEditProps,
  } = props;

  return (
    <Navigation title="Articles" sidebarPosition="articles">
      <div className="flex flex-col h-full">
        <div className="pb-7">
          <h1 className="text-4xl">Create</h1>
        </div>
        <div className="flex flex-col flex-grow">
          <div className="pb-2">
            <div className="flex w-full items-center space-x-2">
              <Label htmlFor="prompt">
                <BrainCircuit />
              </Label>
              <Input
                type="text"
                placeholder="ai prompt"
                id="prompt"
                value={aiPrompt}
                onChange={(event) => setAiPrompt(event.target.value)}
              />
              <Button type="button" onClick={handleGenerateAI}>
                🧠 Generate !
              </Button>
            </div>
          </div>
          <div className="pb-2">
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
            <MDEditor value={body} onChange={setBody} height={"100%"} />
          </div>
          <div className="flex items-center justify-between p-5">
            <Link to={"/articles"}>
              <Button>Cancel</Button>
            </Link>
            <Button onClick={handlePost}>Post</Button>
          </div>
        </div>
      </div>
    </Navigation>
  );
}
