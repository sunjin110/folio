import React from 'react';
import './App.css';

import Home from './pages';
import About from './pages/about';
import Login from './pages/login';
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom';
import Dashboard from './pages/dashboard';
import { TooltipProvider } from '@radix-ui/react-tooltip';
import ArticleDetail from './pages/articles/detail';
import CreateArticle from './pages/articles/create';
import Articles from './pages/articles/list';
import { Toaster } from './components/ui/toaster';

const App: React.FC = () => {
  return (
    <TooltipProvider>
      <BrowserRouter>
        <div className='dark'>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path='/login' element={<Login />} />
            <Route path='/dashboard' element={<Dashboard />} />
            <Route path='/articles' element={<Articles />} />
            <Route path='/articles/:articleId' element={<ArticleDetail />} />
            <Route path='/articles/create' element={<CreateArticle />} />
          </Routes>
        </div>
        <Toaster />
      </BrowserRouter>
    </TooltipProvider>
  );
};

export default App;
