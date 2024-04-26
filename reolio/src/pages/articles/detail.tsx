import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useParams } from "react-router-dom";

export default function ArticleDetail() {
    const { articleId } = useParams();
    return (
        // <div>
        //     <h1>ArticleDetail Page</h1>
        //     <p>ArticleID is : {articleId}</p>
        // </div>
        <Card>
            <CardHeader>
                <CardTitle>ArticleDetail Page</CardTitle>
            </CardHeader>
            <CardContent>
                ArticleID is {articleId}
            </CardContent>
        </Card>
    );
}