import { type Line } from "../line"

export function drawLine(ctx: CanvasRenderingContext2D, line: Line) {
  ctx.beginPath()
  ctx.moveTo(line.path[0].x, line.path[0].y)
  for (const pos of line.path) {
    ctx.lineTo(pos.x, pos.y)
  }
  ctx.stroke()
  ctx.closePath()
}
