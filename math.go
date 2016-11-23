package main

type Vec struct {
	X, Y float32
}

type Rect struct {
	TL, BR Vec
}

func (v Vec) Add(o Vec) Vec {
	return Vec{
		v.X + o.X,
		v.Y + o.Y,
	}
}

func (v Vec) Sub(o Vec) Vec {
	return Vec{
		v.X - o.X,
		v.Y - o.Y,
	}
}

func (v Vec) Mul(o Vec) Vec {
	return Vec{
		v.X * o.X,
		v.Y * o.Y,
	}
}

func (v Vec) Addf(f float32) Vec {
	return Vec{
		v.X + f,
		v.Y + f,
	}
}

func (v Vec) Subf(f float32) Vec {
	return Vec{
		v.X - f,
		v.Y - f,
	}
}

func (v Vec) Mulf(f float32) Vec {
	return Vec{
		v.X * f,
		v.Y * f,
	}
}

func (v Vec) Clamp(min, max float32) Vec {
	return Vec{
		Clamp(v.X, min, max),
		Clamp(v.Y, min, max),
	}
}

func (v Vec) Lerp(o Vec, t float32) Vec {
	return Vec{
		v.X + t*(o.X-v.X),
		v.Y + t*(o.Y-v.Y),
	}
}

func (a Rect) Contains(p Vec) bool {
	return p.X > a.TL.X &&
		p.Y > a.TL.Y &&
		p.X < a.BR.X &&
		p.Y < a.BR.Y
}

func (a Rect) Intersects(b Rect) bool {
	return a.TL.X < b.BR.X &&
		a.BR.X > b.TL.X &&
		a.TL.Y < b.BR.Y &&
		a.BR.Y > b.TL.Y
}

func Clamp(x, min, max float32) float32 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
