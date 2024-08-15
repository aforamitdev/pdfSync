import React from 'react';
import { pdfjs } from "react-pdf"
import ReactDOM from 'react-dom/client';
import './index.css';
import Router from './pages/index.tsx';
import AppProvider from './Providers/AppProvider.tsx';
import { TooltipProvider } from './components/ui/tooltip.tsx';
import { Theme, ThemePanel } from '@radix-ui/themes';

pdfjs.GlobalWorkerOptions.workerSrc = `//unpkg.com/pdfjs-dist@${pdfjs.version}/legacy/build/pdf.worker.min.mjs`;
import 'react-pdf/dist/Page/TextLayer.css';
import 'react-pdf/dist/Page/AnnotationLayer.css';
import '@radix-ui/themes/styles.css';



ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <div className='font-intra '>
      <Theme accentColor="crimson">
        <TooltipProvider delayDuration={0}>
          <AppProvider>
            <Router />
          </AppProvider>
        </TooltipProvider>
        <ThemePanel />

      </Theme>
    </div>
  </React.StrictMode>
);
