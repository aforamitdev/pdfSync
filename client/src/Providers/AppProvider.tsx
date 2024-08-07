import React, { createContext, useContext } from 'react';

export type AppType = {
  isCollapsed: boolean;
  setIsCollapsed: React.Dispatch<React.SetStateAction<boolean>>;
};

export const AppContext = createContext<AppType | any>({});

type Props = {
  children: React.ReactNode;
};
const AppProvider = ({ children }: Props) => {
  const [isCollapsed, setIsCollapsed] = React.useState(false);

  return (
    <AppContext.Provider value={{ isCollapsed, setIsCollapsed }}>
      {children}
    </AppContext.Provider>
  );
};

export const useAppContext = () => useContext<AppType>(AppContext);

export default AppProvider;
