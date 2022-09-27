<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { type Dialogue } from '../../lib/dialogue'
import { type ToolHandlerInterface } from '../../lib/edittools/ToolHandlerInterface'
import { MoveToolHandler } from '../../lib/edittools/MoveToolHandler'
import { DialogueToolHandler } from '../../lib/edittools/DialogueToolHandler'
import { PenToolHandler } from '../../lib/edittools/PenToolHandler'
import { EraserToolHandler } from '../../lib/edittools/EraserToolHandler'
import { type Line } from '../../lib/line'
import DialogueSubTool from './DialogueSubTool.vue'
import MoveSubTool from './MoveSubTool.vue'
import { drawLine } from '../../lib/renderer/path'

const canvascontainer = ref<HTMLCanvasElement>()
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

const canvasScrollX = ref<number>(0)
const canvasScrollY = ref<number>(0)
const canvasScale = ref<number>(0.5)

type EditMode = 'move' | 'pen' | 'dialogue' | 'eraser'
const mode = ref<EditMode>('pen')

onMounted(() => {
  if (canvas.value && workcanvas.value && canvascontainer.value) {
    canvas.value.width = props.pageWidth * canvasScale.value
    canvas.value.height = props.pageHeight * canvasScale.value
    ctx.value = canvas.value.getContext('2d') ?? undefined

    workcanvas.value.width = props.pageWidth * canvasScale.value
    workcanvas.value.height = props.pageHeight * canvasScale.value
    workctx.value = workcanvas.value.getContext('2d') ?? undefined

    canvasScrollX.value =
      canvascontainer.value.clientWidth / 2 - canvas.value.clientWidth / 2
    canvasScrollY.value =
      canvascontainer.value.clientHeight / 2 - canvas.value.clientHeight / 2
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

watch(canvasScale, () => {
  if (canvas.value && workcanvas.value && ctx.value) {
    canvas.value.width = props.pageWidth * canvasScale.value
    canvas.value.height = props.pageHeight * canvasScale.value
    workcanvas.value.width = props.pageWidth * canvasScale.value
    workcanvas.value.height = props.pageHeight * canvasScale.value
    for (const line of lines) {
      drawLine(ctx.value, canvasScale.value, line)
    }
  }
})

const lines: Line[] = []
const dialogues = ref<Dialogue[]>([])
const dialogue_selected = ref<string | null>(null)
const dialogues_dummy = ref(new Map<string, string>())

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
const dialogues_handle_display = computed(() => {
  return dialogues.value.map(dialogue => {
    return {
      id: dialogue.id,
      style: {
        left: `${
          canvasScrollX.value + dialogue.right * canvasScale.value + 4
        }px`,
        top: `${
          canvasScrollY.value + dialogue.top * canvasScale.value - 12 - 4
        }px`,
        width: `12px`,
        height: `12px`
      }
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

const changeScale = (newScale: number) => {
  const oldScale = canvasScale.value
  if (canvascontainer.value && canvas.value) {
    canvasScrollX.value =
      canvascontainer.value.clientWidth / 2 -
      (newScale / oldScale) *
        (canvascontainer.value.clientWidth / 2 - canvasScrollX.value)
    canvasScrollY.value =
      canvascontainer.value.clientHeight / 2 -
      (newScale / oldScale) *
        (canvascontainer.value.clientHeight / 2 - canvasScrollY.value)
    canvasScale.value = newScale
    console.log(newScale)
  }
}
const zoomIn = () => {
  if (canvasScale.value < 5) changeScale(canvasScale.value + 0.25)
}
const zoomOut = () => {
  if (canvasScale.value > 0.25) changeScale(canvasScale.value - 0.25)
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
    if (!canvas.value) throw new Error('canvas not loaded')
    if (!ctx.value) throw new Error('canvas context not loaded')
    return new EraserToolHandler(canvas.value, ctx.value, canvasScale, lines)
  }
}
let handler: ToolHandlerInterface | null

const pointerdown = (e: PointerEvent) => {
  ;(e.target as HTMLElement).setPointerCapture(e.pointerId)
  if (handler) handler.pointerdown(e)
}
const pointermove = (e: PointerEvent) => {
  if (handler) handler.pointermove(e)
}
const pointerup = (e: PointerEvent) => {
  ;(e.target as HTMLElement).releasePointerCapture(e.pointerId)
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
      ref="canvascontainer"
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
      >
        {{ dialogues_dummy.get(dialogue_display.id) }}
      </div>
      <canvas ref="workcanvas" :style="canvasCss"></canvas>
      <div
        v-if="mode == 'dialogue'"
        v-for="dialogue_handle in dialogues_handle_display"
        class="dialoguehandle"
        :style="dialogue_handle.style"
      ></div>
    </div>
    <div class="subtool-container">
      <move-sub-tool
        v-if="mode == 'move'"
        @zoom-in="zoomIn"
        @zoom-out="zoomOut"
      ></move-sub-tool>
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
.canvas-container[data-editmode='move'] {
  cursor: grab;
}
.canvas-container[data-editmode='dialogue'] {
  cursor: crosshair;
}
.canvas-container[data-editmode='pen'] {
  cursor: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjwhLS0gQ3JlYXRlZCB3aXRoIElua3NjYXBlIChodHRwOi8vd3d3Lmlua3NjYXBlLm9yZy8pIC0tPgoKPHN2ZwogICB3aWR0aD0iMTUuOTk5OTk5IgogICBoZWlnaHQ9IjE1Ljk5OTk5OSIKICAgdmlld0JveD0iMCAwIDQuMjMzMzMzIDQuMjMzMzMzIgogICB2ZXJzaW9uPSIxLjEiCiAgIGlkPSJzdmc1IgogICB4bWxuczppbmtzY2FwZT0iaHR0cDovL3d3dy5pbmtzY2FwZS5vcmcvbmFtZXNwYWNlcy9pbmtzY2FwZSIKICAgeG1sbnM6c29kaXBvZGk9Imh0dHA6Ly9zb2RpcG9kaS5zb3VyY2Vmb3JnZS5uZXQvRFREL3NvZGlwb2RpLTAuZHRkIgogICB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiAgIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxzb2RpcG9kaTpuYW1lZHZpZXcKICAgICBpZD0ibmFtZWR2aWV3NyIKICAgICBwYWdlY29sb3I9IiNmZmZmZmYiCiAgICAgYm9yZGVyY29sb3I9IiMwMDAwMDAiCiAgICAgYm9yZGVyb3BhY2l0eT0iMC4yNSIKICAgICBpbmtzY2FwZTpzaG93cGFnZXNoYWRvdz0iMiIKICAgICBpbmtzY2FwZTpwYWdlb3BhY2l0eT0iMC4wIgogICAgIGlua3NjYXBlOnBhZ2VjaGVja2VyYm9hcmQ9IjAiCiAgICAgaW5rc2NhcGU6ZGVza2NvbG9yPSIjZDFkMWQxIgogICAgIGlua3NjYXBlOmRvY3VtZW50LXVuaXRzPSJtbSIKICAgICBzaG93Z3JpZD0iZmFsc2UiIC8+CiAgPGRlZnMKICAgICBpZD0iZGVmczIiIC8+CiAgPGcKICAgICBpbmtzY2FwZTpsYWJlbD0iTGF5ZXIgMSIKICAgICBpbmtzY2FwZTpncm91cG1vZGU9ImxheWVyIgogICAgIGlkPSJsYXllcjEiPgogICAgPGNpcmNsZQogICAgICAgc3R5bGU9ImZpbGw6IzAwMDAwMDtmaWxsLW9wYWNpdHk6MTtzdHJva2U6bm9uZTtzdHJva2Utd2lkdGg6Ny4yOTMxNTtzdHJva2UtbGluZWNhcDpyb3VuZDtzdHJva2UtbGluZWpvaW46cm91bmQ7cGFpbnQtb3JkZXI6c3Ryb2tlIGZpbGwgbWFya2VycyIKICAgICAgIGlkPSJwYXRoODk4IgogICAgICAgY3g9IjIuMTE2NjY2NiIKICAgICAgIGN5PSIyLjExNjY2NjYiCiAgICAgICByPSIxLjU4NzUiIC8+CiAgPC9nPgo8L3N2Zz4K)
      6 6,
    crosshair;
}
.canvas-container[data-editmode='eraser'] {
  cursor: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjwhLS0gQ3JlYXRlZCB3aXRoIElua3NjYXBlIChodHRwOi8vd3d3Lmlua3NjYXBlLm9yZy8pIC0tPgoKPHN2ZwogICB3aWR0aD0iMTIuOTk5OTk5IgogICBoZWlnaHQ9IjEyLjk5OTk5OSIKICAgdmlld0JveD0iMCAwIDMuNDM5NTgzIDMuNDM5NTgzIgogICB2ZXJzaW9uPSIxLjEiCiAgIGlkPSJzdmc1IgogICBpbmtzY2FwZTpleHBvcnQtZmlsZW5hbWU9ImVyYXNlci5zdmciCiAgIGlua3NjYXBlOmV4cG9ydC14ZHBpPSI5NiIKICAgaW5rc2NhcGU6ZXhwb3J0LXlkcGk9Ijk2IgogICB4bWxuczppbmtzY2FwZT0iaHR0cDovL3d3dy5pbmtzY2FwZS5vcmcvbmFtZXNwYWNlcy9pbmtzY2FwZSIKICAgeG1sbnM6c29kaXBvZGk9Imh0dHA6Ly9zb2RpcG9kaS5zb3VyY2Vmb3JnZS5uZXQvRFREL3NvZGlwb2RpLTAuZHRkIgogICB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiAgIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxzb2RpcG9kaTpuYW1lZHZpZXcKICAgICBpZD0ibmFtZWR2aWV3NyIKICAgICBwYWdlY29sb3I9IiNmZmZmZmYiCiAgICAgYm9yZGVyY29sb3I9IiMwMDAwMDAiCiAgICAgYm9yZGVyb3BhY2l0eT0iMC4yNSIKICAgICBpbmtzY2FwZTpzaG93cGFnZXNoYWRvdz0iMiIKICAgICBpbmtzY2FwZTpwYWdlb3BhY2l0eT0iMC4wIgogICAgIGlua3NjYXBlOnBhZ2VjaGVja2VyYm9hcmQ9IjAiCiAgICAgaW5rc2NhcGU6ZGVza2NvbG9yPSIjZDFkMWQxIgogICAgIGlua3NjYXBlOmRvY3VtZW50LXVuaXRzPSJtbSIKICAgICBzaG93Z3JpZD0iZmFsc2UiIC8+CiAgPGRlZnMKICAgICBpZD0iZGVmczIiIC8+CiAgPGcKICAgICBpbmtzY2FwZTpsYWJlbD0iTGF5ZXIgMSIKICAgICBpbmtzY2FwZTpncm91cG1vZGU9ImxheWVyIgogICAgIGlkPSJsYXllcjEiCiAgICAgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuMzk2ODc0OTEsLTAuMzk2ODc0OTEpIj4KICAgIDxjaXJjbGUKICAgICAgIHN0eWxlPSJmaWxsOiNmZmZmZmY7ZmlsbC1vcGFjaXR5OjE7c3Ryb2tlOiMwMjAyMDI7c3Ryb2tlLXdpZHRoOjAuMjY0NTgzO3N0cm9rZS1saW5lY2FwOnJvdW5kO3N0cm9rZS1saW5lam9pbjpyb3VuZDtzdHJva2UtZGFzaGFycmF5Om5vbmU7c3Ryb2tlLW9wYWNpdHk6MTtwYWludC1vcmRlcjpzdHJva2UgZmlsbCBtYXJrZXJzIgogICAgICAgaWQ9InBhdGg4OTgiCiAgICAgICBjeD0iMi4xMTY2NjY2IgogICAgICAgY3k9IjIuMTE2NjY2NiIKICAgICAgIHI9IjEuNTg3NSIgLz4KICA8L2c+Cjwvc3ZnPgo=)
      6 6,
    crosshair;
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
.dialogue:active,
.dialogue:focus,
.dialogue:focus-visible,
.dialogue div:active,
.dialogue div:focus,
.dialogue div:focus-visible {
  outline: none;
}
.dialoguehandle {
  position: absolute;
  background-color: #0099ff;
  border: 2px solid #BBB;
  box-sizing: border-box;
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
