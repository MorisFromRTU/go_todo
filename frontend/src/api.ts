import axios from "axios";
import { Task, NewTask, AuthForm } from "./types/task.types";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

// AUTH
const AUTH_SERVICE = "auth"

export const register = async (registerForm: AuthForm): Promise<void> => {
  await axios.post<AuthForm>(`${API_URL}/${AUTH_SERVICE}/register`, registerForm)
}

export const login = async (loginForm: AuthForm): Promise<void> => {
  await axios.post<AuthForm>(`${API_URL}/${AUTH_SERVICE}/login`, loginForm)
}
// TASKS
const TASKS_SERVICE = "tasks"

export const fetchTasks = async (): Promise<Task[]> => {
  const response = await axios.get<Task[]>(`${API_URL}/${TASKS_SERVICE}`);
  console.log(`${API_URL}/tasks`);
  return response.data;
};

export const addTask = async (task: NewTask): Promise<Task> => {
  const response = await axios.post<Task>(`${API_URL}/${TASKS_SERVICE}`, task);
  return response.data;
};


export const updateTask = async (id: number, updatedTask: Partial<Task>): Promise<Task> => {
  const response = await axios.put<Task>(`${API_URL}/${TASKS_SERVICE}/${id}`, updatedTask);
  return response.data;
};


export const deleteTask = async (id: number): Promise<void> => {
  await axios.delete(`${API_URL}/${TASKS_SERVICE}/${id}`);
};

