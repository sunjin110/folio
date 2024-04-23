import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { generateOAuth2Url } from "@/lib/google/oauth";
import React from "react";
const Login: React.FC = () => {

    const oauth2Url = generateOAuth2Url();
    console.log("oauth2url is ", oauth2Url);
    console.log("process.env", process.env);


    return (
        <Card className="w-[350px]">
            <CardHeader>
                <CardTitle>Login</CardTitle>
            </CardHeader>
            <CardContent>
                <a href={oauth2Url}>
                    <Button>Google Login</Button>
                </a>
            </CardContent>
        </Card>
    );
};

export default Login;
