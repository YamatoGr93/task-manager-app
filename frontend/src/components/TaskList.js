// src/components/TaskList.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';

const TaskList = () => {
    const [tasks, setTasks] = useState([]);

    const fetchTasks = async () => {
        try {
            const response = await axios.get('http://localhost:8080/tasks');
            setTasks(response.data);
        } catch (error) {
            console.error('Error fetching tasks:', error);
        }
    };

    useEffect(() => {
        fetchTasks();
    }, []);

    return (
        <div className="p-4">
            <h2 className="text-lg font-bold mb-2">Tasks</h2>
            <ul>
                {tasks.map(task => (
                    <li key={task.id} className="border-b p-2">
                        <h3 className="font-semibold">{task.title}</h3>
                        <p>{task.description}</p>
                        <p>Due Date: {task.due_date}</p>
                        <p>Status: {task.status}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TaskList;

