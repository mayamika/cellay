import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Link from '@material-ui/core/Link';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from 'react-router-dom';

import {transitions, positions, Provider as AlertProvider} from 'react-alert';
import AlertTemplate from 'react-alert-template-mui';

import Menu from './components/Menu';
import GameGallery from './GameGallery';
import ConnectPage from './Connect';
import GameContainer from './Play';
import StoreProvider from './store';

const useStyles = makeStyles((theme) => ({
  content: {
    display: 'flex',
    flex: '1 1 auto',
    flexFlow: 'column',
    minHeight: '100vh',
    position: 'relative',
  },
  footer: {
    position: 'absolute',
    left: 0,
    bottom: 0,
    right: 0,
  },
}));

// optional configuration
const alertOptions = {
  // you can also just use 'bottom center'
  position: positions.BOTTOM_CENTER,
  timeout: 6000,
  offset: '30px',
  // you can also just use 'scale'
  transition: transitions.SCALE,
};

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://github.com/mayamika/cellay">
        Cellay
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

export default function App() {
  const classes = useStyles();

  return (
    <StoreProvider>
      <AlertProvider template={AlertTemplate} {...alertOptions}>
        <Router>
          <Menu />
          <Container className={classes.content} maxWidth="lg">
            <Switch>
              <Route path="/connect/:sessionId">
                <ConnectPage />
              </Route>
              <Route path="/connect">
                <ConnectPage />
              </Route>
              <Route path="/play">
                <GameContainer />
              </Route>
              <Route path="/">
                <GameGallery />
              </Route>
            </Switch>
            <Box className={classes.footer}>
              <Copyright />
            </Box>
          </Container>
        </Router>
      </AlertProvider>
    </StoreProvider>
  );
}
