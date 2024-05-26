import React from "react";
import "./App.css";

import Home from "./pages";
import About from "./pages/about";
import Login from "./pages/login";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { TooltipProvider } from "@radix-ui/react-tooltip";
import ArticleDetail from "./pages/articles/detail";
import CreateArticle from "./pages/articles/create";
import Articles from "./pages/articles/list";
import { Toaster } from "./components/ui/toaster";
import EditArticle from "./pages/articles/edit";
import Media from "./pages/media/list";
import { NewMediaRepository } from "./infrastructure/repository/media";
import { NewMediaUsecase } from "./usecase/media";
import { Configuration, GolioApi } from "./generate/schema/http";
import MediaDetail from "./pages/media/detial";

const App: React.FC = () => {

  // goli apiの設定
  const golioConfig = new Configuration({
    basePath: process.env.REACT_APP_GOLIO_BASE_URL,
    credentials: "include",
  });
  const golioApi = new GolioApi(golioConfig);

  const mediaRepo = NewMediaRepository(golioApi);
  const mediaUsecase = NewMediaUsecase(mediaRepo);

  return (
    <TooltipProvider>
      <BrowserRouter>
        <div className="dark">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/login" element={<Login />} />
            <Route path="/articles" element={<Articles />} />
            <Route path="/media" element={<Media mediaUsecase={mediaUsecase} />} />
            <Route path="/media/:mediaId" element={<MediaDetail mediaUsecase={mediaUsecase} />} />
            <Route path="/articles/:articleId" element={<ArticleDetail />} />
            <Route path="/articles/create" element={<CreateArticle />} />
            <Route path="/articles/edit/:articleId" element={<EditArticle />} />
          </Routes>
        </div>
        <Toaster />
      </BrowserRouter>
    </TooltipProvider>
  );
};

export default App;
