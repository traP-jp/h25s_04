<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { type ReviewDetail } from '../lib/apis'
import apis from '../lib/apis'

interface Props {
  review: ReviewDetail
}

const { review } = defineProps<Props>()

const imageUrls = ref<Record<string, string>>({})

onMounted(async () => {
  if (review.imageIds && review.imageIds.length > 0) {
    for (const imageId of review.imageIds) {
      try {
        const response = await apis.imagesImageIdGet(imageId, {
          responseType: 'blob',
        })
        imageUrls.value[imageId] = URL.createObjectURL(response.data as Blob)
      } catch (e) {
        console.error(e)
      }
    }
  }
})
</script>

<template>
  <div>
    <div style="display: flex">
      <div>
        <div v-for="image in review.imageIds" :key="image">
          <img
            v-if="imageUrls[image]"
            :class="$style.foodImage"
            :src="imageUrls[image]"
            alt="food"
          />
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
