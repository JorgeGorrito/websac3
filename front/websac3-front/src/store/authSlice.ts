import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export type UserRole = 'admin' | 'guest' | 'program lead' | 'cybersecurity_auditor';

export interface User {
  id: number;
  username: string;
  email: string;
  role: UserRole;
  permissions?: Record<string, string[]>;
}

export interface AuthState {
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  isInitialized: boolean; // New flag to track if auth has been initialized
  error: string | null;
}

const initialState: AuthState = {
  user: null,
  accessToken: null,
  refreshToken: null,
  isAuthenticated: false,
  isLoading: false,
  isInitialized: false, // Start as not initialized
  error: null,
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    loginStart: (state) => {
      state.isLoading = true;
      state.error = null;
    },
    loginSuccess: (state, action: PayloadAction<{ accessToken: string; refreshToken: string; user: User }>) => {
      console.log("🎉 Debug - loginSuccess reducer called with:", action.payload);
      state.isLoading = false;
      state.isAuthenticated = true;
      state.isInitialized = true; // Mark as initialized
      state.accessToken = action.payload.accessToken;
      state.refreshToken = action.payload.refreshToken;
      state.user = action.payload.user;
      state.error = null;
      
      console.log("🎉 Debug - State updated:", { 
        isAuthenticated: state.isAuthenticated, 
        user: state.user,
        isInitialized: state.isInitialized
      });
      
      // Store tokens in localStorage
      if (typeof window !== "undefined") {
        localStorage.setItem("access_token", action.payload.accessToken);
        localStorage.setItem("refresh_token", action.payload.refreshToken);
      }
    },
    loginFailure: (state, action: PayloadAction<string>) => {
      state.isLoading = false;
      state.isAuthenticated = false;
      state.user = null;
      state.accessToken = null;
      state.refreshToken = null;
      state.error = action.payload;
    },
    logout: (state) => {
      state.isAuthenticated = false;
      state.user = null;
      state.accessToken = null;
      state.refreshToken = null;
      state.error = null;
      
      // Clear tokens from localStorage
      if (typeof window !== "undefined") {
        localStorage.removeItem("access_token");
        localStorage.removeItem("refresh_token");
      }
    },
    clearError: (state) => {
      state.error = null;
    },
    updateTokens: (state, action: PayloadAction<{ accessToken: string; refreshToken: string }>) => {
      state.accessToken = action.payload.accessToken;
      state.refreshToken = action.payload.refreshToken;
      
      // Update tokens in localStorage
      if (typeof window !== "undefined") {
        localStorage.setItem("access_token", action.payload.accessToken);
        localStorage.setItem("refresh_token", action.payload.refreshToken);
      }
    },
    initializeAuth: (state) => {
      // Initialize auth state from localStorage on app start
      if (typeof window !== "undefined") {
        const token = localStorage.getItem("access_token");
        if (token) {
          state.accessToken = token;
          state.refreshToken = localStorage.getItem("refresh_token");
          state.isAuthenticated = true;
          // Note: We can't decode the user from token here since it's not available
          // The user will be set when they navigate to a protected route
        }
      }
      state.isInitialized = true; // Always mark as initialized
    },
    setInitialized: (state) => {
      state.isInitialized = true;
      console.log("🔧 Debug - setInitialized reducer called");
    },
  },
});

export const {
  loginStart,
  loginSuccess,
  loginFailure,
  logout,
  clearError,
  updateTokens,
  initializeAuth,
  setInitialized,
} = authSlice.actions;

export default authSlice.reducer;
