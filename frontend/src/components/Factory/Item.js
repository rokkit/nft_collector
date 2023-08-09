import { useEffect, useState } from "react";

import collectionABI from '../../abi/Collection.json'

const Item = (props) => {
  const collectionAddress = props.collectionAddress
  const tokenId = props.tokenId
  const [tokenURI, setTokenURI] = useState('')

  useEffect(async () => {
    if(collectionAddress) {
      const c = await new window.web3.eth.Contract(collectionABI.abi, collectionAddress);
      const txData = await c.methods.tokenURI(tokenId).call()
      setTokenURI(txData)
    }
  }, [collectionAddress])

  return (
    <div>
      <img src={tokenURI} width={500} />
    </div>
  )
}

export default Item