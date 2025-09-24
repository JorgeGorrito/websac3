import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export type ErrorType = "error" | "not_found" | "success";

type ErrorState = {
  isOpen: boolean;
  type: ErrorType;
  title?: string;
  message: string;
  errors: string[];
  statusCode?: number;
  searchTerm?: string;
  onRetry?: () => void;
};

const initialState: ErrorState = {
  isOpen: false,
  type: "error",
  message: "",
  errors: [],
};

export const errorSlice = createSlice({
  name: "error",
  initialState,
  reducers: {
    showError: (state, action: PayloadAction<{ 
      type: ErrorType;
      message: string; 
      errors?: string[]; 
      title?: string;
      statusCode?: number;
      searchTerm?: string;
      onRetry?: () => void;
    }>) => {
      state.isOpen = true;
      state.type = action.payload.type;
      state.message = action.payload.message;
      state.errors = action.payload.errors || [];
      state.title = action.payload.title;
      state.statusCode = action.payload.statusCode;
      state.searchTerm = action.payload.searchTerm;
      state.onRetry = action.payload.onRetry;
    },
    showSuccess: (state, action: PayloadAction<{ 
      message: string; 
      title?: string;
    }>) => {
      state.isOpen = true;
      state.type = "success";
      state.message = action.payload.message;
      state.title = action.payload.title || "Éxito";
      state.errors = [];
      state.statusCode = undefined;
      state.searchTerm = undefined;
      state.onRetry = undefined;
    },
    hideError: (state) => {
      state.isOpen = false;
      state.type = "error";
      state.message = "";
      state.errors = [];
      state.title = undefined;
      state.statusCode = undefined;
      state.searchTerm = undefined;
      state.onRetry = undefined;
    },
  },
});

export const { showError, showSuccess, hideError } = errorSlice.actions;
export default errorSlice.reducer;
