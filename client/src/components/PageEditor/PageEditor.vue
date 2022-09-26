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

type EditMode = 'pen' | 'dialogue' | 'eraser';
const mode = ref<EditMode>('pen');

type Path = {
    x: number,
    y: number,
}[];

const paths :Path[] = [];
const working_path = new Map<number, Path>();

const drawPath = (path: Path) => {
    switch(mode.value){
        case 'pen':
            if(ctx.value){
                ctx.value.beginPath();
                ctx.value.moveTo(path[0].x, path[1].y);
                for(const pos of path){
                    ctx.value.lineTo(pos.x, pos.y);
                }
                ctx.value.stroke();
                ctx.value.closePath();
            }
            break;
    }
}

const pointerdown = (e: PointerEvent) => {
    switch(mode.value){
        case 'pen':
            const bx = canvas.value?.getClientRects().item(0)?.left;
            const by = canvas.value?.getClientRects().item(0)?.top;
            if(!bx || !by)
                return;

            working_path.set(e.pointerId, [{
                x: e.clientX - bx,
                y: e.clientY - by,
            }]);
            break;
    }
}
const pointermove = (e: PointerEvent) => {
    switch(mode.value){
        case 'pen':
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
            break;
    }
}
const pointerup = (e: PointerEvent) => {
    const path = working_path.get(e.pointerId);
    if(!path)
        return;

    paths.push(path);
    working_path.delete(e.pointerId);
}

const selectModeDialogue = () => {
    mode.value = 'dialogue';
}
const selectModePen = () => {
    mode.value = 'pen';
}
const selectModeEraser = () => {
    mode.value = 'eraser';
}

</script>

<template>
<div class="pageeditor">
    <div class="canvas-container">
        <canvas ref="canvas" @pointerdown="pointerdown" @pointermove="pointermove" @pointerup="pointerup"></canvas>
    </div>
    <div class="button-container">
        <button :data-active="mode == 'dialogue'" @click="selectModeDialogue">あ</button>
        <button :data-active="mode == 'pen'" @click="selectModePen">筆</button>
        <button :data-active="mode == 'eraser'" @click="selectModeEraser">消</button>
    </div>
</div>
</template>

<style scoped>
@import '../../font.css';

.pageeditor {
    display: flex;
    flex-direction: column;
}
.canvas-container {
    width: 100%;
    overflow: hidden;
    flex-grow: 1;
}
canvas {
    touch-action: none;
    border: 1px solid;
}
.button-container {
    height: 6rem;
    display: flex;
}
button {
    flex-grow: 1;
    border: 2px solid #888;
    border-bottom: none;
    border-radius: 0;
}
button:focus{
    outline: none;
}
button[data-active="true"]{
    background-color: #3c70ff;
    color: #FFF;
}

</style>
