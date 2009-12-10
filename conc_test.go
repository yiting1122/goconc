package conc

import (
	"testing";
)

func numbers(max int) chan Box {
	out := make(chan Box);
	go func() {
		for i:=0; i<max; i++ {
			out <- i;
		}
		close(out);
	}();
	return out;
}

func TestFor(t *testing.T) {
	var vals [20]int;
	wait := For(numbers(20), 3, func(i Box) {vals[i.(int)] = 1});
	wait();
	
	total := 0;
	for i:=0; i<20; i++ {
		total += vals[i];
	}
	
	if total != 20 {
		t.Fail();
	}
}

func TestMap(t *testing.T) {
	incr := func(a Box) Box {
		return a.(int)+1;
	};
	incrNumbers := Map(incr, numbers(20));
	for i:=0; i<20; i++ {
		j := <- incrNumbers;
		if i+1 != j.(int) {
			t.Fail();
		}
	}
}

func TestReduce(t *testing.T) {
	sum := func(a Box, b Box) Box {
		return a.(int)+b.(int);
	};
	totalSum := Reduce(sum, numbers(10), 0);

	if totalSum.(int) != 45 {
		t.Fail();
	}
}

func TestFilter(t *testing.T) {	
	results := Filter(func(i Box) bool { return i.(int)%2==0 }, numbers(10));
	trueRes := make([]bool, 10);
	for i := range results {
		trueRes[i.(int)] = true;
	}
	for i,v := range trueRes {
		if v != (i%2==0) {
			t.Fail();
		}
	}
}

func TestMapReduce(t *testing.T) {
	
	incr := func(a Box) Box {
		return a.(int)+1;
	};
	sum := func(a Box, b Box) Box {
		return a.(int)+b.(int);
	};
	
	result := Reduce(sum, MapUnordered(incr, numbers(10)), 0);

	if result.(int) != 55 {
		t.Fail();
	}
}