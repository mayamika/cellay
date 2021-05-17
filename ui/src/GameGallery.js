import React from 'react';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import GameList from './components/GameList';

export default function GameGallery() {
  return (
    <Box my={4}>
      <Typography variant="h4" component="h1" gutterBottom>
        Games
      </Typography>
      <GameList />
    </Box>
  );
}
