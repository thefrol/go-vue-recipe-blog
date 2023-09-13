package data

type Recipe struct {
	Name string   `json:"name" yaml:"name" comment:"заголовок рецепта"`
	Text string   `json:"text" yaml:"text" comment:"текст рецепта в mardown разметке"`
	Tags []string `json:"tags" yaml:"tags" comment:"теги"`
}
