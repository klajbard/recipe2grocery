import { reactive } from "vue";
import { Recipe } from "../models";
import { IngredientDetails } from "../models";

interface RecipeStore {
  recipes: Recipe[];
  ingredients: { [key: string]: IngredientDetails[] };
  addRecipe: (recipe: Recipe) => void;
  removeRecipe: (recipe: Recipe) => void;
  addRecipeIngredients: (recipe: Recipe) => void;
  recalculateIngredients: () => void;
}

export const recipesStore = reactive<RecipeStore>({
  recipes: [],
  ingredients: {},
  addRecipe(recipe: Recipe) {
    this.recipes.push(recipe);
    this.addRecipeIngredients(recipe);
  },
  addRecipeIngredients(recipe: Recipe) {
    recipe.ingredients.forEach((ingredient) => {
      const lowecareName = ingredient.name.toLowerCase();
      if (!Object.hasOwn(this.ingredients, lowecareName)) {
        this.ingredients[lowecareName] = [
          {
            amount: ingredient.amount,
            unit: ingredient.unit,
          },
        ];
      } else {
        const ingredientWithSameUnit = this.ingredients[lowecareName].find(
          ({ unit }) => unit === ingredient.unit
        );

        if (ingredientWithSameUnit) {
          ingredientWithSameUnit.amount += ingredient.amount;
        } else {
          this.ingredients[lowecareName].push({
            amount: ingredient.amount,
            unit: ingredient.unit,
          });
        }
      }
    });
  },
  recalculateIngredients() {
    this.ingredients = {};
    this.recipes.forEach((recipe) => this.addRecipeIngredients(recipe));
  },
  removeRecipe(recipe: Recipe) {
    const elementIndex = this.recipes.findIndex(({ id }) => id === recipe.id);
    this.recipes.splice(elementIndex, 1);
    this.recalculateIngredients();
  },
});
