import Centrifuge from 'centrifuge';

const isSecure = (window.location.protocol === 'https:' ||
  process.env.HTTPS === 'true');
const protocol = isSecure ? 'wss' : 'ws';
const host = process.env.HOST || window.location.hostname;
const port = process.env.PORT || window.location.port;
const centrifugeURL = `${protocol}://${host}:${port}/connection/websocket`;

export default class WS {
  constructor(token) {
    this.centrifuge = new Centrifuge(
        centrifugeURL,
        {
          sockjsTransports: ['websocket'],
          debug: true,
        },
    );
    this.centrifuge.setToken(token);
    this.centrifuge.connect();
  }

  subscribe(channel, cb) {
    return this.centrifuge.subscribe(channel, cb);
  }

  disconnect() {
    this.centrifuge.disconnect();
  }
}
