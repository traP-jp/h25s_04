<script lang="ts" setup>
// import type { Eatery } from '../lib/apis'

// interface Props {
//   eateries: Eatery[]
// }

const eateries = [
  { name: 'Test Eatery', description: 'This is a test eatery', id: 1 },
  { name: 'Sample Eatery', description: 'This is a sample eatery', id: 2 },
  { name: 'Demo Eatery', description: 'This is a demo eatery', id: 3 },
]

// const { eateries } = defineProps<Props>()
</script>

<template>
  <ul>
    <li v-for="eatery in eateries" :key="eatery.id">
      {{ eatery.name }} {{ eatery.description }} a
    </li>
  </ul>
</template>

<style lang="scss" module></style>
