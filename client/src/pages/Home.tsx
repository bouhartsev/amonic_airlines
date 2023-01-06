import { Typography } from "@mui/material";
import React from "react";

type Props = {};

const Home = (props: Props) => {
  return (
    <>
      <Typography component="h2" variant="h4">
        What is Amonic Airlines?
      </Typography>
      <Typography>{/* About */}</Typography>
      <Typography component="h2" variant="h4">
        For company
      </Typography>
      <Typography component="h2" variant="h4">
        For people
      </Typography>
      <Typography component="h2" variant="h4">
        Authors
      </Typography>
      <Typography>
        Infinity team:
        <ul>
          <li>Matvey Bouhartsev</li>
          <li>Ilia Sokolov</li>
          <li>Ekaterina Shunaeva</li>
        </ul>
        <a href="mailto:infinity@bouhartsev.top">infinity@bouhartsev.top</a>
      </Typography>
    </>
  );
};

export default Home;
