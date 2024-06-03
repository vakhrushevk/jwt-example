import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Header from './Header/Header';



function Login() {
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (event) => {
    event.preventDefault();

    const data = {
      username: login,
      password: password
    };

    try {
      const response = await fetch('http://localhost:8088/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      });

      if (response.ok) {
        const result = await response.json();
        localStorage.setItem('authToken',result.token);
        navigate('/home')
      } else {
        console.error('Error:', response.statusText);
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <>
    <Header/>
    <form onSubmit={handleSubmit}>
      <div className='login'>
        <input
          type="text"
          placeholder='login'
          value={login}
          onChange={(e) => setLogin(e.target.value)}
        />
      </div>
      <div className='password'>
        <input
          type="text"
          placeholder='password'
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </div>
      <input type="submit" value="Submit" />
    </form>
    </>
  );
}

export default Login;
