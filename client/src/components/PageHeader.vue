<script lang="ts" setup>
import PrimaryButton from './PrimaryButton.vue'
import SecondaryButton from './SecondaryButton.vue'
</script>

<template>
  <header :class="$style.header">
    <div>aaaa</div>
    <div>タイトル</div>
    <div class="buttons">
      <primary-button></primary-button>
      <secondary-button></secondary-button>
    </div>
  </header>
</template>

<style lang="scss" module>
.header {
  display: flex;
  background-color: #ffad3a;
  justify-content: space-between;
  align-items: center;
  padding: 8px;
}
</style>
