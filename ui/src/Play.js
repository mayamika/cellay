import React from 'react';
import Box from '@material-ui/core/Box';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import {Stage, Layer} from 'react-konva';

import API from './api';
import {StoreContext} from './store';
import {withAlert} from 'react-alert';

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

export default withAlert()(PlayPage);
