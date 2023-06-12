package lang

import (
	"bufio"
	"errors"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/gonnagetbetter/architecture-lab-3/painter"
	"golang.org/x/exp/shiny/screen"
)

type State struct {
	BgColor painter.OperationFunc
	Rect    painter.OperationFunc
	Figures []*painter.Figure
}

type Parser struct {
	State State
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation
	
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), ",")
		if len(commands) == 0 {
			continue
		}
		
		for _, val := range commands {
			cmd := strings.Fields(val)
			
			switch cmd[0] {
			case "white":
				p.State.BgColor = painter.OperationFunc(painter.WhiteFill)
			
			case "green":
				p.State.BgColor = painter.OperationFunc(painter.GreenFill)

			case "update":
				if p.State.BgColor != nil {
					res = append(res, p.State.BgColor)
				}
				
				if p.State.Rect != nil {
					res = append(res, p.State.Rect)
				}

				for ind, figureInstance := range p.State.Figures {
					res = append(res, figureInstance.DrawFigure())
					fmt.Printf("FigureInstance %d: X: %d, Y: %d\n", ind, figureInstance.X, figureInstance.Y)
				}

				res = append(res, painter.UpdateOp)

			case "bgrect":
				if len(cmd) != 5 {
					return nil, errors.New("invalid number of arguments for bgrect")
				}

				x1, err1 := strconv.ParseFloat(cmd[1], 64)
				y1, err2 := strconv.ParseFloat(cmd[2], 64)
				x2, err3 := strconv.ParseFloat(cmd[3], 64)
				y2, err4 := strconv.ParseFloat(cmd[4], 64)

				if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
					return nil, errors.New("invalid arguments for bgrect")
				}

				x1Int := int(x1 * 800)
				y1Int := int(y1 * 800)
				x2Int := int(x2 * 800)
				y2Int := int(y2 * 800)

				p.State.Rect = painter.BgRect(x1Int, y1Int, x2Int, y2Int)

			case "figure":
				if len(cmd) != 3 {
					return nil, errors.New("invalid number of arguments for figure")
				}

				x, err1 := strconv.ParseFloat(cmd[1], 64)
				y, err2 := strconv.ParseFloat(cmd[2], 64)

				if err1 != nil || err2 != nil {
					return nil, errors.New("invalid arguments for figure")
				}

				xInt := int(x * 800)
				yInt := int(y * 800)

				p.State.Figures = append(p.State.Figures, &painter.Figure{
					X: xInt,
					Y: yInt,
				})

			case "move":
				if len(cmd) != 3 {
					return nil, errors.New("invalid number of arguments for move")
				}
	
				dx, err1 := strconv.ParseFloat(cmd[1], 64)
				dy, err2 := strconv.ParseFloat(cmd[2], 64)

				if err1 != nil || err2 != nil {
					return nil, errors.New("invalid arguments for move")
				}
				
				dxInt := int(dx * 800)
				dyInt := int(dy * 800)

				for _, figureInstance := range p.State.Figures {
					figureInstance.MoveFigure(dxInt, dyInt)
				}

			case "reset":
				p.State.BgColor = painter.OperationFunc(func(t screen.Texture) {
					t.Fill(t.Bounds(), color.Black, screen.Src)
				})
				p.State.Rect = painter.BgRect(0, 0, 0, 0)
				p.State.Figures = nil
			}
		}
	}

	return res, nil
}
