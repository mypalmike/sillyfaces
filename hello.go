package main

import (
    "github.com/llgcode/draw2d/draw2dimg"
    "github.com/llgcode/draw2d/draw2dkit"
    "image"
    "image/color"
    "math/rand"
    "time"
)

func main() {
    seed := time.Now().UnixNano()
    randSource := rand.NewSource(seed)
    rnd := rand.New(randSource)

    // Initialize the graphic context on an RGBA image
    dest := image.NewRGBA(image.Rect(0, 0, 400, 400))
    gc := draw2dimg.NewGraphicContext(dest)

    draw(gc, rnd);

    // Save to file
    draw2dimg.SaveToPngFile("hello.png", dest)
}

func draw(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
    drawEars(gc, rnd)
    drawFace(gc, rnd)
    drawNose(gc, rnd)
    drawEyes(gc, rnd)
    drawMouth(gc, rnd)
    drawHair(gc, rnd)
}

func drawEars(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
    fillR := uint8(32 + rnd.Intn(224))
    fillG := uint8(32 + rnd.Intn(224))
    fillB := uint8(32 + rnd.Intn(224))

    strokeR := uint8(rnd.Intn(256))
    strokeG := uint8(rnd.Intn(256))
    strokeB := uint8(rnd.Intn(256))

    lineWidth := float64(rnd.Intn(5))

    gc.SetFillColor(color.RGBA{fillR, fillG, fillB, 0xff})
    gc.SetStrokeColor(color.RGBA{strokeR, strokeB, strokeG, 0xff})
    gc.SetLineWidth(lineWidth)

    offsetY := rnd.Intn(60)
    y := float64(200 - offsetY)

    // Ellipse
    radiusW := float64(150 + rnd.Intn(40))
    radiusH := float64(10 + rnd.Intn(80))
    draw2dkit.Ellipse(gc, 200, y, radiusW, radiusH)

    gc.FillStroke()
}

func drawFace(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
    fillR := uint8(32 + rnd.Intn(224))
    fillG := uint8(32 + rnd.Intn(224))
    fillB := uint8(32 + rnd.Intn(224))

    strokeR := uint8(rnd.Intn(256))
    strokeG := uint8(rnd.Intn(256))
    strokeB := uint8(rnd.Intn(256))

    lineWidth := float64(rnd.Intn(5))

    gc.SetFillColor(color.RGBA{fillR, fillG, fillB, 0xff})
    gc.SetStrokeColor(color.RGBA{strokeR, strokeB, strokeG, 0xff})
    gc.SetLineWidth(lineWidth)

    shape := rnd.Intn(3)

    switch shape {
    case 0:
        // Circle
        radius := float64(150 + rnd.Intn(30))
        draw2dkit.Circle(gc, 200, 200, radius)
    case 1:
        // Ellipse
        radiusW := float64(100 + rnd.Intn(80))
        radiusH := float64(160 + rnd.Intn(20))
        draw2dkit.Ellipse(gc, 200, 200, radiusW, radiusH)
    default:
        // Rounded rect
        halfW := 100 + rnd.Intn(80)
        halfH := 160 + rnd.Intn(20)
        x1 := float64(200 - halfW)
        y1 := float64(200 - halfH)
        x2 := float64(200 + halfW)
        y2 := float64(200 + halfH)
        arcWidth := float64(20 + rnd.Intn(40))
        arcHeight := float64(20 + rnd.Intn(40))
        draw2dkit.RoundedRectangle(gc, x1, y1, x2, y2, arcWidth, arcHeight)
    }

    gc.FillStroke()
}

func drawNose(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
    fillR := uint8(32 + rnd.Intn(224))
    fillG := uint8(32 + rnd.Intn(224))
    fillB := uint8(32 + rnd.Intn(224))

    strokeR := uint8(rnd.Intn(256))
    strokeG := uint8(rnd.Intn(256))
    strokeB := uint8(rnd.Intn(256))

    lineWidth := float64(rnd.Intn(5))

    gc.SetFillColor(color.RGBA{fillR, fillG, fillB, 0xff})
    gc.SetStrokeColor(color.RGBA{strokeR, strokeB, strokeG, 0xff})
    gc.SetLineWidth(lineWidth)

    offsetY := rnd.Intn(40)

    radiusX := 5 + rnd.Intn(30)
    radiusY := 5 + rnd.Intn(50)

    topY := float64(200 + offsetY - radiusY)
    bottomY := float64(200 + offsetY + radiusY)
    leftX := float64(200 - radiusX)
    rightX := float64(200 + radiusX)

    gc.MoveTo(float64(200), float64(topY))
    gc.LineTo(float64(leftX), float64(bottomY))
    gc.LineTo(float64(rightX), float64(bottomY))
    gc.LineTo(float64(200), float64(topY))

    gc.FillStroke()
}

func drawEyes(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
    fillR := uint8(rnd.Intn(128))
    fillG := uint8(rnd.Intn(128))
    fillB := uint8(rnd.Intn(128))

    strokeR := uint8(224 + rnd.Intn(32))
    strokeG := uint8(224 + rnd.Intn(32))
    strokeB := uint8(224 + rnd.Intn(32))

    lineWidth := 20 + float64(rnd.Intn(5))

    gc.SetFillColor(color.RGBA{fillR, fillG, fillB, 0xff})
    gc.SetStrokeColor(color.RGBA{strokeR, strokeB, strokeG, 0xff})
    gc.SetLineWidth(lineWidth)

    offsetX := 40 + rnd.Intn(100)
    offsetY := rnd.Intn(50)

    radius := float64(30 + rnd.Intn(30))

    centerX1 := float64(200 - offsetX)
    centerX2 := float64(200 + offsetX)
    centerY := float64(200 - offsetY)

    draw2dkit.Circle(gc, centerX1, centerY, radius)
    gc.FillStroke()

    draw2dkit.Circle(gc, centerX2, centerY, radius)
    gc.FillStroke()
}

func drawMouth(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
    fillR := uint8(200 + rnd.Intn(56))
    fillG := uint8(rnd.Intn(32))
    fillB := uint8(rnd.Intn(32))

    strokeR := uint8(rnd.Intn(256))
    strokeG := uint8(rnd.Intn(256))
    strokeB := uint8(rnd.Intn(256))

    lineWidth := float64(rnd.Intn(5))

    gc.SetFillColor(color.RGBA{fillR, fillG, fillB, 0xff})
    gc.SetStrokeColor(color.RGBA{strokeR, strokeB, strokeG, 0xff})
    gc.SetLineWidth(lineWidth)

    offsetY := 80 + rnd.Intn(50)
    maxHeight := 4 + rnd.Intn(50)
    radius := float64(30 + rnd.Intn(130))

    nzags := rnd.Intn(10)
    deltaX := 2.0 * radius / float64(nzags);
    updown := rnd.Intn(2) * 2 - 1 // -1 or 1
    x := 200 - radius
    y := 200 + offsetY + (rnd.Intn(maxHeight) * updown)

    gc.MoveTo(float64(x), float64(y))
    for zag := 0; zag < nzags; zag++ {
        y = 200 + offsetY + (rnd.Intn(maxHeight) * updown)

        gc.LineTo(float64(x), float64(y))

        updown *= -1
        x += deltaX        
    }

    y = 200 + offsetY + (rnd.Intn(maxHeight) * updown)
    gc.LineTo(float64(x), float64(y))

    gc.FillStroke()
}

func drawHair(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
    drawStraightHair(gc, rnd)
    //hairtype := rnd.Intn(3)

    // switch hairtype {
    //     case 0:
    //         drawCurlyHair(gc, rnd)
    //     case 1:
    //         drawStraightHair(gc, rnd)
    //     default:
    //         // Bald. Do nothing
    // }
}

func drawCurlyHair(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
}

func drawStraightHair(gc *draw2dimg.GraphicContext, rnd *rand.Rand) {
    fillR := uint8(20 + rnd.Intn(32))
    fillG := uint8(20 + rnd.Intn(32))
    fillB := uint8(20 + rnd.Intn(32))

    strokeR := uint8(rnd.Intn(256))
    strokeG := uint8(rnd.Intn(256))
    strokeB := uint8(rnd.Intn(256))

    lineWidth := float64(rnd.Intn(3))

    gc.SetFillColor(color.RGBA{fillR, fillG, fillB, 0xff})
    gc.SetStrokeColor(color.RGBA{strokeR, strokeB, strokeG, 0xff})
    gc.SetLineWidth(lineWidth)

    startX := float64(30 + rnd.Intn(206))
    startY := float64(30 + rnd.Intn(30))

    startLength := float64(30 + rnd.Intn(180))
    deltaLength := float64(10 - rnd.Intn(21))

    length := startLength

    for x := 10; x < 390; x+= 20 {

        gc.MoveTo(startX, startY)
        gc.LineTo(float64(x), length)
        gc.LineTo(float64(x + 10), length + deltaLength)
        gc.LineTo(startX, startY)

        gc.FillStroke()

        length += deltaLength
    }
}
