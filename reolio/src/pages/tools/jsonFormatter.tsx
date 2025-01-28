import { Navigation } from "@/components/organisms/navigation";
import { Textarea } from "@/components/ui/textarea";
import React, { useState } from "react";
import JSONPretty from "react-json-pretty";
import "react-json-pretty/themes/monikai.css";

export default function JsonFormat() {
  const [data, setData] = useState("");
  return (
    <Navigation title="JsonFormat" sidebarPosition="">
      <div className="p-2">
        <div className="pb-7">
          <h1 className="text-4xl">Json Format</h1>
        </div>

        <div className="flex w-full">
          <Textarea
            value={data}
            onChange={(event) => setData(event.target.value)}
          />
        </div>
        <div className="pt-2">
          <JSONPretty data={data} />
        </div>
      </div>
    </Navigation>
  );
}
