import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { FaGoogle } from "react-icons/fa";

export default function Login() {
  return (
    <div className="flex items-center justify-center min-h-screen">
        <Card className="max-w-sm">
        <CardHeader>
            <CardTitle className="text-xl">Sign Up</CardTitle>
            <CardDescription>
            Enter your information to create an account
            </CardDescription>
        </CardHeader>
        <CardContent>
            <div className="grid gap-4">
            <a href={process.env.REACT_APP_GOLIO_BASE_URL + "/auth/google-oauth"}>
                <Button type="submit" className="w-full">
                    <div className="mr-4">
                        <FaGoogle /> 
                    </div>
                    <div className="">
                        Sign up with Google
                    </div>
                </Button>
            </a>
            <Button variant="outline" className="w-full">
                Sign up with GitHub TODO
            </Button>
            </div>
            <div className="mt-4 text-center text-sm">
            Already have an account?{" "}
            <a href="#" className="underline">
                Sign in
            </a>
            </div>
        </CardContent>
        </Card>
    </div>
  )
}
