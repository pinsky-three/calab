export const openConnection = (url: string): (() => void) => {
    const ws = new WebSocket(url);
    const el = document.getElementById("imgcanvas") as HTMLImageElement;

    ws.onmessage = (event: MessageEvent<Blob>) => {
        el.src = URL.createObjectURL(event.data);
    };

    return () => ws.close();
};
