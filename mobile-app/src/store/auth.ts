import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface AuthState {
	username?: string;
	email?: string;
	token?: string;
}

const initialState: AuthState = {};

export const authSlice = createSlice({
	name: 'auth',
	initialState,
	reducers: {
		setAuth: (state, action: PayloadAction<AuthState>) => {
			state.username = action.payload.username;
			state.email = action.payload.email;
			state.token = action.payload.token;
		},
		clearAuth: (state) => {
			state.username = undefined;
			state.email = undefined;
			state.token = undefined;
		},
	},
});

// Action creators are generated for each case reducer function
export const { setAuth, clearAuth } = authSlice.actions;

export default authSlice.reducer;
