import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import Router from './pages/index.tsx';
import AppProvider from './Providers/AppProvider.tsx';
import { TooltipProvider } from './components/ui/tooltip.tsx';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <div className='font-intra '>
      <TooltipProvider delayDuration={0}>
        <AppProvider>
          <Router />
        </AppProvider>
      </TooltipProvider>
    </div>
  </React.StrictMode>
);
