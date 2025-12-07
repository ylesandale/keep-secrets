<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <AppHeader />

    <section class="py-16 md:py-24">
      <UContainer class="max-w-4xl">
        <HeroSection />

        <div
          class="bg-white dark:bg-gray-800 rounded-lg shadow-[0_-4px_6px_-1px_rgba(0,0,0,0.1),0_4px_6px_-1px_rgba(0,0,0,0.1)] dark:shadow-lg p-8 mb-16"
        >
          <SecretForm
            v-if="!secretLink"
            v-model:is-loading="isCreating"
            @submit="handleSubmit"
          />

          <SecretLinkDisplay
            v-else
            :secret-link="secretLink"
            @reset="reset"
          />
        </div>

        <div class="grid md:grid-cols-2 gap-8 mb-16">
          <OneTimeSecrets />
          <EndToEndEncryption />
        </div>
      </UContainer>
    </section>

    <AppFooter />
  </div>
</template>

<script setup lang="ts">
useHead({
  title: 'Keep Secrets - Share secrets securely with one-time links',
  meta: [
    {
      name: 'description',
      content:
        'Share confidential information securely with one-time links that self-destruct after viewing. End-to-end encryption ensures your secrets stay private.',
    },
    {
      property: 'og:title',
      content: 'Keep Secrets - Share secrets securely with one-time links',
    },
    {
      property: 'og:description',
      content:
        'Share confidential information securely with one-time links that self-destruct after viewing. End-to-end encryption ensures your secrets stay private.',
    },
    {
      property: 'og:type',
      content: 'website',
    },
    {
      name: 'twitter:card',
      content: 'summary',
    },
    {
      name: 'twitter:title',
      content: 'Keep Secrets - Share secrets securely',
    },
    {
      name: 'twitter:description',
      content:
        'Share confidential information securely with one-time links that self-destruct after viewing.',
    },
  ],
})

const secretLink = ref('')
const isCreating = ref(false)

const { encrypt } = useEncryption()
const createSecretMutation = useCreateSecret()

const handleSubmit = async (secretText: string) => {
  if (!secretText.trim()) return

  isCreating.value = true
  try {
    const { encrypted, key } = await encrypt(secretText)
    const { id } = await createSecretMutation.mutateAsync(encrypted)
    const baseUrl = window.location.origin
    secretLink.value = `${baseUrl}/secret/${id}#key=${encodeURIComponent(key)}`
  } catch (error) {
    console.error('Failed to create secret:', error)
    alert('Error creating secret. Please try again.')
  } finally {
    isCreating.value = false
  }
}

const reset = () => {
  secretLink.value = ''
  isCreating.value = false
}
</script>
