import { MantineProvider, Container, Title, Button, TextInput, Table, Group, Card, createTheme, ActionIcon } from '@mantine/core';
import { useState, useEffect } from 'react';
import { Notifications } from '@mantine/notifications';
import { IconTrash, IconCheck, IconX } from '@tabler/icons-react';

const theme = createTheme({
  primaryColor: 'teal',
  colors: {
    teal: [
      '#e6fafa',
      '#b3e6e6',
      '#80d2d2',
      '#4cbfbf',
      '#19acac',
      '#139696',
      '#0e8080',
      '#096b6b',
      '#045656',
      '#004040',
    ],
  },
  components: {
    Button: {
      styles: {
        root: {
          transition: 'all 0.2s ease',
          '&:hover': {
            transform: 'translateY(-2px)',
            boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
          },
        },
      },
    },
    Table: {
      styles: {
        table: {
          backgroundColor: '#fff',
          borderRadius: '8px',
          boxShadow: '0 2px 8px rgba(0, 0, 0, 0.1)',
          '& tr:hover': {
            backgroundColor: '#f1fafa',
          },
        },
      },
    },
  },
});

function App() {
  const [tasks, setTasks] = useState([]);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  const fetchTasks = async () => {
    const response = await fetch('http://localhost:8080/api/tasks');
    const data = await response.json();
    setTasks(data);
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const addTask = async () => {
    if (!title || !description) return;
    await fetch('http://localhost:8080/api/tasks', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title, description, status: 'pending' }),
    });
    setTitle('');
    setDescription('');
    fetchTasks();
  };

  const updateTaskStatus = async (id, status) => {
    await fetch(`http://localhost:8080/api/tasks/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ status }),
    });
    fetchTasks();
  };

  const deleteTask = async (id) => {
    await fetch(`http://localhost:8080/api/tasks/${id}`, {
      method: 'DELETE',
    });
    fetchTasks();
  };

  return (
    <MantineProvider theme={theme} withGlobalStyles withNormalizeCSS>
      <Notifications />
      <Container size="lg" style={{ padding: '2rem', background: 'linear-gradient(135deg, #f0f4f8 0%, #e0e7ff 100%)', minHeight: '100vh' }}>
        <Title order={1} align="center" mb="xl" style={{ color: '#2c3e50', fontWeight: 700 }}>
          Task Manager
        </Title>
        <Card shadow="sm" padding="lg" radius="md" withBorder mb="xl">
          <Group position="apart" spacing="xs" grow>
            <TextInput
              placeholder="Task title"
              value={title}
              onChange={(e) => setTitle(e.currentTarget.value)}
              radius="md"
              size="md"
            />
            <TextInput
              placeholder="Task description"
              value={description}
              onChange={(e) => setDescription(e.currentTarget.value)}
              radius="md"
              size="md"
            />
            <Button onClick={addTask} radius="md" size="md" variant="gradient" gradient={{ from: 'teal', to: 'cyan' }}>
              Add Task
            </Button>
          </Group>
        </Card>
        <Table highlightOnHover>
          <thead>
            <tr>
              <th>Title</th>
              <th>Description</th>
              <th>Status</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {tasks.map((task) => (
              <tr key={task._id}>
                <td>{task.title}</td>
                <td>{task.description}</td>
                <td>
                  <span
                    style={{
                      padding: '4px 8px',
                      borderRadius: '12px',
                      backgroundColor: task.status === 'completed' ? '#e6ffed' : '#fff3e6',
                      color: task.status === 'completed' ? '#28a745' : '#e67e22',
                      fontSize: '12px',
                      fontWeight: 500,
                    }}
                  >
                    {task.status.charAt(0).toUpperCase() + task.status.slice(1)}
                  </span>
                </td>
                <td>
                  <Group spacing="xs">
                    <ActionIcon
                      color="teal"
                      variant="light"
                      onClick={() => updateTaskStatus(task._id, task.status === 'pending' ? 'completed' : 'pending')}
                    >
                      {task.status === 'pending' ? <IconCheck size={18} /> : <IconX size={18} />}
                    </ActionIcon>
                    <ActionIcon color="red" variant="light" onClick={() => deleteTask(task._id)}>
                      <IconTrash size={18} />
                    </ActionIcon>
                  </Group>
                </td>
              </tr>
            ))}
          </tbody>
        </Table>
      </Container>
    </MantineProvider>
  );
}

export default App;