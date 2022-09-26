export type Line = {
    id : string
    pageID : string
    color : string
    brushSize : number
    path : {
        x : number,
        y : number,
        force : number,
    }[]
}