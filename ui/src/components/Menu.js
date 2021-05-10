import {useLocation} from 'react-router-dom';
import {Link as RouterLink} from 'react-router-dom';

import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';

import {StoreContext} from '../store';


const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  separator: {
    marginRight: theme.spacing(4),
    marginLeft: theme.spacing(4),
  },
  navButton: {
    marginRight: theme.spacing(2),
    color: 'inherit',
  },
  navBox: {
  },
}));

function renderPath(path) {
  switch (path) {
    case '/play':
      return 'Play';
    case '/':
      return 'Host Game';
    case '/connect':
      return 'Connect';
    default:
      return '?';
  }
}

export default function Menu() {
  const classes = useStyles();
  const location = useLocation();
  const [session] = React.useContext(StoreContext);

  console.log(session);

  function ContinueButton() {
    if (!session.id) {
      return null;
    }
    return (
      <Button
        className={classes.navButton}
        component={RouterLink}
        to="/play">
        Continue
      </Button>
    );
  }
  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6">
            Cellay
          </Typography>
          <Typography variant="h6"className={classes.separator}>|</Typography>
          <Typography variant="h6" className={classes.title}>
            {renderPath(location.pathname)}
          </Typography>

          <Button
            className={classes.navButton}
            component={RouterLink}
            to="/">
            Host
          </Button>
          <ContinueButton />
          <Button
            className={classes.navButton}
            component={RouterLink}
            to="/connect">
            Connect
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}
