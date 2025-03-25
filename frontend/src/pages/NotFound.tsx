import React from "react";
import { Typography } from "@mui/material";

const NotFound: React.FC = () => {
  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <Typography variant="h4">Страница не найдена</Typography>
      <Typography variant="body1">Проверьте URL или вернитесь на <a href="/">главную страницу</a>.</Typography>
    </div>
  );
};

export default NotFound;