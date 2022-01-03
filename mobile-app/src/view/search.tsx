import React from 'react';
import { View, Text } from 'react-native';
import tailwind from 'tailwind-rn';

import { SafeAreaView } from 'react-native-safe-area-context';
import { Searchbar } from 'react-native-paper';
import User from '../components/user';

const users = [
	{
		img: 'https://avatars.githubusercontent.com/u/67632223?s=96&v=4',
		name: 'Kabilan M',
		studId: '191EC175',
		dept: 'ECE',
		batch: '2011-2015',
	},
	{
		img: 'https://avatars.githubusercontent.com/u/67632223?s=96&v=4',
		name: 'Subash T',
		studId: '191IT118',
		dept: 'IT',
		batch: '2011-2015',
	},
];

const Search = () => {
	const [query, setQuery] = React.useState<string>('');
	return (
		<SafeAreaView>
			<View style={tailwind('flex  w-full h-full')}>
				<Searchbar
					autoComplete={''}
					placeholder='user email'
					value={query}
					onChangeText={(val) => setQuery(val)}
				/>
				{users.map((user, id) => (
					<User key={id} {...user} />
				))}
			</View>
		</SafeAreaView>
	);
};

export default Search;
