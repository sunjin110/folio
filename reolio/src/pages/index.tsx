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
          <div>
            ここには自分の経歴や、どのようなことをやってきたかを書いていこうと思いマッスル
          </div>
          <div>
            あとはこのポートフォリオの概要なんかを書いていってもいいかもね、各機能への遷移も置いてあげたら優しいかも
          </div>
          <div>変わっているか確認</div>
        </CardContent>
      </Card>
    </Navigation>
  );
};
export default Home;
