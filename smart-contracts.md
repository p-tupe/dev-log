---
modified: Wed Sep 18 19:30:29 EDT 2024
---
# Blockchain and friends

## Resources

- https://info.etherscan.com/

- https://ethereum.org/en/learn/

- https://wagmi.sh/core/why

- https://www.fon.hum.uva.nl/rob/Courses/InformationInSpeech/CDROM/Literature/LOTwinterschool2006/szabo.best.vwh.net/smart.contracts.html

- https://github.com/ethereumbook/ethereumbook/blob/develop/01what-is.asciidoc

- https://ethereum.org/en/developers/docs/intro-to-ethereum/

- https://stark.mirror.xyz/n2UpRqwdf7yjuiPKVICPpGoUNeDhlWxGqjulrlpyYi0

---

## Blockchain

Consider a ledger, that is, a list of transactions. A ledger that is not owned by anyone; instead it is stored on computers called "nodes" on the internet. Anyone interested can become/create a node and get the ledger.

When a User A sends some money to User B, that transaction is batched along with all the other transactions happening around the same time nearby. This batch is called a "block". The block is sent to all the nearest nodes(?), who check if the transactions in the block follow proper rules (termed "protocols") and connect it ("atomically") to the earier blocks, forming a unique immutable "chain" of blocks. This is called Blockchain.

## Smart Contract

A "transaction" in a blockchain does not have to be a money transfer, though it usually is.

At its core, it is basically a "if X then Y" condition. For example, say User A wants to comissions a digital artwork from User B that they need to use as a brand logo before their product launch in 3 months. A smart contract would be:

```
If
    1. Artwork Received by User A
    2. Date is not more than 3 months from signing date
Then
    Pay User B an amount of X
```

If the logo is received within 3 months, the contract would be fulfilled and User B will automatically be paid.

## Wallets

Wallets don't hold money, instead they hold the keys. Specifically, a private key that proves your identity.

All transactions are on blockchain and signed by each participating party using their keys. Thus, when a User A received amount X from User B, that transaction is added on top of the blockchain and when User A checks their balance, the wallet uses the private key to get all transactions by A and calculated the balance at the end, which now would be previous balance + X.

## Smart Contracts on a Blockchain with Wallets

As mentioned earlier, a contract does not necessarily have to be a money transfer. Expanding on the core of "if X then Y", a smart contract is a collection of data (state) and functions (actions) - like an OOP class.

When a contract is added on a chain, an address is generated for it.

For eg, this contract

```js
artWork = P
signingDate = D
amount = X
fulfilled = false

function checkArtwork(address A) {
     // check if artwork P is received by User A
}

function checkDate() {
    // check if current date is within 3 months of D
}

function transfer(address B) {
    // Transfer Amount X to User B if not already fulfilled
}
```

will genereate an address 0x0Aak...ajad, using which we can check if the contract is fulfilled and can transfer the amount using the address as the namespace.

This means a "blockchain" is similar to a global namespace holding all the addresses, not dissimilar to a virtual runtime holding state/functions.

## Oracles

Since blockchains are by nature unchanging, they don't have access to external world. So, how would a smart contract determine if the the conditions have been fulfilled? Enter oracles. They are "queries" to the outside world that let a smart contract know if a condition has been fulfilled.

## Consensus Mechanisms

The consensus mechanism is the mechanism that ensures all nodes agrees that a given blockchain is in fact correct.

### Proof-of-Stake

### Proof-of-Work

### Proof-of-Authority

## Web3

## Cryptocurrencies

## NFTs

## Gas

## TestNets

## Contracts ABI

### ERC720 vs ERC721
