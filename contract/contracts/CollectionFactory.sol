//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "./Collection.sol";

contract CollectionFactory {
  mapping(address => Collection) _collections;

  event CollectionCreated(
      address indexed collectionAddress,
      string indexed name,
      string indexed symbol
  );

  function createCollection(string calldata name, string calldata symbol)
  public
  returns (address) {
    Collection collection = new Collection(name, symbol, msg.sender);
    _collections[msg.sender] = collection;

    emit CollectionCreated(address(collection), name, symbol);
    return address(collection);
  }
}