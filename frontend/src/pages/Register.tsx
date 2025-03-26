import React from "react";
import AuthForm from "../components/AuthForm";
import { register } from "../api";

const Register: React.FC = () => {
  const handleSubmit = async (username: string, password: string) => {
    await register({username, password});
    window.location.href = "/login";
  };

  return <AuthForm title="Регистрация" buttonText="Зарегистрироваться" onSubmit={handleSubmit} />;
};

export default Register;