<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <AppHeader />

    <section class="py-16 md:py-24">
      <UContainer class="max-w-2xl">
        <div
          class="bg-white dark:bg-gray-800 rounded-lg shadow-[0_-4px_6px_-1px_rgba(0,0,0,0.1),0_4px_6px_-1px_rgba(0,0,0,0.1)] dark:shadow-lg p-8"
        >
          <div class="mb-6">
            <UButton
              to="/"
              variant="ghost"
              color="gray"
              icon="i-heroicons-arrow-left"
              class="mb-4"
            >
              Back to home
            </UButton>
            <h1
              class="text-3xl font-bold text-center mb-8 text-gray-900 dark:text-white"
            >
              View Secret
            </h1>
          </div>

          <SecretLoading v-if="isLoading" />

          <SecretError v-else-if="errorMessage" :error="errorMessage" />

          <SecretView v-else-if="decryptedSecret" :secret="decryptedSecret" />

          <div v-else class="text-center py-8">
            <p class="text-gray-600 dark:text-gray-400">
              Waiting for secret...
            </p>
          </div>
        </div>
      </UContainer>
    </section>

    <AppFooter />
  </div>
</template>

<script setup lang="ts">
const route = useRoute()
const decryptedSecret = ref('')
const errorMessage = ref('')
const hash = ref('')

const { decrypt } = useEncryption()

const id = computed(() => route.params.id as string)

useHead({
  title: 'View Secret - Keep Secrets',
  meta: [
    {
      name: 'description',
      content:
        'View your one-time secret. This secret will be deleted after viewing.',
    },
    {
      name: 'robots',
      content: 'noindex, nofollow',
    },
  ],
})

watch(
  () => route.hash,
  (newHash) => {
    if (newHash) {
      hash.value = newHash
    }
  },
  { immediate: true }
)

const encryptionKey = computed(() => {
  if (!hash.value) return null
  const keyMatch = hash.value.match(/key=([^&]+)/)
  if (!keyMatch) {
    return null
  }
  try {
    return decodeURIComponent(keyMatch[1])
  } catch {
    return null
  }
})

const canFetch = computed(() => {
  return !!id.value && !!encryptionKey.value
})

watch([id, encryptionKey, hash], ([secretId, key, currentHash]) => {
  if (!secretId) {
    errorMessage.value = 'Invalid link'
  } else if (!key || !currentHash) {
    errorMessage.value = 'Encryption key not found in link'
  } else {
    errorMessage.value = ''
  }
})

const { data, isLoading, isError, error } = useGetSecret(id, canFetch)

watch(
  [data, encryptionKey],
  async ([secretData, key]) => {
    if (!secretData || !key) return

    try {
      const decrypted = await decrypt(secretData.encryptedData, key)
      decryptedSecret.value = decrypted
      errorMessage.value = ''
    } catch (err) {
      console.error('Failed to decrypt:', err)
      errorMessage.value =
        err instanceof Error ? err.message : 'Failed to decrypt secret'
      decryptedSecret.value = ''
    }
  },
  { immediate: true }
)

watch([isError, error], ([hasError, err]) => {
  if (hasError && err && !decryptedSecret.value) {
    const errorMsg =
      err instanceof Error ? err.message : 'Failed to retrieve secret'
    if (errorMsg.includes('404') || errorMsg.includes('not found')) {
      errorMessage.value = 'Secret not found or already viewed'
    } else {
      errorMessage.value = errorMsg
    }
  }
})
</script>
