import { View, Text } from 'react-native'
import React from 'react'
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

import login from './screen/login'
import getstart from './screen/getstart';
import signup from './screen/signup';
import profile from './screen/profile';
import home from './screen/home';
import Detail from './screen/detail';
import search from './screen/search';



const Stack = createNativeStackNavigator();

const App = () => {
  return (
    <NavigationContainer>
            <Stack.Navigator>
                <Stack.Screen name="GetStart" component={getstart} options={{headerShown : false}}/>
                <Stack.Screen name="login" component={login} options={{headerShown : false}}/>
                <Stack.Screen name="signup" component={signup} options={{headerShown : false}}/>
                <Stack.Screen name="Profile" component={profile} options={{headerShown : false}}/>
                <Stack.Screen name="Home" component={home} options={{headerShown : false}}/>
                <Stack.Screen name="Detail" component={Detail} options={{headerShown : false}}/>
                <Stack.Screen name="Search" component={search}/>
            </Stack.Navigator>
        </NavigationContainer>
  )
}

export default App