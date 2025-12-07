<template>
  <footer
    class="border-t border-gray-200 dark:border-gray-800 bg-gray-50 dark:bg-gray-900"
  >
    <UContainer>
      <div class="py-12">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div>
            <h3 class="font-bold text-gray-900 dark:text-white mb-4">
              Keep Secrets
            </h3>
            <p class="text-sm text-gray-600 dark:text-gray-400">
              Share confidential information securely with one-time links that
              self-destruct after viewing.
            </p>
          </div>
        </div>

        <div
          class="mt-8 pt-8 border-t border-gray-200 dark:border-gray-800 text-center text-sm text-gray-600 dark:text-gray-400"
        >
          <p>
            © {{ new Date().getFullYear() }} Keep Secrets. All rights reserved.
          </p>
        </div>
      </div>
    </UContainer>
  </footer>
</template>
