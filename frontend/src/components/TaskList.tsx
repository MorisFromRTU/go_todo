import React from "react";
import { List, ListItem, Checkbox, ListItemText, IconButton } from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import { Task } from "../types/task.types";
import { updateTask, deleteTask } from "../api";
import { taskListStyles } from "../styles/taskListStyles";

interface TaskListProps {
  tasks: Task[];
  onUpdateTask: (updatedTask: Task) => void;
  onDeleteTask: (id: number) => void;
}

const TaskList: React.FC<TaskListProps> = ({ tasks, onUpdateTask, onDeleteTask }) => {
  const handleToggle = async (task: Task) => {
    const updatedTask = { ...task, done: !task.done };
    await updateTask(task.id, updatedTask);
    onUpdateTask(updatedTask);
  };

  const handleDelete = async (id: number) => {
    await deleteTask(id);
    onDeleteTask(id);
  };

  return (
    <List sx={taskListStyles.list}>
      {tasks.map((task) => (
        <ListItem key={task.id} sx={taskListStyles.listItem}>
          <Checkbox checked={task.done} onChange={() => handleToggle(task)} />
          <ListItemText
            primary={task.title}
            sx={task.done ? taskListStyles.completed : taskListStyles.listItemText}
          />
          <IconButton onClick={() => handleDelete(task.id)} sx={taskListStyles.deleteButton}>
            <DeleteIcon />
          </IconButton>
        </ListItem>
      ))}
    </List>
  );
};

export default TaskList;