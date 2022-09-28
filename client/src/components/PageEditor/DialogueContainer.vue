<script setup lang="ts">
import { computed, ref } from 'vue'
import { type Dialogue } from '../../lib/dialogue'

const props = defineProps<{
  dialogues: Dialogue[]
  canvasScrollX: number
  canvasScrollY: number
  canvasScale: number
  isActive: boolean
}>()

const emit = defineEmits<{
  (e: 'selectDialogue', id: string): void
  (e: 'updateDialogue', id: string, text: string): void
}>()

const dialogues_dummy = ref(new Map<string, string>())

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
        left: `${
          props.canvasScrollX + dialogue.right * props.canvasScale + 4
        }px`,
        top: `${
          props.canvasScrollY + dialogue.top * props.canvasScale - 12 - 4
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
    if (!tgt.dataset.id) return
    emit('updateDialogue', tgt.dataset.id, tgt.innerText)
  }
}
</script>

<template>
  <div class="dialogue-container">
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
  border: 2px solid #bbb;
  box-sizing: border-box;
}
</style>
