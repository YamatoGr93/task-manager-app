import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { deleteTask, toggleComplete } from '../redux/actions/taskActions';  // Correct the path


const TaskList = () => {
  const tasks = useSelector((state) => state.tasks.tasks);
  const dispatch = useDispatch();

  return (
    <ul>
      {tasks.map((task) => (
        <li key={task.id}>
          <input
            type="checkbox"
            checked={task.completed}
            onChange={() => dispatch(toggleComplete(task.id))}
          />
          {task.task} - {task.priority}
          <button onClick={() => dispatch(deleteTask(task.id))}>Delete</button>
        </li>
      ))}
    </ul>
  );
};

export default TaskList;
