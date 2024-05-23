import { getArticleById, updateArticle } from "@/api/api";
import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/components/ui/use-toast";
import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import MDEditor from "@uiw/react-md-editor";
import { getRandomEmoji } from "@/domain/service/joke";

export default function EditArticle() {
  const { articleId } = useParams();
  const { toast } = useToast();
  const navigate = useNavigate();

  const [title, setTitle] = useState<string>("");
  const [body, setBody] = useState<string | undefined>("");

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

  const handleEdit = async () => {
    try {
      if (!articleId) {
        console.log("articleId is empty");
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
  };

  return (
    <Navigation title="Articles" sidebarPosition="articles">
      <div className="flex flex-col h-full p-2">
        <div className="pb-7">
          <h1 className="text-4xl">Edit</h1>
        </div>
        <div className="flex flex-col flex-grow">
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
