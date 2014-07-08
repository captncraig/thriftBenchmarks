namespace * bench

struct Person{
	1: string Name,
	2: i32 Age,
}

struct NoStrings{
	1: i32 Abc,
	2: i64 Def,
	3: list<i16> Xyz,
}

struct Complex{
	1: string Foo,
	2: i32 Bar,
	3: list<i32> Baz,
	4: map<string,string> Abc,
	5: list<Person> People,
}

struct CalculationData{
	1: i32 a,
	2: i32 b,
}

service BenchmarkService{
	i32 Add(1: CalculationData d)
}