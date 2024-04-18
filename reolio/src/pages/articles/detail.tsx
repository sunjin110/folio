import {  getArticleById } from "@/api/api";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Article } from "@/domain/model/article";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

export default function ArticleDetail() {
    const { articleId } = useParams();

    const [article, setArticle] = useState<Article | null>(null);

    useEffect(() => {
        const fetchArticle = async (id: string) => {
            try {
                setArticle(await getArticleById(id));
            } catch (err) {
                console.error("Error fetcing article:", err);
            }
        };
        if (articleId) {
            fetchArticle(articleId);
        }
    }, [articleId]);


    return (
        <Card>
            <CardHeader>
                <CardTitle>ArticleDetail Page</CardTitle>
            </CardHeader>
            <CardContent>
                ArticleID is {article?.id}, {article?.title}, {article?.body}, {article?.created_at}
            </CardContent>
        </Card>
    );
}
