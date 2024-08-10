'use client';

import React, { useMemo } from 'react';

import Header from '@/components/Header';
import W3qDeployer from '@/components/W3qDeployer';

import { useGlobalContext } from '../GlobalContext';

const HomeComponent: React.FC = () => {
  const { state } = useGlobalContext();
  const { account, chainConfig } = state;

  const contract = useMemo(() => {
    if (chainConfig && chainConfig.chainID) {
      return chainConfig.FileBoxController;
    }

    return null;
  }, [chainConfig]);

  const flatDirectory = useMemo(() => {
    if (chainConfig && chainConfig.chainID) {
      return chainConfig.FlatDirectory;
    }

    return null;
  }, [chainConfig]);

  return (
    <div className="flex min-h-screen w-full flex-col items-center justify-center bg-[url('/images/upload-background.png')] bg-cover bg-center">
      <div className="h-20 w-full justify-between">
        <Header />
      </div>
      <div className="my-4 flex w-full items-center justify-center text-3xl">Web3 Sounds</div>
      <div className="my-4 flex w-full flex-grow items-center justify-center">
        <W3qDeployer
          multiple
          account={account}
          fileContract={contract}
          flatDirectory={flatDirectory}
        />
      </div>
    </div>
  );
};

export default HomeComponent;
