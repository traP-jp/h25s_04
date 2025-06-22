<script setup lang="ts">
import 'leaflet/dist/leaflet.css'
import { ref } from 'vue'
import { LMap, LTileLayer, LMarker } from '@vue-leaflet/vue-leaflet'
import apis, { type Eatery } from '../lib/apis'
import EateryList from '../components/EateryList.vue'

const eateries = ref<Eatery[]>([])
apis
  .eateriesGet()
  .then((res) => {
    eateries.value = res.data.data ?? []
    console.log(res.data)
  })
  .catch((error) => {
    console.error('Error fetching eateries:', error)
  })

const zoom = ref(16)
const center = ref([35.605958, 139.68354]) // 地図の中心座標
const markerPosition = ref([35.605958, 139.68354]) // ピンを立てる座標
</script>

<template>
  <div :class="$style.map">
    <l-map :zoom="zoom" :use-global-leaflet="false" :center="center">
      <l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      <l-marker :lat-lng="markerPosition"></l-marker>
    </l-map>
    <div :class="$style.list">
      <eatery-list :eateries="eateries" />
    </div>
  </div>
</template>

<style lang="scss" module>
.map {
  height: calc(100% - 150px);
  width: 100%;
  position: relative;
}

.list {
  color: black;
  position: absolute; /* 絶対位置に変更 */
  z-index: 1000; /* 地図より前面に表示 */
  top: 100px; /* 必要に応じて調整 */
  left: 20px; /* 必要に応じて調整 */
  background-color: rgba(255, 255, 255, 1);
  padding: 16px;
  border-radius: 12px;
  box-shadow: 16px 16px 20px rgba(0, 0, 0, 0.2);
}
</style>
