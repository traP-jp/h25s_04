<template>
  <div class="honbun">
    <div>
     <span style="color: #000000;">店名</span> <input type="text" v-model="message" placeholder="入力"><br>
     <span style="color: #000000;">説明</span> <input type="text" v-model="message" placeholder="入力"><br>
     <input type="text" v-model="message" placeholder="入力"><br>
     <span style="color: #000000;">住所／座標</span> <input type="text" v-model="message" placeholder="入力"><br>
     <span style="color: #000000;">レビュー本文</span> <textarea v-model="inputText" rows="5" placeholder="入力"></textarea>
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