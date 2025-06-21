<script setup lang="ts">
import { onMounted, ref } from 'vue'
import apis, { type Eatery } from '../lib/apis'
import useParam from '../lib/param'

const eateryId = useParam('eateryId')

const eateryDetail = ref<Eatery>()
onMounted(async () => {
  eateryDetail.value = (await apis.eateriesEateryIdGet(eateryId.value)).data
})
</script>

<template>
  <div :class="$style.shopInfo">
    <div :class="$style.review">
      <div>店名: {{ eateryDetail?.name }}</div>
      <div>説明: {{ eateryDetail?.description }}</div>
      <div></div>
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
