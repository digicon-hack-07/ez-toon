import { type Ref } from 'vue'
import { type ToolHandlerInterface } from './ToolHandlerInterface'
import { type Dialogue } from '../dialogue'

export class DialogueToolHandler implements ToolHandlerInterface {
  draggingData: {
    pointerId: number
    beginX: number
    beginY: number
  } | null = null

  canvas: HTMLElement
  dialogues: Ref<Dialogue[]>
  pageID: string
  canvasScale: Ref<number>

  constructor(canvas: HTMLElement, canvasScale: Ref<number>, dialogues: Ref<Dialogue[]>, pageID: string) {
    this.canvas = canvas
    this.canvasScale = canvasScale
    this.dialogues = dialogues
    this.pageID = pageID
  }
  pointerdown(e: PointerEvent): void {
    const bx = this.canvas.getClientRects().item(0)?.left
    const by = this.canvas.getClientRects().item(0)?.top
    if (bx === undefined || by === undefined) return
    if (
      this.dialogues.value.find(dialogue => {
        return (
          dialogue.left <= (e.clientX - bx) / this.canvasScale.value &&
          (e.clientX - bx) / this.canvasScale.value <= dialogue.right &&
          dialogue.top <= (e.clientY - by) / this.canvasScale.value &&
          (e.clientY - by) / this.canvasScale.value <= dialogue.bottom
        )
      })
    )
      return
    this.draggingData = {
      pointerId: e.pointerId,
      beginX: e.clientX - bx,
      beginY: e.clientY - by
    }
    return
  }
  pointermove(e: PointerEvent): void {
    return
  }
  pointerup(e: PointerEvent): void {
    const bx = this.canvas.getClientRects().item(0)?.left
    const by = this.canvas.getClientRects().item(0)?.top
    if (bx === undefined || by === undefined) return
    if (!this.draggingData || this.draggingData.pointerId != e.pointerId) return
    {
      const left = Math.min(e.clientX - bx, this.draggingData.beginX) / this.canvasScale.value
      const right = Math.max(e.clientX - bx, this.draggingData.beginX) / this.canvasScale.value
      const top = Math.min(e.clientY - by, this.draggingData.beginY) / this.canvasScale.value
      const bottom = Math.max(e.clientY - by, this.draggingData.beginY) / this.canvasScale.value
      if (right - left < 24 || bottom - top < 24) return

      this.dialogues.value.push({
        id: `${
          this.dialogues.value.length
            ? this.dialogues.value.slice(-1)[0].id + 1
            : 0
        }`, // TODO: generate ULID
        pageID: this.pageID,
        dialogue: '',
        left,
        top,
        right,
        bottom,
        fontSize: 24,
        fontName: '',
        color: '#000000'
      })
    }
    this.draggingData = null
    return
  }
}
