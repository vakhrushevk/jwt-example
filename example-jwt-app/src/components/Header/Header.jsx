import React from 'react';
import { useNavigate } from 'react-router-dom';
import './Header.css'

function Header() {
  const navigate = useNavigate();

  return (
    <div class="header">
    <div class="header-right">
      <a onClick={()=>navigate('/home')} class="active" >Home</a>
      <a onClick={()=>navigate('/login')}>Login</a>
      <a onClick={()=>navigate('/auth')}>Auth</a>
      <a onClick={() => navigate('/protected')}>Protected</a>
    </div>
  </div>
  );
}

export default Header;