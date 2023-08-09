import Web3 from 'web3';
import util from 'util';

require('util.promisify').shim();
require("dotenv").config();

const initWeb3 = () => {
  window.wss_web3 = new Web3(new Web3.providers.WebsocketProvider('wss://polygon-mumbai.g.alchemy.com/v2/MN3dkqJ6G_Y6sD90x6F-hD93sfsfZLRy'))
  return new Web3(window.ethereum);
}

const web3 = initWeb3()
window.web3 = web3


export const connectWallet = async () => {
  const addressArray = await web3.eth.requestAccounts()
  if (addressArray.length > 0) {
    const obj = {
      status: "👆🏽",
      address: addressArray[0],
    };
    return obj;
  } else {
    return {
      address: "",
      status: "😥 "
    };
  }
};

export const getCurrentWalletConnected = async () => {
  if (web3) {
    try {
      const addressArray = await web3.eth.getAccounts();
      if (addressArray.length > 0) {
        return {
          address: addressArray[0],
          status: "👆🏽",
        };
      } else {
        return {
          address: "",
          status: "🦊 Connect to Metamask using the top right button.",
        };
      }
    } catch (err) {
      console.log(err)
      return {
        address: "",
        status: "😥 " + err.message,
      };
    }
  } else {
    return {
      address: "",
      status: (
        <span>
          <p>
            {" "}
            🦊{" "}
            <a target="_blank" href={`https://metamask.io/download.html`}>
              You must install Metamask, a virtual Ethereum wallet, in your
              browser.
            </a>
          </p>
        </span>
      ),
    };
  }
};