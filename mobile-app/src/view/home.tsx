import React from 'react';
import { View, Text } from 'react-native';
import tailwind from 'tailwind-rn';
import { Avatar } from 'react-native-paper';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useSelector } from 'react-redux';
import { RootState } from '../store';

const Home = () => {
	const auth = useSelector((state: RootState) => state.auth);
	const name = 'Kabilan M';
	const studId = '191EC175';
	const dept = 'B.E. ECE';
	const batch = '2011-20015';
	const mobile = '+91 9876543210';
	const email = auth.email ? auth.email : 'kabilan.ec11@bitsathy.ac.in';
	const jobTitle = 'Full Stack Developer';
	return (
		<SafeAreaView>
			<View style={tailwind('flex w-full h-full')}>
				<View
					style={tailwind(
						'flex flex-row justify-center bg-blue-300 w-full p-8'
					)}
				>
					<Avatar.Image
						size={100}
						source={{
							uri: 'https://avatars.githubusercontent.com/u/67632223?s=96&v=4',
						}}
					/>
					<View style={tailwind('flex justify-center px-3')}>
						<Text style={tailwind('text-3xl font-bold')}>
							{name}
						</Text>
						<Text style={tailwind('text-2xl font-bold')}>
							{studId}
						</Text>
					</View>
				</View>
				<View style={tailwind('p-3')}>
					<Text style={tailwind('text-xl py-3')}>Dept: {dept}</Text>
					<Text style={tailwind('text-xl py-3')}>Batch: {batch}</Text>
					<Text style={tailwind('text-xl py-3')}>
						Mobile: {mobile}
					</Text>
					<Text style={tailwind('text-xl py-3')}>
						E-Mail ID:{email}
					</Text>
					<Text style={tailwind('text-xl py-3')}>
						Job Title: {jobTitle}
					</Text>
				</View>
			</View>
		</SafeAreaView>
	);
};

export default Home;
