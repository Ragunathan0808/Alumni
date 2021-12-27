import { StatusBar } from 'expo-status-bar';
import { View } from 'react-native';
import Views from './src/view';

export default function App() {
	return (
		<View>
			<StatusBar style='auto' />
			<Views />
		</View>
	);
}
