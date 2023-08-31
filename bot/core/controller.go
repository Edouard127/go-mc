package core

type Controller struct {
	Left, Right, Forward, Back float64
	LeftImpulse, RightImpulse  float64
	Jump, Sneak, Sprint, Drop  bool
}

func (c *Controller) Reset() {
	c.Left, c.Right, c.Forward, c.Back = 0, 0, 0, 0
	c.Jump, c.Sneak, c.Sprint, c.Drop = false, false, false, false
}
