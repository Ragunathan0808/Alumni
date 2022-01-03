import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import {
	createNativeStackNavigator,
	NativeStackScreenProps,
} from '@react-navigation/native-stack';
import Login from './login';
import { useSelector } from 'react-redux';
import { RootState } from '../store';
import SignUp from './signup';
import Home from './home';
import { FontAwesome } from '@expo/vector-icons';
import Search from './search';

export type AppNavParamList = {
	Home: undefined;
	Search: undefined;
};
export type AuthNavParamList = {
	LogIn: undefined;
	SignUp: undefined;
};

export type AuthScreenProp = NativeStackScreenProps<AuthNavParamList>;

const AppNav = createBottomTabNavigator<AppNavParamList>();
const AuthNav = createNativeStackNavigator<AuthNavParamList>();

const RootView = () => {
	const authStore = useSelector((state: RootState) => state.auth);
	console.log('authState: ', authStore);
	const isLoggedIn = authStore.username ? true : false;
	if (isLoggedIn) {
		return (
			<NavigationContainer>
				<AppNav.Navigator
					initialRouteName='Home'
					screenOptions={{
						headerShown: false,
						tabBarShowLabel: false,
					}}
				>
					<AppNav.Screen
						name='Home'
						component={Home}
						options={{
							tabBarIcon: ({ color, size }) => (
								<FontAwesome
									name='home'
									size={size}
									color={color}
								/>
							),
						}}
					/>
					<AppNav.Screen
						name='Search'
						component={Search}
						options={{
							tabBarIcon: ({ color, size }) => (
								<FontAwesome
									name='search'
									size={size}
									color={color}
								/>
							),
						}}
					/>
				</AppNav.Navigator>
			</NavigationContainer>
		);
	}
	return (
		<NavigationContainer>
			<AuthNav.Navigator
				initialRouteName='LogIn'
				screenOptions={{ headerShown: false }}
			>
				<AuthNav.Screen name='LogIn' component={Login} />
				<AuthNav.Screen name='SignUp' component={SignUp} />
			</AuthNav.Navigator>
		</NavigationContainer>
	);
};

export default RootView;
