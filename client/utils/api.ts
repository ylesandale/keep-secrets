import apiClient from './axios'

export interface CreateSecretResponse {
  id: string
}

export interface GetSecretResponse {
  encryptedData: string
}

export const secretApi = {
  create: async (encryptedData: string): Promise<CreateSecretResponse> => {
    const response = await apiClient.post<CreateSecretResponse>(
      '/api/secrets',
      {
        encryptedData,
      }
    )
    return response.data
  },

  get: async (id: string): Promise<GetSecretResponse> => {
    const response = await apiClient.get<GetSecretResponse>(
      `/api/secrets/${id}`
    )
    return response.data
  },
}
