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

function loadImage(data) {
  return new Promise((resolve, reject) => {
    const img = new window.Image();
    img.addEventListener('load', () => resolve(img));
    img.addEventListener('error', (err) => reject(err));
    img.src = 'data:image/png;base64,' + data;
  });
}

class TileLayer {
  constructor(rawLayer, field, image) {
    this.image = image;
    this.tileRealWidth = rawLayer.width;
    this.tileRealHeight = rawLayer.height;
    this.depth = rawLayer.depth;
    this.cols = Math.floor(this.image.width / this.tileRealWidth);
    this.rows = Math.floor(this.image.height / this.tileRealHeight);
    this.tileWidth = field.cellWidth;
    this.tileHeight = field.cellHeight;
    this.scaleX = this.tileWidth / this.tileRealWidth;
    this.scaleY = this.tileHeight / this.tileRealHeight;
  }

  tile(index, x, y) {
    const row = Math.floor(index / this.cols);
    const col = index % this.cols;
    if (row > this.rows) {
      throw new Error('incorrect index');
    }
    const props = {
      image: this.image,
      x: x * this.tileWidth,
      y: y * this.tileHeight,
      width: this.tileWidth,
      height: this.tileHeight,
      crop: {
        x: col * this.tileRealWidth,
        y: row * this.tileRealHeight,
        width: this.tileRealWidth,
        height: this.tileRealHeight,
      },
    };
    return props;
  }
}

function transformAssets(raw) {
  const assets = {};
  const promises = [];
  assets.rows = raw.field.rows;
  assets.cols = raw.field.cols;
  if (raw.background.texture) {
    promises.push(
        loadImage(raw.background.texture).then(
            (image) => assets.background = image,
        ),
    );
  }
  assets.width = raw.background.width;
  assets.height = raw.background.height;
  assets.cellWidth = assets.width / assets.cols;
  assets.cellHeight = assets.height / assets.rows;
  assets.layers = {};
  const order = [];
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
            (image) => assets.layers[name] =
                new TileLayer(layer, assets, image),
        ),
    );
    order.push(name);
  }
  order.sort((first, second) => {
    return raw.layers[first].depth > raw.layers[second].depth;
  });
  assets.order = order;
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

  const stage = React.useRef();

  const aspect = assets.width / assets.height;
  const [canvasSize, setCanvasSize] = React.useState(getWidthHeight(aspect));
  React.useEffect(() => {
    const handleResize = (e) => {
      setCanvasSize(getWidthHeight(aspect));
    };
    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, []);

  const [socket, setSocket] = React.useState(null);

  const [field, setField] = React.useState(null);

  React.useEffect(() => {
    const ws = new WS(session.key);

    const ch = ws.subscribe(session.id, (message) => {
      setField(message.data.Table);
    });

    setSocket(ws);

    return () => {
      ch.unsubscribe();
      ws.disconnect();
    };
  }, []);

  if (!socket || !field) {
    return null;
  }

  const handleClick = (e) => {
    const node = stage.current;
    const transform = node.getAbsoluteTransform().copy().invert();
    const pos = node.getStage().getPointerPosition();
    const {x, y} = transform.point(pos);
    const col = Math.floor(x / assets.cellWidth);
    const row = Math.floor(y / assets.cellHeight);
    socket.send({
      Type: 'click',
      X: col,
      Y: row,
    })
        .catch((error) =>{
          console.log('channel error', error);
        });
  };

  const cellStates = {};
  for (const name of assets.order) {
    const layerStates = [];
    for (let x = 0; x < field[name].length; x++) {
      const rows = field[name][x];
      for (let y = 0; y < rows.length; y++) {
        layerStates.push(assets.layers[name].tile(rows[y], x, y));
      }
    }
    cellStates[name] = layerStates;
  }
  console.log(cellStates);

  return (
    <Stage width={canvasSize.width} height={canvasSize.height}
      scaleX={canvasSize.width / assets.width}
      scaleY={canvasSize.height / assets.height}
      onClick={handleClick}
      ref={stage}>
      <Layer>
        <Image
          image={assets.background}
        />
      </Layer>
      {assets.order.map((name) => {
        return (
          <Layer key={name}>
            {cellStates[name].map((val, index) => {
              return (
                <Image key={index}
                  {...val}
                />
              );
            })}
          </Layer>
        );
      })}
    </Stage>
  );
}
