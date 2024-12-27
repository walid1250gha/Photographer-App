import { View, Text, ActivityIndicator, Image, StyleSheet, Pressable } from 'react-native';
import React, { useState, useEffect } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';

const Search = ({ route, navigation }) => {
  const { token } = route.params;  // รับ token จาก route.params
  const [user, setUser] = useState({});
  const [isLoading, setIsLoading] = useState(true);

  const fetchUser = async () => {
    try {
      if (!token) {
        alert('Token not found. Please log in again.');
        return;
      }

      const response = await fetch('http://192.168.1.5:8080/users/profile', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      });

      const data = await response.json();
      console.log(data);  // ตรวจสอบโครงสร้างข้อมูลที่ได้จาก API

      if (data.status === 'ok') {
        setUser(data.userId);  // เก็บข้อมูลผู้ใช้ใน state
      } else {
        alert('Failed to fetch user data');
      }
    } catch (error) {
      console.error('Error fetching user data:', error);
      alert('Error fetching user data');
    } finally {
      setIsLoading(false); // เปลี่ยนสถานะการโหลดเมื่อเรียบร้อยแล้ว
    }
  };

  useEffect(() => {
    fetchUser();
  }, [token]);  // ใช้ token เป็น dependency

  return (
    <View>
      {isLoading ? (
        <Text>Loading...</Text>
      ) : (
        <View>
          <Text>{user.Fullname || 'No Fullname Available'}</Text>
          <Text>{user.Lastname || 'No Lastname Available'}</Text>
          <Text>{user.Email || 'No Email Available'}</Text>
          <Text>{user.Username || 'No Username Available'}</Text>
          <Text>{user.Password || 'No Password Available'}</Text>
          <Image source={{ uri: user.Img_profile }} style={styles.profileImage} />
        </View>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  profileImage: {
    width: 250,   // Adjust the width as needed
    height: 250,  // Adjust the height as needed
    marginTop: 10,
    borderRadius: 10, // Optional: Adds border radius to make the image circular or rounded
  }, 
  container: {
    position: 'absolute',  
    bottom: 0,          
    left: 0,              
    margin: 16,           
  },
  image: {
    width: 100,         
    height: 100,           
  },
});

export default Search;
