<script lang="ts" setup>
// import type { Eatery } from '../lib/apis'

// interface Props {
//   eateries: Eatery[]
// }

const eateries = [
  { name: 'Test Eatery', description: 'This is a test eatery', id: 1 },
  { name: 'Sample Eatery', description: 'This is a sample eatery', id: 2 },
  { name: 'Demo Eatery', description: 'This is a demo eatery', id: 3 },
  { name: 'Example Eatery', description: 'This is an example eatery', id: 4 },
  { name: 'Mock Eatery', description: 'This is a mock eatery', id: 5 },
  { name: 'Another Eatery', description: 'This is another eatery', id: 6 },
  {
    name: 'Yet Another Eatery',
    description: 'This is yet another eatery',
    id: 7,
  },
  { name: 'More Eaterys', description: 'This is more eateries', id: 8 },
  { name: 'Final Eatery', description: 'This is the final eatery', id: 9 },
  { name: 'Last Eatery', description: 'This is the last eatery', id: 10 },
  { name: 'Extra Eatery', description: 'This is an extra eatery', id: 11 },
  { name: 'Bonus Eatery', description: 'This is a bonus eatery', id: 12 },
  { name: 'Special Eatery', description: 'This is a special eatery', id: 13 },
  { name: 'Unique Eatery', description: 'This is a unique eatery', id: 14 },
  { name: 'Rare Eatery', description: 'This is a rare eatery', id: 15 },
  {
    name: 'Hidden Gem Eatery',
    description: 'This is a hidden gem eatery',
    id: 16,
  },
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

<style lang="scss" module>
ul {
  list-style-type: none;
  padding: 0;
  height: 100%;
  overflow-y: auto;
}

li {
  position: relative; /* ::after を使用するために position を設定 */
}

li:not(:last-child)::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 1px;
  background-color: black; /* 区切り線のスタイル */
}
</style>
