import React from 'react';
import Box from '@material-ui/core/Box';
import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';

import {Stage, Layer, Image} from 'react-konva';

import {useHistory} from 'react-router-dom';
import {useAlert} from 'react-alert';

import {StoreContext} from './store';
import API from './api';
import WS from './ws';


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

/*
function importImage(data) {
  const image = new window.Image();
  image.src = 'data:image/png;base64,' + data;
  return image;
}
*/

function loadImage(data) {
  return new Promise((resolve, reject) => {
    const img = new window.Image();
    img.addEventListener('load', () => resolve(img));
    img.addEventListener('error', (err) => reject(err));
    img.src = 'data:image/png;base64,' + data;
  });
}

class TileLayer {
  constructor(rawLayer, image) {
    this.image = image;
    this.width = rawLayer.width;
    this.height = rawLayer.height;
    this.depth = rawLayer.depth;
    this.cols = Math.floor(this.image.width / this.width);
    this.rows = Math.floor(this.image.height / this.height);
  }

  tile(index, x, y) {
    const row = Math.floor(index / this.cols);
    const col = index % this.cols;
    if (row > this.rows) {
      throw new Error('incorrect index');
    }
    const props = {
      image: this.image,
      x: x * this.width,
      y: y * this.height,
      width: this.width,
      height: this.height,
      crop: {
        x: col * this.width,
        y: row * this.height,
        width: this.width,
        height: this.height,
      },
    };
    return props;
  }
}

function transformAssets(raw) {
  const assets = {};
  const promises = [];
  if (raw.background.texture) {
    promises.push(
        loadImage(raw.background.texture).then(
            (image) => {
              assets.background = image;
              assets.width = image.width;
              assets.height = image.height;
            },
        ),
    );
  } else {
    assets.width = raw.background.width;
    assets.height = raw.background.height;
  }
  assets.layers = {};
  for (const name in raw.layers) {
    if (!Object.prototype.hasOwnProperty.call(raw.layers, name)) {
      continue;
    }
    const layer = raw.layers[name];
    if (!layer.texture) {
      continue;
    }
    promises.push(
        loadImage(layer.texture).then(
            (image) => assets.layers[name] = new TileLayer(layer, image),
        ),
    );
  }
  return Promise.all(promises).then(
      (value) => {
        return assets;
      },
  );
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
          transformAssets(data).then(
              (assets) => setAssets(assets),
          );
        });
    /*
        .catch((error) => {
          resetSession(setSession);
          alertReturnHome(history, alert, 'Assets unavailable');
        });
        */
  };
  // Reset session on exit
  React.useEffect(() => {
    if (!session.id) {
      alertReturnHome(history, alert, 'No session found');
      return;
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
  const session = props.session;
  const assets = props.assets;

  const aspect = assets.width / assets.height;
  const [canvasSize, setCanvasSize] = React.useState(getWidthHeight(aspect));
  React.useEffect(() => {
    const handleResize = (e) => {
      setCanvasSize(getWidthHeight(aspect));
    };
    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, []);

  React.useEffect(() => {
    const ws = new WS(session.key);

    const channel = ws.subscribe(session.id, (message) => {
      console.log(message);
    });

    return () => {
      channel.unsubscribe();
      ws.disconnect();
    };
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
      <Layer>
        <Image
          {...assets.layers.main.tile(1, 2, 1)}
        />
        <Image
          {...assets.layers.main.tile(0, 1, 1)}
        />
      </Layer>
    </Stage>
  );
}
