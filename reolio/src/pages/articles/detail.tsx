import { getArticleById } from "@/api/api";
import { Navigation } from "@/components/organisms/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useToast } from "@/components/ui/use-toast";
import { Article } from "@/domain/model/article";
import MDEditor from "@uiw/react-md-editor";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

export default function ArticleDetail() {
  const { articleId } = useParams();

  const [article, setArticle] = useState<Article | null>(null);
  const { toast } = useToast();
  const navigate = useNavigate();

  useEffect(() => {
    const fetchArticle = async (id: string) => {
      try {
        const output = await getArticleById(id);
        if (output.type === "error") {
          toast({
            title: "ログインし直してください",
            description: output.message,
          });
          navigate("/login");
          return;
        }
        setArticle(output.article);
      } catch (err) {
        console.error("Error fetcing article:", err);
      }
    };
    if (articleId) {
      fetchArticle(articleId);
    }
  }, [articleId]);

  return (
    <Navigation title="Articles" sidebarPosition="articles">
      <div className="p-2">
        <div className="pb-7">
          <h1 className="text-4xl">{article?.title}</h1>
        </div>
        <div>
          {article ? (
            <MDEditor.Markdown
              source={article.body}
              style={{ backgroundColor: "black", color: "#cbd5e1" }}
            />
          ) : (
            <div>Loading...</div>
          )}
        </div>
      </div>
    </Navigation>
  );
}
