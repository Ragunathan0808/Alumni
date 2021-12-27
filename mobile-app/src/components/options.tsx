import React from 'react';
import { Modal, Pressable } from 'react-native';

interface OptionsProps {
	children: React.ReactNode;
	onClose: () => void;
	visible: boolean;
}

const Options = ({ visible = false, onClose, children }: OptionsProps) => {
	return (
		<Modal transparent visible={visible}>
			<Pressable style={{ flex: 1 }} onPress={onClose}>
				{children}
			</Pressable>
		</Modal>
	);
};

export default Options;
