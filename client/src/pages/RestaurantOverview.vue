<script lang="ts" setup>
import RestaurantInfo from '../components/RestaurantInfo.vue'
import { onMounted, ref } from 'vue'
import apis, { type Eatery } from '../lib/apis'
import useParam from '../lib/param'

const eateryId = useParam('eateryId')
</script>

<template>
  <main>
    <RestaurantInfo :eatery-id="eateryId" />
  </main>
</template>

<style lang="scss" module></style>
