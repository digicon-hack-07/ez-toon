import { Ref } from 'vue'
import { Line } from '../line'
import { drawLine, drawPath, PathRenderData } from '../renderer/path'
import { ToolHandlerInterface } from './ToolHandlerInterface'

export class PenToolHandler implements ToolHandlerInterface {
  working_lines: Map<number, Line>
  canvas: HTMLElement
  ctx: CanvasRenderingContext2D
  workctx: CanvasRenderingContext2D
  lines: Line[]

  constructor(
    canvas: HTMLElement,
    ctx: CanvasRenderingContext2D,
    workctx: CanvasRenderingContext2D,
    lines: Line[]
  ) {
    this.canvas = canvas
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
      color: '#000000',
      id: '', // TODO: generate ULID
      pageID: '', // TODO: pageID
      path: [
        {
          x: e.clientX - bx,
          y: e.clientY - by,
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
      x: e.clientX - bx,
      y: e.clientY - by,
      force: e.pressure
    })
    this.workctx.clearRect(
      0,
      0,
      this.workctx.canvas.width,
      this.workctx.canvas.height
    )
    drawLine(this.workctx, line)
    return
  }
  pointerup(e: PointerEvent): void {
    const line = this.working_lines.get(e.pointerId)
    if (!line) return

    this.workctx.clearRect(
      0,
      0,
      this.workctx.canvas.width,
      this.workctx.canvas.height
    )
    drawLine(this.ctx, line)
    this.lines.push(line)
    this.working_lines.delete(e.pointerId)
    return
  }
}
