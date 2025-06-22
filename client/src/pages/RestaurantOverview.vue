<script lang="ts" setup>
import RestaurantInfo from '../components/RestaurantInfo.vue'
import { onMounted, ref } from 'vue'
import apis, { type Eatery } from '../lib/apis'
import useParam from '../lib/param'

const eateryId = useParam('eateryId')
const restaurantInfo = ref<Eatery | null>(null)

onMounted(async () => {
  try {
    const response = await apis.eateriesEateryIdGet(eateryId) // Use dynamic eateryId
    restaurantInfo.value = response.data
  } catch (error) {
    console.error('店舗詳細の取得に失敗しました:', error)
  }
})
</script>

<template>
  <main>
    <RestaurantInfo :eatery-id="eateryId" />
  </main>
</template>

<style lang="scss" module></style>
