//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract Collection is ERC721URIStorage, Ownable {
  address private _factory;
  address private _creator;

  event TokenMinted(
    address indexed collectionAddress,
    address indexed recipient,
    uint256 indexed tokenId,
    string tokenUri
  );

  event TokenMintFailed(
    uint256 indexed tokenId,
    string tokenUri
  );

  modifier onlyFactory(address caller) {
    require(caller == _factory, "Access to contract denied. Contract created in another factory");
    _;
  }

  constructor(
    string memory collectionName,
    string memory collectionSymbol,
    address creator
  )
    ERC721(collectionName, collectionSymbol) {
    _factory = msg.sender;
    _creator = creator;

    transferOwnership(creator);
  }

  // function _baseURI() internal view override returns (string memory) {
  //   return _baseMetadataURI;
  // }

  function mint(uint256 newTokenId, string calldata newTokenURI) public {
    _mint(msg.sender, newTokenId);
    _setTokenURI(newTokenId, newTokenURI);

    emit TokenMinted(address(this), _creator, newTokenId, newTokenURI);
  }
}