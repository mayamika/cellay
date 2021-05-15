import React from 'react';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';

import {Stage, Layer, Text} from 'react-konva';

import {useHistory} from 'react-router-dom';
import {useAlert} from 'react-alert';

import {StoreContext} from './store';
import API from './api';

/*
import {Stage, Layer} from 'react-konva';
import Paper from '@material-ui/core/Paper';

import {StoreContext} from './store';

import Centrifuge from 'centrifuge';

const isSecure = (window.location.protocol === 'https:' ||
  process.env.HTTPS === 'true');
const protocol = isSecure ? 'wss' : 'ws';
const host = process.env.HOST || window.location.hostname;
const port = process.env.PORT || window.location.port;
const centrifugeURL = `${protocol}://${host}:${port}/connection/websocket`;

class PlayPage extends React.Component {
  constructor(props) {
    super(props);
  }

  static contextType = StoreContext

  componentDidMount() {
    const [session] = this.context;
    const alert = this.props.alert;
    const connect = () => {
      console.log('connect');
      const centrifuge = new Centrifuge(
          centrifugeURL,
          {
            sockjsTransports: ['websocket'],
            debug: true,
          },
      );
      // centrifuge.setToken(session.key);
      centrifuge.subscribe(session.id, (message) => {
        console.log(message);
      });
      centrifuge.connect();
    };
    const loadAssets = (gameId) => {
      API.get(`games/${gameId}/assets`)
          .then((res) => {
            const data = res.data;
            this.setState({
              assets: data,
            });
            connect();
          })
          .catch((error) => {
            console.log(error);
            alert.error('Server returned error!');
          });
    };
    API.get(`matches/info/${session.id}`)
        .then((res) => {
          const data = res.data;
          loadAssets(data.gameId);
          console.log(`match info`, data);
        })
        .catch((error) => {
          console.log(error);
          alert.error('Server returned error!');
        });
  }

  componentWillUnmount() {

  }

  render() {
    return (
      <Box my={4}>
        <Typography variant='h4' component='h1' gutterBottom>
          Play with another player
        </Typography>
        <Paper>
          <Stage
            width={0.65 * window.innerWidth}
            height={0.8 * window.innerHeight}
          >
            <Layer>

            </Layer>
          </Stage>
        </Paper>
      </Box>
    );
  }
}
*/

/* state?
 * {
 *   assets: {},
 *   gameState: {},
 *
 * }
 *
 */

function alertReturnHome(history, alert, msg) {
  alert.error(msg, {
    onClose: () => {
      history.push('/');
    },
  });
}

function resetSession(setSession) {
  setSession({
    key: null,
    id: null,
  });
}

function transformAssets(raw) {
  console.log(JSON.stringify(raw));
  return raw;
}


export default function GameContainer(props) {
  const history = useHistory();
  const alert = useAlert();
  const [session, setSession] = React.useContext(StoreContext);

  const [assets, setAssets] = React.useState(null);

  function loadAssets(gameId) {
    API.get(`games/${gameId}/assets`)
        .then((res) => {
          const data = res.data;
          setAssets(transformAssets(data));
        })
        .catch((error) => {
          resetSession(setSession);
          alertReturnHome(history, alert, 'Assets unavailable');
        });
  };
  // Reset session on exit
  React.useEffect(() => {
    if (!session.id) {
      alertReturnHome(history, alert, 'No session found');
    }
    API.get(`matches/info/${session.id}`)
        .then((res) => {
          loadAssets(res.data.gameId);
        })
        .catch((error) => {
          resetSession(setSession);
          alertReturnHome(history, alert, 'Session does not exists');
        });
  }, []);

  return (
    <Box my={4}>
      <Typography variant='h4' component='h1' gutterBottom>
            Play with another player
      </Typography>
      <GameCanvas assets={assets} session={session}/>
    </Box>
  );
}

function GameCanvas(props) {
  if (!props.assets) {
    return (
      <Stage width={window.innerWidth} height={window.innerHeight}>
        <Layer>
          <Text text="Loading assets..." fontSize={25} />
        </Layer>
      </Stage>
    );
  }
  return (
    <Stage width={window.innerWidth} height={window.innerHeight}>
      <Layer>
        <Text text="Loaded" fontSize={25} />
      </Layer>
    </Stage>
  );
}
