<template>
  <div class="space-y-4">
    <UAlert
      color="green"
      variant="soft"
      title="Secret link created!"
      description="Copy the link and share it. It will be available only once."
    />

    <div class="space-y-2">
      <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
        Your link:
      </label>
      <div class="flex gap-2">
        <UInput
          :model-value="secretLink"
          readonly
          class="flex-1 font-mono text-sm"
        />
        <UButton
          :color="copied ? 'green' : 'primary'"
          icon="i-heroicons-clipboard-document"
          @click="copyToClipboard"
        >
          {{ copied ? 'Copied!' : 'Copy' }}
        </UButton>
      </div>
    </div>

    <UButton variant="outline" block @click="handleReset">
      Create new secret
    </UButton>
  </div>
</template>

<script setup lang="ts">
interface Props {
  secretLink: string
}

interface Emits {
  (e: 'reset'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const copied = ref(false)

const copyToClipboard = async () => {
  try {
    await navigator.clipboard.writeText(props.secretLink)
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  } catch (error) {
    console.error('Failed to copy:', error)
  }
}

const handleReset = () => {
  emit('reset')
}
</script>
