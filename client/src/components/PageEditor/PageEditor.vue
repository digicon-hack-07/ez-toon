<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { type PathRenderData, drawPath } from '../../lib/renderer/path'
import { type Dialogue } from '../../lib/dialogue'
import { ToolHandlerInterface } from '../../lib/edittools/ToolHandlerInterface'
import { MoveToolHandler } from '../../lib/edittools/MoveToolHandler'
import { DialogueToolHandler } from '../../lib/edittools/DialogueToolHandler'
import { PenToolHandler } from '../../lib/edittools/PenToolHandler'

const workcanvas = ref<HTMLCanvasElement>()
const canvas = ref<HTMLCanvasElement>()
const workctx = ref<CanvasRenderingContext2D>()
const ctx = ref<CanvasRenderingContext2D>()

interface Props {
  pageID: string
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
  if (canvas.value && workcanvas.value) {
    canvas.value.width = props.pageWidth * canvasScale.value
    canvas.value.height = props.pageHeight * canvasScale.value
    ctx.value = canvas.value.getContext('2d') ?? undefined

    workcanvas.value.width = props.pageWidth * canvasScale.value
    workcanvas.value.height = props.pageHeight * canvasScale.value
    workctx.value = workcanvas.value.getContext('2d') ?? undefined
    handler = getModeHandler()
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

const paths: PathRenderData[] = []
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

class EraserToolHandler implements ToolHandlerInterface {
  pointerdown(e: PointerEvent): void {
    throw new Error('Method not implemented.');
  }
  pointermove(e: PointerEvent): void {
    throw new Error('Method not implemented.');
  }
  pointerup(e: PointerEvent): void {
    throw new Error('Method not implemented.');
  }
}

function getModeHandler(): ToolHandlerInterface {
  switch(mode.value){
  case 'move':
    return new MoveToolHandler(canvasScrollX, canvasScrollY)
  case 'dialogue':
    if(!canvas.value)
      throw new Error('canvas not loaded')
    return new DialogueToolHandler(canvas.value, dialogues, props.pageID)
  case 'pen':
    if(!canvas.value)
        throw new Error('canvas not loaded')
    if(!ctx.value || !workctx.value)
      throw new Error('canvas context not loaded')
    return new PenToolHandler(canvas.value, ctx.value, workctx.value, paths)
  case 'eraser':
    return new EraserToolHandler()
  }
}
let handler: ToolHandlerInterface | null;

const pointerdown = (e: PointerEvent) => {
  if(handler) handler.pointerdown(e)
}
const pointermove = (e: PointerEvent) => {
  if(handler) handler.pointermove(e)
}
const pointerup = (e: PointerEvent) => {
  if(handler) handler.pointerup(e)
}

const changeMode = (new_mode: EditMode) => {
  mode.value = new_mode
  handler = getModeHandler()
}
</script>

<template>
  <div class="pageeditor">
    <div
      class="canvas-container"
      :data-editmode="mode"
      @pointerdown="pointerdown"
      @pointermove="pointermove"
      @pointerup="pointerup"
    >
      <canvas ref="canvas" class="store-canvas" :style="canvasCss"></canvas>
      <div
        v-for="dialogue_display in dialogues_display"
        :key="dialogue_display.id"
        class="dialogue"
        contenteditable="true"
        :style="dialogue_display.style"
      >
        {{ dialogue_display.str }}
      </div>
      <canvas ref="workcanvas" :style="canvasCss"></canvas>
    </div>
    <div class="button-container">
      <button :data-active="mode == 'move'" @click="changeMode('move')">移</button>
      <button :data-active="mode == 'dialogue'" @click="changeMode('dialogue')">あ</button>
      <button :data-active="mode == 'pen'" @click="changeMode('pen')">筆</button>
      <button :data-active="mode == 'eraser'" @click="changeMode('eraser')">消</button>
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
  background-color: #CCC;
}
.store-canvas {
  background-color: #FFF;
}
.dialogue {
  position: absolute;
  border: 1px solid #999;
  outline: none;
  writing-mode: vertical-rl;
  line-height: 1.2;
  word-break: break-all;
  font-family: 'EzTooN-SourceHanSerif', serif;
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
