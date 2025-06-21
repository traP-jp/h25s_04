<template>
        <div :class="$style.tile">
            <div>1</div>
            <div>2</div>
            <div>3</div>
            <div>3</div>
            <div>3</div>
            <div>3</div>
            <div>3</div>
            <div>3</div>
            <div>3</div>
            <div>3</div>
            <div>3</div>
            <div>3</div>
        </div>
  </template>
  
  <style lang="scss" module>
  .tile {
    display: flex;
    flex-wrap: wrap;
    background-color: $color-background;
    height: 80%;
    width: 100%;
    gap: 0.5rem;
    padding: 0.5rem;
  }

.tile div {
    font: 100 24px/100px sans-serif;
    outline: 0.25rem solid $color-secondary;
    background-color: $color-background;
    color: $color-text;
    height: 200px;
    width: 200px;
    text-align: center;
    border-radius: 1rem;
}

  </style>