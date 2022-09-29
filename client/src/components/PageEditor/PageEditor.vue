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
import DialogueContainer from './DialogueContainer.vue'

const canvascontainer = ref<HTMLCanvasElement>()
const workcanvas = ref<HTMLCanvasElement>()
const canvas = ref<HTMLCanvasElement>()
const workctx = ref<CanvasRenderingContext2D>()
const ctx = ref<CanvasRenderingContext2D>()

interface Props {
  pageId: string
}

const props = defineProps<Props>()
const loaded = ref(false)

const pageWidth = ref(707)
const pageHeight = ref(1000)

const canvasScrollX = ref<number>(0)
const canvasScrollY = ref<number>(0)
const canvasScale = ref<number>(0.5)

type EditMode = 'move' | 'pen' | 'dialogue' | 'eraser'
const mode = ref<EditMode>('pen')

type GetPageData = {
  width: number
  height: number
  dialogues: {
    id: string
    dialogue: string
    left: number
    top: number
    right: number
    bottom: number
  }[]
  lines: {
    id: string
    penSize: number
    points: {
      x: number
      y: number
      pressure: number
    }[]
  }[]
}

onMounted(() => {
  fetch(`/api/pages/${props.pageId}`)
    .then(res => res.json())
    .then((res: GetPageData) => {
      pageWidth.value = res.width
      pageHeight.value = res.height
      dialogues.value = res.dialogues.map(dialogue => {
        return {
          pageID: props.pageId,
          id: dialogue.id,
          left: dialogue.left,
          top: dialogue.top,
          right: dialogue.right,
          bottom: dialogue.bottom,
          dialogue: dialogue.dialogue,
          color: '#000000',
          fontName: '',
          fontSize: 24
        }
      })
      lines.push(
        ...res.lines.map(line => {
          return {
            pageID: props.pageId,
            id: line.id,
            color: '#606060',
            brushSize: line.penSize,
            path: line.points.map(point => {
              return {
                x: point.x,
                y: point.y,
                force: point.pressure
              }
            })
          }
        })
      )

      if (canvas.value && workcanvas.value && canvascontainer.value) {
        canvas.value.width = pageWidth.value * canvasScale.value
        canvas.value.height = pageHeight.value * canvasScale.value
        ctx.value = canvas.value.getContext('2d') ?? undefined

        workcanvas.value.width = pageWidth.value * canvasScale.value
        workcanvas.value.height = pageHeight.value * canvasScale.value
        workctx.value = workcanvas.value.getContext('2d') ?? undefined

        canvasScrollX.value =
          canvascontainer.value.clientWidth / 2 - canvas.value.clientWidth / 2
        canvasScrollY.value =
          canvascontainer.value.clientHeight / 2 - canvas.value.clientHeight / 2
        handler = getModeHandler()
      }
      if (ctx.value) {
        for (const line of lines) {
          drawLine(ctx.value, canvasScale.value, line)
        }
      }
      loaded.value = true
    })
})

const canvasCss = computed(() => {
  return {
    left: `${canvasScrollX.value}px`,
    top: `${canvasScrollY.value}px`,
    height: `${Math.floor(pageHeight.value * canvasScale.value)}px`,
    width: `${Math.floor(pageWidth.value * canvasScale.value)}px`,
    display: `${loaded.value ? 'inherit' : 'none'}`
  }
})

watch(canvasScale, () => {
  if (canvas.value && workcanvas.value && ctx.value) {
    canvas.value.width = pageWidth.value * canvasScale.value
    canvas.value.height = pageHeight.value * canvasScale.value
    workcanvas.value.width = pageWidth.value * canvasScale.value
    workcanvas.value.height = pageHeight.value * canvasScale.value
    for (const line of lines) {
      drawLine(ctx.value, canvasScale.value, line)
    }
  }
})

const lines: Line[] = []
const dialogues = ref<Dialogue[]>([])
const dialogue_selected = ref<string | null>(null)
const dialogue_update = (id: string, text: string,
  left: number,
  top: number,
  right: number,
  bottom: number,
  unsettle: boolean) => {
  const dialogue = dialogues.value.find(p => p.id == id)
  if (!dialogue) return
  dialogue.dialogue = text
  dialogue.left = left
  dialogue.right = right
  dialogue.top = top
  dialogue.bottom = bottom
  if(!unsettle)
  fetch(`/api/dialogues/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      dialogue: text,
      top: top,
      bottom: bottom,
      left: left,
      right: right
    })
  })
}
const dialogue_select = (id: string) => {
  dialogue_selected.value = id
}
const dialogue_delete = () => {
  dialogues.value = dialogues.value.filter(p => {
    return p.id != dialogue_selected.value
  })
  fetch(`/api/dialogues/${dialogue_selected.value}`, {
    method: 'DELETE'
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
      props.pageId
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
      lines,
      props.pageId
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
  <div class="pageeditor" :data-editmode="mode">
    <dialogue-container
      v-if="loaded"
      :dialogues="dialogues"
      :canvas-scroll-x="canvasScrollX"
      :canvas-scroll-y="canvasScrollY"
      :canvas-scale="canvasScale"
      :is-active="mode == 'dialogue'"
      @select-dialogue="dialogue_select"
      @update-dialogue="dialogue_update"
    ></dialogue-container>
    <div ref="canvascontainer" class="canvas-container">
      <canvas ref="canvas" class="store-canvas" :style="canvasCss"></canvas>
      <canvas ref="workcanvas" :style="canvasCss"></canvas>
    </div>
    <div
      class="canvas-dummy"
      @pointerdown="pointerdown"
      @pointermove="pointermove"
      @pointerup="pointerup"
    ></div>
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
  overflow: hidden;
}
.canvas-container {
  width: 100%;
  overflow: hidden;
  flex-grow: 1;
  position: relative;
  background-color: #ccc;
  z-index: -1;
}
[data-editmode='move'] {
  cursor: grab;
}
[data-editmode='dialogue'] {
  cursor: crosshair;
}
[data-editmode='pen'] {
  cursor: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjwhLS0gQ3JlYXRlZCB3aXRoIElua3NjYXBlIChodHRwOi8vd3d3Lmlua3NjYXBlLm9yZy8pIC0tPgoKPHN2ZwogICB3aWR0aD0iMTUuOTk5OTk5IgogICBoZWlnaHQ9IjE1Ljk5OTk5OSIKICAgdmlld0JveD0iMCAwIDQuMjMzMzMzIDQuMjMzMzMzIgogICB2ZXJzaW9uPSIxLjEiCiAgIGlkPSJzdmc1IgogICB4bWxuczppbmtzY2FwZT0iaHR0cDovL3d3dy5pbmtzY2FwZS5vcmcvbmFtZXNwYWNlcy9pbmtzY2FwZSIKICAgeG1sbnM6c29kaXBvZGk9Imh0dHA6Ly9zb2RpcG9kaS5zb3VyY2Vmb3JnZS5uZXQvRFREL3NvZGlwb2RpLTAuZHRkIgogICB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiAgIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxzb2RpcG9kaTpuYW1lZHZpZXcKICAgICBpZD0ibmFtZWR2aWV3NyIKICAgICBwYWdlY29sb3I9IiNmZmZmZmYiCiAgICAgYm9yZGVyY29sb3I9IiMwMDAwMDAiCiAgICAgYm9yZGVyb3BhY2l0eT0iMC4yNSIKICAgICBpbmtzY2FwZTpzaG93cGFnZXNoYWRvdz0iMiIKICAgICBpbmtzY2FwZTpwYWdlb3BhY2l0eT0iMC4wIgogICAgIGlua3NjYXBlOnBhZ2VjaGVja2VyYm9hcmQ9IjAiCiAgICAgaW5rc2NhcGU6ZGVza2NvbG9yPSIjZDFkMWQxIgogICAgIGlua3NjYXBlOmRvY3VtZW50LXVuaXRzPSJtbSIKICAgICBzaG93Z3JpZD0iZmFsc2UiIC8+CiAgPGRlZnMKICAgICBpZD0iZGVmczIiIC8+CiAgPGcKICAgICBpbmtzY2FwZTpsYWJlbD0iTGF5ZXIgMSIKICAgICBpbmtzY2FwZTpncm91cG1vZGU9ImxheWVyIgogICAgIGlkPSJsYXllcjEiPgogICAgPGNpcmNsZQogICAgICAgc3R5bGU9ImZpbGw6IzAwMDAwMDtmaWxsLW9wYWNpdHk6MTtzdHJva2U6bm9uZTtzdHJva2Utd2lkdGg6Ny4yOTMxNTtzdHJva2UtbGluZWNhcDpyb3VuZDtzdHJva2UtbGluZWpvaW46cm91bmQ7cGFpbnQtb3JkZXI6c3Ryb2tlIGZpbGwgbWFya2VycyIKICAgICAgIGlkPSJwYXRoODk4IgogICAgICAgY3g9IjIuMTE2NjY2NiIKICAgICAgIGN5PSIyLjExNjY2NjYiCiAgICAgICByPSIxLjU4NzUiIC8+CiAgPC9nPgo8L3N2Zz4K)
      6 6,
    crosshair;
}
[data-editmode='eraser'] {
  cursor: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjwhLS0gQ3JlYXRlZCB3aXRoIElua3NjYXBlIChodHRwOi8vd3d3Lmlua3NjYXBlLm9yZy8pIC0tPgoKPHN2ZwogICB3aWR0aD0iMTIuOTk5OTk5IgogICBoZWlnaHQ9IjEyLjk5OTk5OSIKICAgdmlld0JveD0iMCAwIDMuNDM5NTgzIDMuNDM5NTgzIgogICB2ZXJzaW9uPSIxLjEiCiAgIGlkPSJzdmc1IgogICBpbmtzY2FwZTpleHBvcnQtZmlsZW5hbWU9ImVyYXNlci5zdmciCiAgIGlua3NjYXBlOmV4cG9ydC14ZHBpPSI5NiIKICAgaW5rc2NhcGU6ZXhwb3J0LXlkcGk9Ijk2IgogICB4bWxuczppbmtzY2FwZT0iaHR0cDovL3d3dy5pbmtzY2FwZS5vcmcvbmFtZXNwYWNlcy9pbmtzY2FwZSIKICAgeG1sbnM6c29kaXBvZGk9Imh0dHA6Ly9zb2RpcG9kaS5zb3VyY2Vmb3JnZS5uZXQvRFREL3NvZGlwb2RpLTAuZHRkIgogICB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiAgIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgogIDxzb2RpcG9kaTpuYW1lZHZpZXcKICAgICBpZD0ibmFtZWR2aWV3NyIKICAgICBwYWdlY29sb3I9IiNmZmZmZmYiCiAgICAgYm9yZGVyY29sb3I9IiMwMDAwMDAiCiAgICAgYm9yZGVyb3BhY2l0eT0iMC4yNSIKICAgICBpbmtzY2FwZTpzaG93cGFnZXNoYWRvdz0iMiIKICAgICBpbmtzY2FwZTpwYWdlb3BhY2l0eT0iMC4wIgogICAgIGlua3NjYXBlOnBhZ2VjaGVja2VyYm9hcmQ9IjAiCiAgICAgaW5rc2NhcGU6ZGVza2NvbG9yPSIjZDFkMWQxIgogICAgIGlua3NjYXBlOmRvY3VtZW50LXVuaXRzPSJtbSIKICAgICBzaG93Z3JpZD0iZmFsc2UiIC8+CiAgPGRlZnMKICAgICBpZD0iZGVmczIiIC8+CiAgPGcKICAgICBpbmtzY2FwZTpsYWJlbD0iTGF5ZXIgMSIKICAgICBpbmtzY2FwZTpncm91cG1vZGU9ImxheWVyIgogICAgIGlkPSJsYXllcjEiCiAgICAgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTAuMzk2ODc0OTEsLTAuMzk2ODc0OTEpIj4KICAgIDxjaXJjbGUKICAgICAgIHN0eWxlPSJmaWxsOiNmZmZmZmY7ZmlsbC1vcGFjaXR5OjE7c3Ryb2tlOiMwMjAyMDI7c3Ryb2tlLXdpZHRoOjAuMjY0NTgzO3N0cm9rZS1saW5lY2FwOnJvdW5kO3N0cm9rZS1saW5lam9pbjpyb3VuZDtzdHJva2UtZGFzaGFycmF5Om5vbmU7c3Ryb2tlLW9wYWNpdHk6MTtwYWludC1vcmRlcjpzdHJva2UgZmlsbCBtYXJrZXJzIgogICAgICAgaWQ9InBhdGg4OTgiCiAgICAgICBjeD0iMi4xMTY2NjY2IgogICAgICAgY3k9IjIuMTE2NjY2NiIKICAgICAgIHI9IjEuNTg3NSIgLz4KICA8L2c+Cjwvc3ZnPgo=)
      6 6,
    crosshair;
}
.store-canvas {
  background-color: #fff;
}
.dialogue-container {
  position: relative;
}
.pageeditor[data-editmode='dialogue'] .dialogue-container {
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
.canvas-dummy {
  position: absolute;
  top: 0;
  width: 100%;
  height: 100%;
}
.button-container {
  height: 6rem;
  display: flex;
  z-index: 10000;
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
