import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import App from './App';
import Home from './Home'; // Импортируем ваш компонент Home
import Login from './components/Login'
import Auth from './components/Auth/Auth';
import Protected from './components/Protected/Protected';
import ProtectedPage from './components/Protected/ProtectedPage/ProtectedPage';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <Router>
    <Routes>
      <Route path="/" element={<App />} />
      <Route path="/login" element={<Login />} />
      <Route path="/home" element={<Home />} />
      <Route path='/auth' element={<Auth/>}/>
      <Route 
          path="/protected" 
          element={
            <Protected>
              <ProtectedPage />
            </Protected>
          } 
        />
    </Routes>
  </Router>
);
