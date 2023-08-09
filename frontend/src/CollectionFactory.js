import { useEffect, useState } from "react";
import Wallet from './Wallet'
import axios from 'axios';
import CollectionCreate from './components/Factory/CollectionCreate'

const CollectionFactory = (props) => {
  const [walletAddress, setWallet] = useState("");
  const [status, setStatus] = useState("");

  return (
    <div className="Minter">
      <Wallet onWalletChange={(address) => { setWallet(address) }} />
      <br></br>
      <h1 id="title">🤩 Create a NFT Collection</h1>
      <p id="status" style={{ color: "red" }}>
        {status}
      </p>

      <CollectionCreate walletAddress={walletAddress} />

    </div>
  );
};

export default CollectionFactory;
