<script lang="ts" setup>
import PrimaryButton from './PrimaryButton.vue'
</script>

<template>
  <header :class="$style.header">
    <div>aaaa</div>
    <div>タイトル</div>
    <div class="buttons">
      <primary-button :text="'お店一覧'" />
      <primary-button :text="'新規投稿'" />
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
