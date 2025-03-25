import React from "react";
import { AppBar, Toolbar, Typography, Link } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import { navbarStyles } from "../styles/navbarStyles";

const Navbar: React.FC = () => {
  return (
    <AppBar position="static" sx={navbarStyles.appBar}>
      <Toolbar>
        <Typography variant="h6" component="div" sx={navbarStyles.title}>
          To-Do App
        </Typography>
        <Link
          component={RouterLink}
          to="/"
          color="inherit"
          underline="none"
          sx={navbarStyles.link}
        >
          Home
        </Link>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;