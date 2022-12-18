import { observer } from "mobx-react-lite";
import { useStore } from "stores";
import NavBarAuth from "./NavBarAuth";
import NavBarPublic from "./NavBarPublic";
import { AppBar, Container, Toolbar } from "@mui/material";

const NavBar = () => {
  const { userStore } = useStore();
  return (
    <AppBar position="sticky">
      <Container maxWidth="xl" disableGutters sx={{p: 4}}>
        <Toolbar disableGutters>
          {!userStore.isLogged ? <NavBarPublic /> : <NavBarAuth role={userStore.userData?.role}/>}
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default observer(NavBar);
