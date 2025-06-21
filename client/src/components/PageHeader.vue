<script lang="ts" setup>
import PrimaryButton from './PrimaryButton.vue'
</script>

<template>
  <header :class="$style.header">
    <RouterLink to="/">
      <div>ここにロゴを入れる</div>
    </RouterLink>
    <div>タイトル</div>
    <div :class="$style.buttons">
      <RouterLink to="/eatery">
        <PrimaryButton :text="'お店一覧'" />
      </RouterLink>
      <RouterLink to="/eatery/new">
        <PrimaryButton :text="'新規投稿'" />
      </RouterLink>
    </div>
  </header>
</template>

<style lang="scss" module>
.header {
  display: flex;
  height: 75px;
  background-color: #ffad3a;
  justify-content: space-between;
  align-items: center;
  padding: 8px;
}

.buttons {
  display: flex;
  gap: 8px;
}
</style>
