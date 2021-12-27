import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import _ from 'lodash';
import { InMemoryFriendDB } from '../data/friend.db';
import { IUser } from '../model';
import { IFriendDB } from '../model/friend.db';

export interface FriendsState {
	friends: Array<IUser>;
	error?: string;
}

const initialState: FriendsState = {
	friends: [],
};

export const friendsSlice = createSlice({
	name: 'friends',
	initialState,
	reducers: {
		setError: (state, action: PayloadAction<string>) => {
			state.error = action.payload;
		},
		setFriends: (
			state,
			action: PayloadAction<{ friends: Array<IUser> }>
		) => {
			state.friends = action.payload.friends;
		},
		addFriend: (state, action: PayloadAction<{ newFriend: IUser }>) => {
			const { newFriend } = action.payload;
			const userIndex = _.findIndex(
				state.friends,
				(user: IUser) =>
					user.id === newFriend.id || user.email === newFriend.email
			);
			if (userIndex != -1) return;
			state.friends.push(action.payload.newFriend);
		},
		removeFriend: (state, action: PayloadAction<{ friend: IUser }>) => {
			const { friend } = action.payload;
			state.friends = _.remove(
				state.friends,
				(user: IUser) =>
					user.id === friend.id || user.email === friend.email
			);
		},
	},
});

const { addFriend, removeFriend, setFriends, setError } = friendsSlice.actions;

function friendAction(friendDB: IFriendDB) {
	const loadAllFriendsAsync = () => async (dispatch: Function) => {
		let { friends, error } = await friendDB.getFriends();
		if (error) {
			return dispatch(setError(error.message));
		}
		friends ||= [];
		dispatch(setFriends({ friends }));
	};
	const addFriendAsync = (newFriend: IUser) => async (dispatch: Function) => {
		let { error } = await friendDB.addFriend(newFriend);
		if (error) {
			return dispatch(setError(error.message));
		}
		dispatch(addFriend({ newFriend }));
	};
	const removeFriendAsync = (friend: IUser) => async (dispatch: Function) => {
		let { error } = await friendDB.removeFriend(friend);
		if (error) {
			return dispatch(setError(error.message));
		}
		dispatch(removeFriend({ friend }));
	};
	return { loadAllFriendsAsync, addFriendAsync, removeFriendAsync };
}

export const { loadAllFriendsAsync, addFriendAsync, removeFriendAsync } =
	friendAction(new InMemoryFriendDB());

export default friendsSlice.reducer;
