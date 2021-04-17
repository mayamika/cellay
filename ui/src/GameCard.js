import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export default function ImgMediaCard() {
  const classes = useStyles();

  return (
    <Card className={classes.root}>
      <CardActionArea>
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            Tic Tac Toe
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
            Tic-tac-toe, is a game for two players,
            X and O, who take turns marking the spaces in a 3×3 grid.
            The player who succeeds in placing three of their marks in a
            diagonal, horizontal, or vertical row is the winner.
          </Typography>
        </CardContent>
      </CardActionArea>
      <CardActions>
        <Button size="large" color="primary">
          Play
        </Button>
      </CardActions>
    </Card>
  );
}
