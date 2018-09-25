package pointers

func PointerOperation()  {
	var a	int	=	10;
	var points *int;
	var p *int;

	points	=	&a;

	println("Memory address a is",&a);
	println("value of points is",points)

	println(p==nil);
}
