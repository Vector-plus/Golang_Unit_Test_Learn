package designpatterns

//生成器，辅助生成一个复杂的对象，定义生成该对象的方法，同时可以使用一个主管对象执行生成该对象（例如：获取用户信息可能需要用到多个模块进行信息整合，设计的逻辑可能比较复杂）

//生产一台电脑，cpu，cache

type Computer struct {
	CPUType     string
	CacheType   string
	Description string
}

type IBuilder interface {
	setDescription()
	getComputer() Computer
}

func getBuilder(cpuType, cacheType string) (builder IBuilder, err error) {
	builder = &ComputerBuilder{
		CPUType:   cpuType,
		CacheType: cacheType,
	}
	return
}

type ComputerBuilder struct {
	CPUType     string
	CacheType   string
	Description string
}

func (builder *ComputerBuilder) setDescription() {
	builder.Description = builder.CPUType + "---^^(-_-)^^--" + builder.CacheType
}

func (builder *ComputerBuilder) getComputer() Computer {
	return Computer{
		CPUType:     builder.CPUType,
		CacheType:   builder.CPUType,
		Description: builder.Description,
	}
}

// 主管，可选项
type Director struct {
	builder IBuilder
}

func newDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) setBuilder(builder IBuilder) {
	d.builder = builder
}

func (d *Director) makeComputer() Computer {
	d.builder.setDescription()
	return d.builder.getComputer()
}
