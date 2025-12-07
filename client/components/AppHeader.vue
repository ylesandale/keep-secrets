<template>
  <header class="border-b border-gray-200 dark:border-gray-800">
    <UContainer>
      <div class="flex items-center justify-between h-16">
        <NuxtLink
          to="/"
          class="text-xl font-bold text-gray-900 dark:text-white flex items-center gap-2"
        >
          <UIcon
            name="material-symbols:mail-lock-rounded"
            class="w-5 h-5"
            mode="svg"
          />
          Keep Secrets
        </NuxtLink>

        <UButton
          :icon="themeIcon"
          color="gray"
          variant="ghost"
          aria-label="Toggle color mode"
          @click="toggleColorMode"
        />
      </div>
    </UContainer>
  </header>
</template>

<script setup lang="ts">
const colorMode = useColorMode()

const themeIcon = computed(() =>
  colorMode.value === 'dark' ? 'i-heroicons-sun' : 'i-heroicons-moon'
)

const toggleColorMode = () => {
  colorMode.preference = colorMode.value === 'light' ? 'dark' : 'light'
}
</script>
