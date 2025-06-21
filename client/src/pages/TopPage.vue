<script setup lang="ts">
import EateryTile from '../components/EateryTile.vue'
</script>

<template>
  <main>
    <h1 :class="$style.reviewTitle">最近投稿されたレビュー</h1>
    <EateryTile></EateryTile>
  </main>
</template>

<style lang="scss" module>
.reviewTitle {
  font-size: 36px;
  color: $color-primary-text;
  text-align: left;
}
</style>
