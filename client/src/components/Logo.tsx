import React from "react";
import { Typography } from "@mui/material";

type Props = { sx?: any };

const Logo = (props: Props) => {
  return (
    <Typography
      variant="h5"
      noWrap
      component="a"
      href="/"
      sx={{
        flexGrow: { xs: 1, sm: 0 },
        fontFamily: "monospace",
        fontWeight: 700,
        letterSpacing: ".3rem",
        color: "inherit",
        textDecoration: "none",
        textAlign: "center",
        ...props.sx,
      }}
    >
      LOGO
    </Typography>
  );
};

export default Logo;
