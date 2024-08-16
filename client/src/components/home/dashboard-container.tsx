import { accounts, mails } from './data';
import { Mail } from './components/mail';
import { Routes, Route } from 'react-router-dom';
import Navigator from '../Navigator/Navigator';
import { ResizablePanelGroup } from '../ui/resizable';
import Reader from '../reader/Reader';
import FilesPanel from '../FilesPanel/FilesPanel';
// import { Route } from 'lucide-react';

export default function DashboardContainer() {
  return (
    <>
      <div className='hidden flex-col md:flex'>
        <ResizablePanelGroup
          direction='horizontal'
          onLayout={(sizes: number[]) => {
            document.cookie = `react-resizable-panels:layout:mail=${JSON.stringify(
              sizes
            )}`;
          }}
          className='h-full max-h-screenitems-stretch'
        >
          <Navigator />
          <Routes>
            <Route
              path='/'
              element={
                <Mail accounts={accounts} mails={mails} navCollapsedSize={4} />
              }
            />
            <Route
              path='/shelf'
              element={
                <FilesPanel />
              }
            />
            <Route path="/reader" element={<Reader />}></Route>
          </Routes>
        </ResizablePanelGroup>
      </div>
    </>
  );
}
