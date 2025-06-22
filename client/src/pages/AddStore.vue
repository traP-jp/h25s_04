<script lang="ts" setup>
import PrimaryButton from '../components/PrimaryButton.vue'
import { ref } from 'vue'
import apis from '../lib/apis'

const storeName = ref('')
const description = ref('')
const address = ref('')
const reviewContent = ref('')
const storePhotos = ref<File[]>([])
const fileInputRef = ref<HTMLInputElement | null>(null)
const previewImages = ref<string[]>([])

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    storePhotos.value = Array.from(target.files)
  }
}

const uploadImages = async () => {
  if (storePhotos.value.length === 0) {
    throw new Error('画像が選択されていません')
  }

  try {
    const uploadPromises = storePhotos.value.map((photo) =>
      apis.imagesPost(photo),
    )
    const responses = await Promise.all(uploadPromises)
    return responses.map((response) => response.data.id) // サーバーから返される画像IDを取得
  } catch (error) {
    console.error('画像のアップロードに失敗しました:', error)
    throw new Error('画像のアップロードに失敗しました')
  }
}

const fetchPreviewImages = async (imageIds: string[]) => {
  try {
    const fetchPromises = imageIds.map((id) => apis.imagesImageIdGet(id))
    const responses = await Promise.all(fetchPromises)
    previewImages.value = responses.map((response) =>
      URL.createObjectURL(response.data),
    )
  } catch (error) {
    console.error('プレビュー画像の取得に失敗しました:', error)
    throw new Error('プレビュー画像の取得に失敗しました')
  }
}

const submitStore = async () => {
  try {
    // 画像をアップロードしてIDを取得
    const imageIds = await uploadImages()
    await fetchPreviewImages(imageIds) // 画像ID取得後にプレビュー画像をフェッチ

    // 店舗情報をアップロード
    const eateryResponse = await apis.eateriesPost({
      name: storeName.value,
      description: description.value,
      latitude: 0, // TODO: 必要に応じて座標を設定
      longitude: 0, // TODO: 必要に応じて座標を設定
    })

    const eateryId = eateryResponse.data.id // サーバーから返される店舗IDを取得

    // レビューをアップロード
    const xForwardedUser = 'exampleUserId' // TODO: 実際のログインユーザーIDを取得
    await apis.eateriesEateryIdReviewsPost(eateryId, {
      content: reviewContent.value,
      imageIds: imageIds,
      authorId: xForwardedUser, // 必須フィールドを追加
    })

    alert('店舗とレビューが追加されました！')
  } catch (error) {
    console.error('店舗またはレビューの追加に失敗しました:', error)
    alert('店舗またはレビューの追加に失敗しました。')
  }
}
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
            id="fileInput"
            ref="fileInputRef"
            type="file"
            accept="image/*"
            multiple
            :class="$style.hiddenFileInput"
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
      <div v-if="storePhotos.length > 0" :class="$style.previewContainer">
        <div
          v-for="(photo, index) in storePhotos"
          :key="index"
          :class="$style.previewImageWrapper"
        >
          <img
            v-if="previewImages[index]"
            :src="previewImages[index]"
            :alt="'プレビュー画像' + (index + 1)"
            :class="$style.previewImage"
          />
        </div>
      </div>
      <div v-if="previewImages.length > 0" :class="$style.previewContainer">
        <div
          v-for="(preview, index) in previewImages"
          :key="index"
          :class="$style.previewImageWrapper"
        >
          <img
            :src="preview"
            :alt="`プレビュー画像${index + 1}`"
            :class="$style.previewImage"
          />
        </div>
      </div>
      {{
        storePhotos.length > 0
          ? storePhotos.map((photo) => photo.name).join(', ')
          : '写真が選択されていません'
      }}

      <div>
        <PrimaryButton :text="'投稿'" @click="submitStore" />
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
  width: 200px;
}
.previewContainer {
  display: flex;
  justify-content: center;
  margin-top: 1rem;
}
.previewImageWrapper {
  position: relative;
  display: inline-block;
  margin: 0 0.5rem;
}
.previewImage {
  max-width: 100%;
  max-height: 200px;
  border: 1px solid $color-secondary;
  border-radius: 3px;
}
</style>
