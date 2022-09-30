import React from "react";
import { Link as LinkRouter } from "react-router-dom";
import Image from "components/Image";
import { Button } from "@mui/material";

const NavBarPublic = () => {
  return (
    <>
      <Image path="logo_colors" hasSet/>
      <Button component={LinkRouter} to={"/login"} color="inherit">
        {"Login"}
      </Button>
    </>
  );
};

export default NavBarPublic;
