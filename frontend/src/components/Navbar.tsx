import React from "react";
import { AppBar, Toolbar, Typography, Link } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import { navbarStyles } from "../styles/navbarStyles";

const Navbar: React.FC = () => {
  return (
    <AppBar
      position="static"
      sx={{
        backgroundColor: "#1976d2", // Синий цвет Material-UI
        boxShadow: "none",
      }}
    >
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
          Главная
        </Link>

        <Link
          component={RouterLink}
          to="/login"
          color="inherit"
          underline="none"
          sx={navbarStyles.link}
        >
          Вход
        </Link>

        <Link
          component={RouterLink}
          to="/register"
          color="inherit"
          underline="none"
          sx={navbarStyles.link}
        >
          Регистрация
        </Link>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;