import { createRoot } from 'react-dom/client';
import Dashboard from './components/Dashboard';

const dashboardRoot = document.createElement('dashboard');

if (!dashboardRoot) {
  throw new Error('could not find element id react-header');
}

export function index() {
  const rootElement = document.getElementById('dashboard');
  if (!rootElement) {
    throw new Error(`Could not find element with id`);
  }
  const reactRoot = createRoot(rootElement);
  reactRoot.render(Dashboard());
}
