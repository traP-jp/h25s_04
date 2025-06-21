<script lang="ts" setup>
interface Props {
  text: string
}

const { text } = defineProps<Props>()
</script>

<template>
  <button :class="$style.sButton">{{ text }}</button>
</template>

<style lang="scss" module>
.sButton {
  font-weight: 1000;
  border: 1px solid $color-secondary;
  background-color: #b2bbc7;
}
</style>
