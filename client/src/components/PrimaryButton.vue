<script lang="ts" setup>
interface Props {
  text: string
}

const {text} = defineProps<Props>()
</script>

<template>
  <button class="omise">{{text}}</button>
</template>

<style lang="scss" modele>
.omise {
  font-weight: 1000;
  border: 1px solid $color-secondary;
  background-color: #ffad3a;
}
</style>
