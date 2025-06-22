<script setup lang="ts">
import { onMounted, ref } from 'vue'
import apis, { type Eatery } from '../lib/apis'

interface Props {
  eateryId: string
}

const { eateryId } = defineProps<Props>()
const eateryDetail = ref<Eatery | null>(null)

onMounted(async () => {
  try {
    eateryDetail.value = (await apis.eateriesEateryIdGet(eateryId)).data
  } catch (error) {
    console.error('店舗詳細の取得に失敗しました:', error)
  }
})
</script>

<template>
  <div :class="$style.info">
    <h2>{{ eateryDetail?.name ?? '店名が未設定です' }}</h2>
    <div>
      <p>{{ eateryDetail?.description ?? '説明はありません' }}</p>
    </div>
  </div>
</template>

<style lang="scss" module>
.info {
  padding: 12px;
  justify-items: left;
}

h2 {
  position: relative;
  width: 100%;
  text-align: left;
}

h2::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 1px; /* Adjust thickness of the underline */
  background-color: black; /* Adjust color of the underline */
}
</style>
