<script lang="ts" setup>
interface Props {
  text: string
}

const { text } = defineProps<Props>()
</script>

<template>
  <button :class="$style.pButton">{{ text }}</button>
</template>

<style lang="scss" module>
.pButton {
  color: white;
  font-size: 1.2rem;
  font-weight: 1000;
  border: 1px solid $color-secondary;
  background-color: #ffad3a;
}
</style>
