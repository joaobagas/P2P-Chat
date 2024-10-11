# P2P Chat

## Description
This is a simple P2P Chat that gives you the option of being the Client or the Server and then uses a combination of UDP and TCP to comunicate messages through the console.

## Tasks
- [x] Implement frontend.
- [x] Implement broadcasting (UDP).
- [ ] Implement listening (UDP).
- [ ] Implement sending messages (TCP).
- [ ] Test using command line.
- [ ] Connect frontend with backend.
- [ ] (Bonus) Encrypt Messages.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
