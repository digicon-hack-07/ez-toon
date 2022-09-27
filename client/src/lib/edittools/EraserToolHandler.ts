import { Ref } from 'vue'
import { Line } from '../line'
import { drawLine } from '../renderer/path'
import { type ToolHandlerInterface } from './ToolHandlerInterface'

export class EraserToolHandler implements ToolHandlerInterface {
  lines: Line[]
  canvas: HTMLElement
  canvasScale: Ref<number>
  ctx: CanvasRenderingContext2D
  en: boolean = false

  constructor(
    canvas: HTMLElement,
    ctx: CanvasRenderingContext2D,
    canvasScale: Ref<number>,
    lines: Line[]
  ) {
    this.canvas = canvas
    this.ctx = ctx
    this.canvasScale = canvasScale
    this.lines = lines
  }
  pointerdown(e: PointerEvent): void {
    this.en = true
    return
  }
  pointermove(e: PointerEvent): void {
    if (!this.en) return
    const bx = this.canvas.getClientRects().item(0)?.left
    const by = this.canvas.getClientRects().item(0)?.top
    if (bx === undefined || by === undefined) return

    const nowX = (e.x - bx) / this.canvasScale.value
    const nowY = (e.y - by) / this.canvasScale.value

    const erased = this.lines.filter(line => {
      for (const v of line.path) {
        if ((v.x - nowX) * (v.x - nowX) + (v.y - nowY) * (v.y - nowY) < 25) // TODO: variable size
          return false
      }
      return true
    })
    if (this.lines.length != erased.length) {
      this.lines.length = 0
      for (const line of erased) this.lines.push(line)
      this.ctx.clearRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height)
      for (const line of this.lines) {
        drawLine(this.ctx, this.canvasScale.value, line)
      }
    }
    return
  }
  pointerup(e: PointerEvent): void {
    this.en = false
    return
  }
}
