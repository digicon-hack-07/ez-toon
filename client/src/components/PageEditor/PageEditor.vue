<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { type PathRenderData, drawPath } from '../../lib/renderer/path'
import { type Dialogue } from '../../lib/dialogue'

const canvas = ref<HTMLCanvasElement>()
const ctx = ref<CanvasRenderingContext2D>()

interface Props {
  pageWidth: number
  pageHeight: number
}

const props = withDefaults(defineProps<Props>(), {
  pageWidth: 707,
  pageHeight: 1000
})

const canvasScrollX = ref<number>(10)
const canvasScrollY = ref<number>(20)
const canvasScale = ref<number>(1.0)

type EditMode = 'move' | 'pen' | 'dialogue' | 'eraser'
const mode = ref<EditMode>('pen')

onMounted(() => {
  if (canvas.value) {
    canvas.value.width = props.pageWidth * canvasScale.value
    canvas.value.height = props.pageHeight * canvasScale.value
    ctx.value = canvas.value.getContext('2d') ?? undefined
  }
})

const canvasCss = computed(() => {
  return {
    left: `${canvasScrollX.value}px`,
    top: `${canvasScrollY.value}px`,
    height: `${Math.floor(props.pageHeight * canvasScale.value)}px`,
    width: `${Math.floor(props.pageWidth * canvasScale.value)}px`
  }
})

let draggingData: {
  pointerId: number
  beginX: number
  beginY: number
  tgtBeginX: number
  tgtBeginY: number
} | null = null

const paths: PathRenderData[] = []
const working_path = new Map<number, PathRenderData>()

const dialogue_default_width = 100
const dialogue_default_height = 100
const dialogues = ref<Dialogue[]>([])

const dialogues_display = computed(() => {
  return dialogues.value.map(dialogue => {
    return {
      id: dialogue.id,
      style: {
        left: `${canvasScrollX.value + dialogue.left}px`,
        top: `${canvasScrollY.value + dialogue.top}px`,
        width: `${dialogue.right - dialogue.left}px`,
        height: `${dialogue.bottom - dialogue.top}px`
      },
      str: dialogue.dialogue
    }
  })
})

const pointerdown = (e: PointerEvent) => {
  const bx = canvas.value?.getClientRects().item(0)?.left
  const by = canvas.value?.getClientRects().item(0)?.top
  if (bx === undefined || by === undefined) return
  switch (mode.value) {
  case 'move':
    if (draggingData) break
    draggingData = {
      pointerId: e.pointerId,
      beginX: e.x,
      beginY: e.y,
      tgtBeginX: canvasScrollX.value,
      tgtBeginY: canvasScrollY.value
    }
    break
  case 'pen':
    working_path.set(e.pointerId, [
      {
        x: e.clientX - bx,
        y: e.clientY - by
      }
    ])
    break
  case 'dialogue':
    dialogues.value.push({
      id: `${dialogues.value.length ? dialogues.value.slice(-1)[0].id + 1 : 0}`,  // TODO: generate ULID
      pageID: '', // TODO: pageID
      dialogue: '',
      left: e.clientX - bx - dialogue_default_width,
      top: e.clientY - by,
      right: e.clientX - bx,
      bottom: e.clientY - by + dialogue_default_height
    })
    break
  }
}
const pointermove = (e: PointerEvent) => {
  const bx = canvas.value?.getClientRects().item(0)?.left
  const by = canvas.value?.getClientRects().item(0)?.top
  if (bx === undefined || by === undefined) return

  switch (mode.value) {
  case 'move':
    if (!draggingData || draggingData.pointerId != e.pointerId) break
    canvasScrollX.value = e.x - draggingData.beginX + draggingData.tgtBeginX
    canvasScrollY.value = e.y - draggingData.beginY + draggingData.tgtBeginY
    break
  case 'pen': {
    const path = working_path.get(e.pointerId)
    if (!path) return

    path.push({
      x: e.clientX - bx,
      y: e.clientY - by
    })
    if (ctx.value) drawPath(ctx.value, path)
  } break
  case 'dialogue':
    break
  }
}
const pointerup = (e: PointerEvent) => {
  switch (mode.value) {
  case 'move':
    draggingData = null
    break
  case 'pen':{
    const path = working_path.get(e.pointerId)
    if (!path) return

    paths.push(path)
    working_path.delete(e.pointerId)
  } break
  case 'dialogue':
    break
  }
}

const selectModeMove = () => {
  mode.value = 'move'
}
const selectModeDialogue = () => {
  mode.value = 'dialogue'
}
const selectModePen = () => {
  mode.value = 'pen'
}
const selectModeEraser = () => {
  mode.value = 'eraser'
}
</script>

<template>
  <div class="pageeditor">
    <div class="canvas-container" :data-editmode="mode">
      <div
        v-for="dialogue_display in dialogues_display"
        :key="dialogue_display.id"
        class="dialogue"
        contenteditable="true"
        :style="dialogue_display.style"
      >
        {{ dialogue_display.str }}
      </div>
      <canvas
        ref="canvas"
        :style="canvasCss"
        @pointerdown="pointerdown"
        @pointermove="pointermove"
        @pointerup="pointerup"
        @pointerout="pointerup"
      ></canvas>
    </div>
    <div class="button-container">
      <button :data-active="mode == 'move'" @click="selectModeMove">移</button>
      <button :data-active="mode == 'dialogue'" @click="selectModeDialogue">
        あ
      </button>
      <button :data-active="mode == 'pen'" @click="selectModePen">筆</button>
      <button :data-active="mode == 'eraser'" @click="selectModeEraser">
        消
      </button>
    </div>
  </div>
</template>

<style scoped>
@import '../../font.css';

.pageeditor {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
}
.canvas-container {
  width: 100%;
  overflow: hidden;
  flex-grow: 1;
  position: relative;
}
.dialogue {
  position: absolute;
  border: 1px solid #999;
  outline: none;
  writing-mode: vertical-rl;
  line-height: 1.2;
  word-break: break-all;
  font-family: 'EzTooN-SourceHanSerif', serif;
  z-index: -1;
  text-align: start;
}
.canvas-container[data-editmode='dialogue'] .dialogue {
  z-index: 9999;
}
canvas {
  position: absolute;
  touch-action: none;
  border: 1px solid;
}
.button-container {
  height: 6rem;
  display: flex;
}
button {
  flex-grow: 1;
  border: 2px solid #888;
  border-bottom: none;
  border-radius: 0;
}
button:focus {
  outline: none;
}
button[data-active='true'] {
  background-color: #3c70ff;
  color: #fff;
}
</style>
