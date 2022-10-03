<script setup lang="ts">
const props = defineProps<{
  penSize: number
}>()
const emit = defineEmits<{
  (e: 'changePenSize', size: number): void
}>()

const minPenSize = 1
const maxPenSize = 50

let frameDrag = false
let dragData: {
  beginX: number
  beginY: number
  beginVal: number
} | null = null

const updateByFramePointerEvent = (elem: HTMLElement, x: number, y: number) => {
  const left = elem.getClientRects().item(0)?.left
  const rem = parseFloat(getComputedStyle(document.documentElement).fontSize)
  if (!left) return
  const rawMax = 12 * rem
  const rawValue = Math.min(Math.max(x - left - 2 * rem, 0), rawMax)
  const value = Math.floor(
    (rawValue / rawMax) * (maxPenSize - minPenSize) + minPenSize
  )
  changePenSize(value)
}

const updateByHandlePointerEvent = (x: number, y: number) => {
  if (!dragData) return
  const rem = parseFloat(getComputedStyle(document.documentElement).fontSize)
  const rawDVal =
    ((x - dragData.beginX) / (12 * rem)) * (maxPenSize - minPenSize)
  const value = Math.floor(
    Math.min(Math.max(dragData.beginVal + rawDVal, minPenSize), maxPenSize)
  )
  changePenSize(value)
}

const framePointerDown = (e: PointerEvent) => {
  if (e.currentTarget instanceof HTMLElement) {
    updateByFramePointerEvent(e.currentTarget, e.x, e.y)
    e.currentTarget.setPointerCapture(e.pointerId)
    frameDrag = true
  }
}
const framePointerMove = (e: PointerEvent) => {
  if(!frameDrag) return
  if (e.currentTarget instanceof HTMLElement) {
    updateByFramePointerEvent(e.currentTarget, e.x, e.y)
  }
}
const framePointerUp = (e: PointerEvent) => {
  if (e.currentTarget instanceof HTMLElement) {
    e.currentTarget.releasePointerCapture(e.pointerId)
    updateByFramePointerEvent(e.currentTarget, e.x, e.y)
    frameDrag = false
  }
}

const handlePointerDown = (e: PointerEvent) => {
  dragData = {
    beginVal: props.penSize,
    beginX: e.x,
    beginY: e.y
  }
  if (e.currentTarget instanceof HTMLElement)
    e.currentTarget.setPointerCapture(e.pointerId)
}
const handlePointerMove = (e: PointerEvent) => {
  updateByHandlePointerEvent(e.x, e.y)
}
const handlePointerUp = (e: PointerEvent) => {
  updateByHandlePointerEvent(e.x, e.y)
  dragData = null
  if (e.currentTarget instanceof HTMLElement)
    e.currentTarget.releasePointerCapture(e.pointerId)
}

const changePenSize = (value: number) => {
  emit('changePenSize', value)
}
</script>

<template>
  <div>
    <div
      class="slider"
      @pointerdown="framePointerDown"
      @pointermove.stop="framePointerMove"
      @pointerup.stop="framePointerUp"
    >
      <div
        @pointerdown.stop="handlePointerDown"
        @pointermove.stop="handlePointerMove"
        @pointerup.stop="handlePointerUp"
        :style="{
          left: `calc(1rem + ${
            ((penSize - minPenSize) / (maxPenSize - minPenSize)) * 12
          }rem)`
        }"
        class="slider-handle"
      ></div>
    </div>
  </div>
</template>

<style scoped>
.slider {
  background-color: #eee;
  border-radius: 2rem;
  border: 0.1rem #222 solid;
  width: 16rem;
  height: 4rem;
  margin: 0 auto 1rem;
  position: relative;
  box-sizing: border-box;
}
.slider:before {
  content: '';
  width: 12rem;
  border: 0.1rem #222 solid;
  display: block;
  margin: 1.9rem auto;
  box-sizing: border-box;
}
.slider-handle {
  border: 0.1rem #222 solid;
  width: 2rem;
  height: 2rem;
  border-radius: 1rem;
  position: absolute;
  background-color: #eee;
  top: 1rem;
  left: 1rem;
  box-sizing: border-box;
}
.slider-handle:hover {
  background-color: rgb(156, 233, 255);
}

button {
  border-radius: 3rem;
  border: 2px solid #222;
  height: 6rem;
  width: 6rem;
  position: absolute;
  left: 1rem;
  bottom: 1rem;
  background-color: #666;
  color: #fff;
}
button:active,
button:focus,
button:focus {
  outline: none;
  border: 2px solid #222;
}
</style>
