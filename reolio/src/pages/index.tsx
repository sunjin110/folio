import { Navigation } from "@/components/organisms/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import React from "react";

const Home: React.FC = () => {
  return (
    <Navigation title="title" sidebarPosition="">
      <Card>
        <CardHeader>
          <CardTitle className="text-xl">Index</CardTitle>
        </CardHeader>
        <CardContent>
          <div>unix time: {Math.ceil(Number(new Date().getTime() / 1000))}</div>
        </CardContent>
      </Card>
    </Navigation>
  );
};
export default Home;
