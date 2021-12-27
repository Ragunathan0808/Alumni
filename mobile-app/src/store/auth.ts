import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface AuthState {
	username?: string;
	token?: string;
	email?: string;
}

const initialState: AuthState = {};

export const counterSlice = createSlice({
	name: 'auth',
	initialState,
	reducers: {
		setAuth: (state, action: PayloadAction<AuthState>) => {
			state.username = action.payload.username;
			state.token = action.payload.token;
			state.email = action.payload.email;
		},
		clearAuth: (state) => {
			state.username = undefined;
			state.token = undefined;
			state.email = undefined;
		},
	},
});

// Action creators are generated for each case reducer function
export const { setAuth, clearAuth } = counterSlice.actions;

export default counterSlice.reducer;
