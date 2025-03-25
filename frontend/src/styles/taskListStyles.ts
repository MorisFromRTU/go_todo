import { SxProps } from "@mui/material";

export const taskListStyles: { [key: string]: SxProps } = {
  list: {
    padding: 0,
  },
  listItem: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    padding: "8px 16px",
    borderBottom: "1px solid #ddd",
  },
  listItemText: {
    textDecoration: "none",
  },
  completed: {
    textDecoration: "line-through",
    color: "#888",
  },
  deleteButton: {
    color: "#f44336",
    "&:hover": {
      color: "#d32f2f",
    },
  },
};