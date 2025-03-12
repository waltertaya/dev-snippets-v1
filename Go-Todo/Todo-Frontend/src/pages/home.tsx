import React, { useState } from "react";

interface Task {
  id: number;
  title: string;
  complete: boolean;
}

const ManageTasks: React.FC<{ username: string; initialTasks: Task[] }> = ({
  username,
  initialTasks,
}) => {
  const [tasks, setTasks] = useState<Task[]>(initialTasks);
  const [newTask, setNewTask] = useState("");

  const handleAddTask = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newTask.trim()) return;
    const newTaskObj: Task = {
      id: tasks.length + 1,
      title: newTask,
      complete: false,
    };
    setTasks([...tasks, newTaskObj]);
    setNewTask("");
  };

  const toggleTaskStatus = (taskId: number) => {
    setTasks(
      tasks.map((task) =>
        task.id === taskId ? { ...task, complete: !task.complete } : task
      )
    );
  };

  const deleteTask = (taskId: number) => {
    setTasks(tasks.filter((task) => task.id !== taskId));
  };

  return (
    <div className="flex flex-col items-center bg-gray-200 p-5 min-h-screen">
      <div className="bg-white p-5 rounded-lg shadow-md w-full max-w-xl">
        <h1 className="text-center text-2xl font-bold mb-5">My ToDo List</h1>
        <div className="flex justify-between items-center mb-5">
          <h4 className="text-lg">Hi, {username}</h4>
          <a href="/logout">
            <button className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600">LogOut</button>
          </a>
        </div>
        <form onSubmit={handleAddTask} className="mb-5">
          <div className="flex">
            <input
              type="text"
              id="new-task"
              placeholder="Add a new task"
              name="task"
              value={newTask}
              onChange={(e) => setNewTask(e.target.value)}
              className="flex-1 p-2 border border-gray-300 rounded"
            />
            <button type="submit" className="ml-2 bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">Add</button>
          </div>
        </form>
        <table className="w-full border-collapse mt-5">
          <thead>
            <tr className="bg-gray-100">
              <th className="p-2 text-left">Task</th>
              <th className="p-2 text-left">Status</th>
              <th className="p-2 text-left">Action</th>
            </tr>
          </thead>
          <tbody>
            {tasks.map((task) => (
              <tr key={task.id}>
                <td className="p-2">{task.title}</td>
                <td className="p-2">
                  <span className={task.complete ? "text-green-500" : "text-red-500"}>
                    {task.complete ? "Completed" : "Pending"}
                  </span>
                </td>
                <td className="p-2">
                  <button
                    className="task-actions px-2 py-1 rounded bg-blue-500 text-white hover:bg-blue-600"
                    onClick={() => toggleTaskStatus(task.id)}
                  >
                    {task.complete ? "Done" : "In Progress"}
                  </button>
                  <button
                    className="task-actions px-2 py-1 rounded bg-red-500 text-white hover:bg-red-600 ml-2"
                    onClick={() => deleteTask(task.id)}
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ManageTasks;
