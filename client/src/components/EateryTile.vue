<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { type ReviewSummary } from '../lib/apis'
import apis from '../lib/apis'
import { watch } from 'vue'

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

const reviewsSort = (sortOption: string) => {
  const compare = (a: ReviewSummary, b: ReviewSummary) => {
    if (a.createdAt < b.createdAt) {
      return 1
    } else if (a.createdAt > b.createdAt) {
      return -1
    } else {
      return 0
    }
  }
  if (sortOption == 'newest') {
    reviews.value.sort(compare)
  }
}

const sort = ref('')

watch(sort, reviewsSort)

onMounted(async () => {
  reviews.value = (await apis.reviewsGet()).data.data ?? []
  console.log(reviews.value)
})
</script>

<template>
  <div :class="$style.reviewSort">
    <select v-model="sort">
      <option value="newest">最新順</option>
    </select>
  </div>
  <div :class="$style.tileView">
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
.tileView {
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
