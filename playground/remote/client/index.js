import { openConnection } from "./receiver";

window.onunload = openConnection("ws://localhost:3000");
