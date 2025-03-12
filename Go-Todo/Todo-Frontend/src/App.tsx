import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import ManageTasks from './pages/home';
import LoginForm from './pages/login';
import RegistrationForm from './pages/register';
import LandingPage from './pages/landing';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/register" element={<RegistrationForm />} />
        <Route path="/login" element={<LoginForm />} />
        <Route path="/home" element={<ManageTasks username="John Doe" initialTasks={[]} />} />
      </Routes>
    </Router>
  );
}

export default App;