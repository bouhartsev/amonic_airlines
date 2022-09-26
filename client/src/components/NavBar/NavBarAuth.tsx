import * as React from "react";
import { Link as LinkRouter } from "react-router-dom";
import {
  Menu,
  Avatar,
  Button,
  Tooltip,
  Typography,
  IconButton,
  Box,
  MenuItem,
  Drawer,
  Divider,
  List,
  ListItem,
  ListItemButton,
  ListItemText,
} from "@mui/material";
import { Menu as MenuIcon } from "@mui/icons-material";
import Logo from "components/Logo";

const pages = ["Schedules", "Booking", "Tickets", "Surveys"]; // title and href
const options = ["Dashboard", "Profile", "Logout"]; // AND "USERS" IF ADMIN

type Props = {
  role?: string,
}

const NavBar = (props: Props) => {
  const [mobileOpen, setMobileOpen] = React.useState(false);
  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(
    null
  );

  const handleDrawerToggle = () => {
    setMobileOpen(!mobileOpen);
  };

  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };
  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  const rootContainer = document.body;

  return (
    <>
      <Drawer
        container={rootContainer}
        variant="temporary"
        open={mobileOpen}
        onClose={handleDrawerToggle}
        ModalProps={{
          keepMounted: true, // Better open performance on mobile.
        }}
        sx={{
          display: { sm: "none" },
          "& .MuiDrawer-paper": { boxSizing: "border-box", width: 240 },
        }}
      >
        <Box onClick={handleDrawerToggle} sx={{ textAlign: "center" }}>
          <Logo />
          <Divider />
          <List>
            {pages.map((page) => (
              <ListItem key={page} disablePadding>
                <ListItemButton
                  sx={{ textAlign: "center" }}
                  component={LinkRouter}
                  to={page}
                >
                  <ListItemText primary={page} />
                </ListItemButton>
              </ListItem>
            ))}
          </List>
        </Box>
      </Drawer>
      <IconButton
        size="large"
        aria-label="Navigation menu"
        aria-controls="menu-appbar"
        aria-haspopup="true"
        onClick={handleDrawerToggle}
        color="inherit"
        sx={{ mr: 2, display: { sm: "none" } }}
      >
        <MenuIcon />
      </IconButton>
      <Logo sx={{ mr: 1 }} />
      <Box sx={{ flexGrow: 1, display: { xs: "none", sm: "flex" } }}>
        {pages.map((page) => (
          <Button
            key={page}
            // onClick={handleCloseNavMenu}
            component={LinkRouter}
            to={page}
            sx={{ my: 2, color: "white", display: "block" }}
          >
            {page}
          </Button>
        ))}
      </Box>

      <Box sx={{ flexGrow: 0 }}>
        <Tooltip title="Open options">
          <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
            <Avatar alt="User Name" src="/path/to/photo.png" />
          </IconButton>
        </Tooltip>
        <Menu
          sx={{ mt: "45px" }}
          id="menu-appbar"
          anchorEl={anchorElUser}
          anchorOrigin={{
            vertical: "top",
            horizontal: "right",
          }}
          keepMounted
          transformOrigin={{
            vertical: "top",
            horizontal: "right",
          }}
          open={Boolean(anchorElUser)}
          onClose={handleCloseUserMenu}
        >
          {options.map((setting) => (
            <MenuItem
              key={setting}
              onClick={handleCloseUserMenu}
              component={LinkRouter}
              to={setting}
            >
              <Typography textAlign="center">{setting}</Typography>
            </MenuItem>
          ))}
        </Menu>
      </Box>
    </>
  );
};
export default NavBar;
