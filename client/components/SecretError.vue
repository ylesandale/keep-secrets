<template>
  <div class="space-y-4">
    <UAlert
      color="red"
      variant="soft"
      :title="error"
      description="The secret may have already been viewed or the link is invalid."
    />
  </div>
</template>

<script setup lang="ts">
interface Props {
  error: string
}

defineProps<Props>()
</script>
