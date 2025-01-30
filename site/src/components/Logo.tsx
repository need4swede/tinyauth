// src/components/Logo.tsx
import logo from '../img/logo.png';  // Import the PNG file

export const Logo = () => {
  return (
    <img 
      src={logo} 
      alt="NJESD Logo" 
      style={{ 
        width: '8rem', 
        height: 'auto', 
        margin: '0 auto', 
        display: 'block' 
      }} 
    />
  );
};