import React from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

export default function ConnectPage() {
  const [text, setText] = React.useState();

  function textChange(event) {
    event.preventDefault();
    setText(event.target.value);
  }

  function textKeyDown(event) {
    if (event.key === 'Enter') {
      event.preventDefault();
      textSubmit();
    }
  }

  function buttonClick(event) {
    event.preventDefault();
    textSubmit();
    // Use textFieldInput
  }

  function textSubmit() {
    console.log(text);
  }

  return (
    <Box my={4}>
      <Typography variant="h4" component="h1" gutterBottom>
        Connect another session
      </Typography>
      <Grid
        container
        justify='flex-start'
      >
        <TextField id="standart-basic" label="session"
          onChange={textChange}
          onKeyDown={textKeyDown}
        />
        <Button onClick={buttonClick}>connect</Button>
      </Grid>
    </Box>
  );
}

