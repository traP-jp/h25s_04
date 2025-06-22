<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { type ReviewSummary } from '../lib/apis'
import apis from '../lib/apis'

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
const sort = ref('')

onMounted(async () => {
  reviews.value = (await apis.reviewsGet()).data.data ?? []
  console.log(reviews.value)
})
</script>

<template>
  <div :class="$style.reviewSort">
    <select v-model="sort">
      <option>最新順</option>
      <option>近い順</option>
    </select>
  </div>
  <div>
    <div v-for="review in reviews" :key="review.id" :class="$style.tile">
      <div>
        <img :class="$style.foodImage" :src="review.imageIds[0]" alt="food" />
      </div>
      <div :class="$style.restaurantName">{{ review.eateryName }}</div>
      <div :class="$style.username">{{ review.authorId }}</div>
    </div>
  </div>
</template>

<style lang="scss" module>
.reviewSort {
  display: flex;
  flex-basis: auto;
  gap: 2rem;
}

.tile {
  display: flex;
  flex-direction: column;
  outline: 0.25rem solid $color-secondary;
  background-color: $color-background;
  color: $color-text;
  height: 200px;
  width: 1400px;
  text-align: center;
  border-radius: 1rem;
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
</style>
