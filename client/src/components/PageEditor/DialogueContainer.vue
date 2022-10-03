<script setup lang="ts">
import { computed, ref } from 'vue'
import { type Dialogue } from '../../lib/dialogue'

const handleSize = 2
const handleDisplaySize = 1
const minDialogueSize = 24

const props = defineProps<{
  dialogues: Dialogue[]
  canvasScrollX: number
  canvasScrollY: number
  canvasScale: number
  isActive: boolean
}>()

const emit = defineEmits<{
  (e: 'selectDialogue', id: string): void
  (e: 'updateDialogue', id: string, text: string,
    left: number,
    top: number,
    right: number,
    bottom: number,
    unsettle: boolean): void
}>()

const dialogue_container = ref()

const dialogues_dummy = ref(new Map<string, string>(
  props.dialogues.map(dialogue => {
    return [dialogue.id, dialogue.dialogue]
  })
))

const dialogue_select = (e: FocusEvent) => {
  if (e.target instanceof HTMLElement) {
    const id = e.target.dataset.id
    if (!id) return
    emit('selectDialogue', id)
  }
}

const dialogues_display = computed(() => {
  return props.dialogues.map(dialogue => {
    return {
      id: dialogue.id,
      style: {
        left: `${props.canvasScrollX + dialogue.left * props.canvasScale}px`,
        top: `${props.canvasScrollY + dialogue.top * props.canvasScale}px`,
        width: `${dialogue.right - dialogue.left}px`,
        height: `${dialogue.bottom - dialogue.top}px`,
        fontSize: `${dialogue.fontSize}px`,
        transform: `scale(${props.canvasScale})`
      },
      str: dialogue.dialogue
    }
  })
})
const dialogues_handle_display = computed(() => {
  return props.dialogues.map(dialogue => {
    return {
      id: dialogue.id,
      style: {
        left: `calc(${
          props.canvasScrollX + dialogue.right * props.canvasScale + 4
        }px - ${handleSize / 2 - handleDisplaySize / 2}rem)`,
        top: `calc(${
          props.canvasScrollY + dialogue.top * props.canvasScale - 4
        }px - ${handleDisplaySize / 2 + handleSize / 2}rem)`,
        width: `${handleSize}rem`,
        height: `${handleSize}rem`
      }
    }
  })
})
const dialogues_resizehandle_display = computed(() => {
  return props.dialogues.map(dialogue => {
    return {
      id: dialogue.id,
      style: {
        left: `calc(${
          props.canvasScrollX + dialogue.left * props.canvasScale - 4
        }px - ${handleDisplaySize / 2 + handleSize / 2}rem)`,
        top: `calc(${
          props.canvasScrollY + dialogue.bottom * props.canvasScale + 4
        }px + ${handleDisplaySize / 2 - handleSize / 2}rem)`,
        width: `${handleSize}rem`,
        height: `${handleSize}rem`
      }
    }
  })
})
const dialogue_update = (e: Event) => {
  const tgt = e.target
  if (!tgt) return
  if (e instanceof InputEvent && tgt instanceof HTMLElement) {
    if (!tgt.dataset.id) return
    const id = tgt.dataset.id
    const dialogue = props.dialogues.find(p => p.id == id)
    if (!dialogue) return
    emit('updateDialogue', tgt.dataset.id, tgt.innerText, dialogue.left, dialogue.top, dialogue.right, dialogue.bottom, false)
  }
}

const handlerDragInfo = new Map<
  number,
  {
    beginX: number
    beginY: number
    targetBeginX: number
    targetBeginY: number
    id: string
    targetW: number
    targetH: number
  }
>()

const dialogue_handler_pointerdown = (e: PointerEvent) => {
  const tgt = e.currentTarget
  if (!tgt) return
  if (tgt instanceof HTMLElement) {
    if (!tgt.dataset.id) return
    const id = tgt.dataset.id
    const dialogue = props.dialogues.find(p => p.id == id)
    if (!dialogue) return

    tgt.setPointerCapture(e.pointerId)
    handlerDragInfo.set(e.pointerId, {
      beginX: e.clientX,
      beginY: e.clientY,
      targetBeginX: dialogue.right,
      targetBeginY: dialogue.top,
      id: id,
      targetW: dialogue.right - dialogue.left,
      targetH: dialogue.bottom - dialogue.top
    })
  }
}
const dialogue_handler_pointermove = (e: PointerEvent) => {
  const drag = handlerDragInfo.get(e.pointerId)
  if (!drag) return
  const dialogue = props.dialogues.find(p => p.id == drag.id)
  if (!dialogue) return
  const top = (e.clientY - drag.beginY) / props.canvasScale + drag.targetBeginY
  const right =
    (e.clientX - drag.beginX) / props.canvasScale + drag.targetBeginX
  const left = right - drag.targetW
  const bottom = top + drag.targetH
  emit('updateDialogue', drag.id, dialogue.dialogue, left, top, right, bottom, true)
}
const dialogue_handler_pointerup = (e: PointerEvent) => {
  const tgt = e.currentTarget
  if (!tgt) return
  if (tgt instanceof HTMLElement) {
    const drag = handlerDragInfo.get(e.pointerId)
    if (!drag) return
    const dialogue = props.dialogues.find(p => p.id == drag.id)
    if (!dialogue) return
    const top = (e.clientY - drag.beginY) / props.canvasScale + drag.targetBeginY
    const right =
      (e.clientX - drag.beginX) / props.canvasScale + drag.targetBeginX
    const left = right - drag.targetW
    const bottom = top + drag.targetH
    emit('updateDialogue', drag.id, dialogue.dialogue, left, top, right, bottom, false)

    handlerDragInfo.delete(e.pointerId)
    tgt.releasePointerCapture(e.pointerId)
  }
}
const dialogue_resizehandler_pointerdown = (e: PointerEvent) => {
  const tgt = e.currentTarget
  if (!tgt) return
  if (tgt instanceof HTMLElement) {
    if (!tgt.dataset.id) return
    const id = tgt.dataset.id
    const dialogue = props.dialogues.find(p => p.id == id)
    if (!dialogue) return

    tgt.setPointerCapture(e.pointerId)
    handlerDragInfo.set(e.pointerId, {
      beginX: e.clientX,
      beginY: e.clientY,
      targetBeginX: dialogue.left,
      targetBeginY: dialogue.bottom,
      id: id,
      targetW: dialogue.right - dialogue.left,
      targetH: dialogue.bottom - dialogue.top
    })
  }
}
const dialogue_resizehandler_pointermove = (e: PointerEvent) => {
  const drag = handlerDragInfo.get(e.pointerId)
  if (!drag) return
  const dialogue = props.dialogues.find(p => p.id == drag.id)
  if (!dialogue) return
  const top = dialogue.top
  const right = dialogue.right
  const left = Math.min((e.clientX - drag.beginX) / props.canvasScale + drag.targetBeginX, dialogue.right - minDialogueSize)
  const bottom = Math.max((e.clientY - drag.beginY) / props.canvasScale + drag.targetBeginY, dialogue.top + minDialogueSize)
  emit('updateDialogue', drag.id, dialogue.dialogue, left, top, right, bottom, true)
}
const dialogue_resizehandler_pointerup = (e: PointerEvent) => {
  const tgt = e.currentTarget
  if (!tgt) return
  if (tgt instanceof HTMLElement) {
    const drag = handlerDragInfo.get(e.pointerId)
    if (!drag) return
    const dialogue = props.dialogues.find(p => p.id == drag.id)
    if (!dialogue) return
    const top = dialogue.top
    const right = dialogue.right
    const left = Math.min((e.clientX - drag.beginX) / props.canvasScale + drag.targetBeginX, dialogue.right - minDialogueSize)
    const bottom = Math.max((e.clientY - drag.beginY) / props.canvasScale + drag.targetBeginY, dialogue.top + minDialogueSize)
    emit('updateDialogue', drag.id, dialogue.dialogue, left, top, right, bottom, false)

    handlerDragInfo.delete(e.pointerId)
    tgt.releasePointerCapture(e.pointerId)
  }
}
</script>

<template>
  <div ref="dialogue_container" class="dialogue-container">
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
    <div v-if="isActive">
      <div
        v-for="dialogue_handle in dialogues_handle_display"
        :key="dialogue_handle.id"
        class="dialoguehandle"
        :style="dialogue_handle.style"
        :data-id="dialogue_handle.id"
        @pointerdown="dialogue_handler_pointerdown"
        @pointermove="dialogue_handler_pointermove"
        @pointerup="dialogue_handler_pointerup"
      >
        <div class="dialoguehandle-disp"></div>
      </div>
      <div
        v-for="dialogue_handle in dialogues_resizehandle_display"
        :key="dialogue_handle.id"
        class="dialoguehandle"
        :style="dialogue_handle.style"
        :data-id="dialogue_handle.id"
        @pointerdown="dialogue_resizehandler_pointerdown"
        @pointermove="dialogue_resizehandler_pointermove"
        @pointerup="dialogue_resizehandler_pointerup"
      ></div>
    </div>
  </div>
</template>

<style scoped>
.dialogue {
  position: absolute;
  border: 2px solid #999; /* 小さいと縮小時にブラウザのレンダリングがバグりやすい */
  writing-mode: vertical-rl;
  line-height: 1.2;
  word-break: break-all;
  font-family: 'EzTooN-SourceHanSerif', serif;
  text-align: start;
  transform-origin: top left;
  box-sizing: border-box;
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
  touch-action: none;
  cursor: move;
  
  background-color: #0099ff;
  border: 2px solid #bbb;
  box-sizing: border-box;
}
/* .dialoguehandle::before {
  content: "";
  display: block;
  background-color: #0099ff;
  border: 2px solid #bbb;
  box-sizing: border-box;
  width: 1rem;
  height: 1rem;
  margin: 0.5rem auto;
  touch-action: none;
} */
</style>
