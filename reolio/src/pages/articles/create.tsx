import { createArticle } from "@/api/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@radix-ui/react-label";
import { MouseEventHandler, useState } from "react";

export default function CreateArticle() {
    const [title, setTitle] = useState("");
    const [body, setBody] = useState("");

    const handlePost = async () => {
        console.log("start create post", title, body);
        try {
            await createArticle(title, body);
            console.log("success");
        } catch (err) {
            console.error("failed create article", err);
        }
    };

    return (
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
                        <Button>Cancel</Button>
                        <Button onClick={handlePost}>Post</Button>
                    </div>
                </CardContent>
            </Card>
    );
}
