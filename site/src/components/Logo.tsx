// src/components/Logo.tsx
import logo from '../img/logo.svg';

export const Logo = () => {
  return (
    <img 
      src={logo} 
      alt="Minas Tirith Logo" 
      style={{ 
        width: '8rem', 
        height: 'auto', 
        margin: '0 auto', 
        display: 'block' 
      }} 
    />
  );
};