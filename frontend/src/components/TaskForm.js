// src/components/TaskForm.js
import React, { useState } from 'react';
import axios from 'axios';

const TaskForm = ({ onTaskAdded }) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [dueDate, setDueDate] = useState('');
    const [status, setStatus] = useState('pending');

    const handleSubmit = async (e) => {
        e.preventDefault();
        const newTask = { title, description, due_date: dueDate, status };
        
        try {
            const response = await axios.post('http://localhost:8080/tasks', newTask);
            onTaskAdded(response.data);
            setTitle('');
            setDescription('');
            setDueDate('');
            setStatus('pending');
        } catch (error) {
            console.error('Error adding task:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="p-4 border rounded shadow-md">
            <h2 className="text-lg font-bold mb-2">Add Task</h2>
            <input
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Title"
                required
                className="mb-2 p-2 border rounded w-full"
            />
            <textarea
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="Description"
                required
                className="mb-2 p-2 border rounded w-full"
            />
            <input
                type="date"
                value={dueDate}
                onChange={(e) => setDueDate(e.target.value)}
                required
                className="mb-2 p-2 border rounded w-full"
            />
            <select
                value={status}
                onChange={(e) => setStatus(e.target.value)}
                className="mb-2 p-2 border rounded w-full"
            >
                <option value="pending">Pending</option>
                <option value="completed">Completed</option>
            </select>
            <button type="submit" className="bg-blue-500 text-white rounded p-2">Add Task</button>
        </form>
    );
};

export default TaskForm;
