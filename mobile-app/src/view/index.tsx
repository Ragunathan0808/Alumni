import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { Provider, useSelector } from 'react-redux';

import Login from './login';
import SignUp from './signup';
import Home from './home';
import { RootState, store } from '../store';

export type AuthStackParamList = {
	Login: undefined;
	SignUp: undefined;
};

export type UserStackParamList = {
	Home: undefined;
};

export const AuthStack = createNativeStackNavigator<AuthStackParamList>();
export const UserStack = createNativeStackNavigator<UserStackParamList>();

const Views = () => {
	const authStore = useSelector((store: RootState) => store.auth);
	const isUserLoggedIn = authStore.username !== undefined;
	return (
		<Provider store={store}>
			<NavigationContainer>
				{!isUserLoggedIn ? (
					<AuthStack.Navigator initialRouteName='Login'>
						<AuthStack.Screen name='Login' component={Login} />
						<AuthStack.Screen name='SignUp' component={SignUp} />
					</AuthStack.Navigator>
				) : (
					<UserStack.Navigator initialRouteName='Home'>
						<UserStack.Screen name='Home' component={Home} />
					</UserStack.Navigator>
				)}
			</NavigationContainer>
		</Provider>
	);
};

export default Views;
