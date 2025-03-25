import { SxProps } from "@mui/material";

export const navbarStyles: { [key: string]: SxProps } = {
  appBar: {
    backgroundColor: "#1976d2",
    boxShadow: "none",
  },
  title: {
    flexGrow: 1,
    fontWeight: "bold",
    color: "#fff",
  },
  link: {
    marginLeft: 2,
    fontFamily: "'Montserrat'",
    fontWeight: "bold",
    "&:hover": {
      textDecoration: "underline",
    },
  },
};