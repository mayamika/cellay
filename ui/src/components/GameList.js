import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';

import API from '../api';
import GameCard from './GameCard';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  card: {
    // Provide some spacing between cards
    margin: 16,

    // Use flex layout with column direction for components in the card
    // (CardContent and CardActions)
    display: 'flex',
    flexDirection: 'column',

    // Justify the content so that CardContent will always be
    // at the top of the card,
    // and CardActions will be at the bottom
    justifyContent: 'space-between',
  },
}));

const GameList = () => {
  const [games, setGames] = React.useState([]);

  React.useEffect(() => {
    API.get(`games`)
        .then((res) => {
          const data = res.data;
          console.log(data);
          setGames(data);
        });
  }, []);

  const classes = useStyles();

  function FormRow(props) {
    console.log(props.row);
    return (
      <Grid container item xs={12} spacing={3}>
        {props.row.map((game, index) => {
          return (
            <Grid item xs={4} key={index}>
              <GameCard game={game} />
            </Grid>
          );
        })}
      </Grid>
    );
  }

  function formRows(games, cols) {
    const rows = [];
    if (games === undefined) {
      return rows;
    }
    for (let i = 0; i < games.length; i += cols) {
      const endIndex = Math.min(i + cols, games.length);
      rows.push(games.slice(i, endIndex));
    }
    return rows;
  }

  const gameRows = formRows(games.games, 3);
  console.log(gameRows);

  return (
    <div className={classes.root}>
      <Grid container spacing={1} alignItems="stretch">
        {gameRows.map((row, index) => {
          return (
            <FormRow key={index} row={row} />
          );
        })}
      </Grid>
    </div>
  );
};
export default GameList;
