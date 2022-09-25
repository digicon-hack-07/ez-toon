export type StringRenderData = {
    str: string,
    top: number,
    bottom: number,
    left: number,
    right: number,
    fontSize: number,
    horizontal?: boolean,
    fontType?: string,
};

export function drawString(ctx: CanvasRenderingContext2D, dat: StringRenderData){
    const fontname = dat.fontType ?? 'EzTooN-SourceHanSerif';
    const fontcolor = '#000000';
    const lineheight = 1.35;

    const hps_base = dat.right - dat.fontSize * (1 + (lineheight - 1) / 2);
    let hps =  hps_base;
    const vps_base = dat.top;
    let vps = vps_base;

    ctx.textBaseline = "bottom";
    ctx.strokeRect(dat.left, dat.top, dat.right - dat.left, dat.bottom - dat.top);

    for(let i = 0; i < dat.str.length; i++){
        ctx.font = `${dat.fontSize}px '${fontname}'`;
        ctx.fillStyle  = fontcolor;
        const chr = dat.str.charAt(i);
        const next = dat.str.charAt(i+1);
        const isHalfChar = chr.match(/[A-Za-z0-9"'\.\, \[\]:;!\?]/);
        const w = isHalfChar ? ctx.measureText(chr).width : dat.fontSize;

        if (chr == '\n' || vps + w > dat.bottom) {
            hps -= dat.fontSize * lineheight;
            vps = vps_base;
            if(chr == '\n')
                continue;
        }
        if (isHalfChar) {
            ctx.save();
            ctx.setTransform(0, 1, -1, 0, 0, 0);
            ctx.fillText(chr, vps, -hps);
            ctx.restore();
            vps += w;
        } else {
            ctx.fillText(chr, hps, vps + dat.fontSize); // 通常描画
            if (chr.match(/、|。/) && next.match(/」|』/)){
                vps += 0.5 * dat.fontSize;
            } else {
                vps += dat.fontSize;
            }
        }
    }
}
