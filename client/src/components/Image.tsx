import React from "react";
import { Box } from "@mui/material";

type Props = { path: string; hasSet?: boolean; sx?: object; alt?: string };

const Logo = (props: Props) => {
  const path = props.path ? props.path : "placeholder";
  return (
    <Box
      component="img"
      src={require(`assets/img/${path}.png`)}
      srcSet={
        !!props.hasSet
          ? require(`assets/img/${path}@2x.png`) +
            " 2x, " +
            require(`assets/img/${path}@4x.png`) +
            " 4x"
          : ""
      }
      alt={props.alt ? props.alt : "Unloaded Image"}
      sx={{
        display: "block",
        height: "50px",
        width: "auto",
        margin: "auto",
        ...props.sx,
      }}
    />
  );
};

export default Logo;
