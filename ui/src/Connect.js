import React from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

import {useAlert} from 'react-alert';
import {useHistory} from 'react-router-dom';

import API from './api';
import {StoreContext} from './store';

export default function ConnectPage() {
  const alert = useAlert();
  const history = useHistory();
  const [, setSession] = React.useContext(StoreContext);

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
    API.get(`matches/info/${text}`)
        .then((res) => {
          const data = res.data;
          setSession({
            id: text,
            key: data.key,
            gameName: data.gameName,
          });
          console.log(data);
          history.push(`/play`);
        })
        .catch((error) => {
          alert.error('session is invalid');
        });

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
