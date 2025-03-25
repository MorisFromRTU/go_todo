import { useEffect, useState } from "react";
import { Box, Typography, Paper} from "@mui/material";
import TaskForm from "../components/TaskForm";
import TaskList from "../components/TaskList";
import { fetchTasks, addTask, updateTask, deleteTask } from "../api";
import { Task, NewTask } from "../types/task.types";

const Home = () => {
  const [tasks, setTasks] = useState<Task[]>([]);

  useEffect(() => {
    const loadTasks = async () => {
      const data = await fetchTasks();
      setTasks(data);
    };
    loadTasks();
  }, []);

  const handleAddTask = async (newTask: NewTask) => {
    const addedTask = await addTask(newTask);
    setTasks((prevTasks) => [...prevTasks, addedTask]);
  };

  const handleUpdateTask = async (updatedTask: Task) => {
    await updateTask(updatedTask.id, updatedTask);
    setTasks((prevTasks) =>
      prevTasks.map((task) => (task.id === updatedTask.id ? updatedTask : task))
    );
  };

  const handleDeleteTask = async (id: number) => {
    await deleteTask(id);
    setTasks((prevTasks) => prevTasks.filter((task) => task.id !== id));
  };

  return (
    <Box
      sx={{
        padding: { xs: "16px", sm: "24px" },
        minHeight: "100vh",
        backgroundColor: "#f5f5f5",
      }}
    >
      {/* Заголовок */}
      <Typography
        variant="h4"
        sx={{
          fontWeight: "bold",
          marginBottom: "24px",
          color: "#333",
          textAlign: "center",
        }}
      >
        Список задач
      </Typography>

      {/* Форма добавления задач */}
      <Paper
        elevation={3}
        sx={{
          padding: "24px",
          marginBottom: "24px",
          borderRadius: "12px",
          maxWidth: "600px",
          margin: "0 auto",
        }}
      >
        <TaskForm onAddTask={handleAddTask} />
      </Paper>

      {/* Список задач */}
      <Paper
        elevation={3}
        sx={{
          padding: "24px",
          borderRadius: "12px",
          maxWidth: "600px",
          margin: "0 auto",
        }}
      >
        <TaskList
          tasks={tasks}
          onUpdateTask={handleUpdateTask}
          onDeleteTask={handleDeleteTask}
        />
      </Paper>
    </Box>
  );
};

export default Home;