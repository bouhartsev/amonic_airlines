import React from "react";
import { Link as LinkRouter } from "react-router-dom";
import Logo from "components/Logo";
import { Button } from "@mui/material";

const NavBarPublic = () => {
  return (
    <>
      <Logo sx={{ flexGrow: 1 }}/>
      <Button component={LinkRouter} to={"/login"}  color="inherit">
        {"Login"}
      </Button>
    </>
  );
};

export default NavBarPublic;
