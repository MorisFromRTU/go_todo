import axios from "axios";
import { Task, NewTask } from "./types/task.types";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

export const fetchTasks = async (): Promise<Task[]> => {
  const response = await axios.get<Task[]>(`${API_URL}/tasks`);
  return response.data;
};

export const addTask = async (task: NewTask): Promise<Task> => {
  const response = await axios.post<Task>(`${API_URL}/tasks`, task);
  return response.data;
};


export const updateTask = async (id: number, updatedTask: Partial<Task>): Promise<Task> => {
  const response = await axios.put<Task>(`${API_URL}/tasks/${id}`, updatedTask);
  return response.data;
};


export const deleteTask = async (id: number): Promise<void> => {
  await axios.delete(`${API_URL}/tasks/${id}`);
};