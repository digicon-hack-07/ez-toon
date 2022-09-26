import { Ref } from 'vue'
import { ToolHandlerInterface } from './ToolHandlerInterface'

export class MoveToolHandler implements ToolHandlerInterface {
  draggingData: {
    pointerId: number
    beginX: number
    beginY: number
    tgtBeginX: number
    tgtBeginY: number
  } | null = null
  canvasScrollX: Ref<number>
  canvasScrollY: Ref<number>

  constructor(canvasScrollX: Ref<number>, canvasScrollY: Ref<number>) {
    this.canvasScrollX = canvasScrollX
    this.canvasScrollY = canvasScrollY
  }
  pointerdown(e: PointerEvent): void {
    if (this.draggingData) return
    this.draggingData = {
      pointerId: e.pointerId,
      beginX: e.x,
      beginY: e.y,
      tgtBeginX: this.canvasScrollX.value,
      tgtBeginY: this.canvasScrollY.value
    }
    return
  }
  pointermove(e: PointerEvent): void {
    if (!this.draggingData || this.draggingData.pointerId != e.pointerId) return
    this.canvasScrollX.value =
      e.x - this.draggingData.beginX + this.draggingData.tgtBeginX
    this.canvasScrollY.value =
      e.y - this.draggingData.beginY + this.draggingData.tgtBeginY
    return
  }
  pointerup(e: PointerEvent): void {
    if (!this.draggingData || this.draggingData.pointerId != e.pointerId) return
    this.canvasScrollX.value =
      e.x - this.draggingData.beginX + this.draggingData.tgtBeginX
    this.canvasScrollY.value =
      e.y - this.draggingData.beginY + this.draggingData.tgtBeginY
    this.draggingData = null
    return
  }
}
