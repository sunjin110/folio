import { createArticle } from "@/api/api";
import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { useToast } from "@/components/ui/use-toast";
import { Label } from "@radix-ui/react-label";
import { MouseEventHandler, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function CreateArticle() {
    const [title, setTitle] = useState("");
    const [body, setBody] = useState("");

    const { toast } = useToast();
    const navigate = useNavigate();

    const handlePost = async () => {
        try {
            const output = await createArticle(title, body);
            if (output.type === 'error') {
                toast({title: "ログインし直してください", description: output.message});
                navigate("/login");
                return;
            }
            toast({
                title: 'Success',
                description: 'posted your article!'
            });
            navigate("/articles");
        } catch (err) {
            console.error("failed create article", err);
        }
    };

    return (
            <Navigation title="Articles" sidebarPosition='articles'>
                <Card className="min-h-screen">
                    <CardHeader>
                        <CardTitle>Create Article</CardTitle>
                    </CardHeader>
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
                                <Button>Cancel</Button>
                            </Link>
                            <Button onClick={handlePost}>Post</Button>
                        </div>
                    </CardContent>
                </Card>
            </Navigation>
    );
}
