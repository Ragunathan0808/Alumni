import React from 'react';
import { View, Text } from 'react-native';
import { Avatar } from 'react-native-paper';
import { batch } from 'react-redux';
import tailwind from 'tailwind-rn';

export default function User({
	img,
	name,
	studId,
	dept,
	batch,
}: {
	img: string;
	name: string;
	studId: string;
	dept: string;
	batch: string;
}) {
	return (
		<View style={tailwind('flex flex-row w-full p-3')}>
			<Avatar.Image
				size={80}
				source={{
					uri: img,
				}}
			/>
			<View style={tailwind('flex justify-center px-3')}>
				<Text style={tailwind('text-xl font-bold')}>{name}</Text>
				<Text style={tailwind('text-xl font-bold')}>{dept}</Text>
				<Text style={tailwind('text-xl font-bold')}>{batch}</Text>
			</View>
		</View>
	);
}
