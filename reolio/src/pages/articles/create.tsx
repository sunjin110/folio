import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@radix-ui/react-label";

export default function CreateArticle() {
    return (
            <Card className="min-h-screen">
                <CardHeader>
                    <CardTitle>Create Article</CardTitle>
                </CardHeader>
                <CardContent>
                    <div>
                        <Label htmlFor="title">Title</Label>
                        <Input id="title" type="text" placeholder="article title" required />
                    </div>
                    <div>
                        <Label htmlFor="body">Body</Label>
                        <Textarea placeholder="article body" required />
                    </div>
                    <div className="flex items-center justify-between p-5">
                        <Button>Cancel</Button>
                        <Button>Post</Button>
                    </div>
                </CardContent>
            </Card>
    );
}