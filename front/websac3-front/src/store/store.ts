import { configureStore } from "@reduxjs/toolkit";
import { api } from "@/services/api";
import authReducer from "./authSlice";
import errorReducer from "./errorSlice";

export const store = configureStore({
  reducer: {
    [api.reducerPath]: api.reducer,
    auth: authReducer,
    error: errorReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        // Ignore these action types because they contain Blob data (CSV files)
        ignoredActions: [
          'api/executeQuery/fulfilled',
          'api/executeQuery/pending',
        ],
        // Ignore these paths in the state because they contain Blob data
        ignoredPaths: [
          'api.queries.getDegreeProgramTemplate',
          'api.queries.getCourseTemplate',
          'api.queries.getCourseTopicTemplate',
        ],
      },
    }).concat(api.middleware),
});

export type AppStore = typeof store;
export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;


