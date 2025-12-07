<template>
  <div class="bg-gray-50 dark:bg-gray-800 rounded-lg p-8">
    <h2 class="text-2xl font-bold text-gray-900 dark:text-white mb-4">
      End-to-End Encryption
    </h2>
    <p class="text-gray-600 dark:text-gray-400 mb-6">
      We encrypt your secret directly on your device. The encryption key is
      never stored; instead, it becomes part of the link itself. Without the
      full link, no one, including us, can decrypt your secret.
    </p>
    <div class="space-y-3">
      <div class="flex items-center gap-2">
        <UIcon
          name="material-symbols:lock-outline"
          class="w-5 h-5 text-primary-600 dark:text-primary-400"
          mode="svg"
        />
        <span class="text-sm text-gray-600 dark:text-gray-400"
          >The encryption key is never sent to the server.</span
        >
      </div>
      <div class="flex items-center gap-2">
        <UIcon
          name="material-symbols:shield-outline"
          class="w-5 h-5 text-primary-600 dark:text-primary-400"
          mode="svg"
        />
        <span class="text-sm text-gray-600 dark:text-gray-400"
          >Swiss-level privacy protection.</span
        >
      </div>
      <div class="flex items-center gap-2">
        <UIcon
          name="material-symbols:visibility-off-outline"
          class="w-5 h-5 text-primary-600 dark:text-primary-400"
          mode="svg"
        />
        <span class="text-sm text-gray-600 dark:text-gray-400"
          >No traces left behind.</span
        >
      </div>
    </div>
  </div>
</template>

