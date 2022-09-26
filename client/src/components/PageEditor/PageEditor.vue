<script setup lang="ts">

import { onMounted, ref } from 'vue';

const canvas = ref<HTMLCanvasElement>();
const ctx = ref<CanvasRenderingContext2D>();

const pageWidth = 1000;
const pageHeight = 1000;

onMounted(() => {
    if(canvas.value){
        const width = canvas.value.getClientRects().item(0)?.width;
        const height = canvas.value.getClientRects().item(0)?.height;
        if(width && height){
            canvas.value.width = pageWidth;
            canvas.value.height = pageHeight;
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
            if(bx === undefined || by === undefined)
                return;

            working_path.set(e.pointerId, [{
                x: e.clientX - bx,
                y: e.clientY - by,
            }]);
            break;
        case 'dialogue':
            
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
            if(bx === undefined || by === undefined)
                return;

            path.push({
                x: e.clientX - bx,
                y: e.clientY - by,
            });
            drawPath(path);
            break;
        case 'dialogue':
            break;
    }
}
const pointerup = (e: PointerEvent) => {
    switch(mode.value){
        case 'pen':
            const path = working_path.get(e.pointerId);
            if(!path)
                return;

            paths.push(path);
            working_path.delete(e.pointerId);
        case 'dialogue':
            break;
    }
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
    <div class="canvas-container" :data-editmode="mode">
        <div class="dialogue" contenteditable="true">aaa</div>
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
    width: 100vw;
    height: 100vh;
}
.canvas-container {
    width: 100%;
    overflow: hidden;
    flex-grow: 1;
    position: relative;
}
.dialogue {
    position: absolute;
    border: 1px solid #999;
    outline: none;
    writing-mode: vertical-rl;
    line-height: 1.2;
    word-break: break-all;
    font-family: 'EzTooN-SourceHanSerif', serif;
    z-index: -1;
}
.canvas-container[data-editmode="dialogue"] .dialogue {
    z-index: auto;
}
canvas {
    touch-action: none;
    border: 1px solid;
    z-index: -1;
}
.canvas-container[data-editmode="pen"] canvas {
    z-index: auto;
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
