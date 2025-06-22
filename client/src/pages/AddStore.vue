<script lang="ts" setup>
import PrimaryButton from '../components/PrimaryButton.vue'
import { ref } from 'vue'

const storeName = ref('')
const description = ref('')
const address = ref('')
const reviewContent = ref('')

</script>

<template>
  <main>
    <div :class="$style.inputMenu">
        <div :class="$style.inputGroup">
          <label :class="$style.infoLabel">店名</label>
          <input
            v-model="storeName"
            class="input"
            type="text"
            size="20"
            placeholder="入力"
          />
        </div>
        <div :class="$style.inputGroup">
          <label :class="$style.infoLabel">営業日/営業時間</label>
          <input
            v-model="description"
            class="input"
            type="text"
            size="20"
            placeholder="入力"
          />
        </div>
        <div :class="$style.inputGroup">
          <label :class="$style.infoLabel">座標</label>
          <input
            v-model="address"
            class="input"
            type="text"
            size="20"
            placeholder="入力"
          />
        </div>
        <div :class="$style.inputGroup">
          <label :class="$style.infoLabel">レビュー本文</label>
          <textarea
            v-model="reviewContent"
            class="input"
            rows="5"
            size="20"
            placeholder="入力"
          ></textarea>
        </div>
        <div :class="$style.inputGroup">
          <label :class="$style.infoLabel">お店の写真</label
          ><input class="fileInput" type="file" accept="image/*" />
        </div>
      
        <div>
          <PrimaryButton :text="'投稿'" />
        </div>
    </div>
  </main>
</template>

<style lang="scss" module>

.inputMenu {
  display: flex;
  flex-direction: column;
  background-color: $color-background;
  gap: 1rem;
}
.inputGroup {
  display: flex;
  flex-direction: row;
  justify-content: center;
  gap: 1rem;
  width: 100%;
}
.input {
  color: $color-text;
  background-color: $color-background;
  border: 1px solid $color-secondary;
  border-radius: 3px;
  gap: 1rem;
  width: 200px;
  align-items: center;
}
.fileInput {
  border: 1px solid $color-secondary;
  border-radius: 3px;
  width: 200px;
}

.infoLabel {
  color: $color-text;
  font-size: 16px;
  width: 200px;
  text-align: left;
}
</style>
