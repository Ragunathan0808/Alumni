import { StatusBar } from 'expo-status-bar';
import { SafeAreaProvider } from 'react-native-safe-area-context';

import useCachedResources from './src/hooks/useCachedResources';

import RootView from './src/view';
import { Provider } from 'react-redux';
import { store } from './src/store';

import { Provider as PaperProvider, DefaultTheme } from 'react-native-paper';

export default function App() {
	const isLoadingComplete = useCachedResources();

	if (!isLoadingComplete) {
		return null;
	} else {
		return (
			<Provider store={store}>
				<PaperProvider theme={DefaultTheme}>
					<SafeAreaProvider>
						<RootView />
						<StatusBar />
					</SafeAreaProvider>
				</PaperProvider>
			</Provider>
		);
	}
}
