<template>
  <UForm :state="state" class="space-y-4" @submit="onSubmit">
    <UFormGroup label="Secret" name="secret">
      <UTextarea
        v-model="state.secret"
        placeholder="Enter your secret..."
        :rows="10"
        required
        class="w-full"
      />
    </UFormGroup>

    <UButton type="submit" :loading="isLoading" block size="lg" color="primary">
      Create secret link
    </UButton>
  </UForm>
</template>

<script setup lang="ts">
interface Emits {
  (e: 'submit', value: string): void
}

const emit = defineEmits<Emits>()

const state = reactive({
  secret: '',
})

const isLoading = defineModel<boolean>('isLoading', { default: false })

const onSubmit = () => {
  if (!state.secret.trim()) {
    return
  }
  emit('submit', state.secret)
  state.secret = ''
}
</script>
