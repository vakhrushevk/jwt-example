import React, { useState } from 'react';
import './App.css';
import { useNavigate } from 'react-router-dom';
import Header from './components/Header/Header';



function App() {
  const navigate = useNavigate();
  //navigate('/login')



  return (
    <>
    <Header />
      
    </>
  );
}

export default App;
