# Collaborative Mapping Smart Contracts

## Description
This repository contains the IOTA and Ethereum smart contract implementation of collaborative mapping. The collaborative mapping is a demonstration of how IOTA's smart contracts with the asynchronous requests between chains provide partition tolerance.

- - -
## Ethereum Smart Contract

### Setup and Test
For testing Ethereum smart contracst, you can use [Remix](https://remix.ethereum.org/) IDE. Remix is totally online and you don't need to install anything. You can deploy and test the smart contract on the IDE.

If you have own your private network. You can deploy the smart contract on your own network.

Multiple robots can register on the smart contract. After mapping the area they must submit the created map to the smart contract. By calling the `getGlobalMap` everyone can receive the merged maps.
- - -

## IOTA Smart Contract
*NOTE: IOTA 2.0 is heavily under development and in every version APIs chaing alot. The given instruction might not work with newer versions.*
### Setup and Test
You will need at least two `wasp` node connected to one or multiple `goshimmer` nodes. Each of the `controllerswarm` and `agentswarm` smart contracts should be deployed on two different chain. These chains should be created by two disjoint `wasp` committies. Also the committe members should be connected to two disjoint `goshimmer` nodes.

For setting up everything on your local network:
+ You should bring up a `goshimmer` network. You can use the [script](https://github.com/iotaledger/goshimmer/blob/develop/tools/docker-network/run.sh) provided in `goshimmer` github repository. You need at least two `goshimmer` nodes.
+ In this script, you should also enable `faucet` and `txstream` plugins for the `goshimmer` nodes. You should change the [`dockercompose`](https://github.com/iotaledger/goshimmer/blob/develop/tools/docker-network/docker-compose.yml) file in the `peer_replica`.
```
--node.enablePlugins=bootstrap,"webAPIToolsEndpoint",faucet,txstream
```
+ Then you should setup at least two `wasp` nodes. You can use this [link](https://wiki.iota.org/smart-contracts/guide/chains_and_nodes/docker_standalone). Just make sure that you bring it up in the same network that is created by `goshimmer` script, so the `wasp` node you create can communicate with `goshimmer` ones.
+ By using `wasp-cli` you should create a new chain and deploy the smart contract. For creating the new chain and a tutorial on how you can deploy smart contracts you can follow the instructions given [here](https://wiki.iota.org/smart-contracts/guide/chains_and_nodes/wasp-cli).
+ For compiling the smart contract, you should clone the `wasp` repository. It is tested with `v0.2.2`.
+ Then you should copy the `controllerswarm` and `agentswarm` folders to `contracts/swarm` folder of the repository. And follow the [instructions](https://wiki.iota.org/smart-contracts/guide/chains_and_nodes/setting-up-a-chain#deploying-a-wasm-contract).

