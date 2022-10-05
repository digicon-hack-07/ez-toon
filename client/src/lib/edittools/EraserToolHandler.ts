import { type Ref } from 'vue'
import { type Line } from '../line'
import { drawLine } from '../renderer/path'
import { type ToolHandlerInterface } from './ToolHandlerInterface'

function lineIntersect(
  l1ax: number,
  l1ay: number,
  l1bx: number,
  l1by: number,
  l2ax: number,
  l2ay: number,
  l2bx: number,
  l2by: number
) {
  if ((l1ax == l1bx && l1ay == l1by) || (l2ax == l2bx && l2ay == l2by))
    return false
  const s = (l1ax - l1bx) * (l2ay - l1ay) - (l1ay - l1by) * (l2ax - l1ax)
  const t = (l1ax - l1bx) * (l2by - l1ay) - (l1ay - l1by) * (l2bx - l1ax)
  if (s * t > 0) return false

  const p = (l2ax - l2bx) * (l1ay - l2ay) - (l2ay - l2by) * (l1ax - l2ax)
  const q = (l2ax - l2bx) * (l1by - l2ay) - (l2ay - l2by) * (l1bx - l2ax)
  if (p * q > 0) return false

  if (s == t || p == q)
    // s = t = 0, p = q = 0
    return (
      Math.min(l1ax, l1bx) <= Math.max(l2ax, l2bx) &&
      Math.min(l2ax, l2bx) <= Math.max(l1ax, l1bx) &&
      Math.min(l1ay, l1by) <= Math.max(l2ay, l2by) &&
      Math.min(l2ay, l2by) <= Math.max(l1ay, l1by)
    )

  return true
}

export class EraserToolHandler implements ToolHandlerInterface {
  lines: Line[]
  canvas: HTMLElement
  canvasScale: Ref<number>
  ctx: CanvasRenderingContext2D
  en = false

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

    const prevX = (e.x - bx - e.movementX) / this.canvasScale.value
    const prevY = (e.y - by - e.movementY) / this.canvasScale.value
    const nowX = (e.x - bx) / this.canvasScale.value
    const nowY = (e.y - by) / this.canvasScale.value

    const erasing = new Set<string>()
    for (const line of this.lines) {
      let pPrev = line.path[0]
      for (const v of line.path) {
        if (
          lineIntersect(prevX, prevY, nowX, nowY, pPrev.x, pPrev.y, v.x, v.y)
        ) {
          erasing.add(line.id)
          break
        }
        pPrev = v
      }
    }

    if (erasing.size > 0) {
      erasing.forEach(id => {
        fetch(`/api/lines/${id}`, {
          method: 'DELETE'
        })
      })
      const erased = this.lines.filter(line => !erasing.has(line.id))
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
