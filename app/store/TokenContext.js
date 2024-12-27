// TokenContext.js
import React, { createContext, useState, useEffect, useContext } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';

const TokenContext = createContext();

// Custom hook เพื่อดึง token จาก context
export const useToken = () => useContext(TokenContext);

// TokenProvider สำหรับจัดการ token
export const TokenProvider = ({ children }) => {
  const [token, setToken] = useState(null);

  // อ่าน token จาก AsyncStorage เมื่อเริ่มต้น
  useEffect(() => {
    const fetchToken = async () => {
      const storedToken = await AsyncStorage.getItem('@token');
      setToken(storedToken);
    };
    fetchToken();
  }, []);

  // ฟังก์ชันสำหรับตั้งค่าและลบ token
  const setAuthToken = async (newToken) => {
    setToken(newToken);
    if (newToken) {
      await AsyncStorage.setItem('@token', newToken); // เก็บ token ใน AsyncStorage
    } else {
      await AsyncStorage.removeItem('@token'); // ลบ token
    }
  };

  return (
    <TokenContext.Provider value={{ token, setAuthToken }}>
      {children}
    </TokenContext.Provider>
  );
};
