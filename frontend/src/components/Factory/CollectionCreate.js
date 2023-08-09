import { useEffect, useState } from "react";

import collectionFactoryABI from '../../abi/CollectionFactory.json'
import collectionABI from '../../abi/Collection.json'
import { address }  from '../../abi/factory_contract.json'

import Item from './Item'

const CollectionCreate = (props) => {
  const [collectionName, setCollectionName] = useState('')
  const [collectionSymbol, setCollectionSymbol] = useState('')
  const [collectionAddress, setCollectionAddress] = useState('')
  const [statusMessage, setStatusMessage] = useState('')
  const [mintStatusMessage, setMintStatusMessage] = useState('')

  const [newTokenId, setNewTokenId] = useState(0)
  const [newTokenURI, setNewTokenURI] = useState('https://v5.airtableusercontent.com/v1/19/19/1691568000000/TVkQeKjTFLLedKQxcuXdRw/ytLaijrhZzHMP5g-XvQP7Bdfdq1cRYHZI9uloTHBTLHjkYr--gzRbEI8mWHlfQW8Ud79qnt4b299f6G_-cnbUA/jj35dRsl0dY0yLQXI0V-2qdrValj84RgyehlaXDV5tc')

  const [mintedTokenId, setMintedTokenId] = useState(null)

  useEffect(async () => {
    const contract = await new window.wss_web3.eth.Contract(collectionFactoryABI.abi, address);
    contract.events.CollectionCreated().on('data', function(event){
      console.log('CollectionCreated', event); // same results as the optional callback above
      setStatusMessage('NFT TokenMinted!')
    })
  }, [])

  useEffect(async () => {
    if(collectionAddress) {
      const contract = await new window.wss_web3.eth.Contract(collectionABI.abi, collectionAddress);
      contract.events.TokenMinted().off()
      contract.events.TokenMinted().on('data', function(event){
        console.log('TokenMinted', event); // same results as the optional callback above
        setMintStatusMessage('NFT TokenMinted!')
      })
    }
  }, [collectionAddress])


  const createCollection = async () => {
    setStatusMessage('Start NFT collection create...')
    const contract = await new window.web3.eth.Contract(collectionFactoryABI.abi, address);

    try {
      const tx = await contract.methods.createCollection(collectionName, collectionSymbol).send({ from: props.walletAddress })
      console.log('txData', tx)
      console.log('get event form tx', tx.events.CollectionCreated.returnValues)

      setCollectionAddress(tx.events.CollectionCreated.returnValues.collectionAddress)

      setStatusMessage("✅ Check out your transaction on Etherscan: https://mumbai.polygonscan.com/tx/" + tx.transactionHash)
    } catch (err) {
      setStatusMessage("😥 Something went wrong: " + err.message)
    }
  }

  const mintItem = async () => {
    setMintStatusMessage('Start mint new NFT in collection...')
    const contract = await new window.web3.eth.Contract(collectionABI.abi, collectionAddress);

    try {
      const tx = await contract.methods.mint(newTokenId, newTokenURI).send({ from: props.walletAddress })
      console.log('txData', tx.events.TokenMinted.returnValues)
      setMintedTokenId(tx.events.TokenMinted.returnValues.tokenId)
      setMintStatusMessage("✅ Check out your transaction on Etherscan: https://mumbai.polygonscan.com/tx/" + tx.transactionHash)
    } catch (err) {
      setMintStatusMessage("😥 Something went wrong: " + err.message)
    }
  }

  const switchChain = () => {
    window.ethereum.request({
      method: 'wallet_addEthereumChain',
      params: [{
      chainId: '0x38',
      chainName: 'Binance Smart Chain',
      nativeCurrency: {
          name: 'Binance Coin',
          symbol: 'BNB',
          decimals: 18
      },
      rpcUrls: ['https://bsc-dataseed.binance.org/'],
      blockExplorerUrls: ['https://bscscan.com']
      }]
      })
      .catch((error) => {
      console.log(error)
      })
  }

  return (
    <div>
      <h4>Create new collection</h4>

      <h4>Collection name</h4>
      <input value={collectionName} onChange={(e) => setCollectionName(e.target.value)} placeholder="Crypto crips" />

      <h4>Collection symbol</h4>
      <input value={collectionSymbol} onChange={(e) => setCollectionSymbol(e.target.value)} placeholder="CCRS" />

      <p style={{color: 'green'}}>{statusMessage}</p>

      <button style={{marginBottom: '14px', marginTop: '5px'}} onClick={createCollection}>
        Create!
      </button>

      <h4>Collection address</h4>
      <input value={collectionAddress} onChange={(e) => setCollectionAddress(e.target.value)} placeholder="0x00000" />

      <h4>TokenId</h4>
      <input value={newTokenId} onChange={(e) => setNewTokenId(e.target.value)} placeholder="0" />

      <h4>token URI</h4>
      <input value={newTokenURI} onChange={(e) => setNewTokenURI(e.target.value)} placeholder="https://example.com/nft/" />
      <p style={{color: 'green'}}>{mintStatusMessage}</p>
      <button onClick={mintItem} style={{marginBottom: '14px', marginTop: '5px'}}>Mint collection item!</button>
      <br />

      {mintedTokenId && <Item collectionAddress={collectionAddress} tokenId={mintedTokenId} />}

      <h4>Settings</h4>
      <button onClick={switchChain} style={{marginBottom: '14px', marginTop: '5px'}}>Add Mumbai network</button>
    </div>
  )
}

export default CollectionCreate