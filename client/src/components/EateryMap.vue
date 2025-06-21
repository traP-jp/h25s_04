<script setup lang="ts">
import 'leaflet/dist/leaflet.css'
import { ref } from 'vue'
import { LMap, LTileLayer, LMarker } from '@vue-leaflet/vue-leaflet'

const zoom = ref(16)
const center = ref([35.605958, 139.68354]) // 地図の中心座標
const markerPosition = ref([35.605958, 139.68354]) // ピンを立てる座標
</script>

<template>
  <div style="height: 600px; width: 800px">
    <l-map v-model:zoom="zoom" :use-global-leaflet="false" :center="center">
      <l-tile-layer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      ></l-tile-layer>
      <l-marker :lat-lng="markerPosition"></l-marker>
    </l-map>
  </div>
</template>

<style lang="scss" module>
/* 必要に応じてスタイルを追加 */
</style>
