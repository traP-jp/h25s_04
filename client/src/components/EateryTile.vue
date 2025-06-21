<script lang="ts" setup>
import {ref} from 'vue'

const sort = ref('')

</script>

<template>
    <div :class="$style.reviewsort">
        <div :class="$style.reviewtitle">最近投稿されたレビュー</div>
        <select v-model="sort">
            <option>最新順</option>
            <option>近い順</option>
        </select>
    </div>
    <div :class="$style.tileview">
        <div :class="$style.tile">
            <div>
                <img :class="$style.foodimage" src="https://q.trap.jp/api/v3/public/icon/Pugma" alt="food" />
            </div>
            <div :class="$style.restaurantname">ぷぐま</div>
            <div :class="$style.username">偉大な先輩</div>
        </div>
    </div>
  </template>
  
  <style lang="scss" module>
  .reviewtitle {
    font-size: 36px;
    color: $color-primary-text;
    text-align: left;
  }
  .reviewsort {
    display: flex;
    flex-basis: auto;
    gap: 2rem;
  }
  .tileview {
    display: flex;
    flex-wrap: wrap;
    background-color: $color-background;
    height: 80%;
    width: 100%;
    gap: 0.5rem;
    padding: 0.5rem;
  }

.tile {
    display: flex;
    flex-direction: column;
    outline: 0.25rem solid $color-secondary;
    background-color: $color-background;
    color: $color-text;
    height: 200px;
    width: 200px;
    text-align: center;
    border-radius: 1rem;
}

.foodimage {
    width: 200px;
    height: 150px;
    object-fit: cover;
}

.restaurantname {
    line-height: 1.0;
    font-size: 24px;
    color: $color-text;
}
.username {
    line-height: 1.0;
    font-size: 16px;
    color: $color-text;
}

  </style>