<script setup lang="ts">

const emit = defineEmits<{
  (e: 'zoomIn'): void
  (e: 'zoomOut'): void
}>()
const zoomIn = (e: MouseEvent) => { emit("zoomIn") }
const zoomOut = (e: MouseEvent) => { emit("zoomOut") }

</script>

<template>
  <button style="bottom: 8rem;" @click="zoomIn">＋</button>
  <button style="bottom: 1rem;" @click="zoomOut">－</button>
</template>

<style scoped>
button {
  border-radius: 1.5rem;
  border: 2px solid #222;
  height: 6rem;
  width: 6rem;
  position: absolute;
  left: 1rem;
  background-color: #666;
  color: #FFF;
}
button:active, button:focus, button:focus {
  outline: none;
  border: 2px solid #222;
}

</style>
