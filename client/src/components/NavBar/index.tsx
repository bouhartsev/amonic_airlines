import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import NavBarAuth from "./NavBarAuth";
import NavBarPublic from "./NavBarPublic";
import { AppBar, Container, Toolbar } from "@mui/material";

const NavBar = () => {
  const { userStore } = useStore();
  return (
    <AppBar position="sticky">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          {/* If not auth */}
          {!userStore.isLogged ? <NavBarPublic /> : <NavBarAuth />}
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default observer(NavBar);
