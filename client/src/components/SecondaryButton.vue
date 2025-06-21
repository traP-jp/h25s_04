<script lang="ts" setup>
interface Props {
  text: string
}

const { text } = defineProps<Props>()
</script>

<template>
  <button class="sinki">{{ text }}</button>
</template>

<style lang="scss" modele>
.sinki {
  font-weight: 1000;
  border: 1px solid $color-secondary;
  background-color: #b2bbc7;
}
</style>
