<script lang="ts" setup>
interface Props {
  text: string
}

const { text } = defineProps<Props>()
</script>

<template>
  <button class="sbutton">{{ text }}</button>
</template>

<style lang="scss" module>
.sbutton {
  font-weight: 1000;
  border: 1px solid $color-secondary;
  background-color: #b2bbc7;
}
</style>
