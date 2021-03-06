package main

import (
	"sync"
	"github.com/fresh4less/neopixel-display/neopixeldisplay"
)

type InteractivePanel struct {
	frame neopixeldisplay.ColorFrame
	hue		uint32
	width	int
	height	int
	x	int
	y	int
	mu *sync.Mutex
}

func NewInteractivePanel(width, height int) *InteractivePanel {

	ip := new(InteractivePanel)
	
	ip.frame = neopixeldisplay.MakeColorFrame(width,height,neopixeldisplay.MakeColor(0,0,0))
	ip.hue = 0
	ip.width = width
	ip.height = height
	ip.x = 0
	ip.y = 0
	
	ip.mu = &sync.Mutex{}
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColorHue(ip.hue), neopixeldisplay.Error)
	
	return ip
}

func (ip *InteractivePanel) GetColorFrame() neopixeldisplay.ColorFrame {
	ip.mu.Lock()
	defer ip.mu.Unlock()
	
	return ip.frame
}

func (ip *InteractivePanel) IncrementHue() {
	ip.mu.Lock()
	defer ip.mu.Unlock()
	
	ip.hue = (ip.hue + 5) % 256
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColorHue(ip.hue), neopixeldisplay.Error)
}

func (ip *InteractivePanel) DecrementHue() {
	ip.mu.Lock()
	defer ip.mu.Unlock()
	
	ip.hue = ip.hue - 5
	if ip.hue < 0 {
		ip.hue = 255
	}
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColorHue(ip.hue), neopixeldisplay.Error)
}

func (ip *InteractivePanel) MoveRight() {
	ip.mu.Lock()
	defer ip.mu.Unlock()
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColor(0,0,0), neopixeldisplay.Error)
	
	ip.x = (ip.x + 1) % ip.width
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColorHue(ip.hue), neopixeldisplay.Error)
}

func (ip *InteractivePanel) MoveLeft() {
	ip.mu.Lock()
	defer ip.mu.Unlock()
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColor(0,0,0), neopixeldisplay.Error)
	
	ip.x = ip.x - 1
	if ip.x < 0 {
		ip.x = ip.width - 1
	}
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColorHue(ip.hue), neopixeldisplay.Error)
}

func (ip *InteractivePanel) MoveDown() {
	ip.mu.Lock()
	defer ip.mu.Unlock()
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColor(0,0,0), neopixeldisplay.Error)
	
	ip.y = (ip.y + 1) % ip.height
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColorHue(ip.hue), neopixeldisplay.Error)
}

func (ip *InteractivePanel) MoveUp() {
	ip.mu.Lock()
	defer ip.mu.Unlock()
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColor(0,0,0), neopixeldisplay.Error)
	
	ip.y = ip.y - 1
	if ip.y < 0 {
		ip.y = ip.height - 1
	}
	
	ip.frame.Set(ip.x,ip.y,neopixeldisplay.MakeColorHue(ip.hue), neopixeldisplay.Error)
}
