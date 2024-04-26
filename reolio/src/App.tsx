import React from 'react';
import './App.css';

import Home from './pages';
import About from './pages/about';
import Login from './pages/login';
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom';
import Dashboard from './pages/dashboard';
import { TooltipProvider } from '@radix-ui/react-tooltip';

const App: React.FC = () => {
  return (
    <TooltipProvider>
      <BrowserRouter>
        <div className='dark'>
        {/* <nav>
            <ul>
              <li>
                <Link to="/">Home</Link>
              </li>
              <li>
                <Link to="/about">About</Link>
              </li>
              <li>
                <Link to="/login">Login</Link>
              </li>
            </ul>
          </nav> */}
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path='/login' element={<Login />} />
            <Route path='/dashboard' element={<Dashboard />} />
          </Routes>
        </div>
      </BrowserRouter>
    </TooltipProvider>
  );
};

export default App;
