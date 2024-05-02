import React from 'react';
import { createRoot } from 'react-dom/client';
import Dashboard from './components/Dashboard';

const dashboardRoot = document.getElementById('dashboard');

if (!dashboardRoot) {
  throw new Error('could not find element id react-header');
}
const reactRoot = createRoot(dashboardRoot);

console.log('dashboard');
reactRoot.render(
  <React.StrictMode>
    <Dashboard />
  </React.StrictMode>
);
