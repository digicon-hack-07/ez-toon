<script setup lang="ts">

import { onMounted, ref } from 'vue';

const canvas = ref<HTMLCanvasElement>();
const ctx = ref<CanvasRenderingContext2D>();

onMounted(() => {
    if(canvas.value){
        const width = canvas.value.getClientRects().item(0)?.width;
        const height = canvas.value.getClientRects().item(0)?.height;
        if(width && height){
            canvas.value.width = width;
            canvas.value.height = height;
        }        
        ctx.value = canvas.value.getContext("2d") ?? undefined;
    }
});

type Path = {
    x: number,
    y: number,
}[];

const paths :Path[] = [];
const working_path = new Map<number, Path>();

const drawPath = (path: Path) => {
    if(ctx.value){
        ctx.value.beginPath();
        ctx.value.moveTo(path[0].x, path[1].y);
        for(const pos of path){
            ctx.value.lineTo(pos.x, pos.y);
        }
        ctx.value.stroke();
        ctx.value.closePath();
    }
}

const pointerdown = (e: PointerEvent) => {
    const bx = canvas.value?.getClientRects().item(0)?.left;
    const by = canvas.value?.getClientRects().item(0)?.top;
    if(!bx || !by)
        return;

    working_path.set(e.pointerId, [{
        x: e.clientX - bx,
        y: e.clientY - by,
    }]);
}
const pointermove = (e: PointerEvent) => {
    const path = working_path.get(e.pointerId);
    if(!path)
        return;

    const bx = canvas.value?.getClientRects().item(0)?.left;
    const by = canvas.value?.getClientRects().item(0)?.top;
    if(!bx || !by)
        return;

    path.push({
        x: e.clientX - bx,
        y: e.clientY - by,
    });
    drawPath(path);
}
const pointerup = (e: PointerEvent) => {
    const path = working_path.get(e.pointerId);
    if(!path)
        return;

    paths.push(path);
    working_path.delete(e.pointerId);
}
    
</script>

<template>
<div>
    <canvas ref="canvas" @pointerdown="pointerdown" @pointermove="pointermove" @pointerup="pointerup"></canvas>
</div>
</template>

<style scoped>

canvas {
    height: 500px;
    width: 500px;
    touch-action: none;
    border: 1px solid;
}

</style>
