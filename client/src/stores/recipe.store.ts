import { reactive } from "vue";
import { Ingredient } from "../models";

interface RecipeStore {
  ingredients: Ingredient<number | null>[];
  addIngredient: () => void;
  removeIngredient: (index: number) => void;
  description: { value: string }[];
}

export const recipeStore = reactive<RecipeStore>({
  ingredients: [{ name: "", amount: null, unit: "" }],
  description: [{ value: "" }],
  addIngredient() {
    this.ingredients.push({ name: "", amount: null, unit: "" });
  },
  removeIngredient(index) {
    this.ingredients.splice(index, 1);
  },
});
