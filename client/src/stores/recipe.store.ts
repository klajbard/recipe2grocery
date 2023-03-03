import { reactive } from "vue";
import { Ingredient } from "../models";

interface RecipeStore {
  ingredients: Ingredient[];
  addIngredient: () => void;
  removeIngredient: (index: number) => void;
  description: { value: string }[];
}

export const recipeStore = reactive<RecipeStore>({
  ingredients: [{ name: "", amount: 0, unit: "" }],
  description: [{ value: "" }],
  addIngredient() {
    this.ingredients.push({ name: "", amount: 0, unit: "" });
  },
  removeIngredient(index) {
    this.ingredients.splice(index, 1);
  },
});
