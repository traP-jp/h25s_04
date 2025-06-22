<script lang="ts" setup>
import { type ReviewDetail } from '../lib/apis'

interface Props {
  review: ReviewDetail
}

const { review } = defineProps<Props>()

// const reviews = ref<ReviewDetail[]>([
//   {
//     imageIds: ['https://q.trap.jp/api/v3/public/icon/Pugma'],
//     eateryName: 'ぷぐま',
//     authorId: '偉大な先輩',
//     id: '0',
//     createdAt: '2023-10-01T00:00:00Z',
//     updatedAt: '2023-10-01T00:00:00Z',
//     eateryId: '0',
//     content: 'ぷぐまのレビュー',
//   },
// ])
</script>

<template>
  <div>
    <div style="display: flex">
      <div>
        <div v-for="image in review.imageIds" :key="image">
          <img :class="$style.foodImage" :src="image" alt="food" />
        </div>
      </div>
      <div>
        <div :class="$style.restaurantName">{{ review.eateryName }}</div>
        <div :class="$style.username">{{ review.authorId }}</div>
        <div>
          <div :class="$style.reviewsummary">{{ review.content }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" module>
.tile {
  display: flex;
  flex-direction: column;
  flex-basis: auto;
  color: $color-text;
  text-align: left;
}

.foodImage {
  width: 200px;
  height: 150px;
  object-fit: cover;
}

.restaurantName {
  line-height: 1;
  font-size: 24px;
  color: $color-text;
}
.username {
  line-height: 1;
  font-size: 16px;
  color: $color-text;
}

.reviewsummary {
  font-size: 16px;
  color: $color-text;
}
</style>
