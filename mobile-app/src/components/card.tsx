import React from 'react';
import { View, Text, Image } from 'react-native-ui-lib';
import { createAvatar } from '@dicebear/avatars';
import * as style from '@dicebear/avatars-avataaars-sprites';
import { SvgXml } from 'react-native-svg';
import { trim } from '../../service/util';
import { Pressable } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import { HomeScreenProp } from '..';

interface CardProps {
	username: string;
}

const ChatCard = ({ username }: CardProps) => {
	const avatart = createAvatar(style, {
		seed: username,
		radius: 50,
		backgroundColor: '#c5e8e5',
	});
	const navigation = useNavigation<HomeScreenProp>();

	const message = trim('hello, how are you doing today buddy?');
	return (
		<Pressable onPress={() => navigation.navigate('chat', { username })}>
			<View
				width='100%'
				height={70}
				bg-white
				paddingV-8
				flex-0
				row
				style={{ borderBottomWidth: 0.5, borderColor: 'gray' }}
			>
				<SvgXml
					width={50}
					height={50}
					style={{ borderRadius: 25 }}
					xml={avatart}
				/>
				<View paddingL-12>
					<Text text60>{username}</Text>
					<Text text70>{message}</Text>
				</View>
			</View>
		</Pressable>
	);
};

export default ChatCard;
