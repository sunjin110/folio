import { Navigation } from "@/components/organisms/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import React from "react";
import { Link } from "react-router-dom";

const Home: React.FC = () => {
  return (
    <Navigation title="title" sidebarPosition="">
      <Card>
        <CardHeader>
          <CardTitle className="text-xl">Index</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="p-5">
            unix time: {Math.ceil(Number(new Date().getTime() / 1000))}
          </div>

          <div className="p-5">
            <Link to={"/tools/json"}>
              <Button>Json Formatter</Button>
            </Link>
          </div>
        </CardContent>
      </Card>
    </Navigation>
  );
};
export default Home;
