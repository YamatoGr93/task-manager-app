import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { addTask } from '../redux/actions/taskActions';  // Correct the path


const TaskForm = () => {
  const [task, setTask] = useState('');
  const [priority, setPriority] = useState('medium');
  const dispatch = useDispatch();

  const handleSubmit = (e) => {
    e.preventDefault();
    dispatch(addTask({ id: Date.now(), task, priority, completed: false }));
    setTask('');
    setPriority('medium');
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        placeholder="Add a new task"
        value={task}
        onChange={(e) => setTask(e.target.value)}
        required
      />
      <select value={priority} onChange={(e) => setPriority(e.target.value)}>
        <option value="high">High</option>
        <option value="medium">Medium</option>
        <option value="low">Low</option>
      </select>
      <button type="submit">Add Task</button>
    </form>
  );
};

export default TaskForm;
