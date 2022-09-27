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
    this.dialogues.value.push({
      id: `${
        this.dialogues.value.length
          ? this.dialogues.value.slice(-1)[0].id + 1
          : 0
      }`, // TODO: generate ULID
      pageID: this.pageID,
      dialogue: '',
      left: (e.clientX - bx) / this.canvasScale.value,
      top: (e.clientY - by) / this.canvasScale.value,
      right: (e.clientX - bx) / this.canvasScale.value,
      bottom: (e.clientY - by) / this.canvasScale.value,
      fontSize: 24,
      fontName: '',
      color: '#000000'
    })
    return
  }
  pointermove(e: PointerEvent): void {
    const bx = this.canvas.getClientRects().item(0)?.left
    const by = this.canvas.getClientRects().item(0)?.top
    if (bx === undefined || by === undefined) return
    if (!this.draggingData || this.draggingData.pointerId != e.pointerId) return
    {
      const left = Math.min(e.clientX - bx, this.draggingData.beginX) / this.canvasScale.value
      const right = Math.max(e.clientX - bx, this.draggingData.beginX) / this.canvasScale.value
      const top = Math.min(e.clientY - by, this.draggingData.beginY) / this.canvasScale.value
      const bottom = Math.max(e.clientY - by, this.draggingData.beginY) / this.canvasScale.value

      const dialogue = this.dialogues.value.slice(-1)[0]
      dialogue.left = left
      dialogue.top = top
      dialogue.right = right
      dialogue.bottom = bottom
    }
    return
  }
  pointerup(e: PointerEvent): void {
    if (!this.draggingData || this.draggingData.pointerId != e.pointerId) return
    const dialogue = this.dialogues.value.slice(-1)[0]
    if(dialogue.right - dialogue.left < 24 && dialogue.bottom - dialogue.top < 24)
      this.dialogues.value.pop()
    this.draggingData = null
    return
  }
}
