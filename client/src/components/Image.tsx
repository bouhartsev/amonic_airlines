import React from "react";
import { Box } from "@mui/material";

type Props = { path: string, hasSet?: boolean, sx?: object };

const Logo = (props: Props) => {
  const path = (props.path)?props.path:"placeholder";
  return (
    // <Typography
    //   variant="h5"
    //   noWrap
    //   component="a"
    //   href="/"
    //   sx={{
    //     flexGrow: { xs: 1, sm: 0 },
    //     fontFamily: "monospace",
    //     fontWeight: 700,
    //     letterSpacing: ".3rem",
    //     color: "inherit",
    //     textDecoration: "none",
    //     textAlign: "center",
    //     ...props.sx,
    //   }}
    // >
    //   LOGO
    // </Typography>
    <Box component="img"
      src={require(`assets/img/${path}.png`)}
      srcSet={(!!props.hasSet)?require(`assets/img/${path}@2x.png`)+" 2x, "+require(`assets/img/${path}@4x.png`)+" 4x":""}
      alt="Amonic Airlines Logo"
      sx={{
        height: "50px",
        width: "auto",
        margin: "auto",
        ...props.sx,
      }}
    />
  );
};

export default Logo;
