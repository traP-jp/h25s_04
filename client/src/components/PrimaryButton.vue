<script lang="ts" setup>
interface Props {
  text: string
}

const { text } = defineProps<Props>()
</script>

<template>
  <button class="pbutton">{{ text }}</button>
</template>

<style lang="scss" module>
.pbutton {
  font-weight: 1000;
  border: 1px solid $color-secondary;
  background-color: #ffad3a;
}
</style>
