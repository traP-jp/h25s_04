<script setup lang="ts">
import PageHeader from './components/PageHeader.vue'
</script>

<template>
  <page-header></page-header>
  <main :class="$style.main">
    <router-view />
  </main>
</template>

<style lang="scss" module>
.main {
  height: calc(100vh - 75px);
  padding: 16px;
  background-color: $color-background;
}
</style>
