import shape = require("./IShape")
import circle = require("./Circle")
import triangle = require("./Triangle")

function drawAllShapes(shapeToDraw: shape.IShape){
    shapeToDraw.draw()
}


drawAllShapes(new circle.Circle())
drawAllShapes(new triangle.Triangle())
