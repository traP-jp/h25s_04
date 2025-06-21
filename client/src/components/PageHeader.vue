<script lang="ts" setup>
import PrimaryButton from './PrimaryButton.vue'
</script>

<template>
  <header :class="$style.header">
    <RouterLink to="/"><div>aaaa</div></RouterLink>
    <div>タイトル</div>
    <div class="buttons">
      <RouterLink to="/eatery"
        ><PrimaryButton>お店一覧</PrimaryButton></RouterLink
      >
      <RouterLink to="/eatery/new"
        ><PrimaryButton>新規投稿</PrimaryButton></RouterLink
      >
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
