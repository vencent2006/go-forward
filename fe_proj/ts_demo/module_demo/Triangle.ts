import shape = require("./IShape")
export class Triangle implements shape.IShape {
    draw() {
        console.log("Triangle Draw")
    }
}
