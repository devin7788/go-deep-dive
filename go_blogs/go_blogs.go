package go_blogs

func RecordGoBlogs() {
	var httpUrl = []string{
		"https://www.golanglearn.com/",
		"https://colobu.com/",
	}
	for _, v := range httpUrl {
		println(v)
	}

}
