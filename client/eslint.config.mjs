import eslint from '@nuxt/eslint-config'

export default eslint({
  ignores: ['node_modules', '.nuxt', '.output', 'dist', '*.min.js', '*.d.ts'],
})
