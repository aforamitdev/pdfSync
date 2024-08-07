import { ResizablePanel } from '../ui/resizable';
import { cn } from '@/lib/utils';
import { AccountSwitcher } from '../home/components/account-switcher';
import { Separator } from '@radix-ui/react-separator';
import { Nav } from '../home/components/nav';
import {
  AlertCircle,
  Archive,
  ArchiveX,
  File,
  Inbox,
  MessagesSquare,
  Send,
  ShoppingCart,
  Trash2,
  Users2,
} from 'lucide-react';
import { useAppContext } from '@/Providers/AppProvider';

const Navigator = () => {
  const { isCollapsed, setIsCollapsed } = useAppContext();

  const accounts = [
    {
      label: 'Alicia Koch',
      email: 'alicia@example.com',
      icon: (
        <svg role='img' viewBox='0 0 24 24' xmlns='http://www.w3.org/2000/svg'>
          <title>Vercel</title>
          <path d='M24 22.525H0l12-21.05 12 21.05z' fill='currentColor' />
        </svg>
      ),
    },
    {
      label: 'Alicia Koch',
      email: 'alicia@gmail.com',
      icon: (
        <svg role='img' viewBox='0 0 24 24' xmlns='http://www.w3.org/2000/svg'>
          <title>Gmail</title>
          <path
            d='M24 5.457v13.909c0 .904-.732 1.636-1.636 1.636h-3.819V11.73L12 16.64l-6.545-4.91v9.273H1.636A1.636 1.636 0 0 1 0 19.366V5.457c0-2.023 2.309-3.178 3.927-1.964L5.455 4.64 12 9.548l6.545-4.91 1.528-1.145C21.69 2.28 24 3.434 24 5.457z'
            fill='currentColor'
          />
        </svg>
      ),
    },
    {
      label: 'Alicia Koch',
      email: 'alicia@me.com',
      icon: (
        <svg role='img' viewBox='0 0 24 24' xmlns='http://www.w3.org/2000/svg'>
          <title>iCloud</title>
          <path
            d='M13.762 4.29a6.51 6.51 0 0 0-5.669 3.332 3.571 3.571 0 0 0-1.558-.36 3.571 3.571 0 0 0-3.516 3A4.918 4.918 0 0 0 0 14.796a4.918 4.918 0 0 0 4.92 4.914 4.93 4.93 0 0 0 .617-.045h14.42c2.305-.272 4.041-2.258 4.043-4.589v-.009a4.594 4.594 0 0 0-3.727-4.508 6.51 6.51 0 0 0-6.511-6.27z'
            fill='currentColor'
          />
        </svg>
      ),
    },
  ];
  return (
    <ResizablePanel
      defaultSize={20}
      collapsedSize={4}
      collapsible={true}
      minSize={15}
      maxSize={20}
      onCollapse={() => {
        setIsCollapsed(true);
        document.cookie = `react-resizable-panels:collapsed=${JSON.stringify(
          true
        )}`;
      }}
      onResize={() => {
        setIsCollapsed(false);
        document.cookie = `react-resizable-panels:collapsed=${JSON.stringify(
          false
        )}`;
      }}
      className={cn(
        isCollapsed && 'min-w-[50px] transition-all duration-300 ease-in-out'
      )}
    >
      <div
        className={cn(
          'flex h-[56px] items-center justify-center',
          isCollapsed ? 'h-[56px]' : 'px-2'
        )}
      >
        <AccountSwitcher isCollapsed={isCollapsed} accounts={accounts} />
      </div>
      <Separator />
      <Nav
        isCollapsed={isCollapsed}
        links={[
          {
            title: 'Inbox',
            label: '128',
            icon: Inbox,
            variant: 'default',
          },
          {
            title: 'Drafts',
            label: '9',
            icon: File,
            variant: 'ghost',
          },
          {
            title: 'Sent',
            label: '',
            icon: Send,
            variant: 'ghost',
          },
          {
            title: 'Junk',
            label: '23',
            icon: ArchiveX,
            variant: 'ghost',
          },
          {
            title: 'Trash',
            label: '',
            icon: Trash2,
            variant: 'ghost',
          },
          {
            title: 'Archive',
            label: '',
            icon: Archive,
            variant: 'ghost',
          },
        ]}
      />
      <Separator />
      <Nav
        isCollapsed={isCollapsed}
        links={[
          {
            title: 'Social',
            label: '972',
            icon: Users2,
            variant: 'ghost',
          },
          {
            title: 'Updates',
            label: '342',
            icon: AlertCircle,
            variant: 'ghost',
          },
          {
            title: 'Forums',
            label: '128',
            icon: MessagesSquare,
            variant: 'ghost',
          },
          {
            title: 'Shopping',
            label: '8',
            icon: ShoppingCart,
            variant: 'ghost',
          },
          {
            title: 'Promotions',
            label: '21',
            icon: Archive,
            variant: 'ghost',
          },
        ]}
      />
    </ResizablePanel>
  );
};

export default Navigator;
