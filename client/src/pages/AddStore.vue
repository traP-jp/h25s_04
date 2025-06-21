<template>
  <main class="inputMenu">
    <div>
      <div>
        <span style="color: #000000">店名</span>
        <input
          v-model="storeName"
          class="input"
          type="text"
          placeholder="入力"
        />
      </div>
      <div>
        <span style="color: #000000">説明</span>
        <input
          v-model="description"
          class="input"
          type="text"
          placeholder="入力"
        />
      </div>
      <div>
        <input
          v-model="message3"
          class="input"
          type="text"
          placeholder="入力"
        />
      </div>
      <div>
        <span style="color: #000000">住所／座標</span>
        <input v-model="address" class="input" type="text" placeholder="入力" />
      </div>
      <div>
        <span style="color: #000000">レビュー本文</span>
        <textarea
          v-model="reviewContent"
          class="input"
          rows="5"
          placeholder="入力"
        ></textarea>
      </div>
      <span style="color: #000000">お店の写真</span
      ><input
        type="file"
        accept="image/*"
        style="border: 1px solid #b2bbc7; border-radius: 3px"
      />
    </div>
    <div>
      <PrimaryButton :text="'投稿'" />
    </div>
  </main>
</template>

<script lang="ts" setup>
import PrimaryButton from '../components/PrimaryButton.vue'
import { ref } from 'vue'

const storeName = ref('')
const description = ref('')
const message3 = ref('')
const address = ref('')
const reviewContent = ref('')
</script>

<style>
.inputMenu {
  background-color: #ffffff;
}
.input {
  color: #000000;
  background-color: #ffffff;
  border: 1px solid #b2bbc7;
  border-radius: 3px;
}
</style>
