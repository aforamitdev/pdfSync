import { ResizablePanel } from '../ui/resizable';
import { cn } from '@/lib/utils';
import { AccountSwitcher } from '../home/components/account-switcher';
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
  Container,
  Landmark
} from 'lucide-react';
import { useAppContext } from '@/Providers/AppProvider';
import { Separator } from '../ui/separator';
import { Link, useNavigate } from "react-router-dom";

const Navigator = () => {
  const { isCollapsed, setIsCollapsed } = useAppContext();
  const navigate = useNavigate();

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
    }
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
      {/* <div onClick={()=>navigate({pathname:"/asas"})}>click</div> */}
      <div
        className={cn(
          'flex h-[56px] items-center justify-center',
          isCollapsed ? 'h-[56px]' : 'px-2'
        )}
      >
        {JSON.stringify(isCollapsed)}
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
            link: "/"

          },

          {
            title: 'Store',
            label: '',
            icon: Landmark,
            variant: 'ghost',
            link: "/shelf"
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
            link: "/social"
          }

        ]}
      />
    </ResizablePanel>
  );
};

export default Navigator;
