<script setup lang="ts">
import { ref, watch } from 'vue'
import apis, { type Eatery } from '../lib/apis'

interface Props {
  eateryId: string
}

const { eateryId } = defineProps<Props>()
const eateryDetail = ref<Eatery | null>(null)

const fetchEateryDetail = async () => {
  try {
    eateryDetail.value = (await apis.eateriesEateryIdGet(eateryId)).data
  } catch (error) {
    console.error('店舗詳細の取得に失敗しました:', error)
  }
}

watch(() => eateryId, fetchEateryDetail, { immediate: true })
</script>

<template>
  <div :class="$style.shopInfo">
    <div :class="$style.review">
      <div>店名: {{ eateryDetail?.name }}</div>
      <div>説明: {{ eateryDetail?.description }}</div>
    </div>
  </div>
</template>

<style lang="scss" module>
.shopInfo {
  display: flex;
  flex-direction: column;
  background-color: $color-background;
}
.review {
  display: flex;
  flex-direction: column;
}
.review div {
  font-size: 32px;
  color: $color-text;
}
</style>
