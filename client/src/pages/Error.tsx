import React from "react";
import { Box, Typography } from "@mui/material";
import { Link } from "react-router-dom";

const handledErrors: { [code: string]: { status: string; message: string } } = {
  "400": {
    status: "Bad Request",
    message: "Unhandled client error.",
  },
  "403": {
    status: "Forbidden",
    message: "You don't have access rights to the content.",
  },
  "404": {
    status: "Not Found",
    message:
      "The page is not found. The URL is not recognized, try to change it.",
  },
};

const Error = ({ code = "400" }) => {
  return (
    <Box sx={{ textAlign: "center" }}>
      <Typography component="h1" variant="h2" gutterBottom>
        {code} | {handledErrors[code]?.status}
      </Typography>
      <Typography variant="subtitle1" gutterBottom>
        {handledErrors[code]?.message}<br />
        <Link to="/">Back to the HOME PAGE</Link>
      </Typography>
    </Box>
  );
};

export default Error;
