package examples

import (
	"generator/blocks"
	"generator/generator"
)

func HelloWorld() {

	G := generator.NewGenerator()


	Println := blocks.NewFunction().SetName("fmt.Println").AddInput(blocks.NewVariable().SetValue(blocks.NewString("Hello, World!"))).Executing()

	Hello := blocks.NewFunction().SetName("HelloWorld").AppendBlock(Println)

	Main := blocks.NewFunction().SetName("main").AppendBlock(Hello.Executing())

	S := blocks.NewSkeleton().AppendBlocks(Hello, Main)

	I := blocks.NewImport().SetSource("fmt")

	H := blocks.NewHeading().SetPackageName("main").AppendImports(I)

	F := blocks.NewFile().SetFileName("Hello World").SetHeading(H).SetSkeleton(S)


	G.GenerateFile(F)

}
