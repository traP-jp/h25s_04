<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { type ReviewSummary } from '../lib/apis'
import apis from '../lib/apis'
import review from './Review.vue'

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
  <div v-for="review in reviews" :key="review.id" :class="$style.tile">
    <review></review>
  </div>
</template>

<style lang="scss" module>
.reviewSort {
  display: flex;
  flex-basis: auto;
  gap: 2rem;
}
</style>