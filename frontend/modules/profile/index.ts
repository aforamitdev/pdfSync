import { createRoot } from 'react-dom/client';
import Profile from './components/Profile';

const dashboardRoot = document.createElement('dashboard-hook');

if (!dashboardRoot) {
  throw new Error('could not find element id react-header');
}

export function index() {
  const rootElement = document.getElementById('dashboard-hook');
  if (!rootElement) {
    throw new Error(`Could not find element with id`);
  }
  const reactRoot = createRoot(rootElement);
  reactRoot.render(Profile());
}
