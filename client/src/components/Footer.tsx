import React from "react";
import { Box, Typography } from "@mui/material";

const Footer = () => {
  return (
    <Box
      component="footer"
      sx={{
        py: 4,
        px: 2,
        textAlign: "center",
      }}
    >
      <Typography variant="body1">Copyright © Infinity, 2022</Typography>
    </Box>
  );
};

export default Footer;
