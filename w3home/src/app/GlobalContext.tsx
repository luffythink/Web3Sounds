import type { ReactNode } from 'react';

import React, { createContext, useContext, useReducer } from 'react';

interface State {
  chainConfig: any;
  account: string;
}

type Action = { type: 'SET_CHAIN_CONFIG'; payload: any } | { type: 'SET_ACCOUNT'; payload: string };

const initialState: State = {
  account: '',
  chainConfig: null
};

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case 'SET_CHAIN_CONFIG':
      return { ...state, chainConfig: action.payload };
    case 'SET_ACCOUNT':
      return { ...state, account: action.payload };
    default:
      return state;
  }
}

const GlobalContext = createContext<{ state: State; dispatch: React.Dispatch<Action> }>({
  dispatch: () => null,
  state: initialState
});

export const GlobalProvider = ({ children }: { children: ReactNode }) => {
  const [state, dispatch] = useReducer(reducer, initialState);

  return <GlobalContext.Provider value={{ dispatch, state }}>{children}</GlobalContext.Provider>;
};

export const useGlobalContext = () => useContext(GlobalContext);
