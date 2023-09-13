package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thefrol/go-vue-recipe-blog/internal/data"
)

type RecipesResponse struct {
	Recipes []data.Recipe `json:"recipes"`
}

func RecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes := []data.Recipe{
		{
			Name: "Быстро сырники",
			Text: "баночка обезжиренного йогурта 120г, 1 яйцо, 2 столовые ложки муки.",
			Tags: []string{"Пароварка"},
		},
		{
			Name: "Банан и яйцо",
			Text: "Один банан и одно яйцо взбить блендером и пожарить на сковородке. Охрененно ещё сверху намазать творогом",
			Tags: []string{"Сковорода", "Блины"},
		},
		{
			Name: "Быстро блинчики",
			Text: "Стакан молока, яйцо, 2 столовые ложки муки. Взбить в ешейкере пожарить.",
			Tags: []string{"Сковородка"},
		},
		// TODO
		// нейросеть считает каллорийность этих блюд по рецепту))
	}

	response := RecipesResponse{Recipes: recipes}
	bb, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "appliation/json")
	w.Write([]byte(bb))
}
