<template>
  <div class="honbun">
    <div>
     <span style="color: #000000;">店名</span> <input type="text" style="color: #000000; background-color: #FFFFFF; border: 1px solid #B2BBC7; border-radius: 3px;" v-model="message" placeholder="店名を入力"><br>
     <span style="color: #000000;">説明</span> <input type="text" style="color: #000000; background-color: #FFFFFF; border: 1px solid #B2BBC7; border-radius: 3px" v-model="message" placeholder="営業日を入力"><br>
     <input type="text" style="color: #000000; background-color: #FFFFFF; border: 1px solid #B2BBC7; border-radius: 3px" v-model="message" placeholder="営業時間を入力"><br>
     <span style="color: #000000;">住所／座標</span> <input type="text"  style="color: #000000; background-color: #FFFFFF; border: 1px solid #B2BBC7; border-radius: 3px" v-model="message" placeholder="住所／座標を入力"><br>
     <span style="color: #000000;">レビュー本文</span> <textarea v-model="inputText" rows="5"  style="color: #000000; background-color: #FFFFFF; border: 1px solid #B2BBC7; border-radius: 3px" placeholder="レビューを入力"></textarea><br>
     <span style="color: #000000; ">お店の写真</span><input type="file" accept="image/*" style="border: 1px solid #B2BBC7; border-radius: 3px;">
   </div>
   <div>
    <post-button></post-button>
   </div>
  </div>
</template>

<script lang="ts" setup>
import { isBreakOrContinueStatement } from 'typescript';
import PostButton from './PostButton.vue'

const message = ""
const inputText = ""
const textColor = "#000000"

</script>

<style>
.honbun {
  background-color: #FFFFFF;
}
</style>