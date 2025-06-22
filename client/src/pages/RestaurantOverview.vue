<script lang="ts" setup>
import RestaurantInfo from '../components/RestaurantInfo.vue'
import ReviewList from '../components/ReviewList.vue'
import useParam from '../lib/param'

const eateryId = useParam('eateryId')
</script>

<template>
  <main>
    <RestaurantInfo :eatery-id="eateryId"/>
    <review-list :eatery-id="eateryId"/>
  </main>
</template>

<style lang="scss" module></style>
