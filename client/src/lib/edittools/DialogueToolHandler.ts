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

  constructor(canvas: HTMLElement, dialogues: Ref<Dialogue[]>, pageID: string) {
    this.canvas = canvas
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
          dialogue.left <= e.clientX - bx &&
          e.clientX - bx <= dialogue.right &&
          dialogue.top <= e.clientY - by &&
          e.clientY - by <= dialogue.bottom
        )
      })
    )
      return
    this.draggingData = {
      pointerId: e.pointerId,
      beginX: e.x - bx,
      beginY: e.y - by
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
      const left = Math.min(e.clientX - bx, this.draggingData.beginX)
      const right = Math.max(e.clientX - bx, this.draggingData.beginX)
      const top = Math.min(e.clientY - by, this.draggingData.beginY)
      const bottom = Math.max(e.clientY - by, this.draggingData.beginY)
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
