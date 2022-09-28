import { type Ref } from 'vue'
import { type Line } from '../line'
import { drawLine } from '../renderer/path'
import { type ToolHandlerInterface } from './ToolHandlerInterface'

export class PenToolHandler implements ToolHandlerInterface {
  working_lines: Map<number, Line>
  canvas: HTMLElement
  ctx: CanvasRenderingContext2D
  workctx: CanvasRenderingContext2D
  lines: Line[]
  canvasScale: Ref<number>

  constructor(
    canvas: HTMLElement,
    canvasScale: Ref<number>,
    ctx: CanvasRenderingContext2D,
    workctx: CanvasRenderingContext2D,
    lines: Line[]
  ) {
    this.canvas = canvas
    this.canvasScale = canvasScale
    this.ctx = ctx
    this.workctx = workctx
    this.working_lines = new Map<number, Line>()
    this.lines = lines
  }
  pointerdown(e: PointerEvent): void {
    const bx = this.canvas.getClientRects().item(0)?.left
    const by = this.canvas.getClientRects().item(0)?.top
    if (bx === undefined || by === undefined) return
    this.working_lines.set(e.pointerId, {
      brushSize: 5,
      color: '#606060',
      id: `tmp${e.pointerId}`,
      pageID: '', // TODO: pageID
      path: [
        {
          x: (e.clientX - bx) / this.canvasScale.value,
          y: (e.clientY - by) / this.canvasScale.value,
          force: e.pressure
        }
      ]
    })
    return
  }
  pointermove(e: PointerEvent): void {
    const bx = this.canvas.getClientRects().item(0)?.left
    const by = this.canvas.getClientRects().item(0)?.top
    if (bx === undefined || by === undefined) return
    const line = this.working_lines.get(e.pointerId)
    if (!line) return

    line.path.push({
      x: (e.clientX - bx) / this.canvasScale.value,
      y: (e.clientY - by) / this.canvasScale.value,
      force: e.pressure
    })
    this.workctx.clearRect(
      0,
      0,
      this.workctx.canvas.width,
      this.workctx.canvas.height
    )
    drawLine(this.workctx, this.canvasScale.value, line)
    return
  }
  pointerup(e: PointerEvent): void {
    const line = this.working_lines.get(e.pointerId)
    if (!line) return
    if (line.path.length > 2){
      this.workctx.clearRect(
        0,
        0,
        this.workctx.canvas.width,
        this.workctx.canvas.height
      )
      drawLine(this.ctx, this.canvasScale.value, line)
      fetch(`/api/lines`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          // TODO
        })
      })
      .then(res => res.json())
      .then(res => {
        line.id = res.id
      })
      this.lines.push(line)
    }
    this.working_lines.delete(e.pointerId)
    return
  }
}
