import { createArticle } from "@/api/api";
import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useToast } from "@/components/ui/use-toast";
import { Label } from "@radix-ui/react-label";
import MDEditor from "@uiw/react-md-editor";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function CreateArticle() {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState<string | undefined>("");

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
      navigate("/articles");
    } catch (err) {
      console.error("failed create article", err);
    }
  };

  return (
    <Navigation title="Articles" sidebarPosition="articles">
      <div className="flex flex-col h-full">
        <div className="pb-7">
          <h1 className="text-4xl">Create</h1>
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
