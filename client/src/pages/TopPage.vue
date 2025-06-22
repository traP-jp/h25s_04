<script setup lang="ts">
import EateryTile from '../components/EateryTile.vue'
</script>

<template>
  <h1 :class="$style.reviewTitle">最近投稿されたレビュー</h1>
  <EateryTile></EateryTile>
</template>

<style lang="scss" module>
.reviewTitle {
  font-size: 36px;
  text-align: left;
}
</style>
