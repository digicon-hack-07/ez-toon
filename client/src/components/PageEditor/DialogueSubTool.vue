<script setup lang="ts">

const emit = defineEmits<{
  (e: 'deleteDialogue'): void
}>()
const deleteDialogue = (e: MouseEvent) => { emit("deleteDialogue") }

</script>

<template>
  <button @click="deleteDialogue">消</button>
</template>

<style scoped>
button {
  border-radius: 3rem;
  border: 2px solid #222;
  height: 6rem;
  width: 6rem;
  position: absolute;
  left: 1rem;
  bottom: 1rem;
  background-color: #666;
  color: #FFF;
}
button:active, button:focus, button:focus {
  outline: none;
  border: 2px solid #222;
}

</style>
