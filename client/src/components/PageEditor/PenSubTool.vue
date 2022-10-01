<script setup lang="ts">

const props = defineProps<{
  penSize: number
}>()
const emit = defineEmits<{
  (e: 'changePenSize', size: number): void
}>()

const changePenSize = (e: Event) => {
  if(e.target instanceof HTMLInputElement)
    emit("changePenSize", e.target.valueAsNumber)
}

</script>

<template>
  <div><input type="range" min="1" max="50" :value="penSize" @input="changePenSize" /></div>
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
