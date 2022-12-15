import React from "react";
import { Link as LinkRouter } from "react-router-dom";
import Image from "components/Image";
import { Button } from "@mui/material";

const NavBarPublic = () => {
  return (
    <>
      <a href="/" style={{flexGrow:"1"}}><Image path="logo_colors" hasSet alt="Amonic Airlines Logo" /></a>
      <Button component={LinkRouter} to={"/login"} color="inherit">
        {"Login"}
      </Button>
    </>
  );
};

export default NavBarPublic;
