<script lang="ts" setup>
import type { Eatery } from '../lib/apis'

interface Props {
  eateries: Eatery[]
}

const { eateries } = defineProps<Props>()
</script>

<template>
  <ul>
    <li v-for="eatery in eateries" :key="eatery.id">
      <h2>{{ eatery.name ?? '店名' }}</h2>
      <p>
        {{ eatery.description ?? '説明がありません。' }}
      </p>
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
