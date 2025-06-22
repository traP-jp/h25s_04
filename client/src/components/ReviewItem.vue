<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { type ReviewSummary } from '../lib/apis'
import apis from '../lib/apis'
import review from './ReviewItem.vue'

const reviews = ref<ReviewSummary[]>([
  {
    imageIds: ['https://q.trap.jp/api/v3/public/icon/Pugma'],
    eateryName: 'ぷぐま',
    authorId: '偉大な先輩',
    id: '0',
    createdAt: '2023-10-01T00:00:00Z',
    updatedAt: '2023-10-01T00:00:00Z',
    eateryId: '0',
    summary: 'ぷぐまのレビュー',
  },
])

onMounted(async () => {
  reviews.value = (await apis.reviewsGet()).data.data ?? []
  console.log(reviews.value)
})
</script>

<template>
  <div>
    <div style="display: flex">
      <div>
        <img :class="$style.foodImage" :src="review.imageIds[0]" alt="food" />
      </div>
      <div>
        <div :class="$style.restaurantName">{{ review.eateryName }}</div>
        <div :class="$style.username">{{ review.authorId }}</div>
      </div>
    </div>
  </div>
  <div>
    <div :class="$style.reviewsummary">{{ review.summary }}</div>
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
