import axios from 'axios'

const getApiBase = (): string => {
  if (import.meta.env.DEV) {
    return ''
  }
  const config = useRuntimeConfig()
  return config.public.apiBase || ''
}

const apiClient = axios.create({
  headers: {
    'Content-Type': 'application/json',
  },
})

apiClient.interceptors.request.use((config) => {
  if (!config.baseURL) {
    config.baseURL = getApiBase()
  }
  return config
})

export default apiClient
