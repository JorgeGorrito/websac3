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

// Formation Level types
export type FormationLevelItem = {
  id: number;
  name: string;
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
  formation_level: {
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
  formation_level_id: number;
  graduate_profile: string;
  name: string;
  professional_profile: string;
  program_focus: string;
  snies: number;
  total_credits: number;
};

export type ProfessionalRoleItem = {
  id: number;
  name: string;
};

export type EvaluateDegreeProgramRequest = {
  degree_program_id: number;
  professional_role_id: number;
};

export type ReportItem = {
  id: number;
  created_at: string;
  lang: string;
  score: number; // Can be decimal (0.0875) or integer (87)
};

export type TopicReport = {
  id: number;
  name: string;
  topic_id: number;
  learn_hours_expected: number;
  learn_hours_actual: number;
};

export type KnowledgeAreaReport = {
  id: number;
  name: string;
  total_learn_hours_expected: number;
  total_learn_hours_actual: number;
  score_expected: number;
  score_got: number;
  topic_reports: TopicReport[];
};

export type UnexpectedKnowledgeAreaReport = {
  id: number;
  name: string;
  total_learn_hours: number;
  topic_reports: UnexpectedTopicReport[];
};

export type UnexpectedTopicReport = {
  id: number;
  name: string;
  topic_id: number;
  learn_hours_actual: number;
  topic: {
    id: number;
    name: string;
  };
};

export type ReportDetail = {
  id: number;
  created_at: string;
  degree_program_id: number;
  degree_program: {
    id: number;
    name: string;
    snies: number;
    total_credits: number;
    duration_value: number;
    duration_unit: {
      id: number;
      name: string;
    };
    entry_profile: string;
    graduate_profile: string;
    professional_profile: string;
    program_focus: string;
    created_by: number;
    higher_education_institution: {
      name: string;
      snies: number;
      department: string;
      municipality: string;
      ownership: string;
      institutional_category: string;
    };
  };
  professional_role: {
    id: number;
    name: string;
  };
  score: number;
  knowledge_area_reports: KnowledgeAreaReport[];
  unexpected_knowledge_area_reports?: UnexpectedKnowledgeAreaReport[];
};

export type CreateUserFromTokenResponse = {
  message: string;
};

// Expert Consultation types
export type CreateExpertConsultationRequest = {
  degree_program_id: number;
  report_id: number;
  request_message: string;
  requester_id: number;
};

export type CreateExpertConsultationResponse = {
  id: number;
  message: string;
};

// Access Request types
export type AccessRequestItem = {
  id: number;
  name: string;
  lastname: string;
  email: string;
  identification_number: string;
  identification_type: string;
  job_position: string;
  department_name: string;
  municipality_name: string;
  higher_education_institution_name: string;
  higher_education_institution_snies: number;
  higher_education_institution_ownership: string;
  status_id: number;
  status_name: string;
};

// Role types
export type RoleItem = {
  id: number;
  name: string;
};

// Access Request Actions
export type ApproveAccessRequestRequest = {
  id: number;
  role_id: number;
};

export type RejectAccessRequestRequest = {
  id: number;
};

// User types
export type UserItem = {
  id: number;
  email: string;
  full_name: string;
  is_active: boolean;
};

// User Actions
export type ActivateUserRequest = {
  id: number;
};

export type DeactivateUserRequest = {
  id: number;
};

// Topic types
export type TopicItem = {
  id: number;
  name: string;
  knowledge_area_id: number;
  knowledge_area_name: string;
};

// Course Type and Nature types
export type CourseTypeItem = {
  id: number;
  name: string;
};

export type CourseNatureItem = {
  id: number;
  name: string;
};

// Course types
export type CourseTopic = {
  topic_id: number;
  study_hours: number;
};

export type CourseTopicItem = {
  id: number;
  topic_id: number;
  study_hours: number;
  topic_name: string;
  knowledge_area_name: string;
};

export type CreateCourseRequest = {
  code: string;
  name: string;
  credits: number;
  period_number: number;
  type_id: number;
  nature_id: number;
  is_cybersecurity: boolean;
  contains_cybersecurity_topics: boolean;
  degree_program_id: number;
  course_topics: CourseTopic[];
};

export type UpdateCourseRequest = {
  id: number;
  code: string;
  name: string;
  credits: number;
  period_number: number;
  type_id: number;
  nature_id: number;
  is_cybersecurity: boolean;
  contains_cybersecurity_topics: boolean;
  degree_program_id: number;
  course_topics: CourseTopic[];
};

export type CourseItem = {
  id: number;
  code: string;
  name: string;
  credits: number;
  period_number: number;
  type_id: number;
  nature_id: number;
  is_cybersecurity: boolean;
  contains_cybersecurity_topics: boolean;
  degree_program_id: number;
  nature_name: string;
  type_name: string;
  degree_program_name: string;
  created_by: number;
  creator_name: string;
  status?: 'completed' | 'current' | 'pending' | 'failed';
  grade?: number;
};

export type IdentificationTypeItem = {
  id: number;
  name: string;
};

export type HigherEducationInstitutionItem = {
  snies: number;
  name: string;
};

// Get language from localStorage or default to 'es'
const getLanguage = (): string => {
  if (typeof window !== "undefined") {
    return window.localStorage?.getItem("language") || "es";
  }
  return "es";
};

// Function to set language
export const setLanguage = (language: string): void => {
  if (typeof window !== "undefined") {
    window.localStorage.setItem("language", language);
  }
};

// Function to get current language
export const getCurrentLanguage = (): string => {
  return getLanguage();
};


// Custom base query with automatic token refresh
const baseQueryWithReauth = async (args: any, api: any, extraOptions: any) => {
  const language = getLanguage();
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL?.replace(/\/$/, "") || `http://localhost:8110/api/v1/${language}`;
  
  let result = await fetchBaseQuery({
    baseUrl,
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
      try {
        // Try to refresh the token
        const refreshResult = await fetchBaseQuery({
          baseUrl,
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
          // Try different possible field names for the tokens
          const responseData = refreshResult.data as any;
          
          // Check if tokens are nested in a 'result' object
          const tokenData = responseData.result || responseData;
          let access_token = tokenData.access_token || tokenData.accessToken || tokenData.access;
          let refresh_token = tokenData.refresh_token || tokenData.refreshToken || tokenData.refresh;
          
          if (!access_token) {
            throw new Error("access_token is undefined in refresh response");
          }
          
          // Store the new tokens
          if (typeof window !== "undefined") {
            window.localStorage.setItem("access_token", access_token);
            window.localStorage.setItem("refresh_token", refresh_token);
          }

          // Retry the original request with the new token
          result = await fetchBaseQuery({
            baseUrl,
            credentials: "include",
            prepareHeaders: (headers) => {
              // Use the new token directly, not from localStorage
              headers.set("Authorization", `Bearer ${access_token}`);
              headers.set("Accept", "application/json");
              return headers;
            },
          })(args, api, extraOptions);
          
          // Don't redirect to login after successful refresh, even if retry fails
          // The error will be handled by the normal error handling below
        } else {
          // Refresh failed, clear tokens and redirect to login
          if (typeof window !== "undefined") {
            window.localStorage.removeItem("access_token");
            window.localStorage.removeItem("refresh_token");
            window.location.href = "/login";
          }
        }
      } catch (refreshError) {
        // Refresh request itself failed, clear tokens and redirect to login
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
    // Try both 'Errors' (capital E) and 'errors' (lowercase e) for compatibility
    const errorsArray = errorData.Errors || errorData.errors;
    
    if (errorsArray && Array.isArray(errorsArray) && errorsArray.length > 0) {
      const errorDetails = errorsArray;
      const errorMessage = errorDetails[0];
      const statusCode = typeof result.error.status === 'number' ? result.error.status : undefined;
      
      // Determine error type based on status code
      let errorType: "error" | "not_found" = "error";
      
      if (statusCode === 404) {
        errorType = "not_found";
      }
      
      // Don't show error modal for 404 on certain endpoints that can have empty results
      const url = args?.url || '';
      const shouldShow404Modal = !(
        statusCode === 404 && (
          url.includes('/access-request') || // Access requests can be empty
          url.includes('/users') || // Users list can be empty
          url.includes('/degree-program') || // Degree programs can be empty
          url.includes('/reports') // Reports can be empty
        )
      );
      
      // Dispatch error to Redux store only if we should show the modal
      if (shouldShow404Modal && typeof window !== "undefined") {
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
  tagTypes: ["AccessRequest", "DegreeProgram", "Topic", "CourseType", "CourseNature", "Course", "User"],
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
    approveAccessRequest: builder.mutation<{ message: string }, ApproveAccessRequestRequest>(
      {
        query: ({ id, role_id }) => ({ 
          url: `/access-request/${id}/approve`, 
          method: "POST", 
          body: { role_id } 
        }),
        invalidatesTags: (_result, _error, { id }) => [
          { type: "AccessRequest", id },
          { type: "AccessRequest", id: "LIST" },
        ],
      }
    ),
    rejectAccessRequest: builder.mutation<{ message: string }, RejectAccessRequestRequest>(
      {
        query: ({ id }) => ({ url: `/access-request/${id}/reject`, method: "POST" }),
        invalidatesTags: (_result, _error, { id }) => [
          { type: "AccessRequest", id },
          { type: "AccessRequest", id: "LIST" },
        ],
      }
    ),
    // Admin: Users
    listUsers: builder.query<
      PaginatedPage<UserItem>,
      { current_page?: number; items_per_page?: number; filters?: Record<string, string> }
    >({
      query: ({ current_page = 1, items_per_page = 10, filters } = {}) => {
        const params = new URLSearchParams();
        params.set("current_page", String(current_page));
        params.set("items_per_page", String(items_per_page));
        
        // Add filters to query params
        if (filters) {
          Object.entries(filters).forEach(([key, value]) => {
            if (value.trim()) {
              params.set(key, value.trim());
            }
          });
        }
        
        return { url: `/users?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<UserItem>>) => response.result,
      providesTags: (result) =>
        result
          ? [
              ...result.data.map(({ id }) => ({ type: "User" as const, id })),
              { type: "User", id: "LIST" },
            ]
          : [{ type: "User", id: "LIST" }],
    }),
    activateUser: builder.mutation<{ message: string }, ActivateUserRequest>({
      query: ({ id }) => ({ url: `/users/${id}/activate`, method: "PUT" }),
      invalidatesTags: (_result, _error, { id }) => [
        { type: "User", id },
        { type: "User", id: "LIST" },
      ],
    }),
    deactivateUser: builder.mutation<{ message: string }, DeactivateUserRequest>({
      query: ({ id }) => ({ url: `/users/${id}/deactivate`, method: "PUT" }),
      invalidatesTags: (_result, _error, { id }) => [
        { type: "User", id },
        { type: "User", id: "LIST" },
      ],
    }),
    // Director: Degree Programs
    listDurationUnits: builder.query<DurationUnitItem[], void>({
      query: () => ({ url: "/duration-unit" }),
      transformResponse: (response: ApiResponse<{ data: DurationUnitItem[] }>) => response.result.data,
    }),
    listFormationLevels: builder.query<FormationLevelItem[], void>({
      query: () => ({ url: "/formation-level" }),
      transformResponse: (response: ApiResponse<PaginatedPage<FormationLevelItem>>) => response.result.data,
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
    // Topics
    listTopics: builder.query<
      PaginatedPage<TopicItem>,
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
        return { url: `/topic?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<TopicItem>>) => response.result,
      providesTags: (result) =>
        result
          ? [
              ...result.data.map((item) => ({ type: "Topic" as const, id: item.id })),
              { type: "Topic" as const, id: "LIST" },
            ]
          : [{ type: "Topic" as const, id: "LIST" }],
    }),
    // Course Types
    listCourseTypes: builder.query<
      PaginatedPage<CourseTypeItem>,
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
        return { url: `/course-types?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<CourseTypeItem>>) => response.result,
      providesTags: (result) =>
        result
          ? [
              ...result.data.map((item) => ({ type: "CourseType" as const, id: item.id })),
              { type: "CourseType" as const, id: "LIST" },
            ]
          : [{ type: "CourseType" as const, id: "LIST" }],
    }),
    // Course Natures
    listCourseNatures: builder.query<
      PaginatedPage<CourseNatureItem>,
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
        return { url: `/course-natures?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<CourseNatureItem>>) => response.result,
      providesTags: (result) =>
        result
          ? [
              ...result.data.map((item) => ({ type: "CourseNature" as const, id: item.id })),
              { type: "CourseNature" as const, id: "LIST" },
            ]
          : [{ type: "CourseNature" as const, id: "LIST" }],
    }),
            // Courses
            listCourses: builder.query<
              PaginatedPage<CourseItem>,
              { 
                current_page?: number; 
                items_per_page?: number; 
                filters?: Record<string, string>;
                degree_program_id?: number;
              }
            >({
              query: ({ current_page = 1, items_per_page = 50, filters, degree_program_id } = {}) => {
                const params = new URLSearchParams();
                params.set("current_page", String(current_page));
                params.set("items_per_page", String(items_per_page));
                if (degree_program_id) {
                  params.set("degree_program_id", String(degree_program_id));
                }
                if (filters) {
                  const nested = new URLSearchParams();
                  for (const [key, value] of Object.entries(filters)) {
                    if (value !== undefined && value !== null && value !== "") {
                      nested.append(key, String(value));
                    }
                  }
                  params.set("filters", nested.toString());
                }
                return { url: `/course?${params.toString()}` };
              },
              transformResponse: (response: ApiResponse<PaginatedPage<CourseItem>>) => response.result,
              providesTags: (result) =>
                result
                  ? [
                      ...result.data.map((item) => ({ type: "Course" as const, id: item.id })),
                      { type: "Course" as const, id: "LIST" },
                    ]
                  : [{ type: "Course" as const, id: "LIST" }],
            }),
            // Courses by degree program (specific endpoint)
            listCoursesByDegreeProgram: builder.query<
              PaginatedPage<CourseItem>,
              { 
                degree_program_id: number;
                current_page?: number; 
                items_per_page?: number; 
                filters?: Record<string, string>;
              }
            >({
              query: ({ degree_program_id, current_page = 1, items_per_page = 50, filters } = {}) => {
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
                return { url: `/degree-program/${degree_program_id}/courses?${params.toString()}` };
              },
              transformResponse: (response: ApiResponse<PaginatedPage<CourseItem>>) => response.result,
              providesTags: (result, error, { degree_program_id }) =>
                result
                  ? [
                      ...result.data.map((item) => ({ type: "Course" as const, id: item.id })),
                      { type: "Course" as const, id: `PROGRAM_${degree_program_id}` },
                    ]
                  : [{ type: "Course" as const, id: `PROGRAM_${degree_program_id}` }],
            }),
            createCourse: builder.mutation<ApiResponse<string>, CreateCourseRequest>({
              query: (body) => ({ url: "/course", method: "POST", body }),
              invalidatesTags: (result, error, { degree_program_id }) => [
                { type: "Course", id: "LIST" },
                { type: "Course", id: `PROGRAM_${degree_program_id}` },
              ],
            }),
            // Get course by ID
            getCourseById: builder.query<CourseItem, { course_id: number }>({
              query: ({ course_id }) => ({ url: `/course/${course_id}` }),
              transformResponse: (response: ApiResponse<CourseItem>) => response.result,
              providesTags: (result, error, { course_id }) => [
                { type: "Course", id: course_id },
              ],
            }),
            // Get course topics
            getCourseTopics: builder.query<CourseTopicItem[], { course_id: number }>({
              query: ({ course_id }) => ({ url: `/course/${course_id}/topics` }),
              transformResponse: (response: ApiResponse<any>) => {
                // Handle different possible response structures
                if (Array.isArray(response.result)) {
                  return response.result;
                } else if (response.result && Array.isArray(response.result.data)) {
                  return response.result.data;
                } else if (response.result && response.result.topics && Array.isArray(response.result.topics)) {
                  return response.result.topics;
                }
                return [];
              },
              providesTags: (result, error, { course_id }) => [
                { type: "Course", id: course_id },
                { type: "Course", id: `${course_id}_topics` },
              ],
            }),
            // Update course
            updateCourse: builder.mutation<ApiResponse<string>, { course_id: number; data: UpdateCourseRequest }>({
              query: ({ course_id, data }) => ({ url: `/course/${course_id}`, method: "PUT", body: data }),
              invalidatesTags: (result, error, { course_id }) => [
                { type: "Course", id: course_id },
                { type: "Course", id: "LIST" },
              ],
            }),
            // Delete course
            deleteCourse: builder.mutation<ApiResponse<string>, { course_id: number }>({
              query: ({ course_id }) => ({ url: `/course/${course_id}`, method: "DELETE" }),
              invalidatesTags: (result, error, { course_id }) => [
                { type: "Course", id: course_id },
                { type: "Course", id: "LIST" },
              ],
            }),
    // Professional Roles
    listProfessionalRoles: builder.query<PaginatedPage<ProfessionalRoleItem>, void>({
      query: () => ({ url: "/professional-roles" }),
      transformResponse: (response: ApiResponse<PaginatedPage<ProfessionalRoleItem>>) => response.result,
    }),
    // System Roles
    listRoles: builder.query<PaginatedPage<RoleItem>, { current_page?: number; items_per_page?: number } | void>({
      query: (args) => {
        const params = new URLSearchParams();
        params.set("current_page", String(args?.current_page ?? 1));
        params.set("items_per_page", String(args?.items_per_page ?? 100));
        return { url: `/roles?${params.toString()}` };
      },
      transformResponse: (response: ApiResponse<PaginatedPage<RoleItem>>) => response.result,
    }),
    // Degree Program Evaluation
    evaluateDegreeProgram: builder.mutation<ApiResponse<string>, EvaluateDegreeProgramRequest>({
      query: (body) => ({ url: "/degree-program/evaluate", method: "POST", body }),
    }),
    // Degree Program Reports
    listDegreeProgramReports: builder.query<
      PaginatedPage<ReportItem>,
      { degree_program_id: number; current_page?: number; items_per_page?: number }
    >({
      query: ({ degree_program_id, current_page = 1, items_per_page = 10 }) => {
        const params = new URLSearchParams();
        params.set("current_page", String(current_page));
        params.set("items_per_page", String(items_per_page));
        return { url: `/degree-program/${degree_program_id}/reports?${params.toString()}` };
      },
      transformResponse: (response: any) => {
        // Handle both formats: ApiResponse format and direct format
        if (response.Result) {
          return response.Result;
        }
        return response.result;
      },
    }),
    // Report Detail
    getReportDetail: builder.query<ReportDetail, number>({
      query: (report_id) => ({ url: `/report/${report_id}` }),
      transformResponse: (response: any) => {
        // Handle both formats: ApiResponse format and direct format
        if (response.Result) {
          return response.Result;
        }
        return response.result;
      },
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

    // Expert Consultation
    createExpertConsultation: builder.mutation<CreateExpertConsultationResponse, CreateExpertConsultationRequest>({
      query: (body) => ({ url: "/expert-consultation", method: "POST", body }),
      transformResponse: (response: ApiResponse<CreateExpertConsultationResponse>) => response.result,
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
  // admin users
  useListUsersQuery,
  useActivateUserMutation,
  useDeactivateUserMutation,
  // degree programs
  useListDurationUnitsQuery,
  useListFormationLevelsQuery,
  useListDegreeProgramsQuery,
  useCreateDegreeProgramMutation,
  // topics
  useListTopicsQuery,
  // course types and natures
  useListCourseTypesQuery,
  useListCourseNaturesQuery,
  // courses
  useListCoursesQuery,
  useListCoursesByDegreeProgramQuery,
  useCreateCourseMutation,
  useGetCourseByIdQuery,
  useGetCourseTopicsQuery,
  useUpdateCourseMutation,
  useDeleteCourseMutation,
  // professional roles
  useListProfessionalRolesQuery,
  // system roles
  useListRolesQuery,
  // degree program evaluation
  useEvaluateDegreeProgramMutation,
  // degree program reports
  useListDegreeProgramReportsQuery,
  // report detail
  useGetReportDetailQuery,
  // authentication
  useLoginMutation,
  useRefreshTokenMutation,
  useCreateUserFromTokenMutation,
  // expert consultation
  useCreateExpertConsultationMutation,
} = api;


