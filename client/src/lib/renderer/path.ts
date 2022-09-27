import { type Line } from "../line"

export function drawLine(ctx: CanvasRenderingContext2D, scale: number, line: Line) {
  ctx.lineWidth = line.brushSize * scale
  ctx.strokeStyle = line.color
  ctx.beginPath()
  ctx.moveTo(line.path[0].x * scale, line.path[0].y * scale)
  for (const pos of line.path) {
    ctx.lineTo(pos.x * scale, pos.y * scale)
  }
  ctx.stroke()
  ctx.closePath()
}
