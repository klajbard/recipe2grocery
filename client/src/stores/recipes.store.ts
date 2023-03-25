import { reactive } from 'vue'
import { LSKeys, Recipe } from '../models'
import { IngredientDetails } from '../models'

interface RecipeStore {
  recipes: Recipe[]
  excludes: Set<string>
  ingredients: { [key: string]: IngredientDetails[] }
  addRecipe: (recipe: Recipe) => void
  removeRecipe: (recipe: Recipe) => void
  addRecipeIngredients: (recipe: Recipe) => void
  recalculateIngredients: () => void
  readLS: () => void
  setLS: () => void
}

export const recipesStore = reactive<RecipeStore>({
  recipes: [],
  excludes: new Set(),
  ingredients: {},
  addRecipe(recipe: Recipe) {
    this.recipes.push(recipe)
    this.addRecipeIngredients(recipe)
    this.setLS()
  },
  addRecipeIngredients(recipe: Recipe) {
    recipe.ingredients.forEach((ingredient) => {
      const lowecareName = ingredient.name.toLowerCase()
      if (!Object.hasOwn(this.ingredients, lowecareName)) {
        this.ingredients[lowecareName] = [
          {
            amount: ingredient.amount,
            unit: ingredient.unit,
          },
        ]
      } else {
        const ingredientWithSameUnit = this.ingredients[lowecareName].find(({ unit }) => unit === ingredient.unit)

        if (ingredientWithSameUnit) {
          ingredientWithSameUnit.amount += ingredient.amount
        } else {
          this.ingredients[lowecareName].push({
            amount: ingredient.amount,
            unit: ingredient.unit,
          })
        }
      }
    })
  },
  recalculateIngredients() {
    this.ingredients = {}
    this.recipes.forEach((recipe) => this.addRecipeIngredients(recipe))
  },
  removeRecipe(recipe: Recipe) {
    const elementIndex = this.recipes.findIndex(({ id }) => id === recipe.id)
    this.recipes.splice(elementIndex, 1)
    this.recalculateIngredients()
    this.setLS()
  },
  readLS() {
    const recipes = JSON.parse(
      decodeURIComponent(atob(window.localStorage.getItem(LSKeys.RECIPES) || '[]'))
    ) as Recipe[]

    if (recipes.length) {
      this.recipes = recipes
      this.recalculateIngredients()
    }
  },
  setLS() {
    window.localStorage.setItem(LSKeys.RECIPES, btoa(encodeURIComponent(JSON.stringify(this.recipes))))
  },
})
