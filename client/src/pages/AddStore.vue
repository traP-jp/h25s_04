<template>
  <div class="honbun">
    <div>
      <span style="color: #000000">店名</span>
      <input
        v-model="message"
        type="text"
        style="
          color: #000000;
          background-color: #ffffff;
          border: 1px solid #b2bbc7;
          border-radius: 3px;
        "
        placeholder="店名を入力"
      /><br />
      <span style="color: #000000">説明</span>
      <input
        v-model="message"
        type="text"
        style="
          color: #000000;
          background-color: #ffffff;
          border: 1px solid #b2bbc7;
          border-radius: 3px;
        "
        placeholder="営業日を入力"
      /><br />
      <input
        v-model="message"
        type="text"
        style="
          color: #000000;
          background-color: #ffffff;
          border: 1px solid #b2bbc7;
          border-radius: 3px;
        "
        placeholder="営業時間を入力"
      /><br />
      <span style="color: #000000">住所／座標</span>
      <input
        v-model="message"
        type="text"
        style="
          color: #000000;
          background-color: #ffffff;
          border: 1px solid #b2bbc7;
          border-radius: 3px;
        "
        placeholder="住所／座標を入力"
      /><br />
      <span style="color: #000000">レビュー本文</span>
      <textarea
        v-model="inputText"
        rows="5"
        style="
          color: #000000;
          background-color: #ffffff;
          border: 1px solid #b2bbc7;
          border-radius: 3px;
        "
        placeholder="レビューを入力"
      ></textarea
      ><br />
      <span style="color: #000000">お店の写真</span
      ><input
        type="file"
        accept="image/*"
        style="border: 1px solid #b2bbc7; border-radius: 3px"
      />
    </div>
    <div>
      <post-button></post-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import PostButton from './PostButton.vue'

const message = ''
const inputText = ''
</script>

<style>
.honbun {
  background-color: #ffffff;
}
</style>
