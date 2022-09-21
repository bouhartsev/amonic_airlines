import { observer } from "mobx-react-lite";
import NavBarAuth from "./NavBarAuth";
import NavBarPublic from "./NavBarPublic";
import {
  AppBar,
  Container,
  Toolbar,
} from "@mui/material";

type Props = {};

const NavBar = (props: Props) => {
  return (
    <AppBar position="sticky">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          {/* If not auth */}
          {/* <NavBarPublic /> */}
          <NavBarAuth />
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default observer(NavBar);
