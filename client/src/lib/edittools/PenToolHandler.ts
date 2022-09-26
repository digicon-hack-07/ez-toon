import { Ref } from 'vue'
import { drawPath, PathRenderData } from '../renderer/path'
import { ToolHandlerInterface } from './ToolHandlerInterface'

export class PenToolHandler implements ToolHandlerInterface {
  working_path: Map<number, PathRenderData>
  canvas: HTMLElement
  ctx: CanvasRenderingContext2D
  workctx: CanvasRenderingContext2D
  paths: PathRenderData[]

  constructor(
    canvas: HTMLElement,
    ctx: CanvasRenderingContext2D,
    workctx: CanvasRenderingContext2D,
    paths: PathRenderData[]
  ) {
    this.canvas = canvas
    this.ctx = ctx
    this.workctx = workctx
    this.working_path = new Map<number, PathRenderData>()
    this.paths = paths
  }
  pointerdown(e: PointerEvent): void {
    const bx = this.canvas.getClientRects().item(0)?.left
    const by = this.canvas.getClientRects().item(0)?.top
    if (bx === undefined || by === undefined) return
    this.working_path.set(e.pointerId, [
      {
        x: e.clientX - bx,
        y: e.clientY - by
      }
    ])
    return
  }
  pointermove(e: PointerEvent): void {
    const bx = this.canvas.getClientRects().item(0)?.left
    const by = this.canvas.getClientRects().item(0)?.top
    if (bx === undefined || by === undefined) return
    const path = this.working_path.get(e.pointerId)
    if (!path) return

    path.push({
      x: e.clientX - bx,
      y: e.clientY - by
    })
    this.workctx.clearRect(
      0,
      0,
      this.workctx.canvas.width,
      this.workctx.canvas.height
    )
    drawPath(this.workctx, path)
    return
  }
  pointerup(e: PointerEvent): void {
    const path = this.working_path.get(e.pointerId)
    if (!path) return

    this.workctx.clearRect(
      0,
      0,
      this.workctx.canvas.width,
      this.workctx.canvas.height
    )
    drawPath(this.ctx, path)
    this.paths.push(path)
    this.working_path.delete(e.pointerId)
    return
  }
}
