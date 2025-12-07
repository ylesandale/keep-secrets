import { useMutation, useQuery } from '@tanstack/vue-query'
import { secretApi } from '~/utils/api'
import { toRef, unref, computed, type MaybeRef } from 'vue'

export const useCreateSecret = () => {
  return useMutation({
    mutationFn: secretApi.create,
  })
}

export const useGetSecret = (
  id: MaybeRef<string | null>,
  enabled: MaybeRef<boolean> = true
) => {
  const idRef = toRef(id)
  const enabledRef = toRef(enabled)

  return useQuery({
    queryKey: computed(() => ['secret', unref(idRef)]),
    queryFn: async () => {
      const currentId = unref(idRef)
      if (!currentId) throw new Error('Secret ID is required')
      return await secretApi.get(currentId)
    },
    enabled: computed(() => unref(enabledRef) && !!unref(idRef)),
    retry: false,
  })
}
