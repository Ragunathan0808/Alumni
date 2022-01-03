import React from 'react';
import { View, Text } from 'react-native';
import { Button, TextInput } from 'react-native-paper';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useDispatch } from 'react-redux';
import tailwind from 'tailwind-rn';
import { setAuth } from '../store/auth';

import { useNavigation } from '@react-navigation/native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { AuthNavParamList } from '.';

function Login() {
	const dispatch = useDispatch();

	const navigation =
		useNavigation<NativeStackScreenProps<AuthNavParamList>['navigation']>();

	const login = (email: string, password: string) => {
		const username = 'kabi';
		const token = 'dummy token';
		dispatch(setAuth({ username, email, token }));
	};

	const [email, setEmail] = React.useState<string>('');
	const [password, setPassword] = React.useState<string>('');
	return (
		<SafeAreaView>
			<View
				style={tailwind(
					'flex justify-center items-center w-full h-full'
				)}
			>
				<Text style={tailwind('text-xl font-bold text-indigo-700 p-5')}>
					Welcome Back , login here
				</Text>
				<View
					style={tailwind('flex justify-center items-center w-full')}
				>
					<TextInput
						mode={'outlined'}
						autoComplete={'email'}
						style={tailwind('w-2/3 m-3')}
						label='Email'
						value={email}
						onChangeText={(val) => setEmail(val)}
					/>
					<TextInput
						mode={'outlined'}
						autoComplete={'password'}
						style={tailwind('w-2/3')}
						label='Password'
						value={password}
						onChangeText={(val) => setPassword(val)}
					/>
				</View>
				<Button
					mode='contained'
					style={tailwind('m-3')}
					onPress={() => login(email, password)}
				>
					Log In
				</Button>
				<Text
					style={tailwind('text-blue-600')}
					onPress={() => navigation.navigate('SignUp')}
				>
					Don't have an account?{' '}
				</Text>
			</View>
		</SafeAreaView>
	);
}

export default Login;
