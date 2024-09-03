# Peer to Peer Chat

## Description

This is a simple P2P Chat that gives you the option of being the Client or the Server and then uses a combination of UDP and TCP to comunicate messages through the console.

## How does it work

The Server creates a UDP broadcast and then the Client who is listening for possible broadcasts will get the information and establish a TCP handshake and then both will start comunicating.

## Tasks

- [ ] Implement Server UDP broadcast.
- [ ] Implement Client UDP listening.
- [ ] Implement TCP handshake.
- [ ] Implement sending messages.
- [ ] (Bonus) Encrypt Messages.
