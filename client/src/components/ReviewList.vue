<script lang="ts" setup>
import { ref, watch } from 'vue'
import { type ReviewDetail } from '../lib/apis'
import apis from '../lib/apis'
import ReviewItem from './ReviewItem.vue'

interface Props {
  eateryId: string
}

const { eateryId } = defineProps<Props>()
const reviewsList = ref<ReviewDetail[]>([
  {
    imageIds: ['https://q.trap.jp/api/v3/public/icon/Pugma'],
    eateryName: 'ぷぐま',
    authorId: '偉大な先輩',
    id: '0',
    createdAt: '2023-10-01T00:00:00Z',
    updatedAt: '2023-10-01T00:00:00Z',
    eateryId: '0',
    content: 'ぷぐまのレビュー',
  },
])

const fetchReviewsList = async () => {
  try {
    reviewsList.value = (
      await apis.eateriesEateryIdReviewsGet(eateryId)
    ).data.data
  } catch (error) {
    console.error('店舗詳細の取得に失敗しました:', error)
  }
}
watch(() => eateryId, fetchReviewsList, { immediate: true })
const sort = ref('')
</script>

<template>
  <div :class="$style.reviewSort">
    <select v-model="sort">
      <option>最新順</option>
      <option>近い順</option>
    </select>
  </div>
  <div v-for="review in reviewsList" :key="review.id" :class="$style.tile">
    <ReviewItem :review="review" />
  </div>
</template>

<style lang="scss" module>
.reviewSort {
  display: flex;
  flex-basis: auto;
  gap: 2rem;
}
</style>
