import { SxProps} from "@mui/material";

  export const taskFormStyles: { [key: string]: SxProps } = {
  form: {
    accentColor: "#1976d2",
    display: "flex",
    flexDirection: "column",
    gap: 2,
    marginBottom: 3,
  },
  textField: {
    width: "100%",
    "& .Mui-focused": {
      borderColor: "#1976d2",
    },
    "& .MuiInputBase-root": {
      accentColor: "#1976d2",
    },
  },
  button: {
    margin: "12px 0",
    alignSelf: "flex-start",
    backgroundColor: "#1976d2",
    "&:hover": {
      backgroundColor: "#1565c0",
    },
  },
};