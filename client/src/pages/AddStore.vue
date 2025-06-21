<template>
  <div class="honbun">
    <div>
      <div>
       <span style="color: #000000">店名</span>
       <input
         v-model="message1"
         class="nyuuryoku4"
         type="text"
         placeholder="入力"
       />
      </div>
      <div>
       <span style="color: #000000">説明</span>
       <input
         v-model="message2"
         class="nyuuryoku4"
         type="text"
         placeholder="入力"
       />
      </div>
      <div>
       <input
         v-model="message3"
         class="nyuuryoku4"
         type="text"
         placeholder="入力"
       />
      </div>
      <div>
       <span style="color: #000000">住所／座標</span>
       <input
         v-model="message4"
         class="nyuuryoku4"
         type="text"
         placeholder="入力"
       />
      </div>
      <div>
       <span style="color: #000000">レビュー本文</span>
       <textarea
         v-model="inputText"
         class="nyuuryoku4"
         rows="5"
         placeholder="入力"
       ></textarea
       >
      </div>
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
import { ref } from 'vue'

const message1 = ref('')
const message2 = ref('')
const message3 = ref('')
const message4 = ref('')
const inputText = ref('')
</script>

<style>
.honbun {
  background-color: #ffffff;
}
.nyuuryoku4{
  color: #000000;
  background-color: #ffffff;
  border: 1px solid #b2bbc7;
  border-radius: 3px;
}
</style>
