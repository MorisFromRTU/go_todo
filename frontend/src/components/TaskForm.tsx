import React, { useState } from "react";
import { TextField, Button } from "@mui/material";
import { addTask } from "../api";
import { NewTask } from "../types/task.types";
import { taskFormStyles } from "../styles/taskFormStyles";

const TaskForm: React.FC<{ onAddTask: (newTask: NewTask) => void }> = ({ onAddTask }) => {
  const [title, setTitle] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) return;

    setLoading(true);
    try {
      const newTask: NewTask = {
        title,
        done: false,
        userId: 1,
      };

      const addedTask = await addTask(newTask);
      onAddTask(addedTask);
      setTitle("");
    } catch (error) {
      console.error("Ошибка при добавлении задачи:", error);
      alert("Не удалось добавить задачу. Попробуйте позже.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        label="Новая задача"
        sx={taskFormStyles.textField}
      />

      <Button
        type="submit"
        variant="contained"
        color="primary"
        sx={taskFormStyles.button}
        disabled={!title.trim() || loading}
      >
        {loading ? "Добавление..." : "Добавить"}
      </Button>
    </form>
  );
};

export default TaskForm;