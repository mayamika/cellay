import React from 'react';
import {Grid, Typography, Box, Button, TextField} from '@material-ui/core';
import {DropzoneArea} from 'material-ui-dropzone';

import {useAlert} from 'react-alert';
import {useHistory} from 'react-router-dom';

import API from './api';

export default function AddGamePage() {
  return (
    <Box my={4}>
      <Typography variant="h4" component="h1" gutterBottom>
        Add game
      </Typography>
      <AddForm />
    </Box>
  );
}

function AddForm() {
  const alert = useAlert();
  const history = useHistory();

  const [text, setText] = React.useState({});
  const [, setLayers] = React.useState();
  const [, setBackground] = React.useState();

  const changeText = (event) => {
    const target = event.target;
    setText((state) => {
      state[target.name] = target.value;
      return state;
    });
    console.log(text);
  };

  const publishGame = () => {
    API.post(`/games`, {
      name: text['name'],
      description: text['description'],
      code: text['code'],
      field: {
        cols: 3,
        rows: 3,
      },
    })
        .then((res) => {
          history.push(`/`);
        })
        .catch((error) => {
          console.log(error);
          alert.error('add game failed');
        });
  };

  const submit = (event) => {
    event.preventDefault();
    publishGame();
  };

  return (
    <Box my={4}>
      <Grid
        container
        justify="flex-start"
        spacing={2}
      >
        <Grid container item>
          <Grid item xs>
            <TextField
              name="name"
              id="standart-basic"
              label="name"
              onChange={changeText}
            />
          </Grid>
        </Grid>
        <Grid container item>
          <Grid container item xs={6}>
            <Grid item xs>
              <TextField
                name="cols"
                id="standart-number"
                label="cols"
                type="number"
                InputLabelProps={{
                  shrink: true,
                }}
                onChange={changeText}
              />
            </Grid>
            <Grid item xs>
              <TextField
                name="rows"
                id="standart-number"
                label="rows"
                type="number"
                InputLabelProps={{
                  shrink: true,
                }}
                onChange={changeText}
              />
            </Grid>
          </Grid>
        </Grid>
        <Grid container item>
          <Grid item xs={6}>
            <TextField
              name="description"
              id="standart-basic"
              label="description"
              multiline
              fullWidth
              variant="outlined"
              onChange={changeText}
            />
          </Grid>
        </Grid>
        <Grid container item>
          <Grid item xs={6}>
            <TextField
              name="code"
              id="standart-basic"
              label="code"
              multiline
              fullWidth
              variant="outlined"
              onChange={changeText}
            />
          </Grid>
        </Grid>
        <Grid container item>
          <Grid item xs={6}>
            <DropzoneArea
              filesLimit={1}
              clearOnUnmount={true}
              acceptedFiles={['image/png']}
              dropzoneText={'Upload background'}
              onChange={(files) => setBackground(files)}
            />
          </Grid>
        </Grid>
        <Grid container item>
          <Grid item xs={6}>
            <DropzoneArea
              filesLimit={10}
              clearOnUnmount={true}
              acceptedFiles={['image/png']}
              dropzoneText={'Upload layers'}
              onChange={(files) => setLayers(files)}
            />
          </Grid>
        </Grid>
        <Grid container item>
          <Grid item xs={3}>
            <Button onClick={submit}>add</Button>
          </Grid>
        </Grid>
      </Grid>
    </Box>
  );
}
