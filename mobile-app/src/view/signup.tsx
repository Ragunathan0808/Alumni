import React from 'react';
import { View, Text, TextField, Button } from 'react-native-ui-lib';
import { useNavigation } from '@react-navigation/native';
import { AuthScreenProp } from './index';
import { useSelector, useDispatch } from 'react-redux';
import { signUp } from '../store/auth';
import { RootState } from '../store';

const Signup = () => {
	const navigation = useNavigation<AuthScreenProp>();
	const dispatch = useDispatch();
	const [username, setUsername] = React.useState<string>('');
	const [password, setPassword] = React.useState<string>('');
	const [email, setEmail] = React.useState<string>('');
	const auth = useSelector((state: RootState) => state.auth);
	return (
		<View flex center>
			<TextField
				placeholder='Name'
				textColor='white'
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
				onChangeText={(text: any) => setUsername(text)}
			/>
			<TextField
				placeholder='Email'
				textColor='white'
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
				textColor='white'
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
			{auth.error && <Text color='red'> {auth.error} </Text>}
			<Button
				label='Sign Up'
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
				onPress={() => dispatch(signUp(username, password, email))}
			/>
			<Text
				blue40
				center
				marginT-20
				onPress={() => navigation.navigate('signin')}
			>
				Already have an account?
			</Text>
		</View>
	);
};

export default Signup;
