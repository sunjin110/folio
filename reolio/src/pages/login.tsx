import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { FaGoogle as RawFaGoogle } from "react-icons/fa";

const FaGoogle = RawFaGoogle as React.FC;

export default function Login() {
  return (
    <Navigation title="Login" sidebarPosition="">
      <div className="flex items-center justify-center min-h-screen">
        <Card className="max-w-sm">
          <CardHeader>
            <CardTitle className="text-xl">Sign In</CardTitle>
            <CardDescription>
              Enter your information to create an account
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid gap-4">
              <a
                href={
                  process.env.REACT_APP_GOLIO_BASE_URL + "/auth/google-oauth"
                }
              >
                <Button type="submit" className="w-full">
                  <div className="mr-4">
                    <FaGoogle />
                  </div>
                  <div className="">Sign up with Google</div>
                </Button>
              </a>
              <Button variant="outline" className="w-full">
                Sign up with GitHub TODO
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    </Navigation>
  );
}
