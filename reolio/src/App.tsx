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
import { NewArticleRepository } from "./infrastructure/repository/article";
import { NewArticleUsecase } from "./usecase/article";
import { TranslateModalPage } from "./pages/modals/translateModalPage";
import { NewTranslationRepository } from "./infrastructure/repository/translation";
import { NewArticleTagRepository } from "./infrastructure/repository/article_tag";
import EnglishDictionary from "./pages/english_dictionary";
import { NewEnglishDictionaryUsecase } from "./usecase/english_dictionary";
import { NewEnglishDictionaryRepository } from "./infrastructure/repository/english_dictionary";
import JsonFormat from "./pages/tools/jsonFormatter";
import Time from "./pages/tools/time";

const App: React.FC = () => {
  // goli apiの設定
  const golioConfig = new Configuration({
    basePath: process.env.REACT_APP_GOLIO_BASE_URL,
    credentials: "include",
  });
  const golioApi = new GolioApi(golioConfig);

  const mediaRepo = NewMediaRepository(golioApi);
  const articleRepo = NewArticleRepository(golioApi);
  const articleTagRepo = NewArticleTagRepository(golioApi);
  const translationRepo = NewTranslationRepository(golioApi);
  const englishDictionaryRepo = NewEnglishDictionaryRepository(golioApi);
  const mediaUsecase = NewMediaUsecase(mediaRepo);
  const articleUsecase = NewArticleUsecase(articleRepo, articleTagRepo);
  const englishDctionaryUsecase = NewEnglishDictionaryUsecase(
    englishDictionaryRepo,
  );

  return (
    <TooltipProvider>
      <TranslateModalPage translationRepository={translationRepo} />
      <BrowserRouter>
        <div className="dark">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/login" element={<Login />} />
            <Route
              path="/articles"
              element={<Articles articleUsecase={articleUsecase} />}
            />
            <Route
              path="/media"
              element={<Media mediaUsecase={mediaUsecase} />}
            />
            <Route
              path="/media/:mediaId"
              element={<MediaDetail mediaUsecase={mediaUsecase} />}
            />
            <Route path="/articles/:articleId" element={<ArticleDetail />} />
            <Route
              path="/articles/create"
              element={<CreateArticle articleUsecase={articleUsecase} />}
            />
            <Route
              path="/articles/edit/:articleId"
              element={<EditArticle articleUsecase={articleUsecase} />}
            />
            <Route
              path="/english_dictionary"
              element={
                <EnglishDictionary
                  englishDctionaryUsecase={englishDctionaryUsecase}
                />
              }
            />
            <Route path="/tools/json" element={<JsonFormat />} />
            <Route path="/tools/time" element={<Time />} />
          </Routes>
        </div>
        <Toaster />
      </BrowserRouter>
    </TooltipProvider>
  );
};

export default App;
