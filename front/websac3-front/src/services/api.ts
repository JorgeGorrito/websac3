import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export type ApiResponse<T> = {
  errors?: string[];
  httpStatusCode: number;
  result: T;
};

export type PaginatedPage<T> = {
  current_page: number;
  data: T[];
  items_per_page: number;
  total_count: number;
};

export type AccessRequestItem = {
  id: number;
  name: string;
  lastname: string;
  email: string;
  identification_number: string;
  identification_type: string;
  job_position: string;
  higher_education_institution_name: string;
  higher_education_institution_ownership: string;
  higher_education_institution_snies: number;
  municipality_name: string;
  department_name: string;
  status_id: number;
  status_name: string;
};

export type CreatePersonRequest = {
  email: string;
  higher_education_institution_snies: number;
  identification_number: string;
  identification_type_id: number;
  job_position: string;
  lastname: string;
  name: string;
};

export type CreateAccessRequestRequest = {
  person: CreatePersonRequest;
  register_user_url: string;
  validation_email_url: string;
};

export type ValidateEmailRequest = {
  validation_token: string;
};

export type LoginRequest = {
  email: string;
  password: string;
};

export type LoginResponse = {
  access_token: string;
  refresh_token: string;
};

export type RefreshTokenRequest = {
  refresh_token: string;
};

export type RefreshTokenResponse = {
  access_token: string;
  refresh_token: string;
};

export type CreateUserFromTokenRequest = {
  password: string;
  confirm_password: string;
  create_user_token: string;
};

// Duration Unit types
export type DurationUnitItem = {
  id: number;
  name: string;
  description?: string;
};

// Degree Program types
export type DegreeProgramItem = {
  id: number;
  name: string;
  snies: number;
  total_credits: number;
  duration_value: number;
  duration_unit: {
    id: number;
    name: string;
  };
  program_focus: string;
  entry_profile: string;
  graduate_profile: string;
  professional_profile: string;
  created_by: number;
  higher_education_institution: {
    snies: number;
    name: string;
  };
};

export type CreateDegreeProgramRequest = {
  duration_unit_id: number;
  duration_value: number;
  entry_profile: string;
  graduate_profile: string;
  name: string;
  professional_profile: string;
  program_focus: string;
  snies: number;
  total_credits: number;
};

export type CreateUserFromTokenResponse = {
  message: string;
};

export type IdentificationTypeItem = {
  id: number;
  name: string;
};

export type HigherEducationInstitutionItem = {
  snies: number;
  name: string;
};

// Custom base query with automatic token refresh
const baseQueryWithReauth = async (args: any, api: any, extraOptions: any) => {
  let result = await fetchBaseQuery({
    baseUrl:
      process.env.NEXT_PUBLIC_API_BASE_URL?.replace(/\/$/, "") ||
      "http://localhost:8110/api/v1/es",
    credentials: "include",
    prepareHeaders: (headers) => {
      // Attach bearer token if present
      if (typeof window !== "undefined") {
        const token = window.localStorage?.getItem("access_token");
        if (token) headers.set("Authorization", `Bearer ${token}`);
      }
      headers.set("Accept", "application/json");
      return headers;
    },
  })(args, api, extraOptions);

  // If we get a 401, try to refresh the token
  if (result.error && result.error.status === 401) {
    const refreshToken = typeof window !== "undefined" ? window.localStorage?.getItem("refresh_token") : null;
    
    if (refreshToken) {
      // Try to refresh the token
      const refreshResult = await fetchBaseQuery({
        baseUrl:
          process.env.NEXT_PUBLIC_API_BASE_URL?.replace(/\/$/, "") ||
          "http://localhost:8110/api/v1/es",
        credentials: "include",
        prepareHeaders: (headers) => {
          headers.set("Accept", "application/json");
          return headers;
        },
      })({
        url: "/auth/refresh",
        method: "POST",
        body: { refresh_token: refreshToken },
      }, api, extraOptions);

      if (refreshResult.data) {
        const { access_token, refresh_token } = refreshResult.data as RefreshTokenResponse;
        
        // Store the new tokens
        if (typeof window !== "undefined") {
          window.localStorage.setItem("access_token", access_token);
          window.localStorage.setItem("refresh_token", refresh_token);
        }

        // Retry the original request with the new token
        result = await fetchBaseQuery({
          baseUrl:
            process.env.NEXT_PUBLIC_API_BASE_URL?.replace(/\/$/, "") ||
            "http://localhost:8110/api/v1/es",
          credentials: "include",
          prepareHeaders: (headers) => {
            headers.set("Authorization", `Bearer ${access_token}`);
            headers.set("Accept", "application/json");
            return headers;
          },
        })(args, api, extraOptions);
      } else {
        // Refresh failed, clear tokens and redirect to login
        if (typeof window !== "undefined") {
          window.localStorage.removeItem("access_token");
          window.localStorage.removeItem("refresh_token");
          window.location.href = "/login";
        }
      }
    } else {
      // No refresh token, redirect to login
      if (typeof window !== "undefined") {
        window.location.href = "/login";
      }
    }
  }

  // Handle errors by checking if errors array is not empty
  if (result.error && result.error.data && typeof result.error.data === 'object') {
    const errorData = result.error.data as any;
    
    // Check if it's our API error format with non-empty errors array
    if (errorData.errors && Array.isArray(errorData.errors) && errorData.errors.length > 0) {
      const errorDetails = errorData.errors;
      const errorMessage = errorDetails[0];
      const statusCode = typeof result.error.status === 'number' ? result.error.status : undefined;
      
      // Determine error type based on status code
      let errorType: "error" | "not_found" = "error";
      
      if (statusCode === 404) {
        errorType = "not_found";
      }
      
      // Dispatch error to Redux store
      if (typeof window !== "undefined") {
        // Import store dynamically to avoid circular dependency
        import("@/store/store").then(({ store }) => {
          import("@/store/errorSlice").then(({ showError }) => {
            store.dispatch(showError({
              type: errorType,
              message: errorMessage,
              errors: errorDetails,
              statusCode: statusCode
            }));
          });
        });
      }
    }
  }

  return result;
};

export const api = createApi({
  reducerPath: "api",
  baseQuery: baseQueryWithReauth,
  tagTypes: ["AccessRequest", "DegreeProgram"],
  endpoints: (builder) => ({
    // Catalogs
    listIdentificationTypes: builder.query<IdentificationTypeItem[], { current_page?: number; items_per_page?: number } | void>({
      query: (args) => {
        const params = new URLSearchParams();
        params.set("current_page", String(args?.current_page ?? 1));
        params.set("items_per_page", String(args?.items_per_page ?? 10));
        return { url: `/identification-types?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<IdentificationTypeItem>>) => response.result.data,
    }),
    listHigherEducationInstitutions: builder.query<
      HigherEducationInstitutionItem[],
      { nameCont?: string; current_page?: number; items_per_page?: number } | void
    >({
      query: (args) => {
        const params = new URLSearchParams();
        params.set("current_page", String(args?.current_page ?? 1));
        params.set("items_per_page", String(args?.items_per_page ?? 10));
        if (args && args.nameCont) {
          const filters = new URLSearchParams();
          filters.set("name[cont]", args.nameCont);
          params.set("filters", filters.toString());
        }
        return { url: `/higher-education-institution?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<HigherEducationInstitutionItem>>) => response.result.data,
    }),
    // Admin: Access Requests
    listAccessRequests: builder.query<
      PaginatedPage<AccessRequestItem>,
      { current_page?: number; items_per_page?: number; filters?: Record<string, string> }
    >({
      query: ({ current_page = 1, items_per_page = 10, filters } = {}) => {
        const params = new URLSearchParams();
        params.set("current_page", String(current_page));
        params.set("items_per_page", String(items_per_page));
        if (filters) {
          const nested = new URLSearchParams();
          for (const [key, value] of Object.entries(filters)) {
            if (value !== undefined && value !== null && value !== "") {
              nested.append(key, String(value));
            }
          }
          params.set("filters", nested.toString());
        }
        return { url: `/access-request?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<AccessRequestItem>>) => response.result,
      providesTags: (result) =>
        result
          ? [
              ...result.data.map((item) => ({ type: "AccessRequest" as const, id: item.id })),
              { type: "AccessRequest" as const, id: "LIST" },
            ]
          : [{ type: "AccessRequest" as const, id: "LIST" }],
    }),
    createAccessRequest: builder.mutation<ApiResponse<string>, CreateAccessRequestRequest>({
      query: (body) => ({ url: "/access-request", method: "POST", body }),
      // Don't transform the response, return the full ApiResponse
    }),
    validateAccessRequestEmail: builder.mutation<string, ValidateEmailRequest>({
      query: (body) => ({ url: "/access-request/email/validate", method: "POST", body }),
      transformResponse: (response: ApiResponse<string>) => response.result,
    }),
    approveAccessRequest: builder.mutation<{ message: string }, { id: number }>(
      {
        query: ({ id }) => ({ url: `/access-request/${id}/approve`, method: "POST" }),
        invalidatesTags: (_result, _error, { id }) => [
          { type: "AccessRequest", id },
          { type: "AccessRequest", id: "LIST" },
        ],
      }
    ),
    rejectAccessRequest: builder.mutation<{ message: string }, { id: number }>(
      {
        query: ({ id }) => ({ url: `/access-request/${id}/reject`, method: "POST" }),
        invalidatesTags: (_result, _error, { id }) => [
          { type: "AccessRequest", id },
          { type: "AccessRequest", id: "LIST" },
        ],
      }
    ),
    // Director: Degree Programs
    listDurationUnits: builder.query<DurationUnitItem[], void>({
      query: () => ({ url: "/duration-unit" }),
      transformResponse: (response: ApiResponse<{ data: DurationUnitItem[] }>) => response.result.data,
    }),
    listDegreePrograms: builder.query<
      PaginatedPage<DegreeProgramItem>,
      { current_page?: number; items_per_page?: number; filters?: Record<string, string> }
    >({
      query: ({ current_page = 1, items_per_page = 10, filters } = {}) => {
        const params = new URLSearchParams();
        params.set("current_page", String(current_page));
        params.set("items_per_page", String(items_per_page));
        if (filters) {
          const nested = new URLSearchParams();
          for (const [key, value] of Object.entries(filters)) {
            if (value !== undefined && value !== null && value !== "") {
              nested.append(key, String(value));
            }
          }
          params.set("filters", nested.toString());
        }
        return { url: `/degree-program?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<DegreeProgramItem>>) => response.result,
      providesTags: (result) =>
        result
          ? [
              ...result.data.map((item) => ({ type: "DegreeProgram" as const, id: item.id })),
              { type: "DegreeProgram" as const, id: "LIST" },
            ]
          : [{ type: "DegreeProgram" as const, id: "LIST" }],
    }),
    createDegreeProgram: builder.mutation<ApiResponse<string>, CreateDegreeProgramRequest>({
      query: (body) => ({ url: "/degree-program", method: "POST", body }),
      invalidatesTags: [{ type: "DegreeProgram", id: "LIST" }],
    }),
    // Authentication endpoints
    login: builder.mutation<LoginResponse, LoginRequest>({
      query: (body) => ({ url: "/auth", method: "POST", body }),
      transformResponse: (response: ApiResponse<LoginResponse>) => response.result,
    }),
    refreshToken: builder.mutation<RefreshTokenResponse, RefreshTokenRequest>({
      query: (body) => ({ url: "/auth/refresh", method: "POST", body }),
      transformResponse: (response: ApiResponse<RefreshTokenResponse>) => response.result,
    }),
    createUserFromToken: builder.mutation<CreateUserFromTokenResponse, CreateUserFromTokenRequest>({
      query: (body) => ({ url: "/user/create-from-token", method: "POST", body }),
      transformResponse: (response: ApiResponse<CreateUserFromTokenResponse>) => response.result,
    }),
  }),
});

export const {
  // catalogs
  useListIdentificationTypesQuery,
  useListHigherEducationInstitutionsQuery,
  useListAccessRequestsQuery,
  useCreateAccessRequestMutation,
  useValidateAccessRequestEmailMutation,
  useApproveAccessRequestMutation,
  useRejectAccessRequestMutation,
  // degree programs
  useListDurationUnitsQuery,
  useListDegreeProgramsQuery,
  useCreateDegreeProgramMutation,
  // authentication
  useLoginMutation,
  useRefreshTokenMutation,
  useCreateUserFromTokenMutation,
} = api;


