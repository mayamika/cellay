import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Divider from '@material-ui/core/Divider';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import {useHistory} from 'react-router-dom';

import {useAlert} from 'react-alert';

import {StoreContext} from '../store';
import API from '../api';

const useStyles = (theme) => ({
  root: {
    maxWidth: 345,
    height: '100%',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
  },
});


function GameCard(props) {
  const {classes, game} = props;
  const [, setSession] = React.useContext(StoreContext);

  const history = useHistory();
  const alert = useAlert();

  function handleClick() {
    console.log(`GameCard clicked, game id: ${game.id}`);

    API.get(`matches/new/${game.id}`)
        .then((res) => {
          const data = res.data;
          setSession({
            id: data.session,
            key: data.key,
            gameName: game.name,
          });
          console.log(`new match`, data);
          history.push(`/play`);
        })
        .catch((error) => {
          alert.error('Server returned error!');
        });
  }

  return (
    <Card className={classes.root}>
      <CardActionArea>
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            {game.name}
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
            {game.description}
          </Typography>
        </CardContent>
      </CardActionArea>
      <div>
        <Divider light />
        <CardActions>
          <Button size="large" color="primary" onClick={handleClick}>
            Host Game
          </Button>
        </CardActions>
      </div>
    </Card>
  );
}

export default withStyles(useStyles)(GameCard);
