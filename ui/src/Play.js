import React from 'react';
import Box from '@material-ui/core/Box';
import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';

import {Stage, Layer, Image} from 'react-konva';

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

function importImage(data) {
  const image = new window.Image();
  image.src = 'data:image/png;base64,' + data;
  return image;
}

function transformAssets(raw) {
  const assets = {};
  if (raw.background.texture) {
    assets.background = importImage(raw.background.texture);
    assets.width = assets.background.width;
    assets.height = assets.background.height;
  } else {
    assets.width = raw.background.width;
    assets.height = raw.background.height;
  }
  return assets;
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
          setAssets(data);
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

  function GameBox(props) {
    if (!assets) {
      return (
        <Typography variant='h5' component='h1' gutterBottom>
              Loading
        </Typography>
      );
    }
    return (
      <Box display='flex' justifyContent='center'>
        <GameCanvas assets={assets} session={session} />
      </Box>
    );
  }

  return (
    <Box my={4}>
      <Typography variant='h4' component='h1' gutterBottom>
            Play with another player
      </Typography>
      <Container maxWidth='sm'>
        <GameBox />
      </Container>
    </Box>
  );
}

function getWidthHeight(aspect) {
  const width = window.innerWidth * 0.8;
  const height = window.innerHeight * 0.8;
  if (height < width) {
    return {
      width: height * aspect,
      height: height,
    };
  }
  return {
    width: width,
    height: width / aspect,
  };
}

function GameCanvas(props) {
  const assets = transformAssets(props.assets);

  const aspect = assets.width / assets.height;
  console.log(aspect);

  const [canvasSize, setCanvasSize] = React.useState(getWidthHeight(aspect));

  React.useEffect(() => {
    const handleResize = (e) => {
      setCanvasSize(getWidthHeight(aspect));
    };

    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, []);


  return (
    <Stage width={canvasSize.width} height={canvasSize.height}
      scaleX={canvasSize.width / assets.width}
      scaleY={canvasSize.height / assets.height}>
      <Layer>
        <Image
          image={assets.background}
        />
      </Layer>
    </Stage>
  );
}
