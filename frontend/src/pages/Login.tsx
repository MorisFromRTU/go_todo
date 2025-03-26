import React from "react";
import AuthForm from "../components/AuthForm";
import { login } from "../api";

const Login: React.FC = () => {
  const handleSubmit = async (username: string, password: string) => {
    await login({username, password});
    window.location.href = "/";
  };

  return <AuthForm title="Вход" buttonText="Войти" onSubmit={handleSubmit} />;
};

export default Login;