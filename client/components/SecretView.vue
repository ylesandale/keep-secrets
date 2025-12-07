<template>
  <div class="space-y-4">
    <UAlert
      color="green"
      variant="soft"
      title="Secret successfully decrypted"
      description="This secret has been deleted and is no longer available."
    />

    <div class="space-y-2">
      <label class="text-sm font-medium text-gray-700 dark:text-gray-300">
        Content:
      </label>
      <div
        class="bg-gray-50 dark:bg-gray-900 rounded-lg p-4 border border-gray-200 dark:border-gray-700"
      >
        <p class="whitespace-pre-wrap text-gray-900 dark:text-white">
          {{ secret }}
        </p>
      </div>
    </div>

    <UButton to="/" variant="outline" block> Create new secret </UButton>
  </div>
</template>

<script setup lang="ts">
interface Props {
  secret: string
}

defineProps<Props>()
</script>
