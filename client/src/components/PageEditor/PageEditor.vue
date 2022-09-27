<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { type Dialogue } from '../../lib/dialogue'
import { type ToolHandlerInterface } from '../../lib/edittools/ToolHandlerInterface'
import { MoveToolHandler } from '../../lib/edittools/MoveToolHandler'
import { DialogueToolHandler } from '../../lib/edittools/DialogueToolHandler'
import { PenToolHandler } from '../../lib/edittools/PenToolHandler'
import { EraserToolHandler } from '../../lib/edittools/EraserToolHandler'
import { type Line } from '../../lib/line'
import DialogueSubTool from './DialogueSubTool.vue'

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
const canvasScale = ref<number>(0.5)

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

const lines: Line[] = []
const dialogues = ref<Dialogue[]>([])
const dialogue_selected = ref<string | null>(null)
const dialogues_dummy = ref(new Map<string, string>)

const dialogues_display = computed(() => {
  return dialogues.value.map(dialogue => {
    return {
      id: dialogue.id,
      style: {
        left: `${canvasScrollX.value + dialogue.left * canvasScale.value}px`,
        top: `${canvasScrollY.value + dialogue.top * canvasScale.value}px`,
        width: `${dialogue.right - dialogue.left}px`,
        height: `${dialogue.bottom - dialogue.top}px`,
        fontSize: `${dialogue.fontSize}px`,
        transform: `scale(${canvasScale.value})`
      },
      str: dialogue.dialogue
    }
  })
})
const dialogue_update = (e: Event) => {
  const tgt = e.target
  if (!tgt) return
  if (e instanceof InputEvent && tgt instanceof HTMLElement) {
    const dialogue = dialogues.value.find(p => {
      return p.id == tgt.dataset.id
    })
    if (!dialogue || e.data === null) return
    dialogue.dialogue = tgt.innerText
  }
}
const dialogue_select = (e: FocusEvent) => {
  if (e.target instanceof HTMLElement) {
    const id = e.target.dataset.id
    if (!id) return
    dialogue_selected.value = id
  }
}
const dialogue_delete = () => {
  dialogues.value = dialogues.value.filter(p => {
    return p.id != dialogue_selected.value
  })
}

function getModeHandler(): ToolHandlerInterface {
  switch (mode.value) {
    case 'move':
      return new MoveToolHandler(canvasScrollX, canvasScrollY)
    case 'dialogue':
      if (!canvas.value) throw new Error('canvas not loaded')
      return new DialogueToolHandler(
        canvas.value,
        canvasScale,
        dialogues,
        props.pageID
      )
    case 'pen':
      if (!canvas.value) throw new Error('canvas not loaded')
      if (!ctx.value || !workctx.value)
        throw new Error('canvas context not loaded')
      return new PenToolHandler(
        canvas.value,
        canvasScale,
        ctx.value,
        workctx.value,
        lines
      )
    case 'eraser':
      return new EraserToolHandler()
  }
}
let handler: ToolHandlerInterface | null

const pointerdown = (e: PointerEvent) => {
  if (handler) handler.pointerdown(e)
}
const pointermove = (e: PointerEvent) => {
  if (handler) handler.pointermove(e)
}
const pointerup = (e: PointerEvent) => {
  if (handler) handler.pointerup(e)
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
        :data-id="dialogue_display.id"
        @input="dialogue_update"
        @focus="dialogue_select"
      >{{ dialogues_dummy.get(dialogue_display.id) }}</div>
      <canvas ref="workcanvas" :style="canvasCss"></canvas>
    </div>
    <div class="subtool-container">
      <dialogue-sub-tool
        v-if="mode == 'dialogue'"
        @delete-dialogue="dialogue_delete"
      ></dialogue-sub-tool>
    </div>
    <div class="button-container">
      <button :data-active="mode == 'move'" @click="changeMode('move')">
        移
      </button>
      <button :data-active="mode == 'dialogue'" @click="changeMode('dialogue')">
        あ
      </button>
      <button :data-active="mode == 'pen'" @click="changeMode('pen')">
        筆
      </button>
      <button :data-active="mode == 'eraser'" @click="changeMode('eraser')">
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
  touch-action: none;
}
.canvas-container {
  width: 100%;
  overflow: hidden;
  flex-grow: 1;
  position: relative;
  background-color: #ccc;
}
.store-canvas {
  background-color: #fff;
}
.dialogue {
  position: absolute;
  border: 2px solid #999; /* 小さいと縮小時にブラウザのレンダリングがバグりやすい */
  writing-mode: vertical-rl;
  line-height: 1.2;
  word-break: break-all;
  font-family: 'EzTooN-SourceHanSerif', serif;
  text-align: start;
  transform-origin: top left;
}
.dialogue:active, .dialogue:focus, .dialogue:focus-visible,
.dialogue div:active, .dialogue div:focus, .dialogue div:focus-visible {
  outline: none;
}
.canvas-container[data-editmode='dialogue'] .dialogue {
  z-index: 9999;
}
canvas {
  position: absolute;
  border: 1px solid;
}
.subtool-container {
  position: absolute;
  width: 100%;
  bottom: 6rem;
  z-index: 10000;
}
.button-container {
  height: 6rem;
  display: flex;
}
.button-container button {
  flex-grow: 1;
  border: 2px solid #888;
  border-bottom: none;
  border-radius: 0;
}
.button-container button:focus {
  outline: none;
}
.button-container button[data-active='true'] {
  background-color: #3c70ff;
  color: #fff;
}
</style>
