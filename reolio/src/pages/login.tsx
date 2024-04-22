import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
// import { generateOAuth2Url } from "@/lib/google/oauth";
import React from "react";
import { Link } from "react-router-dom";

const Login: React.FC = () => {

    // const oauth2Url = generateOAuth2Url();
    // console.log("oauth2url is ", oauth2Url);


    return (
        <Card className="w-[350px]">
            <CardHeader>
                <CardTitle>Login</CardTitle>
            </CardHeader>
            <CardContent>
                <a href={"/test"}>
                    <Button>Google Login</Button>
                </a>
            </CardContent>
        </Card>
    );
};

export default Login;
