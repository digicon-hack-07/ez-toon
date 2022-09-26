import { Line } from "../line"

export type PathRenderData = {
  x: number
  y: number
}[]

export function drawPath(ctx: CanvasRenderingContext2D, path: PathRenderData) {
  ctx.beginPath()
  ctx.moveTo(path[0].x, path[0].y)
  for (const pos of path) {
    ctx.lineTo(pos.x, pos.y)
  }
  ctx.stroke()
  ctx.closePath()
}

export function drawLine(ctx: CanvasRenderingContext2D, line: Line) {
  ctx.beginPath()
  ctx.moveTo(line.path[0].x, line.path[0].y)
  for (const pos of line.path) {
    ctx.lineTo(pos.x, pos.y)
  }
  ctx.stroke()
  ctx.closePath()
}
