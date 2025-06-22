<script lang="ts" setup>
import PrimaryButton from '../components/PrimaryButton.vue'
import { ref, computed } from 'vue'

const storeName = ref('')
const description = ref('')
const address = ref('')
const reviewContent = ref('')
const storePhoto = ref<File | null>(null)
const fileInputRef = ref<HTMLInputElement | null>(null)

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    storePhoto.value = target.files[0]
  }
}

const previewUrl = computed(() => {
  return storePhoto.value ? URL.createObjectURL(storePhoto.value) : null
})
</script>

<template>
  <main>
    <div :class="$style.inputMenu">
      <div :class="$style.inputGroup">
        <label for="storeName">店名</label>
        <input
          id="storeName"
          v-model="storeName"
          :class="$style.input"
          type="text"
          size="20"
          placeholder="入力"
        />
      </div>
      <div :class="$style.inputGroup">
        <label for="description">営業日/営業時間</label>
        <input
          id="description"
          v-model="description"
          :class="$style.input"
          type="text"
          size="20"
          placeholder="入力"
        />
      </div>
      <div :class="$style.inputGroup">
        <label for="address">座標</label>
        <input
          id="address"
          v-model="address"
          :class="$style.input"
          type="text"
          size="20"
          placeholder="入力"
        />
      </div>
      <div :class="$style.inputGroup">
        <label for="reviewContent">レビュー本文</label>
        <textarea
          id="reviewContent"
          v-model="reviewContent"
          :class="$style.input"
          rows="5"
          size="20"
          placeholder="入力"
        ></textarea>
      </div>
      <div :class="$style.inputGroup">
        <label :class="$style.infoLabel" for="storePhoto">お店の写真</label>
        <div>
          <input
            id="storePhoto"
            ref="fileInputRef"
            :class="$style.hiddenFileInput"
            type="file"
            accept="image/*"
            @change="handleFileChange"
          />
          <button
            type="button"
            :class="$style.customFileButton"
            @click="fileInputRef?.click()"
          >
            ファイルを選択
          </button>
        </div>
      </div>
      <div v-if="previewUrl" :class="$style.previewContainer">
        <img
          :src="previewUrl"
          :alt="'プレビュー画像'"
          :class="$style.previewImage"
        />
      </div>
      {{ storePhoto ? storePhoto.name : '写真が選択されていません' }}

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

  & > * {
    margin: 8px;
  }

  & > label {
    flex: 1;
    color: $color-text;
    font-size: 16px;
    width: 200px;
    text-align: right;
  }

  & > input,
  & > textarea,
  & > div {
    flex: 3;
    padding: 4px;
  }
}
.hiddenFileInput {
  display: none;
}
.customFileButton {
  background-color: $color-primary;
  color: $color-text;
  border: none;
  border-radius: 3px;
  padding: 0.5rem 1rem;
  cursor: pointer;
  font-size: 16px;
}
.customFileButton:hover {
  background-color: $color-primary;
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

.previewContainer {
  display: flex;
  justify-content: center;
  margin-top: 1rem;
}
.previewImage {
  max-width: 100%;
  max-height: 200px;
  border: 1px solid $color-secondary;
  border-radius: 3px;
}
</style>
