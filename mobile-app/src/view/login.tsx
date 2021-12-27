import { useNavigation } from '@react-navigation/native';
import * as React from 'react';
import { Button, Text, TextField, View } from 'react-native-ui-lib';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '../store';

import { AuthStackParamList } from '.';
import { setAuth } from '../store/auth';

function SignIn() {
	const navigation = useNavigation<AuthStackParamList>();
	const dispatch = useDispatch();
	const auth = useSelector((state: RootState) => state.auth);
	const [password, setPassword] = React.useState<string>('');
	const [email, setEmail] = React.useState<string>('');
	const signIn = (email: string, password: string) => {
		dispatch(setAuth({ email, username: 'kabi' }));
	};
	return (
		<View flex center>
			<TextField
				placeholder='Email'
				textColorii='white'
				tintColor='white'
				textInputStyle={{
					color: 'white',
					fontSize: 16,
					fontWeight: 'bold',
					padding: 10,
					borderBottomWidth: 1,
					borderBottomColor: 'white',
				}}
				width={300}
				onChangeText={(text: any) => setEmail(text)}
			/>
			<TextField
				placeholder='Password'
				textColorii='white'
				tintColor='white'
				textInputStyle={{
					color: 'white',
					fontSize: 16,
					fontWeight: 'bold',
					padding: 10,
					borderBottomWidth: 1,
					borderBottomColor: 'white',
				}}
				width={300}
				onChangeText={(text: any) => setPassword(text)}
			/>
			{auth.error && <Text text50>{auth.error}</Text>}
			{auth.username && <Text text50>{auth.username}</Text>}
			<Button
				label='Sign In'
				labelStyle={{
					color: 'white',
					fontSize: 16,
					fontWeight: 'bold',
				}}
				style={{
					backgroundColor: '#00bcd4',
					borderRadius: 10,
					marginTop: 20,
				}}
				onPress={() => {
					dispatch(signIn(email, password));
				}}
			/>
			<Text
				blue40
				center
				marginT-20
				onPress={() => {
					navigation.navigate('signup');
				}}
			>
				Don't have an account? Sign up
			</Text>
		</View>
	);
}

export default SignIn;
