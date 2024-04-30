import { getArticleById, updateArticle } from "@/api/api";
import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { useToast } from "@/components/ui/use-toast";
import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";


export default function EditArticle() {
    const { articleId } = useParams();
    const { toast } = useToast();
    const navigate = useNavigate();

    const [title, setTitle] = useState<string>("");
    const [body, setBody] = useState<string>("");

    useEffect(() => {
        const fetch = async (id: string) => {
            try {
                const output = await getArticleById(id);
                if (output.type === 'error') {
                    toast({title: "please login", description: output.message});
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
    }, [articleId]);

    const handleEdit = async () => {
        try {
            if (!articleId) {
                console.log("articleId is empty");
                return;
            }
            const output = await updateArticle(articleId, title, body);
            if (output.type === 'error') {
                toast({title: 'ログインし直してください', description: output.message});
                navigate("/login");
                return;
            }
            toast({
                title: 'Success',
                description: 'updated your article!'
            });
        } catch (err) {
            console.error("failed cedit article", err);
        }
    };

    return (
        <Navigation title="Articles" sidebarPosition='articles'>
            <Card className="min-h-screen">
                <CardHeader>Edit</CardHeader>
                <CardContent>
                <div>
                    <Label htmlFor="title">Title</Label>
                    <Input id="title" type="text" placeholder="article title" required value={title} onChange={event => setTitle(event.target.value)} />
                </div>
                <div>
                    <Label htmlFor="body">Body</Label>
                    <Textarea placeholder="article body" required value={body} onChange={event => setBody(event.target.value)} />
                </div>
                <div className="flex items-center justify-between p-5">
                    <Link to={"/articles"}>
                        <Button>Canccel</Button>
                    </Link>
                    <Button onClick={handleEdit}>Edit</Button>
                </div>
            </CardContent>
            </Card>
        </Navigation>
    )
}
