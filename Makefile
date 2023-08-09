.PHONY: deploy
deploy:
	cd contract && yarn hardhat run scripts/deploy_collection_factory.ts --network mumbai | grep "deployed to" | sed 's/CollectionFactory deployed to: //' | tee factory_contract.txt | mv factory_contract.txt ..
	cat factory_contract.txt | xargs printf '{"address":"'"%s"'"}' | tee frontend/src/abi/factory_contract.json
	cp contract/artifacts/contracts/CollectionFactory.sol/CollectionFactory.json frontend/src/abi/CollectionFactory.json
	cp contract/artifacts/contracts/Collection.sol/Collection.json frontend/src/abi/Collection.json
