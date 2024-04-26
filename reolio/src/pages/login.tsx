import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import React from "react";
const Login: React.FC = () => {
    return (
        <Card className="w-[350px]">
            <CardHeader>
                <CardTitle>Login</CardTitle>
            </CardHeader>
            <CardContent>
                <a href={process.env.REACT_APP_GOLIO_BASE_URL + "/auth/google-oauth"}>
                    <Button>Google Login</Button>
                </a>
            </CardContent>
        </Card>
    );
};

export default Login;
