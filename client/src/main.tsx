import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import Router from './pages/index.tsx';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
      <div className='font-intra '>
        <Router />
      </div>
  </React.StrictMode>
);
