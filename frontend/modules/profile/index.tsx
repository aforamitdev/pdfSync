import React from 'react';
import { createRoot } from 'react-dom/client';
import Profile from './components/Profile';

const rootElement = document.getElementById('profile');

if (!rootElement) {
  throw new Error(`Could not find element with id`);
}

const reactRoot = createRoot(rootElement);

reactRoot.render(
  <React.StrictMode>
    <Profile />
  </React.StrictMode>
);
