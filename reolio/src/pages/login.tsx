import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import React from "react";
import { Link } from "react-router-dom";

const Login: React.FC = () => {
    return (
        <Card className="w-[350px]">
            <CardHeader>
                <CardTitle>Login</CardTitle>
            </CardHeader>
            <CardContent>
                <a href="">
                    <Button>Google Login</Button>
                </a>
            </CardContent>
        </Card>
    );
};

export default Login;
